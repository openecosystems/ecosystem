# configuration

==================


# Run tests
go test -run TestCreateConfigurationHandler
go test -run TestUpdateConfigurationHandler
go test -run TestDeleteConfigurationHandler

# Run
go run main.go

# List all services
grpcurl -plaintext localhost:6568 list

# List methods
grpcurl -plaintext localhost:6568 list platform.configuration.v2alpha.ConfigurationService

# Create
grpcurl \
--plaintext -v \
-d '{"parent_id": "123"}' \
-rpc-header ctx-organization-slug:test-organization \
-rpc-header ctx-workspace-slug:test-workspace \
-rpc-header ctx-locale:en_US \
-rpc-header auth-user-id:djeannot \
-rpc-header ctx-timezone:"America/New_York" \
localhost:6568 platform.configuration.v2alpha.ConfigurationService/CreateConfiguration

# Latency Test
ghz -c 100 -n 1 --insecure \
--call platform.configuration.v2alpha.ConfigurationService/CreateConfiguration \
-m '{"auth-user-id": "djeannot"}' \
-d '{"spec_context": {"organization_slug" : "test-organization", "workspace_slug": "test-workspace"},"name": "test", "slug": "success", "short_description": "test short description", "description": "test description"}' \
localhost:6568

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
http://144.202.125.179:6477/v2/configurations


curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--data '{"parent_id": "123"}' \
http://144.202.125.179:6477/v2/configurations
