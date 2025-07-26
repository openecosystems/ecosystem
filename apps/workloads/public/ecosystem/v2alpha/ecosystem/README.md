# Manually Deploy Docker Image

```bash
cpln profile update jeannotcompany --login --org jeannotcompany
export CPLN_PROFILE="jeannotcompany"
cpln image docker-login
```

```bash
nx build workloads-private-event-v2alpha-event-multiplexer
nx container workloads-private-event-v2alpha-event-multiplexer
docker tag workloads/private/event/v2alpha/event-multiplexer:latest jeannotcompany.registry.cpln.io/event-v2alpha-event-multiplexer:JC-123
docker push jeannotcompany.registry.cpln.io/event-v2alpha-event-multiplexer:JC-123
cpln --org jeannotcompany --gvc local-gvc workload force-redeployment event-v2alpha-event-multiplexer
```

cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
localhost:6478 list

```bash
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
--header "x-b3-traceid: 34bc6254cbf6c0ac2236ac7a8999ffa6" \
--header "x-spec-ip: 224.567.324.233" \
--header "x-spec-locale: en_US" \
--header "x-spec-timezone: America/New_York" \
--header "x-spec-device-id: B5372DB0-C21E-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-adv-id: 7A3CBEA0-BDF5-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-manufacturer: Apple" \
--header "x-spec-device-model: iPhone16,2" \
--header "x-spec-device-name: maguro" \
--header "x-spec-device-type: ios" \
--header "x-spec-device-token: ff15bc0c20c4aa6cd50854ff165fd265c838e5405bfeb9571066395b8c9da449" \
--header "x-spec-city: Atlanta" \
--header "x-spec-country: United States" \
--header "x-spec-lat: 40.2964197" \
--header "x-spec-long: -76.9411617" \
--header "x-spec-speed: 0" \
--header "x-spec-cellular: true" \
--header "x-spec-carrier: Verizon" \
--header "x-spec-os-name: iPhone OS" \
--header "x-spec-os-version: 16.1.3" \
--data '{"email": "dimy2@jeannot.company", "first_name": "Dimy2", "last_name": "Jeannot2", "phone_number": "770"}' \
http://localhost:6477/v1/communication/preference-center | jq .

curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-validate-only: true" \
--data '{"email": "dimy2@jeannot.company", "first_name": "Dimy", "last_name": "Jeannot", "phone_number": "770"}' \
http://localhost:6477/v1beta/communication/preference-center
```

```bash

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
--header "x-b3-traceid: 34bc6254cbf6c0ac2236ac7a8999ffa6" \
--header "x-spec-ip: 224.567.324.233" \
--header "x-spec-locale: en_US" \
--header "x-spec-timezone: America/New_York" \
--header "x-spec-device-id: B5372DB0-C21E-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-adv-id: 7A3CBEA0-BDF5-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-manufacturer: Apple" \
--header "x-spec-device-model: iPhone16,2" \
--header "x-spec-device-name: maguro" \
--header "x-spec-device-type: ios" \
--header "x-spec-device-token: ff15bc0c20c4aa6cd50854ff165fd265c838e5405bfeb9571066395b8c9da449" \
--header "x-spec-city: Atlanta" \
--header "x-spec-country: United States" \
--header "x-spec-lat: 40.2964197" \
--header "x-spec-long: -76.9411617" \
--header "x-spec-speed: 0" \
--header "x-spec-cellular: true" \
--header "x-spec-carrier: Verizon" \
--header "x-spec-os-name: iPhone OS" \
--header "x-spec-os-version: 16.1.3" \
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
--header "x-b3-traceid: 34bc6254cbf6c0ac2236ac7a8999ffa6" \
--header "x-spec-ip: 224.567.324.233" \
--header "x-spec-locale: en_US" \
--header "x-spec-timezone: America/New_York" \
--header "x-spec-device-id: B5372DB0-C21E-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-adv-id: 7A3CBEA0-BDF5-11E4-8DFC-AA07A5B093DB" \
--header "x-spec-device-manufacturer: Apple" \
--header "x-spec-device-model: iPhone16,2" \
--header "x-spec-device-name: maguro" \
--header "x-spec-device-type: ios" \
--header "x-spec-device-token: ff15bc0c20c4aa6cd50854ff165fd265c838e5405bfeb9571066395b8c9da449" \
--header "x-spec-city: Atlanta" \
--header "x-spec-country: United States" \
--header "x-spec-lat: 40.2964197" \
--header "x-spec-long: -76.9411617" \
--header "x-spec-speed: 0" \
--header "x-spec-cellular: true" \
--header "x-spec-carrier: Verizon" \
--header "x-spec-os-name: iPhone OS" \
--header "x-spec-os-version: 16.1.3" \
http://localhost:6477/v2/configurations/123

```

