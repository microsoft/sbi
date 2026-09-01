# SBI — Secure Base Image Recommendations

[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/microsoft/sbi/badge)](https://scorecard.dev/viewer/?uri=github.com/microsoft/sbi)

📝 **Blog post:** [The Best MCR Base Image for Your Language? SBI Scans Nightly So You Don't Have To](https://medium.com/microsoftazure/the-best-mcr-base-image-for-your-language-sbi-scans-nightly-so-you-dont-have-to-62f17587a45e)

Every night, this project scans configured MCR (Microsoft Container Registry) container base images for vulnerabilities and generates a recommended secure base images report, ranked by language. The default configuration targets MCR images, but the tool supports scanning any container registry.

## 📊 Daily Reports

| Format        | Link                                                                             |
|---------------|----------------------------------------------------------------------------------|
| Markdown      | [docs/daily_recommendations.md](docs/daily_recommendations.md)                   |
| JSON          | [docs/daily_recommendations.json](docs/daily_recommendations.json)               |
| Detailed JSON | [docs/daily_recommendations_detail.json](docs/daily_recommendations_detail.json) |

Reports are regenerated nightly at 02:00 UTC via [GitHub Actions](.github/workflows/nightly-scan.yml) and committed automatically. Images are ranked per language by: fewest critical → fewest high → fewest total vulnerabilities → smallest size.

## How It Works

A nightly [GitHub Actions workflow](.github/workflows/nightly-scan.yml) runs the full pipeline:

1. **Discover** — Enumerate image tags from MCR (Microsoft Container Registry)
2. **Pull & Analyze** — Pull images, generate SBOM with [Syft](https://github.com/anchore/syft), detect language runtimes
3. **Scan** — Run [Trivy](https://github.com/aquasecurity/trivy) vulnerability scanning
4. **Verify** — Runtime verification of detected languages inside containers
5. **Store** — Persist results in a SQLite database (tracked via Git LFS)
6. **Report** — Generate ranked markdown, JSON, and [detailed JSON](docs/detailed-report.md) reports, commit and push to this repo

### What Gets Scanned

Image sources and tag filtering rules are configured in [`config/repositories.json`](config/repositories.json). Currently scans Azure Linux base/distroless images, .NET, Go, and OpenJDK images from MCR.

> **Note:** The daily reports will be updated on the next scheduled nightly run after merge.

## Running Locally

### Prerequisites

- **Go** 1.27+, **Docker**, **[Syft](https://github.com/anchore/syft#installation)**, **[Trivy](https://github.com/aquasecurity/trivy#get-trivy)**

### Quick Start

```bash
# Install
go install github.com/microsoft/sbi@latest

# Scan all configured repositories and generate reports
sbi scan --verbose

# Regenerate reports from existing database
sbi report

# Generate reports with detailed per-image breakdown
sbi report --detailed

# Clear the database
sbi reset-db
```

### Build from Source

```bash
task build
./bin/$(go env GOOS)-$(go env GOARCH)/sbi scan --verbose
```

### CLI Flags

**Global flags** (available on all subcommands):

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--database` | `azure_linux_images.db` | Path to SQLite database |
| `--config-dir` | `config` | Path to configuration directory |
| `--output` | `docs/daily_recommendations.md` | Path to output report file |
| `--top-n` | 10 | Number of top images per language per base OS (0 = all) |
| `--json-top-n` | 20 | Number of top images per language per base OS in JSON report (0 = all) |
| `--detailed` | false | Generate detailed per-image JSON report with packages and vulnerabilities |
| `--verbose, -v` | false | Enable verbose output |
| `--debug, -d` | false | Enable debug output |

**`scan` flags:**

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--max-tags` | 5 | Maximum tags per repository (0 = all) |
| `--comprehensive` | false | Enable comprehensive scanning (secrets + misconfigs) |
| `--update-existing` | false | Rescan existing images |
| `--no-cleanup` | false | Keep Docker images after scanning |

> **Tip:** The `--detailed` flag generates a rich per-image JSON report with packages, CVEs, and languages. See [docs/detailed-report.md](docs/detailed-report.md) for the report schema and jq query examples.

## Configuration

Image sources and tag filtering rules are defined in [`config/repositories.json`](config/repositories.json).

### Adding or modifying repositories

Each entry in the `repositories` array is a group with a description and a list of images to scan. Images can be either **repositories** (all matching tags are discovered and scanned) or **specific image:tag** pairs.

To add a new repository group, add an entry like:

```json
{
  "description": "My custom images",
  "images": [
    "azurelinux/base/core",
    "mcr.microsoft.com/dotnet/aspnet:8.0"
  ]
}
```

- **Repository** (no `:tag`): Value must be a **repository path only** (no registry prefix), for example `azurelinux/base/core`. Tags are auto-discovered from the registry host configured in `defaults.registry` (default: `mcr.microsoft.com`), filtered by `tagFilter` rules, and limited by `maxTags`.
- **Single image** (with `:tag`): Scanned as-is, no tag discovery. Use a full image reference including registry (e.g., `mcr.microsoft.com/dotnet/aspnet:8.0`).

### Tag filtering

The `tagFilter` section controls which discovered tags are included:

| Field | Purpose | Example |
|-------|---------|---------|
| `skipExact` | Tags to skip by exact match | `["latest", "dev", "nightly"]` |
| `excludeKeywords` | Skip tags containing these substrings | `["debug", "test", "arm"]` |
| `excludePatterns` | Skip tags matching these regex patterns | `["(?i)[-.]?(alpha\|beta)"]` |
| `requireDigit` | Only include tags that contain a digit | `true` |

By default, `excludePatterns` skips pre-release tags and date-stamped historical build tags such as `3.0.20250206`. If you need historical comparisons, remove the date-stamped pattern from a custom config.

### Full config example

```json
{
  "defaults": {
    "registry": "mcr.microsoft.com",
    "maxTags": 0
  },
  "tagFilter": {
    "skipExact": ["latest", "dev", "nightly", "edge"],
    "excludeKeywords": ["debug", "test", "arm", "amd"],
    "excludePatterns": [
      "(?i)[-.]?(alpha|beta|rc|preview)[\\d.]*$",
      "^\\d+\\.\\d+\\.\\d{8}$"
    ],
    "requireDigit": true
  },
  "repositories": [
    {
      "description": "Azure Linux base images",
      "images": ["azurelinux/base/python", "azurelinux/base/nodejs"]
    }
  ]
}
```

## Development

Requires [Task](https://taskfile.dev/) for build automation:

```bash
task build        # Build binary
task test         # Run tests
task lint         # Run all linters (go, markdown, yaml)
task vulncheck    # Run Go vulnerability check
task all          # Build + test + lint
```

## Project Structure

```text
*.go                           # CLI entry point and cobra commands (root level)
pkg/
  domain/                      # Domain models (ImageRecord, Language, etc.)
  infrastructure/
    database/                  # SQLite schema and repository
    scanner/                   # Registry, Docker, Syft, Trivy integration
    report/                    # Markdown and JSON report generation
  usecase/                     # Pipeline orchestration
config/                        # Image sources and tag filter config
docs/                          # Generated daily reports
```

## License

MIT

## Trademarks

This project may contain trademarks or logos for projects, products, or services.
Authorized use of Microsoft trademarks or logos is subject to and must follow
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not
cause confusion or imply Microsoft sponsorship. Any use of third-party trademarks
or logos are subject to those third-party's policies.
