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

package domain

import "time"

// ImageRecord represents a container image with its metadata and vulnerability information.
type ImageRecord struct {
	ID                         int64
	Name                       string
	Registry                   string
	Repository                 string
	Tag                        string
	Digest                     string
	SizeBytes                  int64
	Layers                     int
	CreatedDate                string
	ScanTimestamp              string
	BaseOSName                 string
	BaseOSVersion              string
	TotalVulnerabilities       int
	CriticalVulnerabilities    int
	HighVulnerabilities        int
	MediumVulnerabilities      int
	LowVulnerabilities         int
	NegligibleVulnerabilities  int
	UnknownVulnerabilities     int
	VulnerabilityScanTimestamp string
	VulnerabilityScanner       string
	SecretsFound               int
	ConfigIssues               int
	LicenseIssues              int
	ComprehensiveScanTimestamp string
	ComprehensiveScanner       string
	Languages                  []Language
	Vulnerabilities            []Vulnerability
	PackageManagers            []PackageManager
	Capabilities               []Capability
	SystemPackages             []SystemPackage
	SecurityFindings           []SecurityFinding
}

// Language represents a programming language detected in an image.
type Language struct {
	ID           int64
	ImageID      int64
	Language     string
	Version      string
	MajorMinor   string
	PackageName  string
	PackageType  string
	Architecture string
	Vendor       string
	Verified     bool
}

// Vulnerability represents a security vulnerability found in an image.
type Vulnerability struct {
	ID              int64
	ImageID         int64
	VulnerabilityID string
	Severity        string
	PackageName     string
	PackageVersion  string
	FixedVersion    string
	Description     string
	CVSSScore       float64
}

// PackageManager represents a package manager found in an image.
type PackageManager struct {
	ID       int64
	ImageID  int64
	Name     string
	Version  string
	Language string
}

// Capability represents a capability detected in an image.
type Capability struct {
	ID         int64
	ImageID    int64
	Capability string
}

// SystemPackage represents a system-level package installed in an image.
type SystemPackage struct {
	ID          int64
	ImageID     int64
	Name        string
	Version     string
	PackageType string
}

// SecurityFinding represents a security finding from comprehensive scanning.
type SecurityFinding struct {
	ID          int64
	ImageID     int64
	FindingType string
	Severity    string
	RuleID      string
	Title       string
	Description string
	FilePath    string
	Category    string
	Message     string
}

// ScanConfig holds configuration for a scanning session.
type ScanConfig struct {
	DBPath            string
	ConfigDir         string
	OutputPath        string
	TopN              int
	TopNJSON          int
	MaxTagsPerRepo    int
	ComprehensiveScan bool
	CleanupImages     bool
	UpdateExisting    bool
	Verbose           bool
}

// RepositoryConfig represents the JSON configuration for image sources and tag filtering.
type RepositoryConfig struct {
	Defaults     ConfigDefaults    `json:"defaults"`
	TagFilter    TagFilterConfig   `json:"tagFilter"`
	Repositories []RepositoryGroup `json:"repositories"`
}

// ConfigDefaults holds default values for scanning.
type ConfigDefaults struct {
	Registry string `json:"registry"`
	MaxTags  int    `json:"maxTags"`
}

// TagFilterConfig defines rules for filtering image tags.
type TagFilterConfig struct {
	SkipExact       []string `json:"skipExact"`
	ExcludeKeywords []string `json:"excludeKeywords"`
	ExcludePatterns []string `json:"excludePatterns"`
	RequireDigit    bool     `json:"requireDigit"`
}

// RepositoryGroup holds a set of related image sources.
// Category can be "base" to mark minimal/no-runtime images.
type RepositoryGroup struct {
	Description string   `json:"description"`
	Category    string   `json:"category,omitempty"`
	Images      []string `json:"images"`
}

// TagsResponse represents the JSON response from a container registry tags API.
type TagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// ImageAnalysis holds the combined results of analyzing an image.
type ImageAnalysis struct {
	Image           ImageRecord
	ScanTime        time.Time
	DockerInspect   DockerInspectResult
	SyftResult      SyftResult
	TrivyResult     TrivyResult
	RuntimeVersions map[string]string
}

// DockerInspectResult holds data from docker inspect.
type DockerInspectResult struct {
	SizeBytes int64
	Layers    int
	Digest    string
	Created   string
}

// SyftResult holds parsed SBOM data from Syft.
type SyftResult struct {
	Languages       []Language
	PackageManagers []PackageManager
	SystemPackages  []SystemPackage
	Capabilities    []Capability
}

// TrivyResult holds parsed vulnerability data from Trivy.
type TrivyResult struct {
	TotalVulnerabilities      int
	CriticalVulnerabilities   int
	HighVulnerabilities       int
	MediumVulnerabilities     int
	LowVulnerabilities        int
	NegligibleVulnerabilities int
	UnknownVulnerabilities    int
	Vulnerabilities           []Vulnerability
	SecretsFound              int
	ConfigIssues              int
	LicenseIssues             int
	BaseOSFamily              string // e.g. "azurelinux", "ubuntu", "debian"
	BaseOSVersion             string // e.g. "3.0", "22.04", "12.13"
}

// RecommendedImage holds a ranked image for the report output.
type RecommendedImage struct {
	Name                    string
	Version                 string
	CriticalVulnerabilities int
	HighVulnerabilities     int
	TotalVulnerabilities    int
	SizeBytes               int64
	CreatedDate             string
	Digest                  string
	BaseOSName              string
	AlternateTags           []string
}
