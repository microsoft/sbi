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
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
)

// primaryLanguageSegments maps repository path segments to their primary language.
// Only exact segment matches are used to avoid false positives.
var primaryLanguageSegments = map[string]string{
	"python":  "python",
	"nodejs":  "node",
	"node":    "node",
	"openjdk": "java",
	"jdk":     "java",
	"jre":     "java",
	"golang":  "go",
	"go":      "go",
	"dotnet":  "dotnet",
	"aspnet":  "dotnet",
	"ruby":    "ruby",
	"php":     "php",
	"rust":    "rust",
}

// InferPrimaryLanguage infers the primary language of an image from its
// repository path. Returns the language name (matching the DB convention)
// or "" if no primary language can be determined.
func InferPrimaryLanguage(imageName string) string {
	repo := imageRepo(imageName)
	segments := strings.Split(strings.ToLower(repo), "/")

	for _, seg := range segments {
		if lang, ok := primaryLanguageSegments[seg]; ok {
			return lang
		}
	}

	return ""
}

// FilterIncidentalImages removes images whose inferred primary language
// differs from the queried language. Images with no inferred primary
// language are kept (safe fallback).
func FilterIncidentalImages(images []domain.RecommendedImage, queriedLang string) []domain.RecommendedImage {
	queriedLang = strings.ToLower(queriedLang)
	result := make([]domain.RecommendedImage, 0, len(images))

	for _, img := range images {
		primary := InferPrimaryLanguage(img.Name)
		if primary == "" || primary == queriedLang {
			result = append(result, img)
		}
	}

	return result
}
