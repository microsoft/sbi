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
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// tagsPageSize is the per-request page size for registry tags/list pagination.
const tagsPageSize = 100

// maxTagPages caps GetTags pagination so a cyclic or unbounded Link header
// cannot loop forever. The HTTP client timeout applies per request only.
var maxTagPages = 100

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

// GetTags fetches available tags for a repository, following Registry API
// pagination via the Link response header (rel="next").
func (s *RegistryScanner) GetTags(repo string) ([]string, error) {
	next := fmt.Sprintf("%s/v2/%s/tags/list?n=%d", s.registryURL, repo, tagsPageSize)
	var all []string
	seen := make(map[string]struct{})

	for page := 0; next != ""; page++ {
		if page >= maxTagPages {
			return nil, fmt.Errorf("tag pagination exceeded %d pages for %s", maxTagPages, repo)
		}
		if _, ok := seen[next]; ok {
			log.Warnf("tag pagination cycle detected for %s at %s", repo, next)
			break
		}
		seen[next] = struct{}{}

		tags, linkHeader, err := s.fetchTagPage(next, repo)
		if err != nil {
			return nil, err
		}
		all = append(all, tags...)

		base, err := url.Parse(next)
		if err != nil {
			return nil, fmt.Errorf("parsing tags URL for %s: %w", repo, err)
		}
		next = nextPageFromLink(base, linkHeader)
	}

	return all, nil
}

func (s *RegistryScanner) fetchTagPage(pageURL, repo string) ([]string, string, error) {
	log.Debugf("Fetching tags from: %s", pageURL)

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		return nil, "", fmt.Errorf("creating tags request for %s: %w", repo, err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("fetching tags for %s: %w", repo, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status %d for %s", resp.StatusCode, repo)
	}

	var tagsResp domain.TagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResp); err != nil {
		return nil, "", fmt.Errorf("decoding tags response: %w", err)
	}

	return tagsResp.Tags, resp.Header.Get("Link"), nil
}

// nextPageFromLink returns the absolute URL for rel="next" from an RFC 5988
// Link header, or "" when there is no next page. Only the parameter list
// after ">" is inspected, so a URL that happens to contain "rel=next" is
// not treated as a next link.
func nextPageFromLink(base *url.URL, linkHeader string) string {
	if linkHeader == "" || base == nil {
		return ""
	}

	for _, part := range strings.Split(linkHeader, ",") {
		part = strings.TrimSpace(part)
		start := strings.Index(part, "<")
		end := strings.Index(part, ">")
		if start < 0 || end <= start {
			continue
		}
		if !linkRelIsNext(part[end+1:]) {
			continue
		}

		ref := strings.TrimSpace(part[start+1 : end])
		resolved, err := base.Parse(ref)
		if err != nil {
			log.Warnf("Invalid pagination Link URL %q: %v", ref, err)
			return ""
		}

		return resolved.String()
	}

	return ""
}

func linkRelIsNext(params string) bool {
	for _, p := range strings.Split(params, ";") {
		p = strings.TrimSpace(p)
		key, val, ok := strings.Cut(p, "=")
		if !ok || !strings.EqualFold(strings.TrimSpace(key), "rel") {
			continue
		}
		val = strings.Trim(strings.TrimSpace(val), `"'`)
		for _, token := range strings.Fields(val) {
			if strings.EqualFold(token, "next") {
				return true
			}
		}
	}
	return false
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

	// Sort newest-first with version-aware ordering so LimitTags keeps
	// higher releases (e.g. 10.0 before 9.0), not reverse lexicographic order.
	sortTagsNewestFirst(filtered)

	return filtered
}

// sortTagsNewestFirst sorts tags descending by leading numeric version
// (major.minor.patch…), then by the remaining suffix. Tags without a leading
// version number sort last.
func sortTagsNewestFirst(tags []string) {
	sort.SliceStable(tags, func(i, j int) bool {
		// Descending: i before j when j < i in ascending version order.
		return tagVersionLess(tags[j], tags[i])
	})
}

