---
title: Join Existing Ecosystem
pcx_content_type: overview
weight: 2
---

Join Existing


## Account Signature, Validation, and Association
Use the oeco CLI to create a new unsigned certificate and private key. 
Upload the public key to the Ecosystem Account Authority to be verified and signed at https://api.oeco.cloud

Once signed, the Account Authority will also append a unique hostname and ip address to your certificate.
Once your public key is signed by an Ecosystem Account Authority, you can now connect to the mesh network.

when you create an account you get a cert and key. upload crt to be signed by ca and get a unique hostname and ip address. user.oeco.mesh
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
