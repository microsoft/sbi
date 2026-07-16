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
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test fixtures built from real trivy analysis of container images (2026-03-10).
// Each fixture contains representative vulnerability/metadata data extracted from
// the full trivy JSON output, keeping tests fast while exercising real-world data.

// ---------- helpers ----------

// trivyCVSS is now a named type in trivy.go; no alias needed in tests.

func makeTrivyVulns(entries []struct {
	id, severity, pkg, installed, fixed, desc string
	cvss                                      map[string]trivyCVSS
}) []trivyVuln {
	vulns := make([]trivyVuln, len(entries))
	for i, e := range entries {
		vulns[i] = trivyVuln{
			VulnerabilityID:  e.id,
			Severity:         e.severity,
			PkgName:          e.pkg,
			InstalledVersion: e.installed,
			FixedVersion:     e.fixed,
			Description:      e.desc,
			CVSS:             e.cvss,
		}
	}
	return vulns
}

// ---------- Azure Linux fixtures ----------

func azureLinuxPythonTrivyOutput() *trivyOutput {
	// From mcr.microsoft.com/azurelinux/base/python:3.12
	// OS: azurelinux 3.0, 0 vulnerabilities
	return &trivyOutput{
		Metadata: trivyMetadata{OS: trivyOS{Family: "azurelinux", Name: "3.0"}},
		Results: []trivyResult{
			{
				Target: "mcr.microsoft.com/azurelinux/base/python:3.12 (azurelinux 3.0)",
				Class:  "os-pkgs",
				Type:   "azurelinux",
			},
		},
	}
}

func azureLinuxDistrolessPythonTrivyOutput() *trivyOutput {
	// From mcr.microsoft.com/azurelinux/distroless/python:3.12
	// OS: azurelinux 3.0, 0 vulnerabilities
	return &trivyOutput{
		Metadata: trivyMetadata{OS: trivyOS{Family: "azurelinux", Name: "3.0"}},
		Results: []trivyResult{
			{
				Target: "mcr.microsoft.com/azurelinux/distroless/python:3.12 (azurelinux 3.0)",
				Class:  "os-pkgs",
				Type:   "azurelinux",
			},
		},
	}
}

func azureLinuxDistrolessNodejsTrivyOutput() *trivyOutput {
	// From mcr.microsoft.com/azurelinux/distroless/nodejs:20
	// OS: azurelinux 3.0, 12 vulns: 0 CRITICAL, 10 HIGH, 0 MEDIUM, 2 LOW
	return &trivyOutput{
		Metadata: trivyMetadata{OS: trivyOS{Family: "azurelinux", Name: "3.0"}},
		Results: []trivyResult{
			{
				Target: "mcr.microsoft.com/azurelinux/distroless/nodejs:20 (azurelinux 3.0)",
				Class:  "lang-pkgs",
				Type:   "npm",
				Vulnerabilities: makeTrivyVulns([]struct {
					id, severity, pkg, installed, fixed, desc string
					cvss                                      map[string]trivyCVSS
				}{
					{"CVE-2024-21538", "HIGH", "cross-spawn", "7.0.3", "7.0.5, 6.0.6", "cross-spawn ReDoS vulnerability",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2025-64756", "HIGH", "glob", "10.3.10", "11.1.0, 10.5.0", "glob ReDoS",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2026-26996", "HIGH", "minimatch", "9.0.3", "10.2.1, 9.0.6", "minimatch ReDoS",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2026-27903", "HIGH", "minimatch", "9.0.3", "10.2.3, 9.0.7", "minimatch ReDoS variant",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2026-27904", "HIGH", "minimatch", "9.0.3", "10.2.3, 9.0.7", "minimatch ReDoS variant 2",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2026-23745", "HIGH", "tar", "6.2.0", "7.5.3", "tar arbitrary file write",
						map[string]trivyCVSS{"nvd": {8.2}}},
					{"CVE-2026-23950", "HIGH", "tar", "6.2.0", "7.5.4", "tar path traversal",
						map[string]trivyCVSS{"nvd": {8.8}}},
					{"CVE-2026-24842", "HIGH", "tar", "6.2.0", "7.5.7", "tar arbitrary file overwrite",
						map[string]trivyCVSS{"nvd": {8.2}}},
					{"CVE-2026-26960", "HIGH", "tar", "6.2.0", "7.5.8", "tar symlink traversal",
						map[string]trivyCVSS{"nvd": {7.1}}},
					{"CVE-2026-29786", "HIGH", "tar", "6.2.0", "7.5.10", "tar extraction vulnerability",
						nil},
					{"CVE-2025-5889", "LOW", "brace-expansion", "2.0.1", "2.0.2, 1.1.12", "brace-expansion ReDoS",
						map[string]trivyCVSS{"nvd": {3.1}}},
					{"CVE-2026-24001", "LOW", "diff", "5.1.0", "8.0.3, 5.2.2", "diff ReDoS",
						map[string]trivyCVSS{"nvd": {7.5}}},
				}),
			},
		},
	}
}

