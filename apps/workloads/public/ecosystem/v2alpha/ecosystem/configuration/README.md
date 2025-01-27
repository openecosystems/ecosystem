# configuration

==================

protoset <(buf build -o -)

# Latency Test
ghz -c 100 -n 1 --insecure \
--call platform.configuration.v2alpha.ConfigurationService/CreateConfiguration \
-m '{"auth-user-id": "djeannot"}' \
-d '{"spec_context": {"organization_slug" : "test-organization", "workspace_slug": "test-workspace"},"name": "test", "slug": "success", "short_description": "test short description", "description": "test description"}' \
localhost:6477

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

curl -X POST \
--header "Content-Type: application/json" \
--header 'Connect-Protocol-Version: 1' \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: corporate" \
--data '{"parent_id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

curl --http2 --verbose \
--header "Content-Type: application/json" \
--header 'Connect-Protocol-Version: 1' \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--data '{"id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/GetConfiguration


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
http://localhost:6477/v2/configurations

cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
-H "x-spec-workspace: workspace123" \
-H "x-spec-organization: organization123" \
-d '{"parent_id": "123"}' \
localhost:6477 platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

buf curl --protocol connect --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"parent_id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

buf curl --protocol connect --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-apikey: 12345678" \
--header "x-spec-debug: true" \
--data '{"parent_id": "123"}' \
http://localhost:7676/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration


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
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/v2/configurations

curl -X POST \
--header "Content-Type: application/json" \
--header 'Connect-Protocol-Version: 1' \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

cd /proto
buf curl --protocol connect --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

cd /proto
buf curl --protocol connect --verbose \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

cd /proto
buf curl --protocol grpc --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

cd /proto
buf curl --protocol grpcweb --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"parent_id": "123"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

curl 'http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration' \
-X 'OPTIONS' \
-H 'Accept: */*' \
-H 'Accept-Language: en-US,en;q=0.9,ru;q=0.8' \
-H 'Access-Control-Request-Headers: connect-protocol-version,content-type,x-spec-organization-slug,x-spec-workspace-slug' \
-H 'Access-Control-Request-Method: POST' \
-H 'Cache-Control: no-cache' \
-H 'Connection: keep-alive' \
-H 'Origin: http://localhost:4200' \
-H 'Pragma: no-cache' \
-H 'Referer: http://localhost:4200/' \
-H 'Sec-Fetch-Mode: cors' \
-H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36' \
--insecure