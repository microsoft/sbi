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

	"github.com/stretchr/testify/assert"
)

func TestHumanSize(t *testing.T) {
	tests := []struct {
		name     string
		bytes    int64
		expected string
	}{
		{"zero", 0, "-"},
		{"negative", -1, "-"},
		{"small MB", 85 * 1024 * 1024, "85.0 MB"},
		{"fractional MB", 85300000, "81.3 MB"},
		{"large MB", 500 * 1024 * 1024, "500.0 MB"},
		{"GB range", 2 * 1024 * 1024 * 1024, "2.00 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanSize(tt.bytes)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatDigest(t *testing.T) {
	tests := []struct {
		name     string
		digest   string
		expected string
	}{
		{"empty", "", ""},
		{"sha256 long", "sha256:abcdef123456789012345678", "sha256:abcdef123456"},
		{"sha256 short", "sha256:abc", "sha256:abc"},
		{"other format long", "md5:abcdef1234567890123456", "md5:abcdef123456789"},
		{"other format short", "abc", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDigest(tt.digest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatPinnedReference(t *testing.T) {
	tests := []struct {
		name     string
		imgName  string
		digest   string
		expected string
	}{
		{"name with tag", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot", "sha256:7be8b46abc123", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:7be8b46abc123"},
		{"name without tag", "mcr.microsoft.com/azurelinux/distroless/python", "sha256:7be8b46abc123", "mcr.microsoft.com/azurelinux/distroless/python@sha256:7be8b46abc123"},
		{"empty name", "", "sha256:7be8b46abc123", ""},
		{"empty digest", "mcr.microsoft.com/azurelinux/distroless/python:3.12", "", ""},
		{"all empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatPinnedReference(tt.imgName, tt.digest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatStableTag(t *testing.T) {
	tests := []struct {
		name     string
		imgName  string
		version  string
		expected string
	}{
		{"name with tag, major.minor.patch version", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot", "3.12.10", "mcr.microsoft.com/azurelinux/distroless/python:3.12"},
		{"name with tag, major.minor.patch-release version", "mcr.microsoft.com/azurelinux/distroless/python:3.12.10-1", "3.12.10-1", "mcr.microsoft.com/azurelinux/distroless/python:3.12"},
		{"name without tag, major.minor version", "mcr.microsoft.com/azurelinux/distroless/python", "3.12", "mcr.microsoft.com/azurelinux/distroless/python:3.12"},
		{"name without tag, major only version", "mcr.microsoft.com/azurelinux/distroless/python", "3", "mcr.microsoft.com/azurelinux/distroless/python:3"},
		{"empty version", "mcr.microsoft.com/azurelinux/distroless/python:3.12", "", ""},
		{"empty name", "", "3.12", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatStableTag(tt.imgName, tt.version)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStripTag(t *testing.T) {
	tests := []struct {
		name     string
		imgName  string
		expected string
	}{
		{"with tag", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot", "mcr.microsoft.com/azurelinux/distroless/python"},
		{"with port and tag", "localhost:5000/myimage:latest", "localhost:5000/myimage"},
		{"without tag", "mcr.microsoft.com/azurelinux/distroless/python", "mcr.microsoft.com/azurelinux/distroless/python"},
		{"with port no tag", "localhost:5000/myimage", "localhost:5000/myimage"},
		{"docker hub short", "python:3.12", "python"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripTag(tt.imgName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatDockerfileFrom(t *testing.T) {
	tests := []struct {
		name     string
		imgName  string
		digest   string
		expected string
	}{
		{"name with tag", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot", "sha256:7be8b46abc123", "FROM mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot@sha256:7be8b46abc123"},
		{"name without tag", "mcr.microsoft.com/azurelinux/distroless/python", "sha256:7be8b46abc123", "FROM mcr.microsoft.com/azurelinux/distroless/python@sha256:7be8b46abc123"},
		{"empty digest", "mcr.microsoft.com/azurelinux/distroless/python:3.12", "", ""},
		{"empty name", "", "sha256:7be8b46abc123", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDockerfileFrom(tt.imgName, tt.digest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDisplayOSName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"azurelinux", "Azure Linux"},
		{"ubuntu", "Ubuntu"},
		{"debian", "Debian"},
		{"alpine", "Alpine"},
		{"", "Other"},
		{"Other", "Other"},
		{"fedora", "Fedora"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, DisplayOSName(tt.input))
		})
	}
}
