.PHONY: generate

projectRoot := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export projectRoot
repoRoot := $(abspath ${projectRoot}/../../../)
protoRoot := $(abspath ${repoRoot}/proto)
version := v1.0.16
versionChannel := v2
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
generate:
	gnostic --pb-out=. ./experience/assets-delivery.json --resolve-refs && \
	gnostic-grpc -input ./experience/assets-delivery.pb -output ./experience/${versionChannel}

endif
