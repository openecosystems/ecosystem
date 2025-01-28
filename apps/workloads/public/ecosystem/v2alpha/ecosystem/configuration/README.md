# configuration

==================

## Curl/HTTPie or simple network semantics
http localhost:6477/v2/configurations \
Content-Type:application/json \
x-spec-workspace-slug:workspace123 \
x-spec-organization-slug:organization123 \
x-spec-workspace-jan:JURISDICTION_USA \
parent_id=123

## Field Mask Support
curl -X GET \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
http://localhost:6477/v2/configurations/123 | jq .

curl -X GET \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-fieldmask: spec_context.organization_slug,configuration.id,configuration.created_at" \
http://localhost:6477/v2/configurations/123 | jq .

# Latency Test
cd proto
ghz -c 10 -n 100 --insecure --protoset <(buf build -o -) \
--call platform.configuration.v2alpha.ConfigurationService/GetConfiguration \
-m '{"x-spec-organization-slug": "organization123", "x-spec-workspace-slug": "workspace123"}' \
localhost:6477

## HTTP2 and Connect
cd /proto
buf curl --protocol connect --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/GetConfiguration | jq .

# HTTP1.1 Connect
cd /proto
buf curl --protocol connect --verbose \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/GetConfiguration | jq .

## HTTP2 and GRPC
cd /proto
buf curl --protocol grpc --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/GetConfiguration | jq .

## HTTP2 and GRPC Web
cd /proto
buf curl --protocol grpcweb --verbose --http2-prior-knowledge \
--schema public \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"id": "123"}' \
http://localhost:6477/platform.configuration.v2alpha.ConfigurationService/GetConfiguration | jq .

## Options
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


## 

http http://api.demo-1.oeco.cloud:6477/v2/configurations \
Content-Type:application/json \
x-spec-workspace-slug:workspace123 \
x-spec-organization-slug:organization123 \
x-spec-workspace-jan:JURISDICTION_USA \
parent_id=123

curl -X GET \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
http://api.demo-1.oeco.cloud:6477/v2/configurations/123 | jq .