// ---------- Ubuntu fixture ----------

func ubuntuJDK21TrivyOutput() *trivyOutput {
	// From mcr.microsoft.com/openjdk/jdk:21-ubuntu
	// OS: ubuntu 22.04, 63 vulns: 0 CRITICAL, 0 HIGH, 13 MEDIUM, 50 LOW
	// Trimmed to representative subset for test speed
	return &trivyOutput{
		Metadata: trivyMetadata{OS: trivyOS{Family: "ubuntu", Name: "22.04"}},
		Results: []trivyResult{
			{
				Target: "mcr.microsoft.com/openjdk/jdk:21-ubuntu (ubuntu 22.04)",
				Class:  "os-pkgs",
				Type:   "ubuntu",
				Vulnerabilities: makeTrivyVulns([]struct {
					id, severity, pkg, installed, fixed, desc string
					cvss                                      map[string]trivyCVSS
				}{
					{"CVE-2025-1180", "MEDIUM", "binutils", "2.38-4ubuntu2.12", "", "binutils memory corruption",
						map[string]trivyCVSS{"redhat": {3.1}}},
					{"CVE-2017-13716", "LOW", "binutils", "2.38-4ubuntu2.12", "", "binutils memory leak",
						map[string]trivyCVSS{"nvd": {5.5}, "redhat": {3.3}}},
					{"CVE-2022-28358", "LOW", "libc6", "2.35-0ubuntu3.4", "", "glibc issue",
						nil},
					{"CVE-2024-2236", "MEDIUM", "libgcrypt20", "1.9.4-3ubuntu3", "", "libgcrypt timing side channel",
						map[string]trivyCVSS{"nvd": {5.9}, "redhat": {5.9}}},
				}),
			},
		},
	}
}

// ---------- Debian fixture ----------

func debianDotnetRuntimeTrivyOutput() *trivyOutput {
	// From mcr.microsoft.com/dotnet/runtime:8.0
	// OS: debian 12.13, 95 vulns: 1 CRITICAL, 2 HIGH, 15 MEDIUM, 76 LOW (+ 1 NEGLIGIBLE)
	// Trimmed to representative subset — includes multi-result (os-pkgs + dotnet lang-pkgs)
	return &trivyOutput{
		Metadata: trivyMetadata{OS: trivyOS{Family: "debian", Name: "12.13"}},
		Results: []trivyResult{
			{
				Target: "mcr.microsoft.com/dotnet/runtime:8.0 (debian 12.13)",
				Class:  "os-pkgs",
				Type:   "debian",
				Vulnerabilities: makeTrivyVulns([]struct {
					id, severity, pkg, installed, fixed, desc string
					cvss                                      map[string]trivyCVSS
				}{
					{"CVE-2023-45853", "CRITICAL", "zlib1g", "1:1.2.13.dfsg-1", "", "zlib integer overflow in MiniZip",
						map[string]trivyCVSS{"ghsa": {9.8}, "nvd": {9.8}, "redhat": {5.3}}},
					{"CVE-2026-0861", "HIGH", "libc-bin", "2.36-9+deb12u13", "", "glibc vulnerability",
						map[string]trivyCVSS{"redhat": {8.1}}},
					{"CVE-2026-0861", "HIGH", "libc6", "2.36-9+deb12u13", "", "glibc vulnerability",
						map[string]trivyCVSS{"redhat": {8.1}}},
					{"CVE-2024-41996", "MEDIUM", "libssl3", "3.0.16-1~deb12u1", "", "OpenSSL Diffie-Hellman issue",
						map[string]trivyCVSS{"nvd": {7.5}}},
					{"CVE-2025-0938", "LOW", "libc-bin", "2.36-9+deb12u13", "", "glibc hostname parsing",
						map[string]trivyCVSS{"nvd": {4.3}, "redhat": {4.3}}},
					{"CVE-2024-41110", "LOW", "tzdata", "2024a-0+deb12u1", "", "timezone data issue",
						nil},
				}),
			},
			{
				Target: "usr/share/dotnet/shared/Microsoft.NETCore.App/8.0.24/Microsoft.NETCore.App.deps.json",
				Class:  "lang-pkgs",
				Type:   "dotnet-core",
				// No vulnerabilities in the .NET packages themselves
			},
		},
	}
}

