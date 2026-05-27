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
	"os"
	"path/filepath"
	"testing"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateDetailJSONReport_FullImage(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "mcr.microsoft.com/azurelinux/base/python:3.12", Registry: "mcr.microsoft.com",
		Repository: "azurelinux/base/python", Tag: "3.12",
		BaseOSName: "azurelinux", BaseOSVersion: "3.0",
		Digest: "sha256:abc123", SizeBytes: 85000000, Layers: 5,
		CreatedDate: "2025-04-15T08:30:00Z",
		TotalVulnerabilities: 3, CriticalVulnerabilities: 1, HighVulnerabilities: 1,
		MediumVulnerabilities: 1,
		Languages: []domain.Language{
			{Language: "python", Version: "3.12.9", MajorMinor: "3.12", PackageName: "python3", PackageType: "rpm", Verified: true},
		},
		Vulnerabilities: []domain.Vulnerability{
			{VulnerabilityID: "CVE-2025-0001", Severity: "CRITICAL", PackageName: "openssl", PackageVersion: "3.0.1", FixedVersion: "3.0.2", CVSSScore: 9.8, Description: "Buffer overflow"},
			{VulnerabilityID: "CVE-2025-0002", Severity: "HIGH", PackageName: "curl", PackageVersion: "7.88.0", FixedVersion: "7.88.1", CVSSScore: 7.5},
			{VulnerabilityID: "CVE-2025-0003", Severity: "MEDIUM", PackageName: "zlib", PackageVersion: "1.2.13", CVSSScore: 5.0},
		},
		SystemPackages: []domain.SystemPackage{
			{Name: "openssl", Version: "3.0.1", PackageType: "rpm"},
			{Name: "curl", Version: "7.88.0", PackageType: "rpm"},
		},
		PackageManagers: []domain.PackageManager{
			{Name: "pip", Version: "24.0", Language: "python"},
		},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report DetailJSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Equal(t, 1, report.ImageCount)
	assert.Equal(t, 1, report.SchemaVersion)
	require.Len(t, report.Images, 1)

	entry := report.Images[0]
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", entry.Name)
	assert.Equal(t, "mcr.microsoft.com", entry.Registry)
	assert.Equal(t, "azurelinux/base/python", entry.Repository)
	assert.Equal(t, "3.12", entry.Tag)
	assert.Equal(t, "sha256:abc123", entry.Digest)
	assert.Equal(t, int64(85000000), entry.SizeBytes)
	assert.Equal(t, "81.1 MB", entry.SizeHuman)
	assert.Equal(t, 5, entry.Layers)
	assert.Equal(t, "2025-04-15T08:30:00Z", entry.CreatedDate)

	assert.Equal(t, "Azure Linux", entry.BaseOS.Name)
	assert.Equal(t, "3.0", entry.BaseOS.Version)

	assert.Equal(t, 3, entry.VulnerabilitySummary.Total)
	assert.Equal(t, 1, entry.VulnerabilitySummary.Critical)
	assert.Equal(t, 1, entry.VulnerabilitySummary.High)
	assert.Equal(t, 1, entry.VulnerabilitySummary.Medium)

	require.Len(t, entry.Languages, 1)
	assert.Equal(t, "python", entry.Languages[0].Language)
	assert.Equal(t, "3.12.9", entry.Languages[0].Version)
	assert.True(t, entry.Languages[0].Verified)

	require.Len(t, entry.Vulnerabilities, 3)
	assert.Equal(t, "CVE-2025-0001", entry.Vulnerabilities[0].ID)
	assert.Equal(t, "CRITICAL", entry.Vulnerabilities[0].Severity)
	assert.Equal(t, 9.8, entry.Vulnerabilities[0].CVSSScore)
	assert.Equal(t, "3.0.2", entry.Vulnerabilities[0].FixedVersion)

	require.Len(t, entry.SystemPackages, 2)
	assert.Equal(t, "openssl", entry.SystemPackages[0].Name)

	require.Len(t, entry.PackageManagers, 1)
	assert.Equal(t, "pip", entry.PackageManagers[0].Name)
}

