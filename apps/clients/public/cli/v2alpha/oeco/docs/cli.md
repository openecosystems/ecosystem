# Sections
## Enclave (Branch) (Repo)
Secure Enclave
https://github.com/facebookincubator/sks
## API Explorer (Prs)
## Ecosystem (Issues)


# Notes
- Notice we do not keep a centralized list of users. 
- It is up to the CA to keep tabs on whom they have issued certificates.

# PUBLIC MOCK-DATA ECOSYSTEM
## DNS
## Lighthouse
## Edge Router
## Multiplexer
- Keep Connectors up to date
  - Before creating a connector
  - Create a connector bucket if it doesn't exist
    - Listens on connectors bucket
    - Fetches all connectors from every bucket every 5 minutes
      - compares time.Now to (time.Now-1h),
        - If falls outside the window, update the connector to inactive
        - If falls within the window, do nothing
## Certificate Authority Connector
## Certificate Connector
## Configuration Connector


oeco enclave
oeco context
oeco organization
oeco package
oeco connector
oeco api
oeco ecosystem

## oeco context
Manage your distributed account. Manage your certificates using a secure enclave, your policies, configurations.

## oeco organization
In the future, you will be able to create your own organization and a private ecosystem with only members of your choosing.

## oeco connector

This will help you explore the different connectors we support using mock-data based prototyping:
These connectors are fake/mock versions that are identical in structure, 
but whose actual values are randomized and synthetically generated.
This allows you to prototype and test before doing mesh execution.


### Connector Details Tab
NOTES:
Upon creating a connector
- Create a connector in the KV as active and provide hostname

**LEFT:**

Jurisdiction
- USA
- UK
- EU
- Australia
- Israel

Connector Type
- Syft
- Tuva
- Claim 837I (x12)
- Claim 837P EDI (x12)
- CareQuality
- The CMS CCLF (a.k.a. Medicare CCLF)
- CMS LDS
- Athena Health
- Health Gorilla
- Elation
- CommonWell
- Care Quality

Data Classification
- HIPPA
- PCI
- Soc2 Type 2
- FedRamp
- HITRUST
- GDPR

Security Model - Mutually Untrusted
- Secure Multi-Party Computation (sMPC) (i)
- Homomorphic Encryption
- Differential Privacy
- Synthetic Data
- Distributed Learning
- Zero-Knowledge Proofs
- Trusted Execution Environments and Secure Enclaves

Security Model - Trusted
- Care Quality
- Legal Contracts

Register and begin listening? Yes, No

--- Binding to user1.api.oeco.mesh
--- Binding to cancer-dev.api.emory.mesh

**RIGHT:**
Connector API: mesh://djeannot.api.oeco.system/v2alpha/gorilla
Jurisdiction: United States
Type: Syft
Data Classification Policies: HIPPA, PCI, GDPR
Security Model: sMPC, Synthetic Data
Connected Since: Nov 16th, 2024
Last Ping: Nov 16th, 2024 at 15:33:04 EDT
Number of API Calls Handled (Show chart in ANSCII) https://github.com/guptarohit/asciigraph 
Security:
  - Certificate 105 days until expiration


### Connector Requests
**LEFT:**

| Time |      |        | Protocol | Method | Path        | Status |   |   |
|------|------|--------|----------|--------|-------------|--------|---|---|
|      | Mesh | HTTP/2 | GRPC     | GET    | /v2/gorilla | 200    |   |   |

**Sidebar**
Response
Protocol
Emissions
Time
Performance
  Response time
  - Client to server
  - Server to client
  - RTT
Security Validation

### Connector Logs

### Connector Packets
Packet Details:



## oeco api
API Explorer allows you to connect with different areas of the application. 
See TCP traffic, Packet traffic, UDP traffic


## oeco ecosystem
Dynamically see economic system members and their capabilities

NOTES: When running the ecosystem, start a pushpin client to listen for discovery changes
- Tabs
  - Active Connectors
    - When it starts, pull the latest available connectors in the network (fetch all keys from the connector bucket), 
    - then start pushpin for differentials
