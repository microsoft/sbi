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
	"database/sql"
	"errors"
	"testing"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func TestScanStatsResult(t *testing.T) {
	tests := []struct {
		name      string
		attempted int
		failed    int
		wantErr   bool
		errSubstr string
	}{
		{"no attempts", 0, 0, false, ""},
		{"all success", 5, 0, false, ""},
		{"partial failure", 5, 2, false, ""},
		{"total failure", 3, 3, true, "all 3 scan operations failed"},
		{"single total failure", 1, 1, true, "all 1 scan operations failed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scanStats{attempted: tt.attempted, failed: tt.failed}
			err := s.result()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errSubstr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestScanStatsRecord_SkipsAreNotAttempts(t *testing.T) {
	s := &scanStats{}

	s.record(true, nil)
	assert.Equal(t, 0, s.attempted)
	assert.Equal(t, 0, s.failed)
	assert.NoError(t, s.result())

	// One skipped image plus every real scan failing is a total outage.
	s.record(false, errors.New("pull failed"))
	s.record(true, nil)
	s.record(false, errors.New("analyze failed"))
	assert.Equal(t, 2, s.attempted)
	assert.Equal(t, 2, s.failed)
	require.Error(t, s.result())
	assert.Contains(t, s.result().Error(), "all 2 scan operations failed")

	s.record(false, nil)
	assert.NoError(t, s.result(), "a real success after failures is partial, not total")
}

func TestScanStatsRecord_SkipWithErrorCountsAsFailure(t *testing.T) {
	s := &scanStats{}
	s.record(true, errors.New("exists check failed"))
	assert.Equal(t, 1, s.attempted)
	assert.Equal(t, 1, s.failed)
	require.Error(t, s.result())
}

type stubRegistry struct {
	tags map[string][]string
	errs map[string]error
}

func (s *stubRegistry) GetTags(repo string) ([]string, error) {
	if err, ok := s.errs[repo]; ok {
		return nil, err
	}
	return s.tags[repo], nil
}

type stubAnalyzer struct {
	fail map[string]error
}

func (s *stubAnalyzer) Analyze(imageName string) (*domain.ImageAnalysis, error) {
	if err, ok := s.fail[imageName]; ok {
		return nil, err
	}
	return &domain.ImageAnalysis{
		Image: domain.ImageRecord{Name: imageName},
	}, nil
}

func setupPipelineDB(t *testing.T) *database.Repository {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	_, err = db.Exec("PRAGMA foreign_keys=ON")
	require.NoError(t, err)
	require.NoError(t, database.CreateTables(db))

	return database.NewRepository(db)
}

func testPipeline(repo *database.Repository, images []string, registry tagDiscoverer, analyzer imageAnalyzer) *Pipeline {
	return &Pipeline{
		config: domain.ScanConfig{MaxTagsPerRepo: 0},
		repoCfg: domain.RepositoryConfig{
			Defaults: domain.ConfigDefaults{Registry: "mcr.microsoft.com"},
			Repositories: []domain.RepositoryGroup{{
				Description: "test",
				Images:      images,
			}},
		},
		repo:     repo,
		analyzer: analyzer,
		registry: registry,
	}
}

func TestScanAll_AllScansFail(t *testing.T) {
	repo := setupPipelineDB(t)
	p := testPipeline(repo, []string{"azurelinux/base/python:3.12"},
		&stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"azurelinux/base/python:3.12": errors.New("pull failed")}},
	)

	err := p.ScanAll()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "all 1 scan operations failed")
}

func TestScanAll_PartialFailureSucceeds(t *testing.T) {
	repo := setupPipelineDB(t)
	p := testPipeline(repo, []string{"good:1.0", "bad:1.0"},
		&stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"bad:1.0": errors.New("analyze failed")}},
	)

	require.NoError(t, p.ScanAll())
}

func TestScanAll_SkipDoesNotMaskTotalFailure(t *testing.T) {
	repo := setupPipelineDB(t)
	require.NoError(t, repo.InsertImage(&domain.ImageRecord{Name: "already:1.0"}))

	p := testPipeline(repo, []string{"already:1.0", "new:1.0"},
		&stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"new:1.0": errors.New("pull failed")}},
	)

	err := p.ScanAll()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "all 1 scan operations failed")
}

