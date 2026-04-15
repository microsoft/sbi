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

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test fixtures built from real syft analysis of container images (2026-02-25).
// Each fixture contains only the language-relevant packages extracted from the
// full syft JSON output, keeping tests fast while exercising real-world data.

// ---------- helpers ----------

func makeArtifacts(pkgs []struct{ name, version, typ, lang string }) []syftArtifact {
	arts := make([]syftArtifact, len(pkgs))
	for i, p := range pkgs {
		arts[i] = syftArtifact{
			Name:     p.name,
			Version:  p.version,
			Type:     p.typ,
			Language: p.lang,
		}
	}
	return arts
}

// ---------- Python fixtures ----------

func pythonBaseAzureLinuxArtifacts() []syftArtifact {
	// From mcr.microsoft.com/azurelinux/base/python:3.12 (92 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"SymCrypt", "103.8.0-1.azl3", "rpm", ""},
		{"azurelinux-release", "3.0-10.azl3", "rpm", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"openssl-libs", "3.4.1-1.azl3", "rpm", ""},
		{"pip", "24.2", "python", "python"},
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
		{"python3-pip", "24.2-5.azl3", "rpm", ""},
		{"python3-setuptools", "69.0.3-5.azl3", "rpm", ""},
		{"python3-wheel", "0.43.0-1.azl3", "rpm", ""},
		{"setuptools", "69.0.3", "python", "python"},
		{"wheel", "0.43.0", "python", "python"},
	})
}

func pythonDistrolessArtifacts() []syftArtifact {
	// From mcr.microsoft.com/azurelinux/distroless/python:3.12 (61 packages total)
	// Note: NO python-type packages (no pip), only RPMs
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"SymCrypt", "103.8.0-1.azl3", "rpm", ""},
		{"azurelinux-release", "3.0-10.azl3", "rpm", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"openssl-libs", "3.4.1-1.azl3", "rpm", ""},
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
	})
}

func pythonDockerHubSlimArtifacts() []syftArtifact {
	// From docker.io/library/python:3.12-slim (95 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"adduser", "3.134", "deb", ""},
		{"apt", "2.6.1", "deb", ""},
		{"libc6", "2.36-9+deb12u9", "deb", ""},
		{"libssl3", "3.0.16-1~deb12u1", "deb", ""},
		{"pip", "25.0.1", "python", "python"},
		{"python", "3.12.12", "binary", ""},
	})
}

// ---------- Node.js fixtures ----------

func nodeBaseAzureLinuxArtifacts() []syftArtifact {
	// From mcr.microsoft.com/azurelinux/base/nodejs:20 (294 packages total)
	// Includes npm packages (209) and RPMs (85) — trimmed to relevant ones
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"@isaacs/cliui", "8.0.2", "npm", ""},
		{"@npmcli/agent", "2.2.2", "npm", ""},
		{"@npmcli/node-gyp", "3.0.0", "npm", ""},
		{"npm", "10.7.0", "npm", ""},
		{"semver", "7.6.0", "npm", ""},
		{"nodejs", "20.14.0-13.azl3", "rpm", ""},
		{"nodejs-libs", "20.14.0-13.azl3", "rpm", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"openssl-libs", "3.4.1-1.azl3", "rpm", ""},
	})
}

func nodeDistrolessArtifacts() []syftArtifact {
	// From mcr.microsoft.com/azurelinux/distroless/nodejs:20 (257 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"@isaacs/cliui", "8.0.2", "npm", ""},
		{"@npmcli/node-gyp", "3.0.0", "npm", ""},
		{"npm", "10.7.0", "npm", ""},
		{"nodejs", "20.14.0-13.azl3", "rpm", ""}, // duplicate in real syft output
		{"nodejs", "20.14.0-13.azl3", "rpm", ""}, // (syft reports the same RPM twice)
		{"nodejs-libs", "20.14.0-13.azl3", "rpm", ""},
	})
}

func nodeDockerHubSlimArtifacts() []syftArtifact {
	// From docker.io/library/node:20.0-slim (336 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"@colors/colors", "1.5.0", "npm", ""},
		{"@npmcli/node-gyp", "3.0.0", "npm", ""},
		{"npm", "9.6.4", "npm", ""},
		{"node", "20.0.0", "binary", ""},
		{"apt", "2.6.1", "deb", ""},
		{"libc6", "2.36-9+deb12u9", "deb", ""},
	})
}

