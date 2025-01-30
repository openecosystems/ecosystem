
# Build
go build -o oeco
go build -o $GOPATH/bin/oeco
GOOS=windows go build -o oeco.exe

# Publish Dry Run
goreleaser --snapshot --skip-publish --rm-dist

# Publish
goreleaser release --skip-publish


oeco configuration createConfiguration --request '{"parent_id": "123"}' -m "spec_context.organization_slug,configuration.id" --validate-only=true