// tagVersionLess reports whether a sorts before b in ascending version order.
func tagVersionLess(a, b string) bool {
	aSegs, aRest := versionPrefix(a)
	bSegs, bRest := versionPrefix(b)

	// Non-version tags sort before versioned tags in ascending order so they
	// appear last when sorted newest-first.
	switch {
	case len(aSegs) == 0 && len(bSegs) == 0:
		return a < b
	case len(aSegs) == 0:
		return true
	case len(bSegs) == 0:
		return false
	}

	n := len(aSegs)
	if len(bSegs) > n {
		n = len(bSegs)
	}

	for i := 0; i < n; i++ {
		av, bv := 0, 0
		if i < len(aSegs) {
			av = aSegs[i]
		}
		if i < len(bSegs) {
			bv = bSegs[i]
		}
		if av != bv {
			return av < bv
		}
	}

	if aRest != bRest {
		return aRest < bRest
	}

	return a < b
}

// versionPrefix parses the leading numeric version from a tag.
// An optional conventional "v"/"V" prefix is skipped when immediately followed
// by a digit (e.g. "v10.0" is treated like "10.0"), so version-aware sort does
// not regress to reverse-lexicographic order for v-prefixed tags.
// Examples: "10.0-noble" → ([10, 0], "-noble"), "v10.0" → ([10, 0], ""),
// "3.12.9" → ([3, 12, 9], ""), "21-azurelinux" → ([21], "-azurelinux"),
// "latest" → (nil, "latest").
func versionPrefix(tag string) (segs []int, rest string) {
	i := 0
	// Skip optional conventional version prefix only when a digit follows
	// (so "version" / "vv1" are not misparsed as versions).
	if len(tag) >= 2 && (tag[0] == 'v' || tag[0] == 'V') && tag[1] >= '0' && tag[1] <= '9' {
		i = 1
	}

	for i < len(tag) {
		if tag[i] < '0' || tag[i] > '9' {
			break
		}

		n := 0
		for i < len(tag) && tag[i] >= '0' && tag[i] <= '9' {
			n = n*10 + int(tag[i]-'0')
			i++
		}
		segs = append(segs, n)

		if i < len(tag) && tag[i] == '.' && i+1 < len(tag) && tag[i+1] >= '0' && tag[i+1] <= '9' {
			i++ // consume dot between version segments
			continue
		}

		break
	}

	if len(segs) == 0 {
		// No version parsed — return the original tag as rest (including any
		// leading character we may have peeked but not treated as a prefix).
		return nil, tag
	}

	return segs, tag[i:]
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
// Digest suffixes (@sha256:...) are stripped before parsing. Tags are taken from the
// last ":" after the last "/", so host:port registries (e.g. localhost:5000) work.
func ExtractRegistryAndRepo(imageName string) (registry, repository, tag string) {
	name := imageName

	// Strip @digest first — digests contain ":" (e.g. sha256:abc) and must not
	// be mistaken for tags.
	if at := strings.Index(name, "@"); at >= 0 {
		name = name[:at]
	}

	// Tag separator is the last ":" after the last "/".
	lastSlash := strings.LastIndex(name, "/")
	lastColon := strings.LastIndex(name, ":")
	nameWithoutTag := name
	if lastColon > lastSlash {
		tag = name[lastColon+1:]
		nameWithoutTag = name[:lastColon]
	}

	// Registry is the first path segment when it looks like a host.
	segments := strings.SplitN(nameWithoutTag, "/", 2)
	if len(segments) == 2 && looksLikeRegistryHost(segments[0]) {
		registry = segments[0]
		repository = segments[1]
	} else {
		// Default to MCR for short names like "azurelinux/base/python:3.12"
		registry = "mcr.microsoft.com"
		repository = nameWithoutTag
	}

	return registry, repository, tag
}

// looksLikeRegistryHost reports whether s is a registry host (domain, host:port,
// or localhost) rather than a repository path segment.
func looksLikeRegistryHost(s string) bool {
	return strings.Contains(s, ".") || strings.Contains(s, ":") || strings.EqualFold(s, "localhost")
}
