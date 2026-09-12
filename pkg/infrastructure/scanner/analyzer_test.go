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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterTags(t *testing.T) {
	tags := []string{"3.12", "3.11", "3.12-rc1", "3.11-beta", "3.10-alpha", "latest", "3.12-preview",
		"3.0.20250206", "3.12.9", "10.0.19041", "3.12.9-8-azl3.0.20260204-arm64", "3.12.9-8-azl3.0.20260204-amd64",
		"3.12.9-8-debug-nonroot"}
	filtered := FilterTags(tags, DefaultTagFilter())

	assert.Contains(t, filtered, "3.12")
	assert.Contains(t, filtered, "3.11")
	assert.Contains(t, filtered, "3.12.9")
	assert.Contains(t, filtered, "10.0.19041")
	assert.NotContains(t, filtered, "latest")
	assert.NotContains(t, filtered, "3.12-rc1")
	assert.NotContains(t, filtered, "3.11-beta")
	assert.NotContains(t, filtered, "3.10-alpha")
	assert.NotContains(t, filtered, "3.12-preview")
	assert.NotContains(t, filtered, "3.0.20250206")
	// Arch-specific, dated, and build-specific tags should be excluded
	assert.NotContains(t, filtered, "3.12.9-8-azl3.0.20260204-arm64")
	assert.NotContains(t, filtered, "3.12.9-8-azl3.0.20260204-amd64")
	assert.NotContains(t, filtered, "3.12.9-8-debug-nonroot")
}

func TestFilterTags_VersionAwareNewestFirst(t *testing.T) {
	// Reverse string sort would put 9.0 before 10.0; version-aware sort must not.
	tags := []string{"1.26", "1.25", "10.0", "9.0", "8.0", "2.0"}
	filtered := FilterTags(tags, DefaultTagFilter())

	require.Equal(t, []string{"10.0", "9.0", "8.0", "2.0", "1.26", "1.25"}, filtered)

	// With max-tags, the newest majors must be selected (not 9.x/8.x only).
	assert.Equal(t, []string{"10.0", "9.0", "8.0"}, LimitTags(filtered, 3))
}

func TestFilterTags_VPrefixedVersionAwareNewestFirst(t *testing.T) {
	// Reverse string sort ranks v9.0 above v10.0; version-aware sort must not.
	// These tags pass RequireDigit and must keep numeric major ordering under --max-tags.
	tags := []string{"v1.26", "v1.25", "v10.0", "v9.0", "v8.0", "v2.0"}
	filtered := FilterTags(tags, DefaultTagFilter())

	require.Equal(t, []string{"v10.0", "v9.0", "v8.0", "v2.0", "v1.26", "v1.25"}, filtered)
	assert.Equal(t, []string{"v10.0", "v9.0", "v8.0"}, LimitTags(filtered, 3))
}