// ---------- Edge case fixtures ----------

func emptyTrivyOutput() *trivyOutput {
	return &trivyOutput{}
}

func noVulnsTrivyOutput() *trivyOutput {
	return &trivyOutput{
		Results: []trivyResult{
			{
				Target: "clean-image:latest",
				Class:  "os-pkgs",
				Type:   "alpine",
			},
		},
	}
}

func nilVulnsTrivyOutput() *trivyOutput {
	return &trivyOutput{
		Results: []trivyResult{
			{
				Target:          "image:1.0",
				Vulnerabilities: nil,
			},
		},
	}
}

// ============================================================================
// parseTrivyResult tests
// ============================================================================

func TestParseTrivyResult_AzureLinuxPython(t *testing.T) {
	result := parseTrivyResult(azureLinuxPythonTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalVulnerabilities)
	assert.Equal(t, 0, result.CriticalVulnerabilities)
	assert.Equal(t, 0, result.HighVulnerabilities)
	assert.Equal(t, 0, result.MediumVulnerabilities)
	assert.Equal(t, 0, result.LowVulnerabilities)
	assert.Empty(t, result.Vulnerabilities)
	assert.Equal(t, 0, result.SecretsFound)
	assert.Equal(t, 0, result.ConfigIssues)
	assert.Equal(t, 0, result.LicenseIssues)
	assert.Equal(t, "azurelinux", result.BaseOSFamily)
	assert.Equal(t, "3.0", result.BaseOSVersion)
}

func TestParseTrivyResult_AzureLinuxDistrolessPython(t *testing.T) {
	result := parseTrivyResult(azureLinuxDistrolessPythonTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalVulnerabilities)
	assert.Empty(t, result.Vulnerabilities)
	assert.Equal(t, "azurelinux", result.BaseOSFamily)
	assert.Equal(t, "3.0", result.BaseOSVersion)
}

func TestParseTrivyResult_AzureLinuxDistrolessNodejs(t *testing.T) {
	result := parseTrivyResult(azureLinuxDistrolessNodejsTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 12, result.TotalVulnerabilities)
	assert.Equal(t, 0, result.CriticalVulnerabilities)
	assert.Equal(t, 10, result.HighVulnerabilities)
	assert.Equal(t, 0, result.MediumVulnerabilities)
	assert.Equal(t, 2, result.LowVulnerabilities)
	assert.Len(t, result.Vulnerabilities, 12)
	assert.Equal(t, "azurelinux", result.BaseOSFamily)
	assert.Equal(t, "3.0", result.BaseOSVersion)
}

func TestParseTrivyResult_UbuntuJDK21(t *testing.T) {
	// Trimmed fixture: 4 vulns (2 MEDIUM, 2 LOW)
	result := parseTrivyResult(ubuntuJDK21TrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 4, result.TotalVulnerabilities)
	assert.Equal(t, 0, result.CriticalVulnerabilities)
	assert.Equal(t, 0, result.HighVulnerabilities)
	assert.Equal(t, 2, result.MediumVulnerabilities)
	assert.Equal(t, 2, result.LowVulnerabilities)
	assert.Len(t, result.Vulnerabilities, 4)
	assert.Equal(t, "ubuntu", result.BaseOSFamily)
	assert.Equal(t, "22.04", result.BaseOSVersion)
}

func TestParseTrivyResult_DebianDotnetRuntime(t *testing.T) {
	// Trimmed fixture: 6 vulns (1 CRITICAL, 2 HIGH, 1 MEDIUM, 2 LOW) across 2 Result sections
	result := parseTrivyResult(debianDotnetRuntimeTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 6, result.TotalVulnerabilities)
	assert.Equal(t, 1, result.CriticalVulnerabilities)
	assert.Equal(t, 2, result.HighVulnerabilities)
	assert.Equal(t, 1, result.MediumVulnerabilities)
	assert.Equal(t, 2, result.LowVulnerabilities)
	assert.Len(t, result.Vulnerabilities, 6)
	assert.Equal(t, "debian", result.BaseOSFamily)
	assert.Equal(t, "12.13", result.BaseOSVersion)
}

