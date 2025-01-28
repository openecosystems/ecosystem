
# Manually Deploy Docker Image

```bash
cpln profile update jeannotcompany --login --org jeannotcompany
export CPLN_PROFILE="jeannotcompany"
cpln image docker-login
```

```bash
nx build workloads-public-communications-v1alpha-preference-center
nx container workloads-public-communications-v1alpha-preference-center
docker tag workloads/public/communications/v1alpha/preference-center:latest jeannotcompany.registry.cpln.io/communications-v1alpha-preference-center:JC-123
docker push jeannotcompany.registry.cpln.io/communications-v1alpha-preference-center:JC-123
cpln --org jeannotcompany --gvc local-gvc workload force-redeployment communications-v1alpha-preference-center
```

```bash
curl \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug:workspace123" \
--header "x-spec-organization-slug:organization123" \
http://localhost:6477/v1/communication/preference-center/options | jq .
```

```bash
curl \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
http://localhost:6477/v1/communication/preference-center/dimy | jq .
```

# Get Preference Options
```bash
cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
-H "x-spec-workspace: workspace123" \
-H "x-spec-organization: organization123" \
-d '{}' \
localhost:6477 platform.communication.v1alpha.PreferenceCenterService/GetPreferenceOptions
```

# Create or Update Preferences
```bash

curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
--header "x-spec-validate-only: true" \
--header "x-spec-principal-id: djeannot" \
--header "x-spec-ip: 123.0.0.0" \
--header "x-spec-locale: en-us" \
--header "x-spec-sent-at: 2022-12-10T04:08:31.581Z" \
--header "x-spec-principal-id:12345678" \
--header "x-spec-principal-email: dimy@jeannotfamily.com" \
--header "x-spec-connection-id: sf-corporate" \
--header "x-b3-traceid: 34bc6254cbf6c0ac2236ac7a8999ffa6" \
--header "x-spec-workspace-jan: JURISDICTION_USA" \
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

# Directly to Ingress
curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug:workspace123" \
--header "x-spec-organization-slug:organization123" \
--header "x-spec-validate-only: true" \
--data '{"email": "dimy2@jeannot.company", "first_name": "Dimy2", "last_name": "Jeannot2"}' \
https://api.communication.dev-1.na-us-1.jeannot.company/v1/communication/preference-center | jq .

# Directly to Edge
curl -X POST \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12345678" \
--header "x-spec-validate-only: true" \
--header "x-spec-debug: true" \
--data '{"email": "dimy2@jeannot.company", "first_name": "Dimy2", "last_name": "Jeannot2"}' \
https://api.communication.dev-1.jeannot.company/v1/communication/preference-center | jq .


```

# Get preferences for a user
```bash
curl \
--header "Content-Type: application/json" \
--header "x-spec-workspace-slug: workspace123" \
--header "x-spec-organization-slug: organization123" \
http://localhost:6477/v1/communication/preference-center/7c25c3a0-0e69-4363-bf05-3df64c705320 | jq .
```



cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
-H "x-spec-workspace: workspace123" \
-H "x-spec-organization: organization123" \
-d '{}' \
localhost:6477 platform.communication.v1alpha.PreferenceCenterService/GetPreferenceOptions


cd proto
ghz --insecure --protoset <(buf build -o -) --call platform.audit.v2alpha.AuditService/Search --total 100000 localhost:6476



curl \
--header "Content-Type: application/json" \
--header "x-spec-apikey: 12987312kj13h2i371298312kj3u12i8371298371923" \
http://localhost:6477/v1/communication/preference-center/options | jq .
