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
	"testing"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicateByDigest_TwoSameDigest(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:3", Digest: "sha256:aaa", Version: "3.12.9"},
		{Name: "mcr.microsoft.com/repo:3.12", Digest: "sha256:aaa", Version: "3.12.9"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:3.12", result[0].Name, "most specific tag should win")
	assert.Equal(t, []string{":3"}, result[0].AlternateTags)
}

func TestDeduplicateByDigest_ThreeSameDigest(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:24", Digest: "sha256:bbb"},
		{Name: "mcr.microsoft.com/repo:24.14", Digest: "sha256:bbb"},
		{Name: "mcr.microsoft.com/repo:24.14.1", Digest: "sha256:bbb"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:24.14.1", result[0].Name, "most specific tag should win")
	assert.Len(t, result[0].AlternateTags, 2)
	assert.Contains(t, result[0].AlternateTags, ":24")
	assert.Contains(t, result[0].AlternateTags, ":24.14")
}

func TestDeduplicateByDigest_NoDuplicates(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:3.12", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/repo:3.11", Digest: "sha256:bbb"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 2)
	assert.Nil(t, result[0].AlternateTags)
	assert.Nil(t, result[1].AlternateTags)
}

func TestDeduplicateByDigest_EmptyDigestPreservesOrder(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "emptyA:1", Digest: ""},
		{Name: "real:2", Digest: "sha256:xxx"},
		{Name: "emptyB:3", Digest: ""},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 3)
	assert.Equal(t, "emptyA:1", result[0].Name)
	assert.Equal(t, "real:2", result[1].Name)
	assert.Equal(t, "emptyB:3", result[2].Name)
}

func TestDeduplicateByDigest_LatestVsNumericTag(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:latest", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/repo:3", Digest: "sha256:aaa"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:3", result[0].Name, "numeric tag should win over latest")
	assert.Equal(t, []string{":latest"}, result[0].AlternateTags)
}

func TestDeduplicateByDigest_PrimaryAppearsLater(t *testing.T) {
	// The more specific tag appears second; group should still use first-seen position.
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:3", Digest: "sha256:aaa", Version: "3.12.9"},
		{Name: "other:1.0", Digest: "sha256:bbb"},
		{Name: "mcr.microsoft.com/repo:3.12", Digest: "sha256:aaa", Version: "3.12.9"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 2)
	// The deduped group should appear at position 0 (first-seen order).
	assert.Equal(t, "mcr.microsoft.com/repo:3.12", result[0].Name)
	assert.Equal(t, "other:1.0", result[1].Name)
}

func TestDeduplicateByDigest_MixedDuplicatesAndUnique(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/python:3", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/python:3.12", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/node:24", Digest: "sha256:bbb"},
		{Name: "mcr.microsoft.com/go:1.21", Digest: "sha256:ccc"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 3)
	assert.Equal(t, "mcr.microsoft.com/python:3.12", result[0].Name)
	assert.Equal(t, []string{":3"}, result[0].AlternateTags)
	assert.Equal(t, "mcr.microsoft.com/node:24", result[1].Name)
	assert.Nil(t, result[1].AlternateTags)
	assert.Equal(t, "mcr.microsoft.com/go:1.21", result[2].Name)
	assert.Nil(t, result[2].AlternateTags)
}

func TestDeduplicateByDigest_SingleImage(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:3.12", Digest: "sha256:aaa"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Nil(t, result[0].AlternateTags)
}

func TestDeduplicateByDigest_EmptySlice(t *testing.T) {
	result := DeduplicateByDigest(nil)
	assert.Nil(t, result)

	result = DeduplicateByDigest([]domain.RecommendedImage{})
	assert.Empty(t, result)
}

func TestDeduplicateByDigest_NonrootSuffix(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:24-nonroot", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/repo:24.14-nonroot", Digest: "sha256:aaa"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:24.14-nonroot", result[0].Name, "more version segments should win")
	assert.Equal(t, []string{":24-nonroot"}, result[0].AlternateTags)
}

