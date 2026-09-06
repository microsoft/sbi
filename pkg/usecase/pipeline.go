//     MIT License
//
//     Copyright (c) Microsoft Corporation.
//
//     Permission is hereby granted, free of charge, to any person obtaining a copy
//     of this software and associated documentation files (the "Software"), to deal
//     in the Software without restriction, including without limitation the rights
//     to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//     copies of the Software, and to permit persons to whom the Software is
//     furnished to do so, subject to the following conditions:
//
//     The above copyright notice and this permission notice shall be included in all
//     copies or substantial portions of the Software.
//
//     THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//     IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//     FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//     AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//     LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//     OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//     SOFTWARE

package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	"github.com/microsoft/sbi/pkg/infrastructure/report"
	"github.com/microsoft/sbi/pkg/infrastructure/scanner"
	log "github.com/sirupsen/logrus"
)

type imageAnalyzer interface {
	Analyze(imageName string) (*domain.ImageAnalysis, error)
}

type tagDiscoverer interface {
	GetTags(repo string) ([]string, error)
}

// Pipeline orchestrates the full scan and report generation workflow.
type Pipeline struct {
	config   domain.ScanConfig
	repoCfg  domain.RepositoryConfig
	repo     *database.Repository
	analyzer imageAnalyzer
	registry tagDiscoverer
}

// NewPipeline creates a new Pipeline.
func NewPipeline(config domain.ScanConfig, repo *database.Repository) (*Pipeline, error) {
	repoCfg, err := LoadRepositoryConfig(config.ConfigDir)
	if err != nil {
		return nil, fmt.Errorf("loading repository config: %w", err)
	}

	// CLI flag overrides config default if explicitly set
	maxTags := config.MaxTagsPerRepo
	if maxTags == 0 && repoCfg.Defaults.MaxTags > 0 {
		maxTags = repoCfg.Defaults.MaxTags
	}

	config.MaxTagsPerRepo = maxTags

	return &Pipeline{
		config:   config,
		repoCfg:  repoCfg,
		repo:     repo,
		analyzer: scanner.NewImageAnalyzer(config.ComprehensiveScan, config.CleanupImages),
		registry: scanner.NewRegistryScanner(repoCfg.Defaults.Registry),
	}, nil
}

// scanStats tracks per-operation outcomes for ScanAll.
type scanStats struct {
	attempted int
	failed    int
}

// result returns an error when every attempt failed. Partial failure is OK.
func (s *scanStats) result() error {
	if s.attempted == 0 {
		return nil
	}
	if s.failed == s.attempted {
		return fmt.Errorf("all %d scan operations failed", s.attempted)
	}
	return nil
}

// record counts a real scan attempt. A clean skip (already in the DB,
// --update-existing unset) is not an attempt. A skip paired with an error
// is treated as a failed attempt so a caller cannot hide a failure.
func (s *scanStats) record(skipped bool, err error) {
	if skipped && err == nil {
		return
	}
	s.attempted++
	if err != nil {
		s.failed++
	}
}

// ScanAll loads repositories from config, discovers tags, and scans all images.
// Individual image failures are logged and scanning continues. Returns an error
// only when every scan attempt failed (total outage), so partial success still
// allows report generation from newly scanned images.
func (p *Pipeline) ScanAll() error {
	stats := &scanStats{}

	for _, group := range p.repoCfg.Repositories {
		repos, singleImages := scanner.ParseImagePatterns(group.Images)
		log.Infof("Group %q: %d repositories, %d single images", group.Description, len(repos), len(singleImages))

		for _, repo := range repos {
			p.scanRepository(repo, group.Category, stats)
		}

		for _, img := range singleImages {
			skipped, err := p.scanSingleImage(img, group.Category)
			stats.record(skipped, err)
			if err != nil {
				log.Errorf("Error scanning image %s: %v", img, err)
			}
		}
	}

	if stats.attempted == 0 {
		log.Warn("No scan operations were attempted")
		return nil
	}

	if stats.failed > 0 && stats.failed < stats.attempted {
		log.Warnf("Scan completed with %d/%d failures", stats.failed, stats.attempted)
	}

	return stats.result()
}

// ScanImage scans a single image by name.
func (p *Pipeline) ScanImage(imageName string) error {
	_, err := p.scanSingleImage(imageName, "")
	return err
}

