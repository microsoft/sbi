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

package database

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) (*sql.DB, *Repository) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec("PRAGMA foreign_keys=ON")
	require.NoError(t, err)

	err = CreateTables(db)
	require.NoError(t, err)

	return db, NewRepository(db)
}

func TestInsertAndQueryImage(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name:                    "mcr.microsoft.com/azurelinux/base/python:3.12",
		Registry:                "mcr.microsoft.com",
		Repository:              "azurelinux/base/python",
		Tag:                     "3.12",
		Digest:                  "sha256:abcdef1234567890",
		SizeBytes:               85000000,
		CreatedDate:             "2025-04-15T08:30:00Z",
		TotalVulnerabilities:    5,
		CriticalVulnerabilities: 0,
		HighVulnerabilities:     1,
		MediumVulnerabilities:   2,
		LowVulnerabilities:      2,
		Languages: []domain.Language{
			{Language: "python", Version: "3.12.1", MajorMinor: "3.12", PackageName: "python3", Verified: true},
		},
	}

	err := repo.InsertImage(img)
	require.NoError(t, err)

	// Test ImageExists
	exists, err := repo.ImageExists("mcr.microsoft.com/azurelinux/base/python:3.12")
	require.NoError(t, err)
	assert.True(t, exists)

	exists, err = repo.ImageExists("nonexistent:latest")
	require.NoError(t, err)
	assert.False(t, exists)

	// Test QueryLanguages
	languages, err := repo.QueryLanguages()
	require.NoError(t, err)
	assert.Equal(t, []string{"python"}, languages)

	// Test QueryTopImages
	images, err := repo.QueryTopImages("python", 10)
	require.NoError(t, err)
	require.Len(t, images, 1)
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", images[0].Name)
	assert.Equal(t, "3.12.1", images[0].Version)
	assert.Equal(t, 0, images[0].CriticalVulnerabilities)
	assert.Equal(t, 1, images[0].HighVulnerabilities)
	assert.Equal(t, 5, images[0].TotalVulnerabilities)
	assert.Equal(t, int64(85000000), images[0].SizeBytes)
	assert.Equal(t, "2025-04-15T08:30:00Z", images[0].CreatedDate)
}

func TestQueryTopImagesRanking(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "img-high-crit:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
			CriticalVulnerabilities: 2, HighVulnerabilities: 0, TotalVulnerabilities: 5, SizeBytes: 100,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "img-low-crit:1.0", Registry: "r", Repository: "repo2", Tag: "1.0",
			CriticalVulnerabilities: 0, HighVulnerabilities: 1, TotalVulnerabilities: 3, SizeBytes: 200,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "img-smallest:1.0", Registry: "r", Repository: "repo3", Tag: "1.0",
			CriticalVulnerabilities: 0, HighVulnerabilities: 0, TotalVulnerabilities: 1, SizeBytes: 50,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	result, err := repo.QueryTopImages("python", 10)
	require.NoError(t, err)
	require.Len(t, result, 3)

	// img-smallest should be first (0 crit, 0 high, 1 total, smallest)
	assert.Equal(t, "img-smallest:1.0", result[0].Name)
	// img-low-crit second (0 crit, 1 high)
	assert.Equal(t, "img-low-crit:1.0", result[1].Name)
	// img-high-crit last (2 crit)
	assert.Equal(t, "img-high-crit:1.0", result[2].Name)
}

func TestClearDatabase(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "test:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
		Languages: []domain.Language{{Language: "go", Version: "1.21"}},
	}

	require.NoError(t, repo.InsertImage(img))

	stats, err := repo.GetDatabaseStats()
	require.NoError(t, err)
	assert.Equal(t, 1, stats["images"])
	assert.Equal(t, 1, stats["languages"])

	require.NoError(t, repo.ClearDatabase())

	stats, err = repo.GetDatabaseStats()
	require.NoError(t, err)
	assert.Equal(t, 0, stats["images"])
	assert.Equal(t, 0, stats["languages"])
}

func TestUpsertImage(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "test:1.0", Registry: "r", Repository: "repo", Tag: "1.0",
		CreatedDate:          "2025-04-15T08:30:00Z",
		TotalVulnerabilities: 5,
		Languages:            []domain.Language{{Language: "python", Version: "3.12"}},
	}

	require.NoError(t, repo.InsertImage(img))

	// Update the same image with new vuln count
	img.CreatedDate = "2025-05-01T12:00:00Z"
	img.TotalVulnerabilities = 3
	img.Languages = []domain.Language{{Language: "python", Version: "3.12.1"}}
	require.NoError(t, repo.InsertImage(img))

	images, err := repo.QueryTopImages("python", 10)
	require.NoError(t, err)
	require.Len(t, images, 1)
	assert.Equal(t, 3, images[0].TotalVulnerabilities)
	assert.Equal(t, "3.12.1", images[0].Version)
	assert.Equal(t, "2025-05-01T12:00:00Z", images[0].CreatedDate)
}

