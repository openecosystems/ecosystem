# decision

==================

## Curl/HTTPie or simple network semantics
http localhost:6477/api/v2 \
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
--schema partner \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"decision_request": {"url": "htttp://url.com"}}' \
http://localhost:6477/kevel.advertisement.v1.DecisionService/GetDecisions

# HTTP1.1 Connect
cd /proto
buf curl --protocol connect --verbose \
--schema partner \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"decision_request": {"url": "htttp://url.com"}}' \
http://localhost:6477/kevel.advertisement.v1.DecisionService/GetDecisions

## HTTP2 and GRPC
cd /proto
buf curl --protocol grpc --verbose --http2-prior-knowledge \
--schema partner \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"decision_request": {"url": "htttp://url.com"}}' \
http://localhost:6477/kevel.advertisement.v1.DecisionService/GetDecisions

## HTTP2 and GRPC Web
cd /proto
buf curl --protocol grpcweb --verbose --http2-prior-knowledge \
--schema partner \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--data '{"decision_request": {"url": "htttp://url.com"}}' \
http://localhost:6477/kevel.advertisement.v1.DecisionService/GetDecisions