https://event-v2alpha-event-multiplexer-j1dg37h9q8pk2.cpln.app

cd proto
ghz --insecure --protoset <(buf build -o -) \
-d '{"email": "dimy2@jeannot.company", "first_name": "Dimy", "last_name": "Jeannot", "phone_number": "770"}' \
--call platform.communication.v1alpha.PreferenceCenterService/CreateOrUpdatePreference --total 10000 localhost:6477

curl \
--header "Content-Type: application/json" \
http://localhost:6477/v1/communication/preference-center/options

curl -X POST \
--header "Content-Type: application/json" \
-d '{}' \
http://localhost:6477/v2/event/unsubscribe

curl -X POST \
--header "Content-Type: application/json" \
-d '{}' \
http://localhost:6477/v2/search

docker run -p 4222:8017 -v /Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany.creds://Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany.creds workloads/private/event/v2alpha/event-multiplexer:latest
docker run -p 4333:8017 -v /Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany.creds://Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany.creds workloads/private/event/v2alpha/event-multiplexer:latest

docker run --name nats-server -p 4222:4222 nats:latest -js
docker run --name leafnode-red -p 4222:4222 -v ./leafnode.conf:/leafnode.conf -v /Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany-red.creds:/ngs.creds nats:latest -c /leafnode.conf
docker run --name leafnode-blue -p 4333:4222 -v ./leafnode.conf:/leafnode.conf -v /Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany-blue.creds:/ngs.creds nats:latest -c /leafnode.conf
docker run --name event-multiplexer-blue -p 4444:8017 -p 6478:6478 -v /Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany-blue.creds://Users/dimyjeannot/workspace/personal/cloud-keys/synadia-ngs-jeannotcompany-blue.creds workloads/private/event/v2alpha/event-multiplexer:latest

curl -X POST --header "Content-Type: application/json" --data '{}' http://localhost:6477/v2/search

curl -X POST --header "Content-Type: application/json" --data '{}' http://localhost:6477/v1beta/communication/preference-center
curl \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug:workspace123" \
--header "x-spec-organization-slug:organization123" \
http://localhost:6477/v1beta/communication/preference-center/options

cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
-H "x-spec-workspace: workspace123" \
-H "x-spec-organization: organization123" \
-d '{}' \
localhost:6477 platform.communication.v1alpha.PreferenceCenterService/GetPreferenceOptions

# Testing CORS

curl 'http://localhost:6477/platform.communication.v1alpha.MarketingEmailService/Subscribe' --verbose \
-X 'OPTIONS' \
-H 'Accept: _/_' \
-H 'Accept-Language: en-US' \
-H 'Access-Control-Request-Headers: content-type' \
-H 'Access-Control-Request-Method: GET' \
-H 'Connection: keep-alive' \
-H 'Origin: http://localhost:8080' \
-H 'Sec-Fetch-Dest: empty' \
-H 'Sec-Fetch-Mode: cors' \
-H 'Sec-Fetch-Site: cross-site' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) FramerElectron/2024.34.1 Chrome/122.0.6261.156 Electron/29.3.0 Safari/537.36'
