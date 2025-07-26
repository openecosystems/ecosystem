---
title: SSL
aliases:
    - SSL
pcx_content_type: definition
summary: >-
    Secure Sockets Layer (SSL) was a widely used cryptographic protocol for providing data security for Internet communications. SSL was superseded by [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls); however, most people still refer to Internet cryptographic protocols as SSL.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/tls
---

# Secure Sockets Layer (SSL)

Secure Sockets Layer (SSL) was a widely used cryptographic protocol for providing data security for Internet communications. SSL was superseded by [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls); however, most people still refer to Internet cryptographic protocols as SSL.

Both `SSL` and [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls) a layer of security on top of a Network Protocol.

There were three main versions of this protocol:

-   SSL 1.0
-   SSL 2.0
-   SSL 3.0

`SSL 1.0` was known to be insecure, and therefore never published.

`SSL 2.0` was in use from 1995. It used the rather weak MD5 Hashing algorithm and was vulnerable to a number of attack vectors, among which Man in the Middle, that led to its decommissioning in 2011.

`SSL 3.0` was a complete redesign of `SSL`. In 2014 it was found to be vulnerable to the `POODLE` attack, which affected all block ciphers, and so it was deprecated in 2015, when it was replaced by [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls).