// ============================================================================
// CVSS score extraction tests
// ============================================================================

func TestParseTrivyResult_CVSSScoreExtraction(t *testing.T) {
	result := parseTrivyResult(debianDotnetRuntimeTrivyOutput())
	require.NotNil(t, result)

	// Build lookup by CVE ID
	vulnMap := make(map[string]float64)
	for _, v := range result.Vulnerabilities {
		vulnMap[v.VulnerabilityID] = v.CVSSScore
	}

	// CVE-2023-45853 has multiple CVSS sources: ghsa=9.8, nvd=9.8, redhat=5.3 → max=9.8
	assert.InDelta(t, 9.8, vulnMap["CVE-2023-45853"], 0.01, "should pick highest V3Score")

	// CVE-2026-0861 has single source: redhat=8.1
	assert.InDelta(t, 8.1, vulnMap["CVE-2026-0861"], 0.01)

	// CVE-2025-0938 has nvd=4.3, redhat=4.3 → 4.3
	assert.InDelta(t, 4.3, vulnMap["CVE-2025-0938"], 0.01)
}

func TestParseTrivyResult_CVSSScoreNilCVSS(t *testing.T) {
	result := parseTrivyResult(ubuntuJDK21TrivyOutput())
	require.NotNil(t, result)

	// CVE-2022-28358 has nil CVSS → score should be 0
	for _, v := range result.Vulnerabilities {
		if v.VulnerabilityID == "CVE-2022-28358" {
			assert.Equal(t, float64(0), v.CVSSScore)
			return
		}
	}
	t.Fatal("CVE-2022-28358 not found in results")
}

func TestParseTrivyResult_CVSSMultipleSourcesPicksHighest(t *testing.T) {
	result := parseTrivyResult(ubuntuJDK21TrivyOutput())

	// CVE-2017-13716 has nvd=5.5, redhat=3.3 → should pick 5.5
	for _, v := range result.Vulnerabilities {
		if v.VulnerabilityID == "CVE-2017-13716" {
			assert.InDelta(t, 5.5, v.CVSSScore, 0.01, "should pick highest V3Score across sources")
			return
		}
	}
	t.Fatal("CVE-2017-13716 not found in results")
}

// ============================================================================
// Vulnerability detail extraction tests
// ============================================================================

func TestParseTrivyResult_VulnerabilityDetails(t *testing.T) {
	result := parseTrivyResult(debianDotnetRuntimeTrivyOutput())
	require.NotNil(t, result)

	// Find the CRITICAL vuln
	var critVuln *struct {
		id, severity, pkg, pkgVer, fixedVer, desc string
		cvss                                      float64
	}
	for _, v := range result.Vulnerabilities {
		if v.Severity == "CRITICAL" {
			critVuln = &struct {
				id, severity, pkg, pkgVer, fixedVer, desc string
				cvss                                      float64
			}{v.VulnerabilityID, v.Severity, v.PackageName, v.PackageVersion, v.FixedVersion, v.Description, v.CVSSScore}
			break
		}
	}

	require.NotNil(t, critVuln, "should have a CRITICAL vulnerability")
	assert.Equal(t, "CVE-2023-45853", critVuln.id)
	assert.Equal(t, "CRITICAL", critVuln.severity)
	assert.Equal(t, "zlib1g", critVuln.pkg)
	assert.Equal(t, "1:1.2.13.dfsg-1", critVuln.pkgVer)
	assert.Contains(t, critVuln.desc, "zlib")
}

func TestParseTrivyResult_DescriptionTruncation(t *testing.T) {
	longDesc := make([]byte, 600)
	for i := range longDesc {
		longDesc[i] = 'a'
	}

	output := &trivyOutput{
		Results: []trivyResult{
			{
				Vulnerabilities: []trivyVuln{
					{VulnerabilityID: "CVE-LONG", Severity: "LOW", Description: string(longDesc)},
				},
			},
		},
	}

	result := parseTrivyResult(output)
	require.Len(t, result.Vulnerabilities, 1)
	assert.Len(t, result.Vulnerabilities[0].Description, 500, "description should be truncated to 500 bytes")
	assert.Equal(t, "...", result.Vulnerabilities[0].Description[497:], "should end with ...")
	assert.True(t, utf8.ValidString(result.Vulnerabilities[0].Description))
}

