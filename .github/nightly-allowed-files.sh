# Nightly scan output files — shared allowlist used by:
#   - nightly-scan.yml (inline path validation)
#   - validate-nightly-pr.yml (PR file check)
NIGHTLY_ALLOWED_FILES=(
  "azure_linux_images.db"
  "docs/daily_recommendations.md"
  "docs/daily_recommendations.json"
  "docs/daily_recommendations_detail.json"
)
