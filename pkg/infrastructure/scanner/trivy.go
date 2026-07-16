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
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// trivyOutput represents the top-level Trivy JSON output.
type trivyOutput struct {
	Results  []trivyResult `json:"Results"`
	Metadata trivyMetadata `json:"Metadata"`
}

type trivyMetadata struct {
	OS trivyOS `json:"OS"`
}

type trivyOS struct {
	Family string `json:"Family"`
	Name   string `json:"Name"`
}

type trivyResult struct {
	Target          string           `json:"Target"`
	Class           string           `json:"Class"`
	Type            string           `json:"Type"`
	Vulnerabilities []trivyVuln      `json:"Vulnerabilities"`
	Secrets         []trivySecret    `json:"Secrets"`
	Misconfigs      []trivyMisconfig `json:"Misconfigurations"`
	Licenses        []trivyLicense   `json:"Licenses"`
}

type trivyCVSS struct {
	V3Score float64 `json:"V3Score"`
}

type trivyVuln struct {
	VulnerabilityID  string               `json:"VulnerabilityID"`
	Severity         string               `json:"Severity"`
	PkgName          string               `json:"PkgName"`
	InstalledVersion string               `json:"InstalledVersion"`
	FixedVersion     string               `json:"FixedVersion"`
	Description      string               `json:"Description"`
	CVSS             map[string]trivyCVSS `json:"CVSS"`
}

type trivySecret struct {
	RuleID   string `json:"RuleID"`
	Category string `json:"Category"`
	Severity string `json:"Severity"`
	Title    string `json:"Title"`
	Match    string `json:"Match"`
}

type trivyMisconfig struct {
	ID       string `json:"ID"`
	Title    string `json:"Title"`
	Severity string `json:"Severity"`
	Message  string `json:"Message"`
}

type trivyLicense struct {
	Name     string `json:"Name"`
	Severity string `json:"Severity"`
}

// RunTrivy executes Trivy vulnerability scan on an image.
func RunTrivy(imageName string, comprehensive bool) (*domain.TrivyResult, error) {
	log.Infof("Running Trivy on: %s (comprehensive=%v)", imageName, comprehensive)

	// Prefer --scanners (Trivy renamed --security-checks). Use "misconfig"
	// rather than the deprecated "config" scanner name.
	scanners := "vuln"
	if comprehensive {
		scanners = "vuln,secret,misconfig"
	}

	cmd := exec.Command("trivy", "image",
		"--format", "json",
		"--scanners", scanners,
		imageName,
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// Trivy may exit non-zero when vulns are found, check if we got JSON output
		if stdout.Len() == 0 {
			return nil, fmt.Errorf("running trivy on %s: %w\n%s", imageName, err, stderr.String())
		}
	}

	var output trivyOutput
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		return nil, fmt.Errorf("parsing trivy output: %w", err)
	}

	return parseTrivyResult(&output), nil
}

func parseTrivyResult(output *trivyOutput) *domain.TrivyResult {
	result := &domain.TrivyResult{
		BaseOSFamily:  output.Metadata.OS.Family,
		BaseOSVersion: output.Metadata.OS.Name,
	}

	for _, r := range output.Results {
		// Process vulnerabilities
		for _, v := range r.Vulnerabilities {
			result.TotalVulnerabilities++

			switch strings.ToUpper(v.Severity) {
			case "CRITICAL":
				result.CriticalVulnerabilities++
			case "HIGH":
				result.HighVulnerabilities++
			case "MEDIUM":
				result.MediumVulnerabilities++
			case "LOW":
				result.LowVulnerabilities++
			case "NEGLIGIBLE":
				result.NegligibleVulnerabilities++
			default:
				result.UnknownVulnerabilities++
			}

			var cvssScore float64
			for _, cvss := range v.CVSS {
				if cvss.V3Score > cvssScore {
					cvssScore = cvss.V3Score
				}
			}

			result.Vulnerabilities = append(result.Vulnerabilities, domain.Vulnerability{
				VulnerabilityID: v.VulnerabilityID,
				Severity:        v.Severity,
				PackageName:     v.PkgName,
				PackageVersion:  v.InstalledVersion,
				FixedVersion:    v.FixedVersion,
				Description:     truncateString(v.Description, 500),
				CVSSScore:       cvssScore,
			})
		}

		// Process secrets
		result.SecretsFound += len(r.Secrets)

		// Process misconfigurations
		result.ConfigIssues += len(r.Misconfigs)

		// Process licenses
		result.LicenseIssues += len(r.Licenses)
	}

	return result
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	return s[:maxLen-3] + "..."
}