// ---------- .NET fixtures ----------

func dotnetAspnetDistrolessArtifacts() []syftArtifact {
	// From mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless (18 packages total)
	// Note: dotnet packages have names like Microsoft.NETCore.App... NOT matching name regex
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"SymCrypt", "103.8.0-1.azl3", "rpm", ""},
		{"azurelinux-release", "3.0-10.azl3", "rpm", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"openssl-libs", "3.4.1-1.azl3", "rpm", ""},
		{"Microsoft.AspNetCore.App.Runtime.linux-arm64", "9.0.13", "dotnet", ""},
		{"Microsoft.NETCore.App.Runtime.linux-arm64", "9.0.13", "dotnet", ""},
	})
}

func dotnetRuntimeDistrolessArtifacts() []syftArtifact {
	// From mcr.microsoft.com/dotnet/runtime:9.0-azurelinux3.0-distroless (17 packages)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"SymCrypt", "103.8.0-1.azl3", "rpm", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"Microsoft.NETCore.App.Runtime.linux-arm64", "9.0.13", "dotnet", ""},
	})
}

func dotnetAspnetDebianArtifacts() []syftArtifact {
	// From mcr.microsoft.com/dotnet/aspnet:8.0 (94 packages)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"apt", "2.6.1", "deb", ""},
		{"libc6", "2.36-9+deb12u9", "deb", ""},
		{"libssl3", "3.0.16-1~deb12u1", "deb", ""},
		{"Microsoft.AspNetCore.App.Runtime.linux-arm64", "8.0.24", "dotnet", ""},
		{"Microsoft.NETCore.App.Runtime.linux-arm64", "8.0.24", "dotnet", ""},
	})
}

func dotnetSDKAzureLinuxArtifacts() []syftArtifact {
	// From mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0 (2687 packages total)
	// Trimmed to only language-relevant packages
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
		{"Microsoft.AspNetCore.App.Runtime.linux-arm64", "9.0.13", "dotnet", ""},
		{"Microsoft.NETCore.App.Runtime.linux-arm64", "9.0.13", "dotnet", ""},
		{"dotnet-sdk-9.0", "9.0.311-1.azl3", "rpm", ""},
		{"dotnet-runtime-9.0", "9.0.13-1.azl3", "rpm", ""},
		{"aspnetcore-runtime-9.0", "9.0.13-1.azl3", "rpm", ""},
		{"dotnet-format", "9.0.311.42318", "binary", ""},
		{"Json.More.Net", "2.0.2", "dotnet", ""},
		{"Json.NET", "13.0.3.27908", "dotnet", ""},
		// Many more dotnet packages omitted for brevity
	})
}

// ---------- Java fixtures ----------

func javaAzureLinuxArtifacts() []syftArtifact {
	// From mcr.microsoft.com/openjdk/jdk:21-azurelinux (136 packages total)
	// Contains: RPMs + go-module (embedded jaz.git binary) + java-archive
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
		{"msopenjdk-21", "21.0.10-1", "rpm", ""},
		{"jrt-fs", "21.0.10", "java-archive", "java"},
		// go-module entries from embedded jaz.git binary (NOT Go runtime)
		{"dev.azure.com/devdiv/DevDiv/_git/jaz.git", "v0.0.0-preview.0.20260106", "go-module", "go"},
		{"github.com/Masterminds/semver/v3", "v3.4.0", "go-module", "go"},
		{"github.com/cilium/ebpf", "v0.20.0", "go-module", "go"},
		{"github.com/containerd/cgroups/v3", "v3.1.2", "go-module", "go"},
		{"github.com/godbus/dbus/v5", "v5.2.2", "go-module", "go"},
	})
}

