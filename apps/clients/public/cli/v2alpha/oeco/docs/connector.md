
## oeco connector

```bash
oeco connector create
```

This will help you explore the different connectors we support using mock-data based prototyping:
These connectors are fake/mock versions that are identical in structure,
but whose actual values are randomized and synthetically generated.
This allows you to prototype and test before doing mesh execution.


### Connector Form
Pre-conditions
- Your current context must be part of an ecosystem
- Your account authority must grant your credential to be part of the "connector group

Post-conditions
- Upon creating a connector
  - Create a connector in the KV as active and provide hostname
- Bind to the channel {connector-name}.api.{ecosystem-name}.mesh

Name
- Give your connector a host name
  - {connector-name}.api.{ecosystem-name}.mesh

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
- CoreCard
- Cencora Financial Services

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

Review
- Connector API: {connector-name}.api.{ecosystem-name}.mesh/v2alpha/syft
- Type: Syft
- Data Classification Policies: 
  - HIPPA 
  - PCI 
  - GDPR
- Security Model: 
  - sMPC
  - Synthetic Data

Register and begin listening? Yes, No




