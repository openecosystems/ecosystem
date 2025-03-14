.PHONY: generate

projectRoot := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export projectRoot
repoRoot := $(abspath ${projectRoot}/../../../../)
protoRoot := $(abspath ${repoRoot}/proto)
version := 1.40.1
PATH_TO_CHECK_FILE = ${projectRoot}/bin/${version}

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
	@echo "Downloading binaries"
	@mkdir -p ${projectRoot}/bin
	@cd ${projectRoot}/bin && curl -OL "https://github.com/fastly/pushpin/releases/download/v${version}/pushpin-${version}.tar.bz2"
	@cd ${projectRoot}/bin && mkdir -p pushpin && tar -xf pushpin-${version}.tar.bz2 --directory pushpin
	@touch ${projectRoot}/bin/${version}
endif
