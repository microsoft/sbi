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

// Test fixtures built from real docker inspect output of container images (2026-03-10).

// ============================================================================
// parseDockerSize tests
// ============================================================================

func TestParseDockerSize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		// Real values from docker images output
		{name: "139MB python image", input: "139MB", expected: 139 * 1024 * 1024},
		{name: "225MB aspnet image", input: "225MB", expected: 225 * 1024 * 1024},
		{name: "350MB jdk image", input: "350MB", expected: 350 * 1024 * 1024},
		{name: "106MB nodejs image", input: "106MB", expected: 106 * 1024 * 1024},

		// Fractional MB
		{name: "fractional MB", input: "85.3MB", expected: 89443532},

		// GB range
		{name: "1.2GB", input: "1.2GB", expected: 1288490188},

		// KB variants
		{name: "512KB", input: "512KB", expected: 512 * 1024},
		{name: "512kB", input: "512kB", expected: 512 * 1024},

		// Bare bytes
		{name: "bare bytes", input: "1024B", expected: 1024},

		// Edge cases
		{name: "empty string", input: "", expected: 0},
		{name: "whitespace", input: "  ", expected: 0},
		{name: "invalid", input: "notanumber", expected: 0},
		{name: "whitespace around", input: "  225MB  ", expected: 225 * 1024 * 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDockerSize(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

// ============================================================================
// parseInspectOutput tests — real data from docker inspect
// ============================================================================

func TestParseInspectOutput_AzureLinuxPython(t *testing.T) {
	// From: mcr.microsoft.com/azurelinux/base/python:3.12
	// Size: 138819547, Layers: 3
	input := []byte(`[{
		"Size": 138819547,
		"Created": "2026-02-04T12:58:30.123456789Z",
		"RootFS": {
			"Type": "layers",
			"Layers": [
				"sha256:aaa",
				"sha256:bbb",
				"sha256:ccc"
			]
		},
		"RepoDigests": [
			"mcr.microsoft.com/azurelinux/base/python@sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd"
		]
	}]`)

	result, err := parseInspectOutput(input, "mcr.microsoft.com/azurelinux/base/python:3.12")
	require.NoError(t, err)
	assert.Equal(t, int64(138819547), result.SizeBytes)
	assert.Equal(t, "2026-02-04T12:58:30.123456789Z", result.Created)
	assert.Equal(t, 3, result.Layers)
	assert.Equal(t, "sha256:7be8b46a4dfad5acc5851ecf76ea800cd2818eaec40832baf5787849645461fd", result.Digest)
}

func TestParseInspectOutput_DotnetAspnet(t *testing.T) {
	// From: mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0
	// Size: 224892468, Layers: 6
	input := []byte(`[{
		"Size": 224892468,
		"Created": "2026-02-10T00:28:08.987654321Z",
		"RootFS": {
			"Type": "layers",
			"Layers": [
				"sha256:l1", "sha256:l2", "sha256:l3",
				"sha256:l4", "sha256:l5", "sha256:l6"
			]
		},
		"RepoDigests": [
			"mcr.microsoft.com/dotnet/aspnet@sha256:7cbe033d86dae06d5501ef3a5053a3d545f34b81f907b42bd894478313dec46c"
		]
	}]`)

	result, err := parseInspectOutput(input, "mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0")
	require.NoError(t, err)
	assert.Equal(t, int64(224892468), result.SizeBytes)
	assert.Equal(t, 6, result.Layers)
	assert.Equal(t, "sha256:7cbe033d86dae06d5501ef3a5053a3d545f34b81f907b42bd894478313dec46c", result.Digest)
}

func TestParseInspectOutput_JDKDistroless(t *testing.T) {
	// From: mcr.microsoft.com/openjdk/jdk:21-distroless
	// Size: 350387685, Layers: 4
	input := []byte(`[{
		"Size": 350387685,
		"Created": "2026-01-14T09:16:30.111111111Z",
		"RootFS": {
			"Type": "layers",
			"Layers": [
				"sha256:a1", "sha256:a2", "sha256:a3", "sha256:a4"
			]
		},
		"RepoDigests": [
			"mcr.microsoft.com/openjdk/jdk@sha256:1102084595d0771ad88d57d0696609df3658ac9b6c2e2a779287d5466d92a66c"
		]
	}]`)

	result, err := parseInspectOutput(input, "mcr.microsoft.com/openjdk/jdk:21-distroless")
	require.NoError(t, err)
	assert.Equal(t, int64(350387685), result.SizeBytes)
	assert.Equal(t, 4, result.Layers)
	assert.Equal(t, "sha256:1102084595d0771ad88d57d0696609df3658ac9b6c2e2a779287d5466d92a66c", result.Digest)
}

func TestParseInspectOutput_NodejsDistroless(t *testing.T) {
	// From: mcr.microsoft.com/azurelinux/distroless/nodejs:20
	// Size: 106088720, Layers: 2
	input := []byte(`[{
		"Size": 106088720,
		"Created": "2026-02-04T12:26:33.222222222Z",
		"RootFS": {
			"Type": "layers",
			"Layers": [
				"sha256:x1", "sha256:x2"
			]
		},
		"RepoDigests": [
			"mcr.microsoft.com/azurelinux/distroless/nodejs@sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356"
		]
	}]`)

	result, err := parseInspectOutput(input, "mcr.microsoft.com/azurelinux/distroless/nodejs:20")
	require.NoError(t, err)
	assert.Equal(t, int64(106088720), result.SizeBytes)
	assert.Equal(t, 2, result.Layers)
	assert.Equal(t, "sha256:6643a866773721bea665e29606618a55d28d42fdefb83e38b2e462fbb763a356", result.Digest)
}

// ============================================================================
// parseInspectOutput edge cases
// ============================================================================

func TestParseInspectOutput_EmptyArray(t *testing.T) {
	_, err := parseInspectOutput([]byte(`[]`), "test-image")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty inspect result")
}

func TestParseInspectOutput_InvalidJSON(t *testing.T) {
	_, err := parseInspectOutput([]byte(`not json`), "test-image")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parsing inspect output")
}

func TestParseInspectOutput_MissingFields(t *testing.T) {
	// Minimal valid JSON with no expected fields
	input := []byte(`[{"Id": "sha256:abc123"}]`)

	result, err := parseInspectOutput(input, "test-image")
	require.NoError(t, err)
	assert.Equal(t, int64(0), result.SizeBytes)
	assert.Equal(t, "", result.Created)
	assert.Equal(t, 0, result.Layers)
	assert.Equal(t, "", result.Digest)
}

func TestParseInspectOutput_NoRepoDigests(t *testing.T) {
	input := []byte(`[{
		"Size": 100000,
		"RepoDigests": []
	}]`)

	result, err := parseInspectOutput(input, "test-image")
	require.NoError(t, err)
	assert.Equal(t, int64(100000), result.SizeBytes)
	assert.Equal(t, "", result.Digest)
}

func TestParseInspectOutput_DigestWithoutAt(t *testing.T) {
	input := []byte(`[{
		"RepoDigests": ["no-at-sign-here"]
	}]`)

	result, err := parseInspectOutput(input, "test-image")
	require.NoError(t, err)
	assert.Equal(t, "", result.Digest)
}

func TestParseInspectOutput_NoLayers(t *testing.T) {
	input := []byte(`[{
		"RootFS": {"Type": "layers"}
	}]`)

	result, err := parseInspectOutput(input, "test-image")
	require.NoError(t, err)
	assert.Equal(t, 0, result.Layers)
}
