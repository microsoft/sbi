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

func TestInferPrimaryLanguage(t *testing.T) {
	tests := []struct {
		name     string
		image    string
		expected string
	}{
		// Python images
		{"azure linux base python", "mcr.microsoft.com/azurelinux/base/python:3.12", "python"},
		{"azure linux distroless python", "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot", "python"},

		// Node.js images
		{"azure linux base nodejs", "mcr.microsoft.com/azurelinux/base/nodejs:24", "node"},
		{"azure linux distroless nodejs", "mcr.microsoft.com/azurelinux/distroless/nodejs:24.14", "node"},

		// Java images
		{"openjdk jdk azurelinux", "mcr.microsoft.com/openjdk/jdk:21-azurelinux", "java"},
		{"openjdk jdk distroless", "mcr.microsoft.com/openjdk/jdk:25-distroless", "java"},
		{"openjdk jdk ubuntu", "mcr.microsoft.com/openjdk/jdk:21-ubuntu", "java"},

		// .NET images
		{"dotnet aspnet azurelinux", "mcr.microsoft.com/dotnet/aspnet:8.0-azurelinux3.0-distroless", "dotnet"},
		{"dotnet runtime noble", "mcr.microsoft.com/dotnet/runtime:10.0-noble", "dotnet"},
		{"dotnet sdk", "mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0", "dotnet"},
		{"dotnet aspnet debian", "mcr.microsoft.com/dotnet/aspnet:8.0", "dotnet"},

		// Go images
		{"golang azurelinux", "mcr.microsoft.com/oss/go/microsoft/golang:1.26-azurelinux3.0", "go"},

		// No match — safe fallback
		{"base core image", "mcr.microsoft.com/azurelinux/base/core:3.0", ""},
		{"distroless base", "mcr.microsoft.com/azurelinux/distroless/base:3.0", ""},
		{"distroless minimal", "mcr.microsoft.com/azurelinux/distroless/minimal:3.0", ""},
		{"custom image", "my-custom-image:latest", ""},
		{"docker hub short name", "ubuntu:22.04", ""},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, InferPrimaryLanguage(tt.image))
		})
	}
}

func TestFilterIncidentalImages(t *testing.T) {
	jdkImage := domain.RecommendedImage{
		Name:    "mcr.microsoft.com/openjdk/jdk:21-azurelinux",
		Version: "3.12.9",
	}
	pythonImage := domain.RecommendedImage{
		Name:    "mcr.microsoft.com/azurelinux/base/python:3.12",
		Version: "3.12.9",
	}
	customImage := domain.RecommendedImage{
		Name:    "my-custom-image:latest",
		Version: "3.12.0",
	}

	t.Run("filters JDK from python results", func(t *testing.T) {
		images := []domain.RecommendedImage{jdkImage, pythonImage}
		result := FilterIncidentalImages(images, "python")
		assert.Len(t, result, 1)
		assert.Equal(t, pythonImage.Name, result[0].Name)
	})

	t.Run("keeps JDK in java results", func(t *testing.T) {
		images := []domain.RecommendedImage{jdkImage}
		result := FilterIncidentalImages(images, "java")
		assert.Len(t, result, 1)
		assert.Equal(t, jdkImage.Name, result[0].Name)
	})

	t.Run("keeps images with no inferred primary", func(t *testing.T) {
		images := []domain.RecommendedImage{customImage, pythonImage}
		result := FilterIncidentalImages(images, "python")
		assert.Len(t, result, 2)
	})

	t.Run("empty input returns empty", func(t *testing.T) {
		result := FilterIncidentalImages(nil, "python")
		assert.Empty(t, result)
	})

	t.Run("all filtered returns empty", func(t *testing.T) {
		images := []domain.RecommendedImage{jdkImage}
		result := FilterIncidentalImages(images, "python")
		assert.Empty(t, result)
	})

	t.Run("dotnet image filtered from python", func(t *testing.T) {
		dotnetImage := domain.RecommendedImage{
			Name: "mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0",
		}
		images := []domain.RecommendedImage{dotnetImage, pythonImage}
		result := FilterIncidentalImages(images, "python")
		assert.Len(t, result, 1)
		assert.Equal(t, pythonImage.Name, result[0].Name)
	})
}
