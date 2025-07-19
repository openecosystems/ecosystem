.PHONY: generate

projectRoot := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export projectRoot
repoRoot := $(abspath ${projectRoot}/../../../../../)
protoRoot := $(abspath ${repoRoot}/proto)
version := v1.9.4
PATH_TO_CHECK_FILE = ${projectRoot}/ca/bin/${version}

clean:
	rm -rf ${projectRoot}/ca/bin

# Check if file exists
ifeq ($(wildcard $(PATH_TO_CHECK_FILE)),)
FILE_EXISTS := false
else
FILE_EXISTS := true
endif

# Define target conditionally
ifeq ($(FILE_EXISTS),true)
generate:
	@echo "Skip Downloading since already on filesystem"
else
generate:
	@echo "Downloading binaries"
	@mkdir -p ${projectRoot}/ca/bin
	@cd ${projectRoot}/ca/bin && curl -OL "https://github.com/slackhq/nebula/releases/download/${version}/nebula-darwin.zip"
	@cd ${projectRoot}/ca/bin && unzip -o nebula-darwin.zip -d "nebula-darwin"
	@cd ${projectRoot}/ca/bin && curl -OL "https://github.com/slackhq/nebula/releases/download/${version}/nebula-linux-amd64.tar.gz"
	@cd ${projectRoot}/ca/bin && mkdir -p nebula-linux-amd64 && tar -xf nebula-linux-amd64.tar.gz --directory nebula-linux-amd64
	@cd ${projectRoot}/ca/bin && curl -OL "https://github.com/slackhq/nebula/releases/download/${version}/nebula-linux-arm64.tar.gz"
	@cd ${projectRoot}/ca/bin && mkdir -p nebula-linux-arm64 && tar -xf nebula-linux-arm64.tar.gz --directory nebula-linux-arm64
	@cd ${projectRoot}/ca/bin && curl -OL "https://github.com/slackhq/nebula/releases/download/${version}/nebula-windows-arm64.zip"
	@cd ${projectRoot}/ca/bin && unzip -o nebula-windows-arm64.zip -d "nebula-windows-arm64"
	@cd ${projectRoot}/ca/bin && curl -OL "https://github.com/slackhq/nebula/releases/download/${version}/nebula-windows-amd64.zip"
	@cd ${projectRoot}/ca/bin && unzip -o nebula-windows-amd64.zip -d "nebula-windows-amd64"
	@touch ${projectRoot}/ca/bin/${version}
endif
