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
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) (*sql.DB, *database.Repository) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec("PRAGMA foreign_keys=ON")
	require.NoError(t, err)

	require.NoError(t, database.CreateTables(db))

	return db, database.NewRepository(db)
}

func TestGenerateJSONReport_MultipleOSes(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			Digest: "sha256:abc123", SizeBytes: 85000000, CreatedDate: "2025-04-15T08:30:00Z",
			CriticalVulnerabilities: 0, HighVulnerabilities: 0, TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName: "ubuntu", BaseOSVersion: "22.04",
			Digest: "sha256:def456", SizeBytes: 120000000,
			CriticalVulnerabilities: 1, HighVulnerabilities: 2, TotalVulnerabilities: 10,
			Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
		},
		{
			Name: "azl-go:1.21", Registry: "r", Repository: "repo3", Tag: "1.21",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			CriticalVulnerabilities: 0, TotalVulnerabilities: 0,
			Languages: []domain.Language{{Language: "go", Version: "1.21.0"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Equal(t, 10, report.TopN)
	require.Len(t, report.Images, 3)

	// Verify flat structure: each entry has language and baseOS
	for _, img := range report.Images {
		assert.NotEmpty(t, img.Language)
		assert.NotEmpty(t, img.BaseOS)
		assert.NotEmpty(t, img.Name)
		assert.Greater(t, img.Rank, 0)
	}

	// Verify python entries have both OSes
	var pythonOSes []string
	for _, img := range report.Images {
		if img.Language == "python" {
			pythonOSes = append(pythonOSes, img.BaseOS)
		}
	}
	assert.Contains(t, pythonOSes, "Azure Linux")
	assert.Contains(t, pythonOSes, "Ubuntu")

	var azlPython *JSONImageEntry
	for i := range report.Images {
		if report.Images[i].Name == "azl-python:3.12" {
			azlPython = &report.Images[i]
		}
	}
	require.NotNil(t, azlPython)
	assert.Equal(t, "2025-04-15T08:30:00Z", azlPython.CreatedDate)
}

func TestGenerateJSONReport_RankResetsPerOSGroup(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "azl-python:3.11", Registry: "r", Repository: "repo2", Tag: "3.11",
			BaseOSName: "azurelinux", TotalVulnerabilities: 5,
			Languages: []domain.Language{{Language: "python", Version: "3.11"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo3", Tag: "3.12",
			BaseOSName: "ubuntu", TotalVulnerabilities: 3,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 3)

	// Azure Linux images: rank 1, 2
	assert.Equal(t, 1, report.Images[0].Rank)
	assert.Equal(t, "Azure Linux", report.Images[0].BaseOS)
	assert.Equal(t, 2, report.Images[1].Rank)
	assert.Equal(t, "Azure Linux", report.Images[1].BaseOS)

	// Ubuntu: rank resets to 1
	assert.Equal(t, 1, report.Images[2].Rank)
	assert.Equal(t, "Ubuntu", report.Images[2].BaseOS)
}

func TestGenerateJSONReport_TopNLimit(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: "azl-py:" + string(rune('a'+i)), Registry: "r", Repository: "repo", Tag: string(rune('a' + i)),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 3, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Len(t, report.Images, 3)
	assert.Equal(t, 3, report.TopN)
}

func TestGenerateJSONReport_ZeroTopNReturnsAll(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: "azl-py:" + string(rune('a'+i)), Registry: "r", Repository: "repo", Tag: string(rune('a' + i)),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 0, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Len(t, report.Images, 5)
	assert.Equal(t, 0, report.TopN)
}

func TestGenerateJSONReport_EmptyDB(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	_, err := os.Stat(outPath)
	assert.True(t, os.IsNotExist(err), "empty DB should not produce a report file")
}

func TestGenerateJSONReport_IncludesHelperFields(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "mcr.microsoft.com/azurelinux/base/python:3.12", Registry: "mcr.microsoft.com",
		Repository: "azurelinux/base/python", Tag: "3.12",
		BaseOSName: "azurelinux", Digest: "sha256:abcdef1234567890",
		SizeBytes: 85000000, CreatedDate: "2025-04-15T08:30:00Z", TotalVulnerabilities: 1,
		Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1)
	entry := report.Images[0]
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12@sha256:abcdef1234567890", entry.PinnedReference)
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", entry.StableTag)
	assert.Equal(t, "FROM mcr.microsoft.com/azurelinux/base/python:3.12@sha256:abcdef1234567890", entry.DockerfileFrom)
	assert.Equal(t, "81.1 MB", entry.SizeHuman)
	assert.Equal(t, "2025-04-15T08:30:00Z", entry.CreatedDate)
	assert.Equal(t, "Azure Linux", entry.BaseOS)
	assert.Equal(t, "python", entry.Language)
}

func TestGenerateJSONReport_UnknownOSAsOther(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "unknown-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
		TotalVulnerabilities: 1,
		Languages:            []domain.Language{{Language: "python", Version: "3.12"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1)
	assert.Equal(t, "Other", report.Images[0].BaseOS)
	assert.Equal(t, "", report.Images[0].CreatedDate)
	assert.Contains(t, string(data), `"createdDate": ""`)
}

func TestGenerateJSONReport_BaseLanguageAppearsLast(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "base-core:3.0", Registry: "r", Repository: "repo", Tag: "3.0",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			CriticalVulnerabilities: 0, TotalVulnerabilities: 0,
			Languages: []domain.Language{{Language: "base", Version: "3.0", PackageType: "base"}},
		},
		{
			Name: "python-img:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName:              "azurelinux",
			CriticalVulnerabilities: 1, TotalVulnerabilities: 5,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 2)
	// python should come first, base last
	assert.Equal(t, "python", report.Images[0].Language)
	assert.Equal(t, "base", report.Images[1].Language)
}

func TestGenerateMarkdownReport_BaseLanguageSection(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "base-core:3.0", Registry: "r", Repository: "repo", Tag: "3.0",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			Digest: "sha256:abc123", SizeBytes: 50000000, CreatedDate: "2025-04-15T08:30:00Z",
			CriticalVulnerabilities: 0, TotalVulnerabilities: 1,
			Languages: []domain.Language{{Language: "base", Version: "3.0", PackageType: "base"}},
		},
		{
			Name: "python-img:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName:              "azurelinux",
			CriticalVulnerabilities: 0, TotalVulnerabilities: 3,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.md")
	require.NoError(t, GenerateReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)
	content := string(data)

	// Python section should appear before Base section
	pythonIdx := strings.Index(content, "## Python")
	baseIdx := strings.Index(content, "## Base / No Runtime")
	assert.Greater(t, pythonIdx, -1, "should contain Python section")
	assert.Greater(t, baseIdx, -1, "should contain Base / No Runtime section")
	assert.Less(t, pythonIdx, baseIdx, "Python should appear before Base / No Runtime")

	// Base section should contain the base image
	assert.Contains(t, content, "base-core:3.0")
	assert.Contains(t, content, "| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |")
	assert.Contains(t, content, "2025-04-15")
	assert.Contains(t, content, "| 1 | `python-img:3.12` | 3.12 | - | 0 | 0 | 3 | - | - | `` | `-` |")
}

func TestDisplayLanguageName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"base", "Base / No Runtime"},
		{"python", "Python"},
		{"go", "Go"},
		{"dotnet", "Dotnet"},
		{"java", "Java"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, DisplayLanguageName(tt.input))
		})
	}
}