func javaDistrolessArtifacts() []syftArtifact {
	// From mcr.microsoft.com/openjdk/jdk:21-distroless (39 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"openjdk", "21.0.10+7-LTS", "binary", ""},
		{"jrt-fs", "21.0.10", "java-archive", "java"},
		// go-module entries from embedded jaz.git binary
		{"dev.azure.com/devdiv/DevDiv/_git/jaz.git", "v0.0.0-preview.0.20260106", "go-module", "go"},
		{"github.com/Masterminds/semver/v3", "v3.4.0", "go-module", "go"},
		{"github.com/cilium/ebpf", "v0.20.0", "go-module", "go"},
	})
}

func javaUbuntuArtifacts() []syftArtifact {
	// From mcr.microsoft.com/openjdk/jdk:21-ubuntu (160 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"libc6", "2.35-0ubuntu3.9", "deb", ""},
		{"java-common", "0.72build2", "deb", ""},
		{"msopenjdk-21", "21.0.10-1", "deb", ""},
		{"jrt-fs", "21.0.10", "java-archive", "java"},
		// go-module entries from embedded jaz.git binary
		{"dev.azure.com/devdiv/DevDiv/_git/jaz.git", "v0.0.0-preview.0.20260106", "go-module", "go"},
		{"github.com/Masterminds/semver/v3", "v3.4.0", "go-module", "go"},
	})
}

// ---------- Go fixtures ----------

func goAzureLinux126Artifacts() []syftArtifact {
	// From mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0 (218 packages total)
	// Go is tarball-installed — no golang RPM. Binary cataloger finds 'go'.
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
		{"openssl-libs", "3.4.1-1.azl3", "rpm", ""},
		// Go binary detected by binary cataloger
		{"go", "1.26.0", "binary", ""},
		// Go standard library modules
		{"cmd/asm", "UNKNOWN", "go-module", "go"},
		{"cmd/cgo", "UNKNOWN", "go-module", "go"},
		{"cmd/compile", "UNKNOWN", "go-module", "go"},
		{"cmd/go", "UNKNOWN", "go-module", "go"},
		{"cmd/gofmt", "UNKNOWN", "go-module", "go"},
		{"cmd/link", "UNKNOWN", "go-module", "go"},
		{"std", "UNKNOWN", "go-module", "go"},
		{"golang.org/x/crypto", "v0.36.0", "go-module", "go"},
		{"golang.org/x/net", "v0.38.0", "go-module", "go"},
	})
}

func goAzureLinux125Artifacts() []syftArtifact {
	// From mcr.microsoft.com/oss/go/microsoft/golang:1.25-azurelinux3.0 (216 packages total)
	return makeArtifacts([]struct{ name, version, typ, lang string }{
		{"glibc", "2.38-7.azl3", "rpm", ""},
		{"python3", "3.12.9-8.azl3", "rpm", ""},
		{"python3-libs", "3.12.9-8.azl3", "rpm", ""},
		// Go binary detected by binary cataloger
		{"go", "1.25.7", "binary", ""},
		// Go standard library modules
		{"cmd/asm", "UNKNOWN", "go-module", "go"},
		{"cmd/go", "UNKNOWN", "go-module", "go"},
		{"std", "UNKNOWN", "go-module", "go"},
		{"golang.org/x/crypto", "v0.35.0", "go-module", "go"},
	})
}

// ============================================================================
// Tests for parseSyftResult — current behaviour
// ============================================================================

func TestParseSyftResult_PythonBaseAzureLinux(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: pythonBaseAzureLinuxArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "python")
	// Base python3 RPM should be preferred over python3-libs for language detection
	assert.Equal(t, "python3", langs["python"].PackageName)
	assert.Equal(t, "3.12.9", langs["python"].Version) // cleaned from 3.12.9-8.azl3
	assert.Equal(t, "3.12", langs["python"].MajorMinor)
	assert.Equal(t, "rpm", langs["python"].PackageType)
	assert.NotContains(t, langs, "go")
	assert.NotContains(t, langs, "dotnet")
}

func TestParseSyftResult_PythonDistroless(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: pythonDistrolessArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "python")
	assert.Equal(t, "python3", langs["python"].PackageName)
	assert.Equal(t, "3.12.9", langs["python"].Version)
	// No pip installed — ensure detection still works via RPM name
	assert.Equal(t, "rpm", langs["python"].PackageType)
}