// GenerateReport generates both the markdown and JSON recommendations reports.
func (p *Pipeline) GenerateReport() error {
	if err := report.GenerateReport(p.repo, p.config.OutputPath, p.config.TopN, &p.repoCfg); err != nil {
		return err
	}

	jsonPath := strings.TrimSuffix(p.config.OutputPath, ".md") + ".json"

	if err := report.GenerateJSONReport(p.repo, jsonPath, p.config.TopNJSON, &p.repoCfg); err != nil {
		return err
	}

	if p.config.Detailed {
		detailPath := strings.TrimSuffix(jsonPath, ".json") + "_detail.json"

		return report.GenerateDetailJSONReport(p.repo, detailPath)
	}

	return nil
}

func (p *Pipeline) scanRepository(repo string, category string, stats *scanStats) {
	log.Infof("Scanning repository: %s", repo)

	tags, err := p.registry.GetTags(repo)
	if err != nil {
		// Count tag discovery failure as a single failed attempt for this repo.
		stats.record(false, err)
		log.Errorf("Error scanning repository %s: getting tags: %v", repo, err)
		return
	}

	filtered := scanner.FilterTags(tags, p.repoCfg.TagFilter)
	limited := scanner.LimitTags(filtered, p.config.MaxTagsPerRepo)
	log.Infof("Repository %s: %d tags found, %d after filtering, %d to scan",
		repo, len(tags), len(filtered), len(limited))

	defaultRegistry := p.repoCfg.Defaults.Registry
	for _, tag := range limited {
		imageName := scanner.BuildFullImageName(defaultRegistry, repo, tag)
		skipped, err := p.scanSingleImage(imageName, category)
		stats.record(skipped, err)
		if err != nil {
			log.Errorf("Error scanning %s: %v", imageName, err)
		}
	}
}

func (p *Pipeline) scanSingleImage(imageName string, category string) (bool, error) {
	if !p.config.UpdateExisting {
		exists, err := p.repo.ImageExists(imageName)
		if err != nil {
			return false, fmt.Errorf("checking image existence: %w", err)
		}

		if exists {
			log.Infof("Skipping existing image: %s", imageName)
			return true, nil
		}
	}

	analysis, err := p.analyzer.Analyze(imageName)
	if err != nil {
		return false, fmt.Errorf("analyzing %s: %w", imageName, err)
	}

	if category == "base" && len(analysis.Image.Languages) == 0 {
		analysis.Image.Languages = append(analysis.Image.Languages, domain.Language{
			Language:    "base",
			Version:     analysis.Image.BaseOSVersion,
			MajorMinor:  analysis.Image.BaseOSVersion,
			PackageName: "Base OS",
			PackageType: "base",
		})
	}

	if err := p.repo.InsertImage(&analysis.Image); err != nil {
		return false, fmt.Errorf("inserting %s: %w", imageName, err)
	}

	log.Infof("Successfully scanned and stored: %s", imageName)

	return false, nil
}

// LoadRepositoryConfig reads the JSON config from the config directory.
func LoadRepositoryConfig(configDir string) (domain.RepositoryConfig, error) {
	path := configDir + "/repositories.json"

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn("No repositories.json found, using defaults")
			return defaultRepositoryConfig(), nil
		}

		return domain.RepositoryConfig{}, fmt.Errorf("reading %s: %w", path, err)
	}

	var cfg domain.RepositoryConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return domain.RepositoryConfig{}, fmt.Errorf("parsing %s: %w", path, err)
	}

	if cfg.Defaults.Registry == "" {
		cfg.Defaults.Registry = "mcr.microsoft.com"
	}

	if len(cfg.TagFilter.SkipExact) == 0 && len(cfg.TagFilter.ExcludeKeywords) == 0 &&
		len(cfg.TagFilter.ExcludePatterns) == 0 && !cfg.TagFilter.RequireDigit {
		cfg.TagFilter = scanner.DefaultTagFilter()
	}

	return cfg, nil
}

func defaultRepositoryConfig() domain.RepositoryConfig {
	return domain.RepositoryConfig{
		Defaults: domain.ConfigDefaults{
			Registry: "mcr.microsoft.com",
			MaxTags:  0,
		},
		TagFilter: scanner.DefaultTagFilter(),
		Repositories: []domain.RepositoryGroup{
			{
				Description: "Azure Linux base images",
				Images: []string{
					"azurelinux/base/python",
					"azurelinux/base/nodejs",
				},
			},
			{
				Description: "Azure Linux distroless images",
				Images: []string{
					"azurelinux/distroless/base",
					"azurelinux/distroless/python",
					"azurelinux/distroless/nodejs",
				},
			},
		},
	}
}
