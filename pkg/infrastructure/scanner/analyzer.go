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

package scanner

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// Runtime verification commands per language.
var runtimeCommands = map[string][]string{
	"python": {"python3", "--version"},
	"node":   {"node", "--version"},
	"java":   {"java", "-version"},
	"go":     {"go", "version"},
	"ruby":   {"ruby", "--version"},
	"php":    {"php", "--version"},
	"dotnet": {"dotnet", "--info"},
	"rust":   {"rustc", "--version"},
}

// Version extraction patterns per language.
var versionPatterns = map[string]*regexp.Regexp{
	"python": regexp.MustCompile(`Python\s+(\d+\.\d+\.\d+)`),
	"node":   regexp.MustCompile(`v(\d+\.\d+\.\d+)`),
	"java":   regexp.MustCompile(`version\s+"?(\d+[\d.]*)`),
	"go":     regexp.MustCompile(`go(\d+\.\d+[\d.]*)`),
	"ruby":   regexp.MustCompile(`ruby\s+(\d+\.\d+\.\d+)`),
	"php":    regexp.MustCompile(`PHP\s+(\d+\.\d+\.\d+)`),
	"dotnet": regexp.MustCompile(`Version:\s+(\d+\.\d+[\d.]*)`),
	"rust":   regexp.MustCompile(`rustc\s+(\d+\.\d+\.\d+)`),
}

// ImageAnalyzer orchestrates the full analysis of a container image.
type ImageAnalyzer struct {
	docker          *DockerClient
	comprehensive   bool
	cleanupImages   bool
}

// NewImageAnalyzer creates a new ImageAnalyzer.
func NewImageAnalyzer(comprehensive, cleanupImages bool) *ImageAnalyzer {
	return &ImageAnalyzer{
		docker:        NewDockerClient(),
		comprehensive: comprehensive,
		cleanupImages: cleanupImages,
	}
}

// Analyze performs the full analysis pipeline on an image: pull → syft → trivy → runtime verify.
func (a *ImageAnalyzer) Analyze(imageName string) (*domain.ImageAnalysis, error) {
	log.Infof("Analyzing image: %s", imageName)
	startTime := time.Now()

	// Pull the image
	if err := a.docker.Pull(imageName); err != nil {
		return nil, fmt.Errorf("pull failed: %w", err)
	}

	if a.cleanupImages {
		defer func() {
			if err := a.docker.Remove(imageName); err != nil {
				log.Warnf("cleanup failed for %s: %v", imageName, err)
			}
		}()
	}

	// Docker inspect
	inspectResult, err := a.docker.Inspect(imageName)
	if err != nil {
		log.Warnf("Inspect failed for %s: %v", imageName, err)
	}

	// Syft SBOM analysis
	syftResult, err := RunSyft(imageName)
	if err != nil {
		log.Warnf("Syft failed for %s: %v", imageName, err)
		syftResult = &domain.SyftResult{}
	}

	// Trivy vulnerability scan
	trivyResult, err := RunTrivy(imageName, a.comprehensive)
	if err != nil {
		log.Warnf("Trivy failed for %s: %v", imageName, err)
		trivyResult = &domain.TrivyResult{}
	}

	// Merge Syft languages with image-name detection (e.g. dotnet from aspnet:8.0)
	languages := mergeLanguages(syftResult.Languages, nil, imageName)

	// Runtime version verification — runs on ALL detected languages including
	// image-name-detected ones so e.g. dotnet aspnet/runtime gets the precise
	// patch version from "dotnet --info" rather than just the tag version.
	runtimeVersions := a.verifyRuntimes(imageName, languages)

	// Apply verified versions back to the language list
	languages = applyRuntimeVersions(languages, runtimeVersions)

	// Get on-disk image size from "docker images" (matches Python behaviour)
	imageSize := inspectResult.SizeBytes
	if diskSize, err := a.docker.GetImageSize(imageName); err == nil && diskSize > 0 {
		imageSize = diskSize
	}

	// Build the image record
	registry, repository, tag := ExtractRegistryAndRepo(imageName)
	analysis := &domain.ImageAnalysis{
		ScanTime:        startTime,
		DockerInspect:   inspectResult,
		SyftResult:      *syftResult,
		TrivyResult:     *trivyResult,
		RuntimeVersions: runtimeVersions,
		Image: domain.ImageRecord{
			Name:                        imageName,
			Registry:                    registry,
			Repository:                  repository,
			Tag:                         tag,
			Digest:                      inspectResult.Digest,
			SizeBytes:                   imageSize,
			Layers:                      inspectResult.Layers,
			CreatedDate:                 inspectResult.Created,
			BaseOSName:                  trivyResult.BaseOSFamily,
			BaseOSVersion:               trivyResult.BaseOSVersion,
			TotalVulnerabilities:        trivyResult.TotalVulnerabilities,
			CriticalVulnerabilities:     trivyResult.CriticalVulnerabilities,
			HighVulnerabilities:         trivyResult.HighVulnerabilities,
			MediumVulnerabilities:       trivyResult.MediumVulnerabilities,
			LowVulnerabilities:          trivyResult.LowVulnerabilities,
			NegligibleVulnerabilities:   trivyResult.NegligibleVulnerabilities,
			UnknownVulnerabilities:      trivyResult.UnknownVulnerabilities,
			VulnerabilityScanTimestamp:  time.Now().UTC().Format(time.RFC3339),
			VulnerabilityScanner:        "trivy",
			SecretsFound:                trivyResult.SecretsFound,
			ConfigIssues:                trivyResult.ConfigIssues,
			LicenseIssues:               trivyResult.LicenseIssues,
			Languages:                   languages,
			Vulnerabilities:             trivyResult.Vulnerabilities,
			PackageManagers:             syftResult.PackageManagers,
			Capabilities:                syftResult.Capabilities,
			SystemPackages:              syftResult.SystemPackages,
		},
	}

	if a.comprehensive {
		analysis.Image.ComprehensiveScanTimestamp = time.Now().UTC().Format(time.RFC3339)
		analysis.Image.ComprehensiveScanner = "trivy"
	}

	elapsed := time.Since(startTime)
	log.Infof("Analysis of %s completed in %s", imageName, elapsed)

	return analysis, nil
}

