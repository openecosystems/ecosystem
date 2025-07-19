

nx release --first-release -p go-protobuf-sdk-v2beta --dry-run


# Step 1: Run NX to Tag and version
nx release --first-release -p go-protobuf-sdk-v2beta

# Step 2: Push the tag and version bump
git push --follow-tags

# Step 3: Run GoReleaser
# Ensure Git is not in a "dirty" state
nx run go-protobuf-sdk-v2beta:release


nx release --first-release -p go-protobuf-sdk-v2beta


nx release --first-release -p go-sdk-v2beta --skip-publish --dry-run

nx release publish --tag="nightly" --dry-run --group="sdks-typescript"