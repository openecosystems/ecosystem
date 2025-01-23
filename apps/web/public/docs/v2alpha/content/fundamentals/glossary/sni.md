---
title: SNI
aliases:
  - SNI
  - Server Name Indication (SNI)
  - Server Name Indication
pcx_content_type: definition
summary: >-
  Server Name Indication (SNI) allows a server to host multiple TLS Certificates for multiple websites using a single [IP address](/fundamentals/glossary/ip-address). `SNI` adds the website hostname in the [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls) handshake to inform the server which website to present when using shared IPs. Open Ecosystems uses `SNI` for all Universal [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) certificates.
hidden: true
has_more: false
links_to:
  - /fundamentals/glossary/ip-address
  - /fundamentals/design-and-architecture/standards-based/data-standards/tls
  - /fundamentals/design-and-architecture/standards-based/data-standards/ssl
---

# Server Name Indication (SNI)

Server Name Indication (SNI) allows a server to host multiple TLS Certificates for multiple websites using a single [IP address](/fundamentals/glossary/ip-address). `SNI` adds the website hostname in the [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls) handshake to inform the server which website to present when using shared IPs. Open Ecosystems uses `SNI` for all Universal [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) certificates.
