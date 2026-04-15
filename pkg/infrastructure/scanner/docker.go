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
	"strconv"
	"strings"

	"github.com/microsoft/sbi/pkg/domain"
	log "github.com/sirupsen/logrus"
)

// DockerClient wraps docker CLI operations.
type DockerClient struct{}

// NewDockerClient creates a new DockerClient.
func NewDockerClient() *DockerClient {
	return &DockerClient{}
}

// Pull pulls a container image.
func (d *DockerClient) Pull(imageName string) error {
	log.Infof("Pulling image: %s", imageName)

	cmd := exec.Command("docker", "pull", imageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pulling image %s: %w\n%s", imageName, err, string(output))
	}

	return nil
}

// Inspect returns metadata about a pulled image.
func (d *DockerClient) Inspect(imageName string) (domain.DockerInspectResult, error) {
	cmd := exec.Command("docker", "inspect", imageName)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return domain.DockerInspectResult{}, fmt.Errorf("inspecting image %s: %w", imageName, err)
	}

	return parseInspectOutput(stdout.Bytes(), imageName)
}

// parseInspectOutput parses the JSON output from docker inspect into a DockerInspectResult.
func parseInspectOutput(data []byte, imageName string) (domain.DockerInspectResult, error) {
	var result domain.DockerInspectResult

	var inspectData []map[string]interface{}
	if err := json.Unmarshal(data, &inspectData); err != nil {
		return result, fmt.Errorf("parsing inspect output: %w", err)
	}

	if len(inspectData) == 0 {
		return result, fmt.Errorf("empty inspect result for %s", imageName)
	}

	data0 := inspectData[0]

	if size, ok := data0["Size"].(float64); ok {
		result.SizeBytes = int64(size)
	}

	if created, ok := data0["Created"].(string); ok {
		result.Created = created
	}

	// Count layers from RootFS
	if rootFS, ok := data0["RootFS"].(map[string]interface{}); ok {
		if layers, ok := rootFS["Layers"].([]interface{}); ok {
			result.Layers = len(layers)
		}
	}

	// Get digest from RepoDigests
	if digests, ok := data0["RepoDigests"].([]interface{}); ok && len(digests) > 0 {
		if d, ok := digests[0].(string); ok {
			if parts := strings.SplitN(d, "@", 2); len(parts) == 2 {
				result.Digest = parts[1]
			}
		}
	}

	return result, nil
}

// GetImageSize returns the image size using docker images command.
func (d *DockerClient) GetImageSize(imageName string) (int64, error) {
	cmd := exec.Command("docker", "images", imageName, "--format", "{{.Size}}")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("getting image size: %w", err)
	}

	sizeStr := strings.TrimSpace(stdout.String())
	if sizeStr == "" {
		return 0, nil
	}

	return parseDockerSize(sizeStr), nil
}

// RunCommand runs a command inside a container and returns its output.
// Returns combined stdout+stderr since some runtimes (e.g. java -version)
// write version info to stderr.
func (d *DockerClient) RunCommand(imageName string, command []string) (string, error) {
	args := append([]string{"run", "--rm", "--entrypoint", ""}, imageName)
	args = append(args, command...)

	cmd := exec.Command("docker", args...)

	var combined bytes.Buffer
	cmd.Stdout = &combined
	cmd.Stderr = &combined

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("running command in %s: %w\n%s", imageName, err, combined.String())
	}

	return strings.TrimSpace(combined.String()), nil
}

// Remove removes a container image.
func (d *DockerClient) Remove(imageName string) error {
	log.Debugf("Removing image: %s", imageName)

	cmd := exec.Command("docker", "rmi", "--force", imageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Warnf("Failed to remove image %s: %s", imageName, string(output))
		return nil // Non-fatal
	}

	return nil
}

// parseDockerSize converts Docker size strings like "85.3MB" to bytes.
func parseDockerSize(s string) int64 {
	s = strings.TrimSpace(s)
	multiplier := int64(1)

	switch {
	case strings.HasSuffix(s, "GB"):
		multiplier = 1024 * 1024 * 1024
		s = strings.TrimSuffix(s, "GB")
	case strings.HasSuffix(s, "MB"):
		multiplier = 1024 * 1024
		s = strings.TrimSuffix(s, "MB")
	case strings.HasSuffix(s, "KB") || strings.HasSuffix(s, "kB"):
		multiplier = 1024
		s = strings.TrimSuffix(strings.TrimSuffix(s, "KB"), "kB")
	case strings.HasSuffix(s, "B"):
		s = strings.TrimSuffix(s, "B")
	}

	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}

	return int64(val * float64(multiplier))
}
