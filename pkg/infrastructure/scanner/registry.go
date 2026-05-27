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
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// DefaultTagFilter returns the default tag filter configuration.
func DefaultTagFilter() domain.TagFilterConfig {
	return domain.TagFilterConfig{
		SkipExact:       []string{"latest", "dev", "nightly", "edge", "rc", "beta", "alpha"},
		ExcludeKeywords: []string{"debug", "test", "experimental", "arm", "amd", "azl"},
		ExcludePatterns: []string{`(?i)[-.]?(alpha|beta|rc|preview|dev|nightly|canary)[\d.]*$`, `^\d+\.\d+\.\d{8}$`},
		RequireDigit:    true,
	}
}

// RegistryScanner discovers image tags from container registries.
type RegistryScanner struct {
	client      *http.Client
	registryURL string
}

// NewRegistryScanner creates a new RegistryScanner with a default registry.
func NewRegistryScanner(defaultRegistry string) *RegistryScanner {
	if defaultRegistry == "" {
		defaultRegistry = "mcr.microsoft.com"
	}

	return &RegistryScanner{
		client:      &http.Client{Timeout: 30 * time.Second},
		registryURL: "https://" + defaultRegistry,
	}
}

// GetTags fetches available tags for an MCR repository.
func (s *RegistryScanner) GetTags(repo string) ([]string, error) {
	url := fmt.Sprintf("%s/v2/%s/tags/list", s.registryURL, repo)
	log.Debugf("Fetching tags from: %s", url)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching tags for %s: %w", repo, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, repo)
	}

	var tagsResp domain.TagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResp); err != nil {
		return nil, fmt.Errorf("decoding tags response: %w", err)
	}

	return tagsResp.Tags, nil
}

// FilterTags removes pre-release, arch-specific, and unwanted tags based on the provided config.
func FilterTags(tags []string, cfg domain.TagFilterConfig) []string {
	skipExact := make(map[string]bool, len(cfg.SkipExact))
	for _, s := range cfg.SkipExact {
		skipExact[strings.ToLower(s)] = true
	}

	var compiledPatterns []*regexp.Regexp
	for _, p := range cfg.ExcludePatterns {
		if re, err := regexp.Compile(p); err == nil {
			compiledPatterns = append(compiledPatterns, re)
		} else {
			log.Warnf("Invalid exclude pattern %q: %v", p, err)
		}
	}

	var filtered []string

	for _, tag := range tags {
		lower := strings.ToLower(tag)

		if skipExact[lower] {
			continue
		}

		excluded := false
		for _, pat := range compiledPatterns {
			if pat.MatchString(tag) {
				excluded = true
				break
			}
		}

		if excluded {
			continue
		}

		keywordHit := false
		for _, kw := range cfg.ExcludeKeywords {
			if strings.Contains(lower, strings.ToLower(kw)) {
				keywordHit = true
				break
			}
		}

		if keywordHit {
			continue
		}

		if cfg.RequireDigit {
			hasDigit := false
			for _, c := range tag {
				if c >= '0' && c <= '9' {
					hasDigit = true
					break
				}
			}

			if !hasDigit {
				continue
			}
		}

		filtered = append(filtered, tag)
	}

	// Sort descending (newest first assuming semver-like ordering)
	sort.Sort(sort.Reverse(sort.StringSlice(filtered)))

	return filtered
}

// LimitTags returns at most maxTags tags. If maxTags is 0, all tags are returned.
func LimitTags(tags []string, maxTags int) []string {
	if maxTags <= 0 || maxTags >= len(tags) {
		return tags
	}

	return tags[:maxTags]
}

// ParseImagePatterns reads repository patterns and categorizes them.
func ParseImagePatterns(patterns []string) (repos []string, singleImages []string) {
	for _, entry := range patterns {
		entry = strings.TrimSpace(entry)
		if entry == "" || strings.HasPrefix(entry, "#") {
			continue
		}

		// Single image if it contains a tag separator ':'
		if strings.Contains(strings.SplitN(entry, "@", 2)[0], ":") {
			singleImages = append(singleImages, entry)
		} else {
			repos = append(repos, entry)
		}
	}

	return repos, singleImages
}

// BuildFullImageName constructs the full image name from registry and repo/tag.
func BuildFullImageName(defaultRegistry, repo, tag string) string {
	if strings.HasPrefix(repo, "mcr.microsoft.com/") || strings.HasPrefix(repo, "docker.io/") || strings.HasPrefix(repo, "ghcr.io/") {
		return fmt.Sprintf("%s:%s", repo, tag)
	}

	if defaultRegistry == "" {
		defaultRegistry = "mcr.microsoft.com"
	}

	return fmt.Sprintf("%s/%s:%s", defaultRegistry, repo, tag)
}

// ExtractRegistryAndRepo splits a full image name into registry, repository, and tag.
func ExtractRegistryAndRepo(imageName string) (registry, repository, tag string) {
	// Split off the tag
	parts := strings.SplitN(imageName, ":", 2)
	nameWithoutTag := parts[0]
	if len(parts) == 2 {
		tag = parts[1]
	}

	// Split into registry and repository
	segments := strings.SplitN(nameWithoutTag, "/", 2)
	if len(segments) == 2 && strings.Contains(segments[0], ".") {
		registry = segments[0]
		repository = segments[1]
	} else {
		// Default to MCR
		registry = "mcr.microsoft.com"
		repository = nameWithoutTag
	}

	return registry, repository, tag
}
