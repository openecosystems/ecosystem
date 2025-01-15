# certificate-authority

==================

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
http://127.0.0.1:6477/v2alpha/cryptography/ca/create

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
--data '{"parent_id": "123", "name": "hello"}' \
http://144.202.125.179:6477/v2alpha/cryptography/ca/create


curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: health-organization" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: corporate" \
--data '{"parent_id": "123", "name": "hello"}' \
http://api.communication.dev-1.na-us-1.jeannot.company:6477/v2alpha/cryptography/ca/create

curl -X POST -vv \
--header "Content-Type: application/json" \
--header 'Origin: https://example.com' \
--header "x-spec-apikey: 12345678" \
--header "x-spec-debug: true" \
--header "x-spec-validate-only: true" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--data '{"parent_id": "123", "name": "hello"}' \
https://api.communication.dev-1.jeannot.company/v2alpha/cryptography/ca/create

curl -X POST -vv \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
--data '{"parent_id": "123", "name": "hello"}' \
https://api.dev-1.oeco.cloud/v2alpha/cryptography/ca/create

curl -X POST -vv \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
--header "x-spec-debug: true" \
--header "x-spec-validate-only: true" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--data '{"parent_id": "123", "name": "hello"}' \
http://api.dev-1.na-us-1.oeco.cloud:6477/v2alpha/cryptography/ca/create


curl -X POST -vv \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
--header "x-spec-debug: true" \
--header "x-spec-validate-only: true" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--data '{"parent_id": "123", "name": "hello"}' \
http://localhost:7676/v2alpha/cryptography/ca/create


curl -v  --request OPTIONS 'https://api.communication.dev-1.jeannot.company/v2alpha/cryptography/ca/create' --header 'Origin: https://example.com' --header 'Access-Control-Request-Method: POST'

