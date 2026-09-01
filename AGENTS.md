# AGENTS.md — LLM Context for SBI

> **SBI** (Secure Base Images) is a CLI tool that scans container base images and, in the default nightly workflow/config, performs vulnerability scanning of images from MCR (Microsoft Container Registry), then generates ranked recommendation reports.

## Project Overview

- **Module:** `github.com/microsoft/sbi`
- **Language:** Go 1.27, built with `task build`
- **CLI framework:** Cobra (`github.com/spf13/cobra`)
- **Database:** SQLite via `modernc.org/sqlite` (pure Go, no CGO)
- **Binary name:** `sbi`

## Repository Structure

```text
*.go (root)                   # Cobra CLI commands (scan, report, reset-db)
pkg/
  domain/                     # Data models (ImageRecord, Language, Vulnerability, etc.)
  infrastructure/
    database/                 # SQLite schema, persistence, ranked queries
      schema.go               # Table definitions (images, languages, vulnerabilities, etc.)
      repository.go           # InsertImage (upsert), QueryTopImages, QueryLanguages, QueryAllImageDetails
    scanner/
      analyzer.go             # Orchestrates: pull -> inspect -> syft -> trivy -> verify
      syft.go                 # SBOM parsing, language detection from packages
      trivy.go                # Vulnerability scanning
      registry.go             # Tag discovery from container registries
      docker.go               # Docker pull, inspect, run commands in containers
    report/
      markdown.go             # Markdown report generation
      json.go                 # JSON summary report generation
      json_detail.go          # Detailed per-image JSON report (packages, CVEs, languages)
  usecase/
    pipeline.go               # Top-level orchestration: config -> discover -> scan -> report
config/
  repositories.json           # Image sources for nightly scans
  smoke-test/repositories.json # Minimal config for CI smoke tests
docs/
  daily_recommendations.md    # Generated nightly report (committed to repo)
  daily_recommendations.json  # Generated JSON report
```

## Build, Test, Lint

