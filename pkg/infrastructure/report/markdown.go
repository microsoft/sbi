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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/microsoft/sbi/pkg/domain"
	"github.com/microsoft/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
)

// GenerateReport produces a markdown recommendations report from the database.
func GenerateReport(repo *database.Repository, outputPath string, topN int, repoCfg *domain.RepositoryConfig) error {
	languages, err := repo.QueryLanguages()
	if err != nil {
		return fmt.Errorf("querying languages: %w", err)
	}

	if len(languages) == 0 {
		log.Warn("No languages found in database; report not generated.")
		return nil
	}

	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	var sb strings.Builder

	sb.WriteString("# Daily Recommended Images by Language\n\n")
	topNLabel := fmt.Sprintf("Top %d", topN)
	if topN <= 0 {
		topNLabel = "All images"
	}
	fmt.Fprintf(&sb, "_Generated: %s. Criteria: lowest critical → high → total vulnerabilities → size. %s per language per base OS._\n\n", ts, topNLabel)
	sb.WriteString("**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).\n\n")

	if repoCfg != nil {
		writeScannedRepos(&sb, repoCfg)
	}

	for _, lang := range languages {
		var section strings.Builder

		oses, err := repo.QueryBaseOSes(lang)
		if err != nil {
			log.Warnf("Failed to query OSes for %s: %v", lang, err)
			continue
		}

		if len(oses) <= 1 {
			osName := "Other"
			if len(oses) == 1 {
				osName = oses[0]
			}

			images, err := repo.QueryTopImagesByOS(lang, osName, 0)
			if err != nil {
				log.Warnf("Failed to query images for %s: %v", lang, err)
				continue
			}

			images = FilterIncidentalImages(images, lang)
			images = DeduplicateByDigest(images)
			if topN > 0 && len(images) > topN {
				images = images[:topN]
			}

			if len(images) > 0 {
				writeImageTable(&section, images)
			}
		} else {
			for _, osName := range oses {
				osImages, err := repo.QueryTopImagesByOS(lang, osName, 0)
				if err != nil {
					log.Warnf("Failed to query images for %s/%s: %v", lang, osName, err)
					continue
				}

				osImages = FilterIncidentalImages(osImages, lang)
				osImages = DeduplicateByDigest(osImages)
				if topN > 0 && len(osImages) > topN {
					osImages = osImages[:topN]
				}

				if len(osImages) == 0 {
					continue
				}

				fmt.Fprintf(&section, "### %s\n\n", DisplayOSName(osName))
				writeImageTable(&section, osImages)
			}
		}

		if section.Len() > 0 {
			fmt.Fprintf(&sb, "## %s\n\n", DisplayLanguageName(lang))
			sb.WriteString(section.String())
		}
	}

	// Ensure output directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(strings.TrimRight(sb.String(), "\n")+"\n"), 0o644); err != nil {
		return fmt.Errorf("writing report: %w", err)
	}

	log.Infof("Wrote daily recommendations to %s", outputPath)

	return nil
}

// writeImageTable writes a ranked markdown table of images.
func writeImageTable(sb *strings.Builder, images []domain.RecommendedImage) {
	sb.WriteString("| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |\n")
	sb.WriteString("|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|\n")

	for idx, img := range images {
		version := img.Version
		if version == "" {
			version = "-"
		}

		pinnedRef := FormatPinnedReference(img.Name, img.Digest)
		if pinnedRef == "" {
			pinnedRef = "-"
		}

		alsoTagged := "-"
		if len(img.AlternateTags) > 0 {
			alsoTagged = strings.Join(img.AlternateTags, ", ")
		}

		fmt.Fprintf(sb, "| %d | `%s` | %s | %s | %d | %d | %d | %s | %s | `%s` | `%s` |\n",
			idx+1, img.Name, version, alsoTagged,
			img.CriticalVulnerabilities, img.HighVulnerabilities, img.TotalVulnerabilities,
			HumanSize(img.SizeBytes), FormatCreatedDate(img.CreatedDate), FormatDigest(img.Digest), pinnedRef)
	}

	sb.WriteString("\n")
}

// DisplayLanguageName converts a raw language string to a human-readable display name.
func DisplayLanguageName(lang string) string {
	if strings.ToLower(lang) == "base" {
		return "Base / No Runtime"
	}

	return strings.Title(lang) //nolint:staticcheck
}

// DisplayOSName converts a raw OS family string to a human-readable display name.
func DisplayOSName(osFamily string) string {
	switch strings.ToLower(osFamily) {
	case "azurelinux":
		return "Azure Linux"
	case "ubuntu":
		return "Ubuntu"
	case "debian":
		return "Debian"
	case "alpine":
		return "Alpine"
	case "", "other":
		return "Other"
	default:
		return strings.Title(osFamily) //nolint:staticcheck
	}
}