func TestTruncateString_UTF8Safe(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		maxLen int
	}{
		{"ascii short", "hello", 10},
		{"ascii exact", "hello", 5},
		{"ascii truncate", "hello world", 8},
		{"cjk mid-rune boundary", "中文漏洞描述", 5},
		{"cjk longer", strings.Repeat("中文", 40), 50},
		{"emoji", "🚨CRITICAL vulnerability description", 10},
		{"empty", "", 10},
		{"max zero", "abc", 0},
		{"max three", "abcdef", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := truncateString(tt.input, tt.maxLen)
			assert.True(t, utf8.ValidString(out), "output must be valid UTF-8: %q", out)
			if tt.maxLen > 0 {
				assert.LessOrEqual(t, len(out), tt.maxLen)
			} else {
				assert.Empty(t, out)
			}
			if len(tt.input) <= tt.maxLen {
				assert.Equal(t, tt.input, out)
			}
		})
	}
}

// ============================================================================
// Multi-result section tests
// ============================================================================

func TestParseTrivyResult_MultipleResultSections(t *testing.T) {
	// Debian .NET has 2 result sections (os-pkgs + dotnet-core lang-pkgs)
	result := parseTrivyResult(debianDotnetRuntimeTrivyOutput())
	require.NotNil(t, result)
	// All vulns are in the first section, second section has 0
	assert.Equal(t, 6, result.TotalVulnerabilities)
}

func TestParseTrivyResult_VulnsAcrossMultipleResults(t *testing.T) {
	output := &trivyOutput{
		Results: []trivyResult{
			{
				Vulnerabilities: makeTrivyVulns([]struct {
					id, severity, pkg, installed, fixed, desc string
					cvss                                      map[string]trivyCVSS
				}{
					{"CVE-001", "HIGH", "pkg1", "1.0", "1.1", "vuln 1", nil},
					{"CVE-002", "LOW", "pkg2", "2.0", "", "vuln 2", nil},
				}),
			},
			{
				Vulnerabilities: makeTrivyVulns([]struct {
					id, severity, pkg, installed, fixed, desc string
					cvss                                      map[string]trivyCVSS
				}{
					{"CVE-003", "CRITICAL", "pkg3", "3.0", "3.1", "vuln 3", nil},
				}),
			},
		},
	}

	result := parseTrivyResult(output)
	assert.Equal(t, 3, result.TotalVulnerabilities)
	assert.Equal(t, 1, result.CriticalVulnerabilities)
	assert.Equal(t, 1, result.HighVulnerabilities)
	assert.Equal(t, 1, result.LowVulnerabilities)
	assert.Len(t, result.Vulnerabilities, 3)
}

// ============================================================================
// Secrets / misconfig / license tests
// ============================================================================

func TestParseTrivyResult_SecretsAndMisconfigs(t *testing.T) {
	output := &trivyOutput{
		Results: []trivyResult{
			{
				Secrets: []trivySecret{
					{RuleID: "aws-secret-key", Severity: "CRITICAL", Title: "AWS secret key"},
					{RuleID: "generic-api-key", Severity: "HIGH", Title: "API key"},
				},
				Misconfigs: []trivyMisconfig{
					{ID: "DS001", Severity: "HIGH", Title: "root user"},
				},
				Licenses: []trivyLicense{
					{Name: "GPL-3.0", Severity: "HIGH"},
					{Name: "AGPL-3.0", Severity: "CRITICAL"},
				},
			},
		},
	}

	result := parseTrivyResult(output)
	assert.Equal(t, 2, result.SecretsFound)
	assert.Equal(t, 1, result.ConfigIssues)
	assert.Equal(t, 2, result.LicenseIssues)
	assert.Equal(t, 0, result.TotalVulnerabilities, "secrets/misconfigs/licenses don't count as vulns")

	require.Len(t, result.SecurityFindings, 3)
	assert.Equal(t, "secret", result.SecurityFindings[0].FindingType)
	assert.Equal(t, "aws-secret-key", result.SecurityFindings[0].RuleID)
	assert.Equal(t, "CRITICAL", result.SecurityFindings[0].Severity)
	assert.Equal(t, "secret", result.SecurityFindings[1].FindingType)
	assert.Equal(t, "generic-api-key", result.SecurityFindings[1].RuleID)
	assert.Equal(t, "misconfiguration", result.SecurityFindings[2].FindingType)
	assert.Equal(t, "DS001", result.SecurityFindings[2].RuleID)
	assert.Equal(t, "root user", result.SecurityFindings[2].Title)
}

