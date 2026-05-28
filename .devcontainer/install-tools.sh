#!/usr/bin/env bash
set -euo pipefail

# Source centralised tool versions (single source of truth)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=../.github/tool-versions.sh
source "${SCRIPT_DIR}/../.github/tool-versions.sh"

echo "=== Installing scanning tools ==="

# GitHub and Git LFS tooling used for issue/PR workflows in the dev container
echo "Installing GitHub CLI and Git LFS..."
ARCH="$(dpkg --print-architecture)"
case "${ARCH}" in
	amd64|arm64)
		;;
	*)
		echo "Unsupported architecture: ${ARCH}" >&2
		exit 1
		;;
esac

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

GH_VERSION_NO_V="${GH_VERSION#v}"
GH_TARBALL="${TMP_DIR}/gh.tar.gz"
curl -sSfL -o "${GH_TARBALL}" \
	"https://github.com/cli/cli/releases/download/${GH_VERSION}/gh_${GH_VERSION_NO_V}_linux_${ARCH}.tar.gz"
tar -xzf "${GH_TARBALL}" -C "${TMP_DIR}"
sudo install -m 0755 "${TMP_DIR}/gh_${GH_VERSION_NO_V}_linux_${ARCH}/bin/gh" /usr/local/bin/gh

GIT_LFS_VERSION_NO_V="${GIT_LFS_VERSION#v}"
GIT_LFS_TARBALL="${TMP_DIR}/git-lfs.tar.gz"
curl -sSfL -o "${GIT_LFS_TARBALL}" \
	"https://github.com/git-lfs/git-lfs/releases/download/${GIT_LFS_VERSION}/git-lfs-linux-${ARCH}-v${GIT_LFS_VERSION_NO_V}.tar.gz"
tar -xzf "${GIT_LFS_TARBALL}" -C "${TMP_DIR}"
GIT_LFS_DIR="$(find "${TMP_DIR}" -maxdepth 1 -type d -name 'git-lfs-*' | head -n 1)"
if [[ -z "${GIT_LFS_DIR}" ]]; then
	echo "Failed to locate extracted git-lfs directory" >&2
	exit 1
fi
sudo bash "${GIT_LFS_DIR}/install.sh"
sudo git lfs install --system

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
gh --version
git lfs version
syft version
trivy version
golangci-lint version