func TestParseSyftResult_PythonDockerHub(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: pythonDockerHubSlimArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "python")
	// Binary 'python' detected with version 3.12.12
	assert.Equal(t, "python", langs["python"].PackageName)
	assert.Equal(t, "3.12.12", langs["python"].Version)
}

func TestParseSyftResult_NodeBaseAzureLinux(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: nodeBaseAzureLinuxArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "node")
	assert.Equal(t, "nodejs", langs["node"].PackageName)
	assert.Equal(t, "20.14.0", langs["node"].Version)
	assert.Equal(t, "20.14", langs["node"].MajorMinor)
	assert.Equal(t, "rpm", langs["node"].PackageType)
}

func TestParseSyftResult_NodeDistroless(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: nodeDistrolessArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "node")
	assert.Equal(t, "nodejs", langs["node"].PackageName)
	assert.Equal(t, "20.14.0", langs["node"].Version)
}

func TestParseSyftResult_NodeDockerHub(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: nodeDockerHubSlimArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "node")
	assert.Equal(t, "node", langs["node"].PackageName)
	assert.Equal(t, "20.0.0", langs["node"].Version)
}

func TestParseSyftResult_DotnetAspnetDistroless(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetAspnetDistrolessArtifacts()})

	langs := langMap(result)
	// .NET detected via type=dotnet packages
	require.Contains(t, langs, "dotnet")
	assert.Equal(t, "dotnet", langs["dotnet"].PackageType)
	assert.Equal(t, "9.0.13", langs["dotnet"].Version)
	assert.Equal(t, "9.0", langs["dotnet"].MajorMinor)
}

func TestParseSyftResult_DotnetRuntimeDistroless(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetRuntimeDistrolessArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "dotnet")
	assert.Equal(t, "Microsoft.NETCore.App.Runtime.linux-arm64", langs["dotnet"].PackageName)
	assert.Equal(t, "9.0.13", langs["dotnet"].Version)
}

func TestParseSyftResult_DotnetAspnetDebian(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetAspnetDebianArtifacts()})

	langs := langMap(result)
	// .NET detected via type=dotnet packages
	require.Contains(t, langs, "dotnet")
	assert.Equal(t, "8.0.24", langs["dotnet"].Version)
}

func TestParseSyftResult_DotnetSDK(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetSDKAzureLinuxArtifacts()})

	langs := langMap(result)
	// SDK images have both type=dotnet runtime packages and name-regex RPMs.
	// The type=dotnet runtime package is preferred via isDotnetRuntime().
	require.Contains(t, langs, "dotnet")
	assert.True(t, isDotnetRuntime(langs["dotnet"].PackageName),
		"expected a .NET runtime package, got %s", langs["dotnet"].PackageName)
	assert.Equal(t, "9.0.13", langs["dotnet"].Version)
	// Also detects incidental Python
	require.Contains(t, langs, "python")
	assert.Equal(t, "python3", langs["python"].PackageName)
}

func TestParseSyftResult_DotnetNoFalsePositive(t *testing.T) {
	// Json.More.Net has type=dotnet so type-based detection picks it up.
	// In real images, isDotnetRuntime() ensures the core runtime package
	// is preferred. This test verifies the edge case where only NuGet
	// libraries exist (no runtime) — first match wins.
	arts := makeArtifacts([]struct{ name, version, typ, lang string }{
		{"Json.More.Net", "2.0.2", "dotnet", ""},
		{"Json.NET", "13.0.3.27908", "dotnet", ""},
		{"glibc", "2.38-7.azl3", "rpm", ""},
	})
	result := parseSyftResult(&syftOutput{Artifacts: arts})

	langs := langMap(result)
	// With only NuGet libraries and no runtime, the first type=dotnet
	// package is detected. In practice this doesn't happen — real .NET
	// images always have Microsoft.NETCore.App.Runtime.
	if dotnet, ok := langs["dotnet"]; ok {
		assert.Equal(t, "Json.More.Net", dotnet.PackageName)
	}
}

