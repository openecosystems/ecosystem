

## Account Signature, Validation, and Association
create an iam user service. when you create an account you get a cert and key. upload crt to be signed by ca and get a unique hostname and ip address. user.oeco.mesh
```mermaid
sequenceDiagram
    autonumber
    participant c as Client
    participant M as Multiplexer
    participant C as IAM Account Connector
    links C: {"Details": "https://docs.openecosystems.com/connector"}
    links M: {"Details": "https://docs.openecosystems.com/multiplexer"}

  Note left of c: On local machine, <br/> get an unsigned <br/> client certificate <br/> and private key
    c->>M: Upload certificate to be <br/> signed and validated <br/>by an Account Authority
    M->>C: Finds available <br/> IAM Account connector, <br/> routes traffic <br/>with codec requested.
    Note right of C: I am available. <br/> I will handle this.
    C-->>M: Validates certificate, <br/> signs, and adds a hostname, <br/> and allocates an <br/> IP address for <br/> mesh addressibility
    M-->>c: Byte buffered response <br/> using the security <br/> model codec
```

## Connector Signature, Validation, and Association
create a connector account. upload crt to be signed by central ca server. ca server provisions an IP address on the network. and a unique hostname and ip. api.ecosystem.mesh/v2alpha/connector. this gets converted to the nats channel: api.ecosystem.b2alpha.connector or mesh.ecosystem.api.v2alha.connector
store this hostname in KV. store ip address in KV. ip is key, value is host. host is key, value is ip
we need a single key to find the next available IP address. ideally not sequentially.
this will auto register with dns.
```mermaid
sequenceDiagram
    autonumber
    participant U as Client
    participant M as Multiplexer
    participant C as IAM ConnectorAccount Connector
    participant KV as KV Storage
    Note left of U: On local machine, <br/> get an unsigned connector certificate <br/>and private key
    U->>M: Upload certificate to be <br/> signed and validated <br/>by an Account Authority
    M->>C: Finds available <br/> IAM ConnectorAccount <br/>connector, routes traffic <br/>with codec requested.
    Note right of C: I am available. <br/> I will handle this.
    C->>KV: Checks for available <br/>IP Address and <br/>uniqueness <br/> of hostname
    KV-->>C: Returns available <br/>IP address
    C-->>M: Validates certificate, <br/> signs, adds a hostname, <br/> and IP address for <br/>mesh addressability
    M-->>U: Byte buffered response <br/> using the security <br/> model codec
```

## Traffic Flow Across Ports
For the Edge Router, there is a single port:
- 443/tcp
  - All internet facing traffic goes through this choke point

For the Event Multiplexer, there are three ports:
- 4222/tcp (Nats over mesh socket)
    - Nats is accessible over the mesh overlay network
- 4242/udp (Nebula)
- 6477/tcp (gRPC/REST/Connect)
    - Serves as API gateway for both internet and mesh traffic
    - Available over the public internet only through the event-router
    - Secured by mTLS


```mermaid
sequenceDiagram
    autonumber
    participant c as Client
    participant R as Edge Router
    participant M as Multiplexer
    participant C as Connector

    c->>R: Makes a public-facing request
    R->>M: Routes request to Multiplexer
    M->>C: Finds available <br/> Connector, <br/> and routes traffic <br/>with requested codec.
    C-->>M: Handles request
    M-->>R: Byte buffered response <br/> using the security <br/> model codec
    R-->>c: Byte buffered response <br/> using the security <br/> model codec

```

## Traffic Flow Across Meshes
Once client connection to mesh is established, all traffic no longer goes through edge-router (Ideally it does, revisit this)

Now, .mesh urls are now available. For example: api.ecosystem.mesh/v2alpha/connector

There are patient, private, and public ecosystems.
An organization can create either a private or public ecosystem.
By default, Open Economic Systems creates a public ecosystem.
- api.oeco.mesh/v2alpha/connector

