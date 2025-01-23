---
title: TLS
aliases:
  - TLS
pcx_content_type: definition
summary: >-
  Transport Layer Security (TLS) is a security protocol that replaces [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) for data privacy and Internet communication security. TLS encrypts communications between web applications and servers such as between a visitor’s browser loading a website.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/ssl
  - /fundamentals/glossary/encryption
  - /fundamentals/glossary/pfs
  - /fundamentals/glossary/symmetric-encryption
  - /fundamentals/glossary/hashing
---

# Transport Layer Security (TLS)

Transport Layer Security (TLS) is a security protocol that replaces [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) for data privacy and Internet communication security. `TLS` encrypts communications between web applications and servers such as between a visitor’s browser loading a website.

`TLS` is used for many Network Protocols, but of those HTTPS is the most publicly visible.

## Handshake

`TLS` starts with a handshake, where the client and the server negotiate how they are going to set up the [Encryption](/fundamentals/glossary/encryption). When that handshake is complete, the connection is secure.

Because the handshake by definition must happen unencrypted, and the intent is to have an encrypted connection, part of the connection is encrypted, and another part is not. The easiest way to implement this is by using _two_ port numbers: one for unencrypted communication, and another one for encrypted communication. Many Network Protocols that apply `TLS` use this approach.

Alternatively client and server could make use of a `STARTTLS` request to indicate when the rest of the communication will be encrypted.

`TLS` applies [PFS](/fundamentals/glossary/pfs). Every new connection will use a different [Symmetric Key](/fundamentals/glossary/symmetric-encryption).

## Versions

At the time of writing there are 4 versions of `TLS`:

- TLS 1.0
- TLS 1.1
- TLS 1.2
- TLS 1.3 (the current version)

`TLS 1.0` and `TLS 1.1` saw the light of day in 1995 and 2006 respectively. Both were deprecated in 2021.

`TLS 1.0` had a lot in common with SSL `3.0`, except that it was no longer vulnerable to `POODLE` attacks. `TLS 1.1` was an improvement on `TLS 1.0` in that it had extra protection against some cryptographic attack vectors.

`TLS 1.2` is mostly similar to `TLS 1.1`, with the exception that it moved from MD5 and SHA1 to SHA256. It also included the ability to negotiate during the handshake what [Hashes](/fundamentals/glossary/hashing) and ciphers each party thought acceptable. `TLS 1.2` also supported AES256, where `TLS 1.1` did not. `TLS 1.2` was also modified to never allow a fallback to SSL `2.0`. `TLS 1.2` is not deprecated, but one is encouraged to migrate to `TLS 1.3`.

`TLS 1.3` is the latest version of `TLS`. It dropped support many ciphers and [Hashing](/fundamentals/glossary/hashing) algorithms that we now consider "weak", it dropped support for features like "compression" and "renegotiation", and it added support for more secure ciphers and [Hashing](/fundamentals/glossary/hashing) algorithms, like `ED25519`.

## References

- https://en.wikipedia.org/wiki/Transport_Layer_Security
