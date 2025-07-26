# Build

go build -o oeco
go build -o $GOPATH/bin/oeco
GOOS=windows go build -o oeco.exe

# Publish Dry Run

goreleaser --snapshot --skip-publish --rm-dist

# Publish

goreleaser release --skip-publish

oeco configuration createConfiguration --request '{"parent_id": "123"}' -m "spec_context.organization_slug,configuration.id" --validate-only=true

# Debugging

## Local Debugging

go install github.com/google/gops@latest
go build -gcflags="all=-N -l" -o $GOPATH/bin/oeco

Goland: Run -> Attach to Process

## Remote Debugging

https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#step-3-start-the-debugging-process-on-your-local-computer
dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 .
