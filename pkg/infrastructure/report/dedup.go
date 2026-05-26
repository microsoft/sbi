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
	"sort"
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
)

// DeduplicateByDigest groups images sharing the same repository and digest,
// keeping the most specific tag as the primary entry and collecting alternate
// tag names. Only images from the same repository are grouped; different
// repositories with the same digest remain separate entries.
// Images with an empty digest are never grouped.
func DeduplicateByDigest(images []domain.RecommendedImage) []domain.RecommendedImage {
	if len(images) <= 1 {
		return images
	}

	type group struct {
		primary int
		indices []int
	}

	// Ordered list of groups preserving first-seen order.
	var groups []group
	// Maps "repo\x00digest" to index in groups slice.
	byKey := make(map[string]int)

	for i, img := range images {
		d := img.Digest
		if d == "" {
			groups = append(groups, group{primary: i, indices: []int{i}})
			continue
		}

		key := imageRepo(img.Name) + "\x00" + d
		if gi, ok := byKey[key]; ok {
			groups[gi].indices = append(groups[gi].indices, i)
			newScore := tagSpecificity(images[i].Name)
			oldScore := tagSpecificity(images[groups[gi].primary].Name)
			if newScore > oldScore || (newScore == oldScore && images[i].Name < images[groups[gi].primary].Name) {
				groups[gi].primary = i
			}
		} else {
			byKey[key] = len(groups)
			groups = append(groups, group{primary: i, indices: []int{i}})
		}
	}

	result := make([]domain.RecommendedImage, 0, len(groups))
	for _, g := range groups {
		img := images[g.primary]
		img.AlternateTags = nil

		for _, idx := range g.indices {
			if idx == g.primary {
				continue
			}
			alt := extractTag(images[idx].Name)
			if alt != "" {
				img.AlternateTags = append(img.AlternateTags, ":"+alt)
			}
		}
		sort.Strings(img.AlternateTags)

		result = append(result, img)
	}

	return result
}

// tagSpecificity scores how specific an image tag is.
// Higher score = more specific. Prefers numeric version tags over
// non-numeric ones, then counts dot-separated version segments.
func tagSpecificity(name string) int {
	tag := extractTag(name)
	if tag == "" {
		return 0
	}

	// Tags starting with a digit are strongly preferred over non-numeric tags
	// (e.g., "3.12" should always beat "latest").
	numericBonus := 0
	if len(tag) > 0 && tag[0] >= '0' && tag[0] <= '9' {
		numericBonus = 10000
	}

	// Extract the leading version portion (before any '-' suffix like "-nonroot").
	versionPart := tag
	if idx := strings.IndexByte(tag, '-'); idx > 0 {
		versionPart = tag[:idx]
	}

	segments := strings.Count(versionPart, ".") + 1

	// Use numericBonus + segments * 1000 + tag length for tie-breaking.
	return numericBonus + segments*1000 + len(tag)
}

// extractTag returns the tag portion of a full image name.
// E.g., "mcr.microsoft.com/repo:3.12-nonroot" → "3.12-nonroot"
// Strips any @sha256:... digest suffix before parsing.
func extractTag(name string) string {
	// Strip @digest suffix if present (e.g., "repo:3.12@sha256:abc..." → "repo:3.12").
	if at := strings.Index(name, "@"); at >= 0 {
		name = name[:at]
	}

	lastSlash := strings.LastIndex(name, "/")
	lastColon := strings.LastIndex(name, ":")

	if lastColon > lastSlash {
		return name[lastColon+1:]
	}

	return ""
}

// imageRepo returns the repository portion of a full image name (everything
// before the tag or digest). E.g., "mcr.microsoft.com/repo:3.12" → "mcr.microsoft.com/repo"
func imageRepo(name string) string {
	// Strip @digest suffix first.
	if at := strings.Index(name, "@"); at >= 0 {
		name = name[:at]
	}

	lastSlash := strings.LastIndex(name, "/")
	lastColon := strings.LastIndex(name, ":")

	if lastColon > lastSlash {
		return name[:lastColon]
	}

	return name
}
