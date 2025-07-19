# Prerequisites
- Ensure Git is not in a "dirty" state

# Initial Release
nx release --first-release -p oeco-sdk-v2beta --dry-run

# Subsequent Releases

# 0. Fail if working directory is dirty
git diff --exit-code || (echo "Uncommitted changes. Aborting." && exit 1)

# 1. Bump version & tag (writes to go.mod, CHANGELOG, etc.)
nx release --group="sdks-public" --yes