func TestDeduplicateByDigest_PreservesOrder(t *testing.T) {
	images := []domain.RecommendedImage{
		{Name: "first:1.0", Digest: "sha256:aaa"},
		{Name: "second:2.0", Digest: "sha256:bbb"},
		{Name: "third:3.0", Digest: "sha256:ccc"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 3)
	assert.Equal(t, "first:1.0", result[0].Name)
	assert.Equal(t, "second:2.0", result[1].Name)
	assert.Equal(t, "third:3.0", result[2].Name)
}

func TestTagSpecificity(t *testing.T) {
	tests := []struct {
		name string
		less string
		more string
	}{
		{"major vs major.minor", "mcr.microsoft.com/repo:3", "mcr.microsoft.com/repo:3.12"},
		{"major.minor vs major.minor.patch", "mcr.microsoft.com/repo:24.14", "mcr.microsoft.com/repo:24.14.1"},
		{"with suffix", "mcr.microsoft.com/repo:24-nonroot", "mcr.microsoft.com/repo:24.14-nonroot"},
		{"v-prefix vs older bare major", "mcr.microsoft.com/repo:9.0", "mcr.microsoft.com/repo:v10.0"},
		{"v-prefix vs latest", "mcr.microsoft.com/repo:latest", "mcr.microsoft.com/repo:v3.12"},
		{"v major.minor vs v major", "mcr.microsoft.com/repo:v3", "mcr.microsoft.com/repo:v3.12"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, tagSpecificity(tt.more), tagSpecificity(tt.less),
				"specificity(%s)=%d specificity(%s)=%d",
				tt.more, tagSpecificity(tt.more), tt.less, tagSpecificity(tt.less))
		})
	}
}

func TestDeduplicateByDigest_VPrefixedNumericPreferred(t *testing.T) {
	// Bare "9.0" must not outrank "v10.0" solely because the latter starts with 'v'.
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:9.0", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/repo:v10.0", Digest: "sha256:aaa"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:v10.0", result[0].Name)
	assert.Equal(t, []string{":9.0"}, result[0].AlternateTags)
}

func TestExtractTag(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"with tag", "mcr.microsoft.com/repo:3.12", "3.12"},
		{"with tag and suffix", "mcr.microsoft.com/repo:3.12-nonroot", "3.12-nonroot"},
		{"no tag", "mcr.microsoft.com/repo", ""},
		{"port in name", "localhost:5000/repo:latest", "latest"},
		{"with digest suffix", "mcr.microsoft.com/repo:3.12@sha256:abcdef", "3.12"},
		{"digest only no tag", "mcr.microsoft.com/repo@sha256:abcdef", ""},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, extractTag(tt.input))
		})
	}
}

func TestDeduplicateByDigest_DeterministicTieBreak(t *testing.T) {
	// When two tags have the same specificity score, the lexicographically
	// smaller name should be chosen as primary for determinism.
	// "zeta" and "beta" have the same length and structure, so scores are equal.
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/repo:zeta", Digest: "sha256:aaa"},
		{Name: "mcr.microsoft.com/repo:beta", Digest: "sha256:aaa"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 1)
	assert.Equal(t, "mcr.microsoft.com/repo:beta", result[0].Name, "lexicographically smaller name should win on tie")
	assert.Equal(t, []string{":zeta"}, result[0].AlternateTags)
}

func TestDeduplicateByDigest_CrossRepoSameDigestNotMerged(t *testing.T) {
	// Images from different repositories with the same digest should NOT be merged.
	images := []domain.RecommendedImage{
		{Name: "mcr.microsoft.com/azurelinux/base/python:3.12", Digest: "sha256:same"},
		{Name: "mcr.microsoft.com/azurelinux/base/nodejs:24", Digest: "sha256:same"},
	}

	result := DeduplicateByDigest(images)

	assert.Len(t, result, 2, "different repos with same digest should remain separate")
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", result[0].Name)
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/nodejs:24", result[1].Name)
	assert.Nil(t, result[0].AlternateTags)
	assert.Nil(t, result[1].AlternateTags)
}

func TestImageRepo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"with tag", "mcr.microsoft.com/repo:3.12", "mcr.microsoft.com/repo"},
		{"with tag and suffix", "mcr.microsoft.com/repo:3.12-nonroot", "mcr.microsoft.com/repo"},
		{"no tag", "mcr.microsoft.com/repo", "mcr.microsoft.com/repo"},
		{"with digest", "mcr.microsoft.com/repo:3.12@sha256:abc", "mcr.microsoft.com/repo"},
		{"digest only", "mcr.microsoft.com/repo@sha256:abc", "mcr.microsoft.com/repo"},
		{"nested path", "mcr.microsoft.com/azurelinux/base/python:3.12", "mcr.microsoft.com/azurelinux/base/python"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, imageRepo(tt.input))
		})
	}
}
