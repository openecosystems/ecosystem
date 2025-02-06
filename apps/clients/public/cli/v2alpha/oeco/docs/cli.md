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
## Multiplexer
- Keep Connectors up to date
  - Before creating a connector
  - Create a connector bucket if it doesn't exist
    - Listens on connectors bucket
    - Fetches all connectors from every bucket every 5 minutes
      - compares time.Now to (time.Now-1h),
        - If falls outside the window, update the connector to inactive
        - If falls within the window, do nothing
## Account Authority Connector
## Certificate Connector
## Configuration Connector
?oeco package


oeco enclave
  - oeco enclave sign
  - oeco enclave find
  - oeco enclave remove
  - oeco enclave attest
oeco context
  - oeco context report
  - oeco context switch
oeco connector
  - oeco connector create
oeco api
  - oeco api configuration create --rest --request='{"name":"example"}'
  - oeco api configuration create --grpc --request='{"name":"example"}'
  - oeco api configuration create --grpc-web --request='{"name":"example"}'
  - oeco api configuration create --graphql --request='{"name":"example"}'
  - oeco api configuration create --connect --request='{"name":"example"}'
oeco ecosystem
  - oeco ecosystem create
  - oeco ecosystem report
oeco dash
  - pages:
    - Ecosystem Overview
    - Connectors Overview in the Ecosystem
    - Requests
    - Wire
    - Logs

