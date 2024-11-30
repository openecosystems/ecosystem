.PHONY: generate

projectRoot := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export projectRoot
repoRoot := $(abspath ${projectRoot}/../../../)
protoRoot := $(abspath ${repoRoot}/proto)

# Clean
clean:
	@echo "Cleaning directories" && \
	ls .

# Generate
generate: #clean download_openapi generate_proto

download_openapi:
	curl -o v3/swagger.json "https://raw.githubusercontent.com/twilio/sendgrid-oai/refs/heads/main/spec/json/tsg_mc_contacts_v3.json"

convert_openapi2_to_openapi3:
	curl -o v3/openapi.json "https://converter.swagger.io/api/convert?url=https://raw.githubusercontent.com/twilio/sendgrid-oai/refs/heads/main/spec/json/tsg_mc_contacts_v3.json" -H "Accept: application/json"

generate_proto:
	gnostic --grpc-out=v3 v3/swagger.json && \
	gnostic --pb-out=v3 v3/swagger.json
