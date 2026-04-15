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

package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
)

// JSONReport is the top-level structure for the JSON report.
type JSONReport struct {
	GeneratedAt  string           `json:"generatedAt"`
	TopN         int              `json:"topN"`
	ScannedRepos []JSONRepoGroup  `json:"scannedRepositories,omitempty"`
	Images       []JSONImageEntry `json:"images"`
}

// JSONRepoGroup represents a group of scanned repositories/images.
type JSONRepoGroup struct {
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

// JSONImageEntry represents a single recommended image in the JSON report.
type JSONImageEntry struct {
	Rank                    int    `json:"rank"`
	Language                string `json:"language"`
	BaseOS                  string `json:"baseOS"`
	Name                    string `json:"name"`
	Version                 string `json:"version"`
	CriticalVulnerabilities int    `json:"criticalVulnerabilities"`
	HighVulnerabilities     int    `json:"highVulnerabilities"`
	TotalVulnerabilities    int    `json:"totalVulnerabilities"`
	SizeBytes               int64  `json:"sizeBytes"`
	SizeHuman               string `json:"sizeHuman"`
	Digest                  string `json:"digest"`
	PinnedReference         string `json:"pinnedReference,omitempty"`
	StableTag               string `json:"stableTag,omitempty"`
	DockerfileFrom          string `json:"dockerfileFrom,omitempty"`
}

// GenerateJSONReport produces a flat JSON recommendations report from the database.
func GenerateJSONReport(repo *database.Repository, outputPath string, topN int, repoCfg *domain.RepositoryConfig) error {
	languages, err := repo.QueryLanguages()
	if err != nil {
		return fmt.Errorf("querying languages: %w", err)
	}

	if len(languages) == 0 {
		log.Warn("No languages found in database; JSON report not generated.")
		return nil
	}

	report := JSONReport{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		TopN:        topN,
		Images:      []JSONImageEntry{},
	}

	if repoCfg != nil {
		for _, group := range repoCfg.Repositories {
			report.ScannedRepos = append(report.ScannedRepos, JSONRepoGroup{
				Description: group.Description,
				Images:      group.Images,
			})
		}
	}

	for _, lang := range languages {
		oses, err := repo.QueryBaseOSes(lang)
		if err != nil {
			log.Warnf("Failed to query OSes for %s: %v", lang, err)
			continue
		}

		for _, osName := range oses {
			images, err := repo.QueryTopImagesByOS(lang, osName, topN)
			if err != nil {
				log.Warnf("Failed to query images for %s/%s: %v", lang, osName, err)
				continue
			}

			for idx, img := range images {
				report.Images = append(report.Images, JSONImageEntry{
					Rank:                    idx + 1,
					Language:                strings.ToLower(lang),
					BaseOS:                  DisplayOSName(img.BaseOSName),
					Name:                    img.Name,
					Version:                 img.Version,
					CriticalVulnerabilities: img.CriticalVulnerabilities,
					HighVulnerabilities:     img.HighVulnerabilities,
					TotalVulnerabilities:    img.TotalVulnerabilities,
					SizeBytes:               img.SizeBytes,
					SizeHuman:               HumanSize(img.SizeBytes),
					Digest:                  img.Digest,
					PinnedReference:         FormatPinnedReference(img.Name, img.Digest),
					StableTag:               FormatStableTag(img.Name, img.Version),
					DockerfileFrom:          FormatDockerfileFrom(img.Name, img.Digest),
				})
			}
		}
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling JSON report: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return fmt.Errorf("writing JSON report: %w", err)
	}

	log.Infof("Wrote JSON recommendations to %s", outputPath)

	return nil
}