func TestUpsertImage_UpdatesBaseOS(t *testing.T) {
	testDB, repo := setupTestDB(t)
	defer func() { _ = testDB.Close() }()

	// Insert image with no OS data
	img := &domain.ImageRecord{
		Name: "python-img:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
		TotalVulnerabilities: 5,
		Languages:            []domain.Language{{Language: "python", Version: "3.12"}},
	}
	require.NoError(t, repo.InsertImage(img))

	var osName string
	err := testDB.QueryRow("SELECT COALESCE(base_os_name, '') FROM images WHERE name = ?", img.Name).Scan(&osName)
	require.NoError(t, err)
	assert.Equal(t, "", osName)

	// Re-insert same image, now with OS data
	img.BaseOSName = "azurelinux"
	img.BaseOSVersion = "3.0"
	require.NoError(t, repo.InsertImage(img))

	var osVersion string
	err = testDB.QueryRow("SELECT base_os_name, base_os_version FROM images WHERE name = ?", img.Name).Scan(&osName, &osVersion)
	require.NoError(t, err)
	assert.Equal(t, "azurelinux", osName)
	assert.Equal(t, "3.0", osVersion)
}

func TestQueryBaseOSes(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName: "ubuntu", BaseOSVersion: "22.04",
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "debian-python:3.12", Registry: "r", Repository: "repo3", Tag: "3.12",
			BaseOSName: "debian", BaseOSVersion: "12",
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	oses, err := repo.QueryBaseOSes("python")
	require.NoError(t, err)
	assert.Equal(t, []string{"azurelinux", "debian", "ubuntu"}, oses)
}

func TestQueryBaseOSes_EmptyOSGroupedAsOther(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux",
			Languages:  []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "unknown-python:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	oses, err := repo.QueryBaseOSes("python")
	require.NoError(t, err)
	assert.Equal(t, []string{"azurelinux", "Other"}, oses)
}

func TestQueryTopImagesByOS(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", CreatedDate: "2025-04-15T08:30:00Z", TotalVulnerabilities: 3,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName: "ubuntu", TotalVulnerabilities: 5,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	result, err := repo.QueryTopImagesByOS("python", "azurelinux", 10)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "azl-python:3.12", result[0].Name)
	assert.Equal(t, "azurelinux", result[0].BaseOSName)
	assert.Equal(t, "2025-04-15T08:30:00Z", result[0].CreatedDate)
}

func TestQueryTopImagesByOS_OtherGroup(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "unknown-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
		TotalVulnerabilities: 2,
		Languages:            []domain.Language{{Language: "python", Version: "3.12"}},
	}
	require.NoError(t, repo.InsertImage(img))

	result, err := repo.QueryTopImagesByOS("python", "Other", 10)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "unknown-python:3.12", result[0].Name)
	assert.Equal(t, "Other", result[0].BaseOSName)
}

func TestQueryTopImagesByOS_Limit(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: fmt.Sprintf("azl-python:%d", i), Registry: "r", Repository: "repo", Tag: fmt.Sprintf("%d", i),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	result, err := repo.QueryTopImagesByOS("python", "azurelinux", 3)
	require.NoError(t, err)
	require.Len(t, result, 3)
	assert.Equal(t, "azl-python:0", result[0].Name)
}

func TestQueryTopImagesByOS_ZeroMeansUnlimited(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: fmt.Sprintf("azl-python:%d", i), Registry: "r", Repository: "repo", Tag: fmt.Sprintf("%d", i),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	result, err := repo.QueryTopImagesByOS("python", "azurelinux", 0)
	require.NoError(t, err)
	assert.Len(t, result, 5)
}