func TestGenerateDetailJSONReport_EmptyDB(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	_, err := os.Stat(outPath)
	assert.True(t, os.IsNotExist(err), "empty DB should not produce a detail report file")
}

func TestGenerateDetailJSONReport_EmptyDBRemovesStaleFile(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, os.WriteFile(outPath, []byte(`{"stale": true}`), 0o644))

	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	_, err := os.Stat(outPath)
	assert.True(t, os.IsNotExist(err), "stale detail report should be removed when DB is empty")
}

func TestGenerateDetailJSONReport_MultipleImages(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "python-img:3.12", Registry: "r", Repository: "repo1", Tag: "3.12",
			BaseOSName: "azurelinux", TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12.9"}},
			Vulnerabilities: []domain.Vulnerability{
				{VulnerabilityID: "CVE-2025-1111", Severity: "HIGH", PackageName: "pkg1", PackageVersion: "1.0"},
				{VulnerabilityID: "CVE-2025-1112", Severity: "LOW", PackageName: "pkg2", PackageVersion: "2.0"},
			},
			SystemPackages: []domain.SystemPackage{{Name: "pkg1", Version: "1.0"}},
		},
		{
			Name: "go-img:1.21", Registry: "r", Repository: "repo2", Tag: "1.21",
			BaseOSName: "azurelinux", TotalVulnerabilities: 0,
			Languages:  []domain.Language{{Language: "go", Version: "1.21.0"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report DetailJSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Equal(t, 2, report.ImageCount)
	require.Len(t, report.Images, 2)

	// Images ordered by name
	assert.Equal(t, "go-img:1.21", report.Images[0].Name)
	assert.Equal(t, "python-img:3.12", report.Images[1].Name)

	// Verify child data attaches to correct image
	assert.Len(t, report.Images[0].Vulnerabilities, 0)
	assert.Len(t, report.Images[0].Languages, 1)
	assert.Equal(t, "go", report.Images[0].Languages[0].Language)

	assert.Len(t, report.Images[1].Vulnerabilities, 2)
	assert.Len(t, report.Images[1].SystemPackages, 1)
}

func TestGenerateDetailJSONReport_VulnSummaryCounts(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "test-img:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
		TotalVulnerabilities: 6, CriticalVulnerabilities: 1, HighVulnerabilities: 2,
		MediumVulnerabilities: 1, LowVulnerabilities: 1, NegligibleVulnerabilities: 0,
		UnknownVulnerabilities: 1,
		Languages: []domain.Language{{Language: "python", Version: "3.12"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report DetailJSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1)
	summary := report.Images[0].VulnerabilitySummary
	assert.Equal(t, 6, summary.Total)
	assert.Equal(t, 1, summary.Critical)
	assert.Equal(t, 2, summary.High)
	assert.Equal(t, 1, summary.Medium)
	assert.Equal(t, 1, summary.Low)
	assert.Equal(t, 0, summary.Negligible)
	assert.Equal(t, 1, summary.Unknown)
}

func TestGenerateDetailJSONReport_EmptyChildSlicesNotNull(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "minimal:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
		Languages: []domain.Language{{Language: "base", Version: "3.0"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	// Empty arrays should serialize as [] not null
	content := string(data)
	assert.Contains(t, content, `"vulnerabilities": []`)
	assert.Contains(t, content, `"systemPackages": []`)
	assert.Contains(t, content, `"packageManagers": []`)
}

func TestGenerateDetailJSONReport_CreatesOutputDir(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "test:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
		Languages: []domain.Language{{Language: "go", Version: "1.21"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "nested", "dir", "detail.json")
	require.NoError(t, GenerateDetailJSONReport(repo, outPath))

	_, err := os.Stat(outPath)
	assert.NoError(t, err, "output file should exist in newly created directory")
}
