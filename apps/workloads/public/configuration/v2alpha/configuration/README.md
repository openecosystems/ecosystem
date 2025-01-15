# configuration

==================

protoset <(buf build -o -)

# gRPC Call
buf curl --protocol connect --http2-prior-knowledge \
--schema <(buf build -o -) \
--header "x-spec-apikey: 12345678" \
--data '{"parent_id": "123"}' \
http://api.dev-1.oeco.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

buf curl --protocol connect --http2-prior-knowledge \
--schema ./public/platform/configuration/v2alpha/configuration.proto \
--header "x-spec-apikey: 12345678" \
--data '{"parent_id": "123"}' \
http://api.dev-1.oeco.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

grpcurl \
-protoset <(buf build -o -) -plaintext \
-rpc-header "x-spec-apikey: 12345678" \
-rpc-header "x-spec-workspace-slug: workspace123" \
-rpc-header "x-spec-organization-slug: organization123" \
-rpc-header "x-spec-workspace-jan: JURISDICTION_USA" \
-d '{"parent_id": "123"}' \
api.dev-1.na-us-1.oeco.cloud:6477 platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

grpcurl \
-protoset <(buf build -o -) -plaintext \
-rpc-header "x-spec-apikey: 12345678" \
-rpc-header "x-spec-workspace-slug: workspace123" \
-rpc-header "x-spec-organization-slug: organization123" \
-rpc-header "x-spec-workspace-jan: JURISDICTION_USA" \
-d '{"parent_id": "123"}' \
api.dev-1.oeco.cloud:443 platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

# Latency Test
ghz -c 100 -n 1 --insecure \
--call platform.configuration.v2alpha.ConfigurationService/CreateConfiguration \
-m '{"auth-user-id": "djeannot"}' \
-d '{"spec_context": {"organization_slug" : "test-organization", "workspace_slug": "test-workspace"},"name": "test", "slug": "success", "short_description": "test short description", "description": "test description"}' \
localhost:6568


# When using the Edge Router, you can run the following
curl -X GET -vv --http1.1 \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
https://api.dev-1.oeco.cloud/v2/configurations/123 | jq .

curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
--data '{"parent_id": "123"}' \
https://api.dev-1.oeco.cloud/v2/configurations


# When running on localhost, you need to add the headers that normally get added from the edge router and other upstream systems
curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: corporate" \
--data '{"parent_id": "123"}' \
http://localhost:6477/v2/configurations | jq .

curl -X GET \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: corporate" \
--header "x-spec-fieldmask: spec_context.organization_slug,configuration.id" \
http://localhost:6477/v2/configurations/123 | jq .

# When calling from the mesh network
curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: corporate" \
--data '{"parent_id": "123"}' \
http://192.168.100.9:6477/v2/configurations

# When debugging the multiplexer directly
curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--data '{"parent_id": "123"}' \
http://144.202.125.179:6477/v2/configurations
