---
title: DNS
pcx_content_type: definition
summary: >-
    The Domain Name System (DNS) is the phonebook of the Internet. DNS translates domain names to IP addresses.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv4
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv6
    - /fundamentals/glossary/caching
    - /fundamentals/glossary/ttl
---

# DNS

<!-- This document is an original CloudFlare Document from which the cloudflare links are removed. -->

The Domain Name System (DNS) is the phonebook of the Internet. DNS translates domain names to IP addresses (see [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) and [IPv6](/fundamentals/design-and-architecture/standards-based/data-standards/ipv6)).

In order to prevent a `DNS Server` to be hit for each and every request, a `Domain Name` is typically resolved once using a `DNS Server`, and the result is then [Cached](/fundamentals/glossary/caching). For how long depends on the record's [TTL](/fundamentals/glossary/ttl).