func TestGenerateJSONReport_DeduplicatesAlternateTags(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "mcr.microsoft.com/azurelinux/base/nodejs:24", Registry: "mcr.microsoft.com",
			Repository: "azurelinux/base/nodejs", Tag: "24",
			BaseOSName: "azurelinux", Digest: "sha256:samedigest",
			SizeBytes: 100000000, TotalVulnerabilities: 1,
			Languages: []domain.Language{{Language: "node", Version: "24.14.0"}},
		},
		{
			Name: "mcr.microsoft.com/azurelinux/base/nodejs:24.14", Registry: "mcr.microsoft.com",
			Repository: "azurelinux/base/nodejs", Tag: "24.14",
			BaseOSName: "azurelinux", Digest: "sha256:samedigest",
			SizeBytes: 100000000, TotalVulnerabilities: 1,
			Languages: []domain.Language{{Language: "node", Version: "24.14.0"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1, "duplicate digests should be merged into one entry")
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/nodejs:24.14", report.Images[0].Name)
	require.Len(t, report.Images[0].AlternateTags, 1)
	assert.Equal(t, ":24", report.Images[0].AlternateTags[0])
}

func TestGenerateMarkdownReport_DeduplicatesWithAlternateColumn(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "mcr.microsoft.com/azurelinux/base/python:3", Registry: "mcr.microsoft.com",
			Repository: "azurelinux/base/python", Tag: "3",
			BaseOSName: "azurelinux", Digest: "sha256:pydigest",
			SizeBytes: 85000000, TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12.9"}},
		},
		{
			Name: "mcr.microsoft.com/azurelinux/base/python:3.12", Registry: "mcr.microsoft.com",
			Repository: "azurelinux/base/python", Tag: "3.12",
			BaseOSName: "azurelinux", Digest: "sha256:pydigest",
			SizeBytes: 85000000, TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12.9"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.md")
	require.NoError(t, GenerateReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)
	content := string(data)

	// Should have only one data row (deduped), not two
	assert.Contains(t, content, "`mcr.microsoft.com/azurelinux/base/python:3.12`")
	assert.Contains(t, content, "| :3 |", "alternate tag should appear in Also Tagged As column")
	// Only one ranked row should exist (rank 1, no rank 2)
	assert.NotContains(t, content, "| 2 | `mcr", "should have only one image row after dedup")
}
