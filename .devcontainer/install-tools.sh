#!/usr/bin/env bash
set -euo pipefail

# Source centralised tool versions (single source of truth)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../.github/tool-versions.sh
source "${SCRIPT_DIR}/../.github/tool-versions.sh"

echo "=== Installing scanning tools ==="

# Syft (SBOM generator)
echo "Installing Syft..."
curl -sSfL "https://raw.githubusercontent.com/anchore/syft/${SYFT_INSTALLER_SHA}/install.sh" | sudo sh -s -- -b /usr/local/bin "${SYFT_VERSION}"

# Trivy (vulnerability scanner)
echo "Installing Trivy..."
curl -sSfL "https://raw.githubusercontent.com/aquasecurity/trivy/${TRIVY_INSTALLER_SHA}/contrib/install.sh" | sudo sh -s -- -b /usr/local/bin "${TRIVY_VERSION}"

# golangci-lint
echo "Installing golangci-lint..."
curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_LINT_INSTALLER_SHA}/install.sh" | sudo sh -s -- -b /usr/local/bin "${GOLANGCI_LINT_VERSION}"

echo "=== Tools installed ==="
syft version
trivy version
golangci-lint version
