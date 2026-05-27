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
	"time"

	"github.com/microsoft/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
)

// DetailJSONReport is the top-level structure for the detailed JSON report.
type DetailJSONReport struct {
	SchemaVersion int                `json:"schemaVersion"`
	GeneratedAt   string             `json:"generatedAt"`
	ImageCount    int                `json:"imageCount"`
	Images        []DetailImageEntry `json:"images"`
}

// DetailImageEntry represents a single image with full detail data.
type DetailImageEntry struct {
	Name                 string                  `json:"name"`
	Registry             string                  `json:"registry,omitempty"`
	Repository           string                  `json:"repository,omitempty"`
	Tag                  string                  `json:"tag,omitempty"`
	Digest               string                  `json:"digest,omitempty"`
	SizeBytes            int64                   `json:"sizeBytes"`
	SizeHuman            string                  `json:"sizeHuman"`
	Layers               int                     `json:"layers"`
	CreatedDate          string                  `json:"createdDate"`
	ScanTimestamp        string                  `json:"scanTimestamp,omitempty"`
	BaseOS               DetailBaseOS            `json:"baseOS"`
	VulnerabilitySummary DetailVulnSummary       `json:"vulnerabilitySummary"`
	Languages            []DetailLanguageEntry   `json:"languages"`
	Vulnerabilities      []DetailVulnEntry       `json:"vulnerabilities"`
	SystemPackages       []DetailSystemPkgEntry  `json:"systemPackages"`
	PackageManagers      []DetailPkgManagerEntry `json:"packageManagers"`
}

// DetailBaseOS holds the OS information for an image.
type DetailBaseOS struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

// DetailVulnSummary holds vulnerability counts by severity.
type DetailVulnSummary struct {
	Total      int `json:"total"`
	Critical   int `json:"critical"`
	High       int `json:"high"`
	Medium     int `json:"medium"`
	Low        int `json:"low"`
	Negligible int `json:"negligible"`
	Unknown    int `json:"unknown"`
}

// DetailLanguageEntry represents a detected language/runtime.
type DetailLanguageEntry struct {
	Language    string `json:"language"`
	Version     string `json:"version"`
	MajorMinor  string `json:"majorMinor,omitempty"`
	PackageName string `json:"packageName,omitempty"`
	PackageType string `json:"packageType,omitempty"`
	Verified    bool   `json:"verified"`
}

// DetailVulnEntry represents a single vulnerability.
type DetailVulnEntry struct {
	ID             string  `json:"id"`
	Severity       string  `json:"severity"`
	CVSSScore      float64 `json:"cvssScore"`
	PackageName    string  `json:"packageName"`
	PackageVersion string  `json:"packageVersion"`
	FixedVersion   string  `json:"fixedVersion,omitempty"`
	Description    string  `json:"description,omitempty"`
}

// DetailSystemPkgEntry represents an OS-level package.
type DetailSystemPkgEntry struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageType string `json:"packageType,omitempty"`
}

// DetailPkgManagerEntry represents a detected package manager.
type DetailPkgManagerEntry struct {
	Name     string `json:"name"`
	Version  string `json:"version,omitempty"`
	Language string `json:"language,omitempty"`
}

// GenerateDetailJSONReport produces a detailed JSON report with per-image
// package inventories, vulnerability breakdowns, and detected languages.
func GenerateDetailJSONReport(repo *database.Repository, outputPath string) error {
	images, err := repo.QueryAllImageDetails()
	if err != nil {
		return fmt.Errorf("querying image details: %w", err)
	}

	if len(images) == 0 {
		log.Warn("No images found in database; detailed JSON report not generated.")
		// Remove any stale detail report from a previous run
		_ = os.Remove(outputPath)
		return nil
	}

	report := DetailJSONReport{
		SchemaVersion: 1,
		GeneratedAt:   time.Now().UTC().Format(time.RFC3339),
		ImageCount:  len(images),
		Images:      make([]DetailImageEntry, 0, len(images)),
	}

	for _, img := range images {
		entry := DetailImageEntry{
			Name:          img.Name,
			Registry:      img.Registry,
			Repository:    img.Repository,
			Tag:           img.Tag,
			Digest:        img.Digest,
			SizeBytes:     img.SizeBytes,
			SizeHuman:     HumanSize(img.SizeBytes),
			Layers:        img.Layers,
			CreatedDate:   img.CreatedDate,
			ScanTimestamp: img.ScanTimestamp,
			BaseOS: DetailBaseOS{
				Name:    DisplayOSName(img.BaseOSName),
				Version: img.BaseOSVersion,
			},
			VulnerabilitySummary: DetailVulnSummary{
				Total:      img.TotalVulnerabilities,
				Critical:   img.CriticalVulnerabilities,
				High:       img.HighVulnerabilities,
				Medium:     img.MediumVulnerabilities,
				Low:        img.LowVulnerabilities,
				Negligible: img.NegligibleVulnerabilities,
				Unknown:    img.UnknownVulnerabilities,
			},
			Languages:       make([]DetailLanguageEntry, 0, len(img.Languages)),
			Vulnerabilities: make([]DetailVulnEntry, 0, len(img.Vulnerabilities)),
			SystemPackages:  make([]DetailSystemPkgEntry, 0, len(img.SystemPackages)),
			PackageManagers: make([]DetailPkgManagerEntry, 0, len(img.PackageManagers)),
		}

		for _, l := range img.Languages {
			entry.Languages = append(entry.Languages, DetailLanguageEntry{
				Language:    l.Language,
				Version:     l.Version,
				MajorMinor:  l.MajorMinor,
				PackageName: l.PackageName,
				PackageType: l.PackageType,
				Verified:    l.Verified,
			})
		}

		for _, v := range img.Vulnerabilities {
			entry.Vulnerabilities = append(entry.Vulnerabilities, DetailVulnEntry{
				ID:             v.VulnerabilityID,
				Severity:       v.Severity,
				CVSSScore:      v.CVSSScore,
				PackageName:    v.PackageName,
				PackageVersion: v.PackageVersion,
				FixedVersion:   v.FixedVersion,
				Description:    v.Description,
			})
		}

		for _, sp := range img.SystemPackages {
			entry.SystemPackages = append(entry.SystemPackages, DetailSystemPkgEntry{
				Name:        sp.Name,
				Version:     sp.Version,
				PackageType: sp.PackageType,
			})
		}

		for _, pm := range img.PackageManagers {
			entry.PackageManagers = append(entry.PackageManagers, DetailPkgManagerEntry{
				Name:     pm.Name,
				Version:  pm.Version,
				Language: pm.Language,
			})
		}

		report.Images = append(report.Images, entry)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling detailed JSON report: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return fmt.Errorf("writing detailed JSON report: %w", err)
	}

	log.Infof("Wrote detailed JSON report to %s (%d images)", outputPath, len(images))

	return nil
}