// ============================================================================
// Edge case tests
// ============================================================================

func TestParseTrivyResult_EmptyOutput(t *testing.T) {
	result := parseTrivyResult(emptyTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalVulnerabilities)
	assert.Empty(t, result.Vulnerabilities)
}

func TestParseTrivyResult_NoVulnerabilities(t *testing.T) {
	result := parseTrivyResult(noVulnsTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalVulnerabilities)
	assert.Empty(t, result.Vulnerabilities)
}

func TestParseTrivyResult_NilVulnerabilities(t *testing.T) {
	result := parseTrivyResult(nilVulnsTrivyOutput())
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalVulnerabilities)
	assert.Empty(t, result.Vulnerabilities)
}

func TestParseTrivyResult_UnknownSeverity(t *testing.T) {
	output := &trivyOutput{
		Results: []trivyResult{
			{
				Vulnerabilities: []trivyVuln{
					{VulnerabilityID: "CVE-UNKNOWN", Severity: "UNKNOWN"},
					{VulnerabilityID: "CVE-WEIRD", Severity: "WEIRD_SEVERITY"},
				},
			},
		},
	}

	result := parseTrivyResult(output)
	assert.Equal(t, 2, result.TotalVulnerabilities)
	assert.Equal(t, 2, result.UnknownVulnerabilities, "unrecognized severities should count as unknown")
}

func TestParseTrivyResult_NegligibleSeverity(t *testing.T) {
	output := &trivyOutput{
		Results: []trivyResult{
			{
				Vulnerabilities: []trivyVuln{
					{VulnerabilityID: "CVE-NEG", Severity: "NEGLIGIBLE"},
				},
			},
		},
	}

	result := parseTrivyResult(output)
	assert.Equal(t, 1, result.TotalVulnerabilities)
	assert.Equal(t, 1, result.NegligibleVulnerabilities)
}

func TestParseTrivyResult_CaseInsensitiveSeverity(t *testing.T) {
	output := &trivyOutput{
		Results: []trivyResult{
			{
				Vulnerabilities: []trivyVuln{
					{VulnerabilityID: "CVE-1", Severity: "critical"},
					{VulnerabilityID: "CVE-2", Severity: "High"},
					{VulnerabilityID: "CVE-3", Severity: "medium"},
					{VulnerabilityID: "CVE-4", Severity: "Low"},
				},
			},
		},
	}

	result := parseTrivyResult(output)
	assert.Equal(t, 4, result.TotalVulnerabilities)
	assert.Equal(t, 1, result.CriticalVulnerabilities)
	assert.Equal(t, 1, result.HighVulnerabilities)
	assert.Equal(t, 1, result.MediumVulnerabilities)
	assert.Equal(t, 1, result.LowVulnerabilities)
}

// ============================================================================
// OS metadata extraction tests
// ============================================================================

func TestParseTrivyResult_OSMetadata_EmptyOutput(t *testing.T) {
	result := parseTrivyResult(emptyTrivyOutput())
	assert.Equal(t, "", result.BaseOSFamily)
	assert.Equal(t, "", result.BaseOSVersion)
}

func TestParseTrivyResult_OSMetadata_NoMetadataField(t *testing.T) {
	// Output with results but no Metadata — simulates older Trivy or non-OS images
	output := &trivyOutput{
		Results: []trivyResult{
			{Target: "image:1.0", Class: "os-pkgs", Type: "alpine"},
		},
	}
	result := parseTrivyResult(output)
	assert.Equal(t, "", result.BaseOSFamily)
	assert.Equal(t, "", result.BaseOSVersion)
}

func TestParseTrivyResult_OSMetadata_AllOSFamilies(t *testing.T) {
	// Verify all three OS families we encounter in MCR
	tests := []struct {
		name, family, version string
	}{
		{"Azure Linux", "azurelinux", "3.0"},
		{"Ubuntu", "ubuntu", "22.04"},
		{"Debian", "debian", "12.13"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &trivyOutput{
				Metadata: trivyMetadata{OS: trivyOS{Family: tt.family, Name: tt.version}},
				Results:  []trivyResult{{Target: "test-image"}},
			}
			result := parseTrivyResult(output)
			assert.Equal(t, tt.family, result.BaseOSFamily)
			assert.Equal(t, tt.version, result.BaseOSVersion)
		})
	}
}
