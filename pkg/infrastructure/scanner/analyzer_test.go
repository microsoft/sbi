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
