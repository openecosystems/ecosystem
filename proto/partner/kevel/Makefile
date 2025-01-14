.PHONY: generate

projectRoot := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export projectRoot
repoRoot := $(abspath ${projectRoot}/../../../)
protoRoot := $(abspath ${repoRoot}/proto)
version := v1.0.16
versionChannel := v1
PATH_TO_CHECK_FILE = ${projectRoot}/bin/cache

clean:
	rm -rf ${projectRoot}/bin

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
generate: download
	gnostic --pb-out=. ./bin/decision/decision.yaml --resolve-refs && \
	gnostic-grpc -input ./bin/decision/decision.pb -output ./advertisement/${versionChannel}

download:
	@echo "Downloading binaries"
	@mkdir -p ${projectRoot}/bin
	@cd ${projectRoot}/bin && curl -L "https://github.com/adzerk/adzerk-api-specification/archive/refs/tags/${version}.tar.gz" | tar zx --strip-components=1 --directory .
	@echo "${version}" > ${projectRoot}/bin/cache
	@mv ${projectRoot}/bin/decision/openapi-3.yaml ${projectRoot}/bin/decision/decision.yaml

endif