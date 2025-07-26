---
id: downstream spec event
title: Internal Spec Overview
---

```json
{
    "version": "",
    "messageId": "022bb90c-bbac-11e4-8dfc-aa07a5b093db", // From channel
    "sentAt": "2022-12-10T04:08:31.581Z", // From client
    "receivedAt": "2022-12-10T04:08:31.909Z",
    "completedAt": "2022-12-10T04:08:32.805Z",
    "specType": "user", // From channel
    "specEventType": "SPEC_EVENT_TYPE_COMMAND", // From channel
    "specEvent": "USER_COMMANDS_CREATE", // From channel
    "principal": {
        "type": "SPEC_PRINCIPAL_TYPE_USER", // authorization filter
        "anonymousId": "072bb90c-bbac-11e4-8dfc-aa07a5b093kj", // authorization filter
        "principalId": "025bb90c-bbac-11e4-8dfc-aa07a5b093nb", // authorization filter
        "principalEmail": "sam@global-company.com", // authorization filter
        "connectionId": "global-saml" // authorization filter
    },
    "spanContext": {
        // From client or proxy or channel
        "traceId": "052bb90c-bbac-11e4-6dfc-aa07a5b093yh",
        "spanId": "072bb90c-bdac-11e4-8dfc-aa07a4b093po",
        "parentSpanId": "",
        "traceFlags": ""
    },
    "context": {
        "organizationSlug": "global-company",
        "workspaceSlug": "local-us-workspace", // From spec context on request
        "workspaceLocation": "usa",
        "ip": "224.567.324.233", // From edge cache, or proxy (read from header)
        "locale": "en_US", // From header ctx-locale or url parameter
        "timezone": "Europe/Amsterdam", // From edge cache
        "userAgent": "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
        "producer": {
            "name": "sf-api-user", // From channel
            "version": "v1.0.1", // From channel
            "build": "324", // From channel
            "namespace": "platform.user.v1" // From channel
        },
        "device": {
            "id": "B5372DB0-C21E-11E4-8DFC-AA07A5B093DB", // From client
            "advertisingId": "7A3CBEA0-BDF5-11E4-8DFC-AA07A5B093DB", // From client
            "manufacturer": "Apple", // From client
            "model": "iPhone12,2", // From client
            "name": "maguro", // From client
            "type": "ios", // From client
            "token": "ff15bc0c20c4aa6cd50854ff165fd265c838e5405bfeb9571066395b8c9da449" // From client
        },
        "location": {
            "city": "San Francisco", // From edge cache
            "country": "United States", // From edge cache
            "latitude": 40.2964197, // From edge cache
            "longitude": -76.9411617, // From edge cache
            "speed": 0 // From edge cache
        },
        "network": {
            "bluetooth": false, // From client
            "cellular": true, // From client
            "wifi": false, // From client
            "carrier": "T-Mobile US" // From client
        },
        "os": {
            "name": "iPhone OS", // From client
            "version": "12.1.3" // From client
        }
    },
    "routineContext": {
        "routineId": "052bb90c-bbac-11e4-6dfc-aa07a5b093yh", // From channel
        "routineData": [] // From channel
    },
    "data": {} // From channel
}
```

---

id: upstream spec event
title: External Spec Overview

---

```json
{
    "version": 1,
    "messageId": "022bb90c-bbac-11e4-8dfc-aa07a5b093db",
    "sentAt": "2022-12-10T04:08:31.581Z",
    "receivedAt": "2022-12-10T04:08:31.909Z",
    "completedAt": "2022-12-10T04:08:32.805Z",
    "specType": "user",
    "specEventType": "SPEC_EVENT_TYPE_COMMAND",
    "specEvent": "USER_COMMANDS_CREATE",
    "context": {
        "organizationSlug": "global-company",
        "workspaceSlug": "local-us-workspace",
        "workspaceLocation": "usa",
        "locale": "en_US",
        "timezone": "Europe/Amsterdam"
    },
    "data": {}
}
```
