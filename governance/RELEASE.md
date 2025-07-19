# Prerequisites
- Ensure Git is not in a "dirty" state

# Initial Release
nx release --first-release -p go-protobuf-sdk-v2beta --dry-run
nx release --first-release -p go-sdk-v2beta --dry-run

# Subsequent Releases


# 0. Fail if working directory is dirty
git diff --exit-code || (echo "Uncommitted changes. Aborting." && exit 1)

# 1. Bump version & tag (writes to go.mod, CHANGELOG, etc.)
nx release --group="sdks-public" --yes

# 2. Commit and push changes and tag
git push origin HEAD --follow-tags

# 3. Checkout the tag (ensure GoReleaser sees the correct commit)
TAG=$(git describe --tags --abbrev=0)
git checkout $TAG

# 4. Run GoReleaser from the tagged commit
nx run-many -t distribute --projects=tag:type:sdk