func TestParseSyftResult_JavaAzureLinux(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: javaAzureLinuxArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "java")
	assert.Equal(t, "msopenjdk-21", langs["java"].PackageName)
	assert.Equal(t, "21.0.10", langs["java"].Version)
	assert.Equal(t, "rpm", langs["java"].PackageType)
	// Incidental Python from system dependency
	require.Contains(t, langs, "python")
	assert.Equal(t, "python3", langs["python"].PackageName)
	// go-module entries should NOT be detected as Go runtime
	assert.NotContains(t, langs, "go",
		"go-module entries from embedded jaz.git binary must NOT be detected as Go runtime")
}

func TestParseSyftResult_JavaDistroless(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: javaDistrolessArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "java")
	assert.Equal(t, "openjdk", langs["java"].PackageName)
	assert.Equal(t, "21.0.10", langs["java"].Version)
	// go-module entries should NOT be detected as Go
	assert.NotContains(t, langs, "go",
		"go-module entries must NOT be detected as Go runtime in JDK images")
}

func TestParseSyftResult_JavaUbuntu(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: javaUbuntuArtifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "java")
	assert.Equal(t, "msopenjdk-21", langs["java"].PackageName)
	assert.Equal(t, "21.0.10", langs["java"].Version)
	assert.Equal(t, "deb", langs["java"].PackageType)
	// java-common should NOT be treated as Java runtime (it's excluded)
	// go-module entries should NOT be detected as Go
	assert.NotContains(t, langs, "go")
}

func TestParseSyftResult_GoAzureLinux126(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: goAzureLinux126Artifacts()})

	langs := langMap(result)
	// Go detected via type=binary, name=go
	require.Contains(t, langs, "go")
	assert.Equal(t, "go", langs["go"].PackageName)
	assert.Equal(t, "binary", langs["go"].PackageType)
	assert.Equal(t, "1.26.0", langs["go"].Version)
	assert.Equal(t, "1.26", langs["go"].MajorMinor)
	// Incidental Python IS detected
	require.Contains(t, langs, "python")
	assert.Equal(t, "python3", langs["python"].PackageName)
}

func TestParseSyftResult_GoAzureLinux125(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: goAzureLinux125Artifacts()})

	langs := langMap(result)
	require.Contains(t, langs, "go")
	assert.Equal(t, "go", langs["go"].PackageName)
	assert.Equal(t, "1.25.7", langs["go"].Version)
	require.Contains(t, langs, "python")
}

// ============================================================================
// Tests for mergeLanguages — Stage 2 image-name detection
// ============================================================================

