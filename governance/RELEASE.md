# Prerequisites
- Ensure Git is not in a "dirty" state

# Initial Release
nx release --first-release -p go-protobuf-sdk-v2beta --dry-run

# Subsequent Releases

# Step 1: Run NX to Tag and version
nx release -p go-protobuf-sdk-v2beta --yes

# Step 2: Push the tag and version bump
git push --follow-tags

# Step 3: Run GoReleaser
nx run go-protobuf-sdk-v2beta:distribute

