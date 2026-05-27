# Detailed JSON Report

The `--detailed` flag generates a separate `*_detail.json` report with per-image package inventories, vulnerability breakdowns, and detected languages. This report covers **all** scanned images (not top-N filtered) and is designed for downstream tooling and drill-down analysis.

## Generating the Report

```bash
# Generate all reports including the detailed JSON
sbi scan --detailed --verbose

# Or regenerate from an existing database
sbi report --detailed
```

The detailed report is written alongside the other reports as `docs/daily_recommendations_detail.json` (derived from the `--output` path).

## Report Structure

```json
{
  "schemaVersion": 1,
  "generatedAt": "2026-05-27T11:27:25Z",
  "imageCount": 122,
  "images": [
    {
      "name": "mcr.microsoft.com/azurelinux/base/python:3.12",
      "registry": "mcr.microsoft.com",
      "repository": "azurelinux/base/python",
      "tag": "3.12",
      "digest": "sha256:...",
      "sizeBytes": 145752064,
      "sizeHuman": "139.0 MB",
      "layers": 3,
      "createdDate": "2026-05-17T10:17:59Z",
      "scanTimestamp": "2026-05-27T05:43:37Z",
      "baseOS": { "name": "Azure Linux", "version": "3.0" },
      "vulnerabilitySummary": {
        "total": 1, "critical": 0, "high": 1,
        "medium": 0, "low": 0, "negligible": 0, "unknown": 0
      },
      "languages": [
        { "language": "python", "version": "3.12.9", "majorMinor": "3.12",
          "packageName": "python3", "packageType": "rpm", "verified": true }
      ],
      "vulnerabilities": [
        { "id": "CVE-2026-4046", "severity": "HIGH", "cvssScore": 5.3,
          "packageName": "glibc", "packageVersion": "2.38-19.azl3",
          "fixedVersion": "2.38-20.azl3", "description": "..." }
      ],
      "systemPackages": [
        { "name": "openssl", "version": "3.0.1", "packageType": "rpm" }
      ],
      "packageManagers": [
        { "name": "pip", "version": "24.0", "language": "python" }
      ]
    }
  ]
}
```

> **Optional fields:** Some fields are omitted when empty or unknown, including `registry`, `repository`, `tag`, `digest`, `scanTimestamp`, `baseOS.version`, `languages[].majorMinor`, `languages[].packageName`, `languages[].packageType`, `vulnerabilities[].fixedVersion`, and `vulnerabilities[].description`. Array fields (`languages`, `vulnerabilities`, `systemPackages`, `packageManagers`) are always present as empty arrays `[]`, never null.
>
> **Scope:** The current detailed report includes system packages, package managers, detected languages, and CVE vulnerabilities. Comprehensive scan findings (secrets, misconfigurations) and capabilities are not included.

## Querying with jq

### Get full details for a specific image

```bash
jq '.images[] | select(.name | contains("python:3.12"))' docs/daily_recommendations_detail.json
```

### Summary view with package counts

```bash
jq '.images[] | select(.name | contains("python:3.12")) | {
  name, baseOS, vulnerabilitySummary, languages,
  systemPackageCount: (.systemPackages | length),
  packageManagerCount: (.packageManagers | length)
}' docs/daily_recommendations_detail.json
```

### List all images with critical vulnerabilities

```bash
jq '.images[] | select(.vulnerabilitySummary.critical > 0) |
  {name, critical: .vulnerabilitySummary.critical}' docs/daily_recommendations_detail.json
```

### Show CVEs with available fixes

```bash
jq '.images[] | select(.name | contains("python:3.12")) |
  .vulnerabilities[] | select((.fixedVersion // "") != "") |
  {id, severity, packageName, fixedVersion}' docs/daily_recommendations_detail.json
```

### Count system packages per image

```bash
jq '.images[] | {name, packages: (.systemPackages | length)} |
  select(.packages > 0)' docs/daily_recommendations_detail.json
```

### Find images with the most vulnerabilities

```bash
jq '[.images[] | {name, total: .vulnerabilitySummary.total}] |
  sort_by(-.total) | .[0:10]' docs/daily_recommendations_detail.json
```

### List all HIGH/CRITICAL CVEs across all images

```bash
jq '[.images[] as $img | $img.vulnerabilities[] |
  select(.severity == "CRITICAL" or .severity == "HIGH") |
  {image: $img.name, id, severity, cvssScore, packageName, fixedVersion}] |
  sort_by(-.cvssScore)' docs/daily_recommendations_detail.json
```

## Git LFS

The detailed report can be large (several MB for hundreds of images). The `docs/*_detail.json` pattern is tracked via Git LFS in `.gitattributes` to avoid repository bloat.

## Schema Versioning

The report includes a `schemaVersion` field (currently `1`) to support forward-compatible changes. Downstream consumers should check this field to detect breaking schema changes.