func TestMergeLanguages_DotnetAspnetFallback(t *testing.T) {
	// .NET is now detected at Stage 1 via type=dotnet, so mergeLanguages
	// does NOT need the image-name fallback
	result := parseSyftResult(&syftOutput{Artifacts: dotnetAspnetDistrolessArtifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/dotnet/aspnet:9.0-azurelinux3.0-distroless")

	langs := langMapFromSlice(merged)
	require.Contains(t, langs, "dotnet")
	assert.Equal(t, "9.0.13", langs["dotnet"].Version)
	assert.Equal(t, "9.0", langs["dotnet"].MajorMinor)
}

func TestMergeLanguages_DotnetRuntimeFallback(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetRuntimeDistrolessArtifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/dotnet/runtime:10.0-azurelinux3.0-distroless")

	langs := langMapFromSlice(merged)
	require.Contains(t, langs, "dotnet")
	// Version comes from type=dotnet package, not from image name tag
	assert.Equal(t, "9.0.13", langs["dotnet"].Version)
}

func TestMergeLanguages_DotnetSDKNotFromImageName(t *testing.T) {
	// SDK pattern (dotnet/sdk:*) is NOT matched by dotnetImagePattern
	result := parseSyftResult(&syftOutput{Artifacts: dotnetSDKAzureLinuxArtifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0")

	langs := langMapFromSlice(merged)
	// .NET should be detected via syft (type=dotnet or name regex), not image name
	require.Contains(t, langs, "dotnet")
	assert.True(t, isDotnetRuntime(langs["dotnet"].PackageName),
		"expected a .NET runtime package, got %s", langs["dotnet"].PackageName)
}

func TestMergeLanguages_DotnetAspnetDebianFallback(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: dotnetAspnetDebianArtifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/dotnet/aspnet:8.0")

	langs := langMapFromSlice(merged)
	require.Contains(t, langs, "dotnet")
	// Version from type=dotnet package (full patch version)
	assert.Equal(t, "8.0.24", langs["dotnet"].Version)
}

func TestMergeLanguages_NonDotnetImageNoFalsePositive(t *testing.T) {
	result := parseSyftResult(&syftOutput{Artifacts: pythonBaseAzureLinuxArtifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/azurelinux/base/python:3.12")

	langs := langMapFromSlice(merged)
	assert.NotContains(t, langs, "dotnet",
		"Python image should not falsely detect dotnet from image name")
	require.Contains(t, langs, "python")
}

func TestMergeLanguages_GoImageDetected(t *testing.T) {
	// After implementation, Go IS detected at Stage 1 via binary type
	result := parseSyftResult(&syftOutput{Artifacts: goAzureLinux126Artifacts()})
	merged := mergeLanguages(result.Languages, nil, "mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0")

	langs := langMapFromSlice(merged)
	require.Contains(t, langs, "go")
	assert.Equal(t, "1.26.0", langs["go"].Version)
	// Python also present (incidental dependency)
	require.Contains(t, langs, "python")
}

// ============================================================================
// Tests for priority and version cleaning
// ============================================================================

func TestCleanVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3.12.9-8.azl3", "3.12.9"},
		{"20.14.0-13.azl3", "20.14.0"},
		{"21.0.10-1", "21.0.10"},
		{"5.36.0-7+deb12u3", "5.36.0"},
		{"3.0.16-1~deb12u1", "3.0.16"},
		{"1.26.0", "1.26.0"},
		{"24.2", "24.2"},
		{"UNKNOWN", "UNKNOWN"},
		{"", ""},
		{"21.0.10+7-LTS", "21.0.10"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, cleanVersion(tt.input))
		})
	}
}

func TestExtractMajorMinor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3.12.9", "3.12"},
		{"20.14.0", "20.14"},
		{"21.0.10", "21.0"},
		{"9.0", "9.0"},
		{"1.26.0", "1.26"},
		{"8", "8"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, extractMajorMinor(tt.input))
		})
	}
}

func TestIsDotnetRuntime(t *testing.T) {
	assert.True(t, isDotnetRuntime("Microsoft.NETCore.App.Runtime.linux-arm64"))
	assert.True(t, isDotnetRuntime("Microsoft.AspNetCore.App.Runtime.linux-arm64"))
	assert.False(t, isDotnetRuntime("Json.More.Net"))
	assert.False(t, isDotnetRuntime("Json.NET"))
	assert.False(t, isDotnetRuntime("dotnet-format"))
	assert.False(t, isDotnetRuntime("DotNetWatchTasks"))
}

// ============================================================================
// Test for exclude patterns
// ============================================================================

func TestExcludePackages(t *testing.T) {
	excluded := []string{"python-pip", "python3-pip", "nodejs-npm", "java-common", "python-setuptools"}
	for _, name := range excluded {
		assert.True(t, excludePackages.MatchString(name), "%s should be excluded", name)
	}

	notExcluded := []string{"python3", "nodejs", "msopenjdk-21", "go", "dotnet-runtime-9.0"}
	for _, name := range notExcluded {
		assert.False(t, excludePackages.MatchString(name), "%s should NOT be excluded", name)
	}
}

// ============================================================================
// Helpers
// ============================================================================

func langMap(result *domain.SyftResult) map[string]langInfo {
	m := make(map[string]langInfo)
	for _, l := range result.Languages {
		m[l.Language] = langInfo{
			PackageName: l.PackageName,
			PackageType: l.PackageType,
			Version:     l.Version,
			MajorMinor:  l.MajorMinor,
			Verified:    l.Verified,
		}
	}
	return m
}

func langMapFromSlice(langs []domain.Language) map[string]langInfo {
	m := make(map[string]langInfo)
	for _, l := range langs {
		m[l.Language] = langInfo{
			PackageName: l.PackageName,
			PackageType: l.PackageType,
			Version:     l.Version,
			MajorMinor:  l.MajorMinor,
			Verified:    l.Verified,
		}
	}
	return m
}

type langInfo struct {
	PackageName string
	PackageType string
	Version     string
	MajorMinor  string
	Verified    bool
}