func TestQueryLanguages_BaseSortsLast(t *testing.T) {
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
			CriticalVulnerabilities: 5, TotalVulnerabilities: 20,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "go-img:1.21", Registry: "r", Repository: "repo3", Tag: "1.21",
			CriticalVulnerabilities: 2, TotalVulnerabilities: 10,
			Languages: []domain.Language{{Language: "go", Version: "1.21"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	languages, err := repo.QueryLanguages()
	require.NoError(t, err)
	require.Len(t, languages, 3)

	// "base" should be last despite having 0 vulnerabilities
	assert.Equal(t, "go", languages[0])
	assert.Equal(t, "python", languages[1])
	assert.Equal(t, "base", languages[2])
}

func TestInsertAndQuerySecurityFindingsAndCapabilities(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name:         "mcr.microsoft.com/example:1.0",
		Registry:     "mcr.microsoft.com",
		Repository:   "example",
		Tag:          "1.0",
		SecretsFound: 1,
		ConfigIssues: 1,
		Languages: []domain.Language{
			{Language: "python", Version: "3.12"},
		},
		Capabilities: []domain.Capability{
			{Capability: "ssl"},
			{Capability: "http_client"},
		},
		SecurityFindings: []domain.SecurityFinding{
			{
				FindingType: "secret",
				Severity:    "CRITICAL",
				RuleID:      "aws-secret-key",
				Title:       "AWS secret key",
				Description: "AKIA...",
				Category:    "AWS",
			},
			{
				FindingType: "misconfiguration",
				Severity:    "HIGH",
				RuleID:      "DS001",
				Title:       "root user",
				Message:     "Dockerfile runs as root",
			},
		},
	}
	require.NoError(t, repo.InsertImage(img))

	details, err := repo.QueryAllImageDetails()
	require.NoError(t, err)
	require.Len(t, details, 1)

	got := details[0]
	require.Len(t, got.SecurityFindings, 2)
	assert.Equal(t, "secret", got.SecurityFindings[0].FindingType)
	assert.Equal(t, "aws-secret-key", got.SecurityFindings[0].RuleID)
	assert.Equal(t, "misconfiguration", got.SecurityFindings[1].FindingType)
	assert.Equal(t, "DS001", got.SecurityFindings[1].RuleID)

	require.Len(t, got.Capabilities, 2)
	assert.Equal(t, "ssl", got.Capabilities[0].Capability)
	assert.Equal(t, "http_client", got.Capabilities[1].Capability)

	// Upsert clears and rewrites related rows
	img.SecurityFindings = []domain.SecurityFinding{
		{FindingType: "secret", Severity: "LOW", RuleID: "generic-api-key", Title: "API key"},
	}
	img.Capabilities = []domain.Capability{{Capability: "compression"}}
	img.SecretsFound = 1
	img.ConfigIssues = 0
	require.NoError(t, repo.InsertImage(img))

	details, err = repo.QueryAllImageDetails()
	require.NoError(t, err)
	require.Len(t, details, 1)
	require.Len(t, details[0].SecurityFindings, 1)
	assert.Equal(t, "generic-api-key", details[0].SecurityFindings[0].RuleID)
	require.Len(t, details[0].Capabilities, 1)
	assert.Equal(t, "compression", details[0].Capabilities[0].Capability)
}

// TestInsertImage_RollbackOnClearFailure ensures a failure while clearing
// related tables rolls back the whole upsert and releases the connection.
//
// This guards against the classic Go bug where `if _, err := tx.Exec(...)`
// shadows the outer err used by deferred rollback, leaving the transaction
// open (and, with MaxOpenConns(1), blocking all further DB use).
func TestInsertImage_RollbackOnClearFailure(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "rollback.db")
	db, err := sql.Open("sqlite", dbPath)
	require.NoError(t, err)
	defer func() { _ = db.Close() }()

	// Match production OpenDB: a single connection so a leaked tx is observable.
	db.SetMaxOpenConns(1)
	_, err = db.Exec("PRAGMA foreign_keys=ON")
	require.NoError(t, err)
	require.NoError(t, CreateTables(db))

	repo := NewRepository(db)
	img := &domain.ImageRecord{
		Name:                 "test:1.0",
		Registry:             "r",
		Repository:           "repo",
		Tag:                  "1.0",
		TotalVulnerabilities: 5,
		Languages: []domain.Language{
			{Language: "python", Version: "3.12"},
		},
	}
	require.NoError(t, repo.InsertImage(img))

	// Force the related-table clear path to fail on the next upsert.
	_, err = db.Exec("DROP TABLE languages")
	require.NoError(t, err)

	img.TotalVulnerabilities = 99
	err = repo.InsertImage(img)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "clearing languages")

	// Upsert must not have been committed.
	var total int
	err = db.QueryRow(`SELECT total_vulnerabilities FROM images WHERE name = ?`, img.Name).Scan(&total)
	require.NoError(t, err)
	assert.Equal(t, 5, total, "failed clear must roll back the image upsert")

	// Connection must be released so further writes can proceed.
	require.NoError(t, CreateTables(db)) // recreates languages (IF NOT EXISTS)
	img.TotalVulnerabilities = 7
	img.Languages = []domain.Language{{Language: "python", Version: "3.12.1"}}
	require.NoError(t, repo.InsertImage(img))

	err = db.QueryRow(`SELECT total_vulnerabilities FROM images WHERE name = ?`, img.Name).Scan(&total)
	require.NoError(t, err)
	assert.Equal(t, 7, total)
}