func TestScanAll_AllSkippedIsNotAnError(t *testing.T) {
	repo := setupPipelineDB(t)
	require.NoError(t, repo.InsertImage(&domain.ImageRecord{Name: "already:1.0"}))

	p := testPipeline(repo, []string{"already:1.0"},
		&stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"already:1.0": errors.New("should not be called")}},
	)

	require.NoError(t, p.ScanAll())
}

func TestScanAll_TagDiscoveryFailure(t *testing.T) {
	repo := setupPipelineDB(t)
	p := testPipeline(repo, []string{"azurelinux/base/python"},
		&stubRegistry{errs: map[string]error{"azurelinux/base/python": errors.New("404")}},
		&stubAnalyzer{},
	)

	err := p.ScanAll()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "all 1 scan operations failed")
}

func TestScanAll_RepositoryTagsAllFail(t *testing.T) {
	repo := setupPipelineDB(t)
	reg := &stubRegistry{tags: map[string][]string{
		"azurelinux/base/python": {"3.12", "3.11"},
	}}
	az := &stubAnalyzer{fail: map[string]error{
		"mcr.microsoft.com/azurelinux/base/python:3.12": errors.New("fail 3.12"),
		"mcr.microsoft.com/azurelinux/base/python:3.11": errors.New("fail 3.11"),
	}}
	p := testPipeline(repo, []string{"azurelinux/base/python"}, reg, az)

	err := p.ScanAll()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "all 2 scan operations failed")
}

func TestScanAll_RepositoryPartialSuccess(t *testing.T) {
	repo := setupPipelineDB(t)
	reg := &stubRegistry{tags: map[string][]string{
		"azurelinux/base/python": {"3.12", "3.11"},
	}}
	az := &stubAnalyzer{fail: map[string]error{
		"mcr.microsoft.com/azurelinux/base/python:3.11": errors.New("fail 3.11"),
	}}
	p := testPipeline(repo, []string{"azurelinux/base/python"}, reg, az)

	require.NoError(t, p.ScanAll())
}

func TestScanAll_UpdateExistingRescansSkippedImage(t *testing.T) {
	repo := setupPipelineDB(t)
	require.NoError(t, repo.InsertImage(&domain.ImageRecord{Name: "already:1.0"}))

	called := 0
	az := &countingAnalyzer{fn: func(name string) (*domain.ImageAnalysis, error) {
		called++
		return &domain.ImageAnalysis{Image: domain.ImageRecord{Name: name}}, nil
	}}
	p := testPipeline(repo, []string{"already:1.0"}, &stubRegistry{}, az)
	p.config.UpdateExisting = true

	require.NoError(t, p.ScanAll())
	assert.Equal(t, 1, called, "existing image must be rescanned when --update-existing is set")
}

type countingAnalyzer struct {
	fn func(string) (*domain.ImageAnalysis, error)
}

func (c *countingAnalyzer) Analyze(imageName string) (*domain.ImageAnalysis, error) {
	return c.fn(imageName)
}

func TestScanAll_EmptyConfigIsNotAnError(t *testing.T) {
	repo := setupPipelineDB(t)
	p := testPipeline(repo, nil, &stubRegistry{}, &stubAnalyzer{})
	p.repoCfg.Repositories = nil

	require.NoError(t, p.ScanAll())
}

func TestScanImage_ReturnsAnalyzeError(t *testing.T) {
	repo := setupPipelineDB(t)
	p := testPipeline(repo, nil, &stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"img:1.0": errors.New("boom")}},
	)

	err := p.ScanImage("img:1.0")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "analyzing img:1.0")
}

func TestScanImage_SkipExisting(t *testing.T) {
	repo := setupPipelineDB(t)
	require.NoError(t, repo.InsertImage(&domain.ImageRecord{Name: "img:1.0"}))
	p := testPipeline(repo, nil, &stubRegistry{},
		&stubAnalyzer{fail: map[string]error{"img:1.0": errors.New("should not analyze")}},
	)

	require.NoError(t, p.ScanImage("img:1.0"))
}