// verifyRuntimes runs version commands inside the container for detected languages.
func (a *ImageAnalyzer) verifyRuntimes(imageName string, languages []domain.Language) map[string]string {
	versions := make(map[string]string)
	checked := make(map[string]bool)

	for _, lang := range languages {
		langKey := strings.ToLower(lang.Language)
		if checked[langKey] {
			continue
		}

		checked[langKey] = true

		cmd, ok := runtimeCommands[langKey]
		if !ok {
			continue
		}

		output, err := a.docker.RunCommand(imageName, cmd)
		if err != nil {
			log.Debugf("Runtime verification for %s failed in %s: %v", langKey, imageName, err)
			continue
		}

		if pattern, ok := versionPatterns[langKey]; ok {
			if matches := pattern.FindStringSubmatch(output); len(matches) > 1 {
				versions[langKey] = matches[1]
				log.Debugf("Verified %s version: %s in %s", langKey, matches[1], imageName)
			}
		}
	}

	return versions
}

// dotnetImagePattern detects .NET from image names like mcr.microsoft.com/dotnet/aspnet:8.0
var dotnetImagePattern = regexp.MustCompile(`^(?i)mcr\.microsoft\.com/dotnet/(?:aspnet|runtime):(\d+\.\d+)`)

// mergeLanguages combines Syft-detected languages with image-name-based detection
// (e.g. dotnet from mcr.microsoft.com/dotnet/aspnet:8.0).
func mergeLanguages(syftLangs []domain.Language, runtimeVersions map[string]string, imageName string) []domain.Language {
	result := make([]domain.Language, len(syftLangs))
	copy(result, syftLangs)

	// Apply any runtime-verified versions passed in
	result = applyRuntimeVersions(result, runtimeVersions)

	// Detect dotnet from image name (fallback for aspnet/runtime images)
	hasDotnet := false
	for _, l := range result {
		if strings.ToLower(l.Language) == "dotnet" {
			hasDotnet = true
			break
		}
	}

	if !hasDotnet {
		if m := dotnetImagePattern.FindStringSubmatch(imageName); len(m) > 1 {
			ver := m[1]
			result = append(result, domain.Language{
				Language:    "dotnet",
				Version:     ver,
				MajorMinor:  ver,
				PackageName: "Microsoft .NET Runtime",
				PackageType: "container_runtime",
			})
		}
	}

	return result
}

// applyRuntimeVersions updates language entries with versions obtained by
// running the runtime inside the container (e.g. "dotnet --info" → "8.0.24").
func applyRuntimeVersions(langs []domain.Language, runtimeVersions map[string]string) []domain.Language {
	if len(runtimeVersions) == 0 {
		return langs
	}

	for i := range langs {
		langKey := strings.ToLower(langs[i].Language)
		if version, ok := runtimeVersions[langKey]; ok && version != "" {
			langs[i].Verified = true
			langs[i].Version = version
			langs[i].MajorMinor = extractMajorMinor(version)
		}
	}

	return langs
}
