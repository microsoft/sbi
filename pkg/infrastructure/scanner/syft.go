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
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// syftOutput represents the top-level Syft JSON output.
type syftOutput struct {
	Artifacts []syftArtifact `json:"artifacts"`
}

type syftArtifact struct {
	Name      string                 `json:"name"`
	Version   string                 `json:"version"`
	Type      string                 `json:"type"`
	Language  string                 `json:"language"`
	Metadata  map[string]interface{} `json:"metadata"`
	Locations []struct {
		Path string `json:"path"`
	} `json:"locations"`
}

// Language detection patterns matching the Python implementation.
var languagePatterns = map[string]*regexp.Regexp{
	"python": regexp.MustCompile(`(?i)(python|cpython)[\d.-]*$`),
	"node":   regexp.MustCompile(`(?i)(node|nodejs)[\d.-]*$`),
	"java":   regexp.MustCompile(`(?i)(java|openjdk|jdk|jre)[\d.-]*$`),
	"go":     regexp.MustCompile(`(?i)^golang$|^go[\d.]`),
	"ruby":   regexp.MustCompile(`(?i)^ruby[\d.-]*$`),
	"php":    regexp.MustCompile(`(?i)^php[\d.-]*$`),
	"dotnet": regexp.MustCompile(`(?i)^(dotnet|aspnet)[\w.-]*$`),
	"rust":   regexp.MustCompile(`(?i)^rust[\d.-]*$`),
	"lua":    regexp.MustCompile(`(?i)^lua[\d.-]*$`),
}

// Packages to exclude from language detection (lib packages, not runtimes).
var excludePackages = regexp.MustCompile(`(?i)(python-pip|python3-pip|nodejs-npm|java-common|python-setuptools)`)

// Package manager patterns.
var packageManagerPatterns = map[string]string{
	"pip":   "python",
	"pip3":  "python",
	"npm":   "node",
	"yarn":  "node",
	"cargo": "rust",
	"gem":   "ruby",
	"go":    "go",
}

// Capability detection keywords.
var capabilityKeywords = map[string][]string{
	"ssl":         {"openssl", "libssl", "ca-certificates"},
	"http_client": {"curl", "wget", "libcurl"},
	"database":    {"libpq", "sqlite", "mysql", "mariadb"},
	"compression": {"zlib", "bzip2", "xz", "gzip"},
	"xml":         {"libxml", "expat"},
	"json":        {"json"},
}

// RunSyft executes Syft on an image and returns the parsed SBOM.
func RunSyft(imageName string) (*domain.SyftResult, error) {
	log.Infof("Running Syft on: %s", imageName)

	cmd := exec.Command("syft", imageName, "-o", "json")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("running syft on %s: %w\n%s", imageName, err, stderr.String())
	}

	var output syftOutput
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		return nil, fmt.Errorf("parsing syft output: %w", err)
	}

	return parseSyftResult(&output), nil
}

func parseSyftResult(output *syftOutput) *domain.SyftResult {
	result := &domain.SyftResult{}

	langDetected := make(map[string]domain.Language)

	for _, artifact := range output.Artifacts {
		if excludePackages.MatchString(artifact.Name) {
			continue
		}

		cleanVer := cleanVersion(artifact.Version)

		// Detect languages via name regex (first match wins per language)
		for lang, pattern := range languagePatterns {
			if pattern.MatchString(artifact.Name) {
				if _, exists := langDetected[lang]; !exists {
					langDetected[lang] = domain.Language{
						Language:    lang,
						Version:     cleanVer,
						MajorMinor:  extractMajorMinor(cleanVer),
						PackageName: artifact.Name,
						PackageType: artifact.Type,
					}
				}

				break
			}
		}

		// Detect .NET via package type (catches distroless images where
		// package names like Microsoft.NETCore.App.Runtime.* don't match
		// the name regex). Prefer the core runtime package.
		if artifact.Type == "dotnet" {
			prev, exists := langDetected["dotnet"]
			if !exists || (isDotnetRuntime(artifact.Name) && !isDotnetRuntime(prev.PackageName)) {
				langDetected["dotnet"] = domain.Language{
					Language:    "dotnet",
					Version:     cleanVer,
					MajorMinor:  extractMajorMinor(cleanVer),
					PackageName: artifact.Name,
					PackageType: artifact.Type,
				}
			}
		}

		// Detect Go via binary cataloger (Go is tarball-installed, not
		// RPM-packaged, so the name regex won't find it).
		if artifact.Type == "binary" && strings.ToLower(artifact.Name) == "go" {
			if _, exists := langDetected["go"]; !exists {
				langDetected["go"] = domain.Language{
					Language:    "go",
					Version:     cleanVer,
					MajorMinor:  extractMajorMinor(cleanVer),
					PackageName: artifact.Name,
					PackageType: artifact.Type,
				}
			}
		}

		// Detect package managers
		if pmLang, ok := packageManagerPatterns[artifact.Name]; ok {
			result.PackageManagers = append(result.PackageManagers, domain.PackageManager{
				Name:     artifact.Name,
				Version:  artifact.Version,
				Language: pmLang,
			})
		}

		// Collect system packages (RPM, DEB, APK types)
		if isSystemPackageType(artifact.Type) {
			result.SystemPackages = append(result.SystemPackages, domain.SystemPackage{
				Name:        artifact.Name,
				Version:     artifact.Version,
				PackageType: artifact.Type,
			})
		}

		// Detect capabilities
		for cap, keywords := range capabilityKeywords {
			for _, kw := range keywords {
				if strings.Contains(strings.ToLower(artifact.Name), kw) {
					result.Capabilities = append(result.Capabilities, domain.Capability{
						Capability: cap,
					})

					break
				}
			}
		}
	}

	for _, lang := range langDetected {
		result.Languages = append(result.Languages, lang)
	}

	return result
}

// cleanVersion strips RPM/DEB build suffixes, keeping only the semver-like portion.
// e.g. "3.12.9-8.azl3" → "3.12.9", "5.36.0-7+deb12u3" → "5.36.0"
var semverPrefix = regexp.MustCompile(`^(\d+(?:\.\d+)*)`)

func cleanVersion(v string) string {
	if m := semverPrefix.FindString(v); m != "" {
		return m
	}
	return v
}

// isDotnetRuntime returns true for the core .NET runtime packages
// (Microsoft.NETCore.App.Runtime.* or Microsoft.AspNetCore.App.Runtime.*),
// as opposed to NuGet libraries or SDK tools.
func isDotnetRuntime(name string) bool {
	lower := strings.ToLower(name)
	return strings.Contains(lower, "microsoft.netcore.app.runtime") ||
		strings.Contains(lower, "microsoft.aspnetcore.app.runtime")
}

func extractMajorMinor(version string) string {
	parts := strings.SplitN(version, ".", 3)
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}

	return version
}

func isSystemPackageType(pkgType string) bool {
	switch strings.ToLower(pkgType) {
	case "rpm", "deb", "apk":
		return true
	}

	return false
}