func TestSortTagsNewestFirst(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "major versions not lexical",
			input: []string{"9.0", "10.0", "8.0"},
			want:  []string{"10.0", "9.0", "8.0"},
		},
		{
			name:  "v-prefixed majors not lexical",
			input: []string{"v9.0", "v10.0", "v8.0"},
			want:  []string{"v10.0", "v9.0", "v8.0"},
		},
		{
			name:  "mixed bare and v-prefixed majors",
			input: []string{"v9.0", "10.0", "v8.0"},
			want:  []string{"10.0", "v9.0", "v8.0"},
		},
		{
			name:  "patch vs minor",
			input: []string{"3.12", "3.12.9", "3.11"},
			want:  []string{"3.12.9", "3.12", "3.11"},
		},
		{
			name:  "suffix after version",
			input: []string{"8.0-noble", "10.0-azurelinux3.0", "9.0"},
			want:  []string{"10.0-azurelinux3.0", "9.0", "8.0-noble"},
		},
		{
			name:  "v-prefixed with suffix",
			input: []string{"v8.0-noble", "v10.0-azurelinux3.0", "v9.0"},
			want:  []string{"v10.0-azurelinux3.0", "v9.0", "v8.0-noble"},
		},
		{
			name:  "single segment",
			input: []string{"21", "25", "17"},
			want:  []string{"25", "21", "17"},
		},
		{
			name:  "equal version different suffix",
			input: []string{"8.0-noble", "8.0-azurelinux3.0"},
			// Ascending rest: "-azurelinux3.0" < "-noble"; newest-first reverses that.
			want: []string{"8.0-noble", "8.0-azurelinux3.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append([]string(nil), tt.input...)
			sortTagsNewestFirst(got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVersionPrefix(t *testing.T) {
	tests := []struct {
		tag      string
		wantSegs []int
		wantRest string
	}{
		{"10.0", []int{10, 0}, ""},
		{"10.0-noble", []int{10, 0}, "-noble"},
		{"v10.0", []int{10, 0}, ""},
		{"V10.0", []int{10, 0}, ""},
		{"v10.0-noble", []int{10, 0}, "-noble"},
		{"v9.0", []int{9, 0}, ""},
		{"3.12.9", []int{3, 12, 9}, ""},
		{"21-azurelinux", []int{21}, "-azurelinux"},
		{"latest", nil, "latest"},
		{"version", nil, "version"}, // leading 'v' but not a version prefix
		{"v", nil, "v"},
		{"", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			segs, rest := versionPrefix(tt.tag)
			assert.Equal(t, tt.wantSegs, segs)
			assert.Equal(t, tt.wantRest, rest)
		})
	}
}

func TestLimitTags(t *testing.T) {
	tags := []string{"a", "b", "c", "d", "e"}

	assert.Len(t, LimitTags(tags, 3), 3)
	assert.Len(t, LimitTags(tags, 0), 5)
	assert.Len(t, LimitTags(tags, 10), 5)
}

func TestParseImagePatterns(t *testing.T) {
	patterns := []string{
		"azurelinux/base/python",
		"docker.io/library/python:3.12-slim",
		"# comment",
		"",
		"mcr.microsoft.com/dotnet/aspnet:8.0",
		"azurelinux/distroless/node",
	}

	repos, singles := ParseImagePatterns(patterns)

	assert.Equal(t, []string{"azurelinux/base/python", "azurelinux/distroless/node"}, repos)
	assert.Equal(t, []string{"docker.io/library/python:3.12-slim", "mcr.microsoft.com/dotnet/aspnet:8.0"}, singles)
}

func TestExtractRegistryAndRepo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		registry string
		repo     string
		tag      string
	}{
		{"MCR full", "mcr.microsoft.com/azurelinux/base/python:3.12", "mcr.microsoft.com", "azurelinux/base/python", "3.12"},
		{"Docker Hub", "docker.io/library/python:3.12-slim", "docker.io", "library/python", "3.12-slim"},
		{"Short MCR", "azurelinux/base/python:3.12", "mcr.microsoft.com", "azurelinux/base/python", "3.12"},
		{"No tag", "mcr.microsoft.com/azurelinux/base/python", "mcr.microsoft.com", "azurelinux/base/python", ""},
		{"Tag with digest", "mcr.microsoft.com/dotnet/aspnet:8.0@sha256:abcdef", "mcr.microsoft.com", "dotnet/aspnet", "8.0"},
		{"Digest only", "mcr.microsoft.com/dotnet/aspnet@sha256:abcdef", "mcr.microsoft.com", "dotnet/aspnet", ""},
		{"Registry with port", "localhost:5000/myrepo:1.0", "localhost:5000", "myrepo", "1.0"},
		{"Localhost no port", "localhost/myrepo:1.0", "localhost", "myrepo", "1.0"},
		{"Nested path with digest", "mcr.microsoft.com/azurelinux/base/python:3.12-nonroot@sha256:deadbeef", "mcr.microsoft.com", "azurelinux/base/python", "3.12-nonroot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg, repo, tag := ExtractRegistryAndRepo(tt.input)
			assert.Equal(t, tt.registry, reg)
			assert.Equal(t, tt.repo, repo)
			assert.Equal(t, tt.tag, tag)
		})
	}
}

func TestBuildFullImageName(t *testing.T) {
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", BuildFullImageName("mcr.microsoft.com", "azurelinux/base/python", "3.12"))
	assert.Equal(t, "docker.io/library/python:3.12", BuildFullImageName("mcr.microsoft.com", "docker.io/library/python", "3.12"))
}