All commands use [Task](https://taskfile.dev/) (see `Taskfile.yml`).
**Always run `task lint` before creating a PR** — CI runs linting on all PRs and will fail on errors.

```bash
task build          # Build binary to bin/{OS}-{ARCH}/sbi
task test           # Unit tests with coverage
task test:short     # Fast tests without coverage
task lint           # All linters: golangci-lint, markdown, YAML, govulncheck
task lint:go        # golangci-lint with auto-fix
task lint:md        # Markdown linting (uses markdownlint-cli)
task lint:yaml      # YAML validation
task all            # deps -> lint -> test -> build
task scan           # Build + run full nightly scan
task report         # Generate reports from existing database
task clean          # Remove build/coverage artifacts
```

## CLI Usage

```bash
# Install
go install github.com/microsoft/sbi@latest

# Scan images and generate report
sbi scan \
  --database azure_linux_images.db \
  --config-dir config \
  --output docs/daily_recommendations.md \
  --max-tags 10 \
  --top-n 10 \
  --comprehensive \
  --verbose

# Generate report from existing DB
sbi report --database azure_linux_images.db

# Reset database
sbi reset-db --database azure_linux_images.db
```

**Key flags:**

- `--top-n N` — number of top images per language per base OS in markdown report (default 10, 0 = all)
- `--json-top-n N` — number of top images per language per base OS in JSON report (default 20, 0 = all)
- `--max-tags N` — limit tags per repository (0 = all)
- `--comprehensive` — enable secrets + misconfiguration scanning
- `--detailed` — generate a detailed per-image JSON report with full package/vulnerability/language data
- `--update-existing` — rescan images already in the database
- `--no-cleanup` — keep Docker images after scanning

## Scan Pipeline Flow

```text
config/repositories.json
  |
  |-- Repository entries (no ':') -> GetTags() from registry API -> FilterTags() -> LimitTags()
  |-- Single image entries (has ':') -> scan directly
  |
  v
For each image:
  1. docker pull
  2. docker inspect          -> size, layers, digest, created date
  3. syft <image> -o json    -> SBOM: languages, packages, capabilities
  4. trivy image <image>     -> vulnerabilities (+ secrets/misconfig if --comprehensive)
  5. mergeLanguages()        -> combine Syft results + image-name detection
  6. verifyRuntimes()        -> run version commands inside container
  7. applyRuntimeVersions()  -> update with verified versions
  8. InsertImage() (upsert)  -> store in SQLite
  |
  v
Generate reports:
  - docs/daily_recommendations.md
  - docs/daily_recommendations.json
```

## Image Configuration Format

`config/repositories.json` supports two entry types:

```json
{
  "defaults": { "registry": "mcr.microsoft.com", "maxTags": 0 },
  "tagFilter": {
    "skipExact": ["latest", "dev", "nightly"],
    "excludeKeywords": ["debug", "test", "experimental", "arm", "amd", "azl"],
    "excludePatterns": ["(?i)[-.]?(alpha|beta|rc|preview)...", "^\\d+\\.\\d+\\.\\d{8}$"],
    "requireDigit": true
  },
  "repositories": [
    {
      "description": "Tag discovery repos",
      "images": ["azurelinux/base/python"]
    },
    {
      "description": "Pinned single images",
      "images": ["mcr.microsoft.com/dotnet/aspnet:8.0"]
    },
    {
      "description": "Base / minimal images (no runtime)",
      "category": "base",
      "images": ["azurelinux/base/core"]
    }
  ]
}
```

- **No colon** -> repository: tags are discovered via `GET /v2/{repo}/tags/list`, then filtered
- **Has colon** -> single image: scanned directly, no tag discovery
- **`"category": "base"`** -> marks a group as containing minimal/no-runtime images. During scanning, images in these groups that have no detected language runtime receive a synthetic `"base"` language entry, so they appear in a **"Base / No Runtime"** report section. Images in the group that do have detected languages are unaffected.
- Default tag filtering excludes date-stamped historical build tags such as `3.0.20250206`; remove `^\\d+\\.\\d+\\.\\d{8}$` from a custom config to scan historical tags.

## Language Detection (3-stage pipeline)

### Stage 1: SBOM Package Detection (`syft.go`)

Syft scans all packages in the image. Package **names** are matched against regex patterns:

| Language | Regex | Matches |
|----------|-------|---------|
| Python | `(?i)(python&#124;cpython)[\d.-]*$` | `python3`, `cpython` |
| Node.js | `(?i)(node&#124;nodejs)[\d.-]*$` | `nodejs`, `node18` |
| Java | `(?i)(java&#124;openjdk&#124;jdk&#124;jre)[\d.-]*$` | `openjdk-21`, `jre` |
| Go | `(?i)^golang$&#124;^go[\d.]` | `golang`, `go1.21` |
| Ruby | `(?i)^ruby[\d.-]*$` | `ruby3.2` |
| PHP | `(?i)^php[\d.-]*$` | `php8.2` |
| .NET | `(?i)^(dotnet&#124;aspnet)[\w.-]*$` | `dotnet-runtime-8.0`, `aspnetcore-runtime-9.0` |
| Rust | `(?i)^rust[\d.-]*$` | `rust` |
| Lua | `(?i)^lua[\d.-]*$` | `lua5.4` |

**Excluded packages** (not treated as runtimes): `python-pip`, `python3-pip`, `nodejs-npm`, `java-common`, `python-setuptools`

**Resolution strategy** — for name-regex matches, **first match wins** per language. In practice, the name regexes are specific enough that only the correct runtime package matches (e.g., `python3-libs` does not match the Python regex, only `python3` does).

In addition to name-regex matching, Stage 1 uses **package type detection** for two languages:

- **`.NET` via `type=dotnet`** — catches distroless/runtime images where package names like `Microsoft.NETCore.App.Runtime.linux-arm64` don't match the name regex. When multiple `type=dotnet` packages exist, the `isDotnetRuntime()` helper prefers core runtime packages (`Microsoft.NETCore.App.Runtime.*`, `Microsoft.AspNetCore.App.Runtime.*`) over NuGet libraries.
- **Go via `type=binary, name=go`** — Go is tarball-installed in MCR images, not RPM-packaged, so the name regex (`golang`, `go1.x`) won't find it. The binary cataloger detects the `go` binary directly.

**Version cleaning:** `^(\d+(?:\.\d+)*)` strips build suffixes — `3.12.9-8.azl3` -> `3.12.9`

### Stage 2: Image-Name Detection (`analyzer.go`)

Fallback for .NET images that were not detected by Stage 1. Currently only .NET is detected this way:

| Pattern | Matches | Detects |
|---------|---------|---------|
| `(?i)mcr\.microsoft\.com/dotnet/(?:aspnet&#124;runtime):(\d+\.\d+)` | `mcr.microsoft.com/dotnet/aspnet:8.0` | .NET 8.0 |

Only triggers if .NET was **not** already found by Stage 1. In practice, Stage 1's `type=dotnet` detection now handles most .NET images (including distroless), so Stage 2 is rarely the primary detection path.

Note: The SDK pattern (`dotnet/sdk:*`) is **not** matched by `dotnetImagePattern`. SDK images are detected by Stage 1 via `type=dotnet` runtime packages or `dotnet-sdk-*` RPMs.

### Stage 3: Runtime Verification (`analyzer.go`)

For each detected language that has a defined runtime verification command (Python, Node.js, Java, Go, Ruby, PHP, .NET, Rust), runs the actual runtime command **inside the container**:

| Language | Command | Version Pattern |
|----------|---------|-----------------|
| Python | `python3 --version` | `Python\s+(\d+\.\d+\.\d+)` |
| Node.js | `node --version` | `v(\d+\.\d+\.\d+)` |
| Java | `java -version` | `version\s+"?(\d+[\d.]*)` |
| Go | `go version` | `go(\d+\.\d+[\d.]*)` |
| Ruby | `ruby --version` | `ruby\s+(\d+\.\d+\.\d+)` |
| PHP | `php --version` | `PHP\s+(\d+\.\d+\.\d+)` |
| .NET | `dotnet --info` | `Version:\s+(\d+\.\d+[\d.]*)` |
| Rust | `rustc --version` | `rustc\s+(\d+\.\d+\.\d+)` |

Runtime verification overrides the Syft/image-name version with the precise version and sets `Verified=true`. Note: runtime verification only runs for languages already detected by Stage 1 or 2 — it does not independently discover new languages.

## Known Edge Cases

1. **Go build images** — Go is tarball-installed, not RPM-packaged. Syft won't detect `golang` as a system package, but Stage 1's `type=binary, name=go` detection finds the Go binary directly.

2. **JDK images include Python** — Azure Linux JDK images ship `python3` as a system dependency. The image will appear in **both** Java and Python rankings. This is intentional — Python CVEs in the image are real security concerns.

3. **Go modules in non-Go images** — Syft detects `go-module` entries (compiled Go binaries embedded in images, e.g., `jaz.git` in JDK images). These are **not** treated as "Go runtime" — only `type=binary, name=go` triggers Go detection, not `go-module` type entries.

4. **.NET distroless images** — Syft reports .NET packages with `type=dotnet` and names like `Microsoft.NETCore.App.Runtime.linux-arm64`. Stage 1's type-based detection handles these directly. The `isDotnetRuntime()` helper ensures the core runtime package is preferred over NuGet libraries. Stage 2 (`dotnetImagePattern`) serves as a fallback if Stage 1 misses detection.

5. **.NET SDK images** — The SDK image name pattern (`dotnet/sdk:10.0-azurelinux3.0`) is **not** matched by `dotnetImagePattern` which only looks for `aspnet` or `runtime`. SDK images are detected by Stage 1 via `type=dotnet` runtime packages or `dotnet-sdk-*` RPMs.

## Report Ranking

Images are ranked **per language, per base OS** using this sort order (ascending = fewer/smaller is better):

1. **Critical vulnerabilities** (ascending)
2. **High vulnerabilities** (ascending)
3. **Total vulnerabilities** (ascending)
4. **Image size in bytes** (ascending)

### Report Structure

Reports are grouped by **Language → Base OS → ranked table**:

- When a language has images from multiple OSes (e.g., Azure Linux + Debian), each OS gets a `### {OS Name}` sub-heading with its own ranked table.
- When a language has images from only one OS, the table is shown directly under the language heading (no OS sub-heading).
- Images with unknown/empty OS are grouped as "Other" (sorted last).
- The **"Base / No Runtime"** section (language = `"base"`) always appears last, after all runtime language sections. It contains minimal images (e.g., `azurelinux/base/core`) suitable for deploying static binaries (Go, Rust).

### Incidental Runtime Filtering

At report generation time, images are filtered so that recommendation sections only include images whose **primary runtime** matches the section language. This prevents images with incidental secondary runtimes (e.g., Python installed as a system dependency in JDK images) from appearing in unrelated language sections.

**How it works:** The primary language is inferred from the image's repository path segments. For example, `openjdk/jdk` → Java, `azurelinux/base/python` → Python, `dotnet/aspnet` → .NET. If no primary language can be inferred from the image name, the image is kept in all sections (safe fallback).

**Mapped segments:** `python` → Python, `nodejs`/`node` → Node.js, `openjdk`/`jdk`/`jre` → Java, `golang`/`go` → Go, `dotnet`/`aspnet` → .NET, `ruby` → Ruby, `php` → PHP, `rust` → Rust.

All detected languages are still stored in the database — the filtering is applied only when generating reports. The filter runs before deduplication and top-N limiting.

**Language display names:** `"base"` → "Base / No Runtime", all others → title case (e.g., `"python"` → "Python").

**OS display names:** `azurelinux` → "Azure Linux", `ubuntu` → "Ubuntu", `debian` → "Debian", `alpine` → "Alpine".

### Markdown Report

The markdown report shows the top N images per language per OS (default: 10) with columns: Rank, Image, Version, Crit, High, Total, Size, Created, Digest, Pinned Reference.

### JSON Report

The JSON report uses a flat `images` array — each entry has `language` and `baseOS` fields for easy filtering:

```json
{
  "images": [
    { "rank": 1, "language": "dotnet", "baseOS": "Azure Linux", "name": "...", ... },
    { "rank": 1, "language": "dotnet", "baseOS": "Debian", "name": "...", ... }
  ]
}
```

JSON top-N defaults to 20 (configurable via `--json-top-n`, 0 = all). Rank resets per language+OS group. Each entry includes `createdDate`, `pinnedReference`, `stableTag`, and `dockerfileFrom`.

`RecommendedImage` includes image name, language version, vulnerability counts, size, created date, digest, and base OS for report generation.

### Detailed JSON Report

When `--detailed` is passed, a separate `*_detail.json` file is generated with per-image package inventories, vulnerability breakdowns (CVE ID, severity, CVSS, fix availability), system packages, package managers, and detected languages. This report covers all scanned images (not top-N filtered), uses a flat `images` array with one entry per DB image, and includes a `schemaVersion` field for forward compatibility. See [docs/detailed-report.md](docs/detailed-report.md) for schema and jq query examples.

## Database Schema

| Table | Purpose |
|-------|---------|
| `images` | Image metadata, vulnerability counts, scan timestamps |
| `languages` | Detected languages per image (language, version, major_minor, verified) |
| `vulnerabilities` | Individual CVEs (severity, CVSS score, fixed_version) |
| `system_packages` | OS-level packages (RPM/DEB/APK) |
| `package_managers` | Detected package managers (pip, npm, etc.) |
| `capabilities` | Container capabilities |
| `security_findings` | Secrets and misconfigurations (from --comprehensive) |

`InsertImage()` performs an **upsert** (`ON CONFLICT(name) DO UPDATE`) and clears related child tables on re-scan.

## CI/CD Workflows

| Workflow | Trigger | Purpose |
|----------|---------|---------|
| `ci.yml` | Push/PR to main | Build + unit tests + govulncheck |
| `lint.yml` | Push/PR to main | golangci-lint, markdown lint, YAML lint |
| `nightly-scan.yml` | Daily 02:00 UTC | Full scan -> commit reports to main |
| `scan-smoke-test.yml` | PR to main | Minimal scan to validate pipeline |
