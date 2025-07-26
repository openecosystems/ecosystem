---
date_created: 2022-12-11T17:01:41
title: HTTPS
aliases:
    - HTTPS
pcx_content_type: definition
summary: >-
    Hypertext Transfer Protocol Secure (HTTPS) is an extension of the Hypertext Transfer Protocol ([HTTP](/fundamentals/design-and-architecture/standards-based/data-standards/#http)). It is used for secure communication over a computer network, and is widely used on the Internet
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/#http
    - /fundamentals/design-and-architecture/standards-based/data-standards/http
    - /fundamentals/design-and-architecture/standards-based/data-standards/ssl
    - /fundamentals/design-and-architecture/standards-based/data-standards/tls
    - /fundamentals/glossary/certificate
    - /fundamentals/glossary/man-in-the-middle-attack
---

# HTTPS

Hypertext Transfer Protocol Secure (HTTPS) is an extension of the Hypertext Transfer Protocol (HTTP). It is used for secure communication over a computer network, and is widely used on the Internet

`HTTPS` is the same as [HTTP](/fundamentals/design-and-architecture/standards-based/data-standards/http), but then on top of [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) (or better: [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls)).

One extra feature that is added besides [TLS](/fundamentals/design-and-architecture/standards-based/data-standards/tls), is that servers and/or clients need to provide [Certificates](/fundamentals/glossary/certificate) that can be used to unambiguously verify that a connection was established with an intended party, and not with a [Man in the Middle](/fundamentals/glossary/man-in-the-middle-attack).

## Sources

-   https://en.wikipedia.org/wiki/HTTPS