// writeScannedRepos appends the scanned repositories section to the report.
func writeScannedRepos(sb *strings.Builder, cfg *domain.RepositoryConfig) {
	sb.WriteString("## Scanned Repositories and Images\n\n")

	var totalImages int
	for _, group := range cfg.Repositories {
		totalImages += len(group.Images)
	}

	fmt.Fprintf(sb, "This report includes analysis from **%d configured sources** across %d groups (see [repositories.json](../config/repositories.json)):\n\n",
		totalImages, len(cfg.Repositories))

	for _, group := range cfg.Repositories {
		if group.Description != "" {
			fmt.Fprintf(sb, "**%s:**\n\n", group.Description)
		}

		for _, img := range group.Images {
			fmt.Fprintf(sb, "- `%s`\n", img)
		}

		sb.WriteString("\n")
	}
}

// HumanSize converts bytes to a human-readable size string.
func HumanSize(numBytes int64) string {
	if numBytes <= 0 {
		return "-"
	}

	mb := float64(numBytes) / (1024 * 1024)
	if mb < 1024 {
		return fmt.Sprintf("%.1f MB", mb)
	}

	gb := mb / 1024

	return fmt.Sprintf("%.2f GB", gb)
}

// FormatCreatedDate returns the image creation date in compact report form.
func FormatCreatedDate(createdDate string) string {
	if createdDate == "" {
		return "-"
	}

	parsed, err := time.Parse(time.RFC3339Nano, createdDate)
	if err != nil {
		return "-"
	}

	return parsed.UTC().Format("2006-01-02")
}

// FormatDigest returns a shortened digest string for display.
func FormatDigest(digest string) string {
	if digest == "" {
		return ""
	}

	if strings.HasPrefix(digest, "sha256:") {
		hashPart := digest[7:]
		if len(hashPart) > 12 {
			return fmt.Sprintf("sha256:%s", hashPart[:12])
		}

		return digest
	}

	if len(digest) > 19 {
		return digest[:19]
	}

	return digest
}

// FormatPinnedReference returns a copy-friendly pinned image reference with digest.
// Format: {name}@{digest} for supply chain security.
// Note: name already includes the tag (e.g., "registry/image:tag").
func FormatPinnedReference(name, digest string) string {
	if name == "" || digest == "" {
		return ""
	}

	return fmt.Sprintf("%s@%s", name, digest)
}

// stripTag removes the tag portion from a full image name.
// E.g., "mcr.microsoft.com/azurelinux/distroless/python:3.12-nonroot" -> "mcr.microsoft.com/azurelinux/distroless/python"
func stripTag(name string) string {
	// Find the last colon that's after the last slash (to avoid stripping port numbers)
	lastSlash := strings.LastIndex(name, "/")
	lastColonIdx := strings.LastIndex(name, ":")

	// If colon exists and is after the last slash, it's a tag
	if lastColonIdx > lastSlash {
		return name[:lastColonIdx]
	}

	return name
}

// FormatStableTag returns a stable tag reference (major.minor) for auto-updates.
// Format: {base_name}:{major.minor} where base_name is the image name without its current tag.
func FormatStableTag(name, version string) string {
	if name == "" || version == "" {
		return ""
	}

	baseName := stripTag(name)

	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return fmt.Sprintf("%s:%s.%s", baseName, parts[0], parts[1])
	}

	return fmt.Sprintf("%s:%s", baseName, version)
}

// FormatDockerfileFrom returns a FROM line for direct Dockerfile use.
// Format: FROM {name}@{digest}
func FormatDockerfileFrom(name, digest string) string {
	pinnedRef := FormatPinnedReference(name, digest)
	if pinnedRef == "" {
		return ""
	}

	return fmt.Sprintf("FROM %s", pinnedRef)
}

// FormatRecommendedImages formats a list of recommended images for a given language (used externally).
func FormatRecommendedImages(lang string, images []domain.RecommendedImage) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "## %s\n\n", DisplayLanguageName(lang))
	sb.WriteString("| Rank | Image | Version | Also Tagged As | Crit | High | Total | Size | Created | Digest | Pinned Reference |\n")
	sb.WriteString("|------|-------|---------|----------------|------|------|-------|------|---------|--------|------------------|\n")

	for idx, img := range images {
		version := img.Version
		if version == "" {
			version = "-"
		}

		pinnedRef := FormatPinnedReference(img.Name, img.Digest)
		if pinnedRef == "" {
			pinnedRef = "-"
		}

		alsoTagged := "-"
		if len(img.AlternateTags) > 0 {
			alsoTagged = strings.Join(img.AlternateTags, ", ")
		}

		fmt.Fprintf(&sb, "| %d | `%s` | %s | %s | %d | %d | %d | %s | %s | `%s` | `%s` |\n",
			idx+1, img.Name, version, alsoTagged,
			img.CriticalVulnerabilities, img.HighVulnerabilities, img.TotalVulnerabilities,
			HumanSize(img.SizeBytes), FormatCreatedDate(img.CreatedDate), FormatDigest(img.Digest), pinnedRef)
	}

	return sb.String()
}
