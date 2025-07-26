---
title: DNS Record
aliases:
    - DNS Record
    - DNS Records
pcx_content_type: definition
summary: >-
    DNS Records are instructions that live in authoritative [DNS servers](/fundamentals/glossary/dns-server) and provide information about a domain including what IP address is associated with that domain and how to handle requests for that domain.
hidden: true
has_more: false
links_to:
    - /fundamentals/glossary/dns-server
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv4
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv6
---

<!-- This document is an original CloudFlare Document from which the cloudflare links are removed. -->

# DNS record

DNS Records are instructions that live in authoritative [DNS servers](/fundamentals/glossary/dns-server) and provide information about a domain including what IP address is associated with that domain and how to handle requests for that domain.

There are a few different types of records, each of which have different meanings, among which:

-   `A`, returns a [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) address.
-   `AAAA`, returns an [IPv6](/fundamentals/design-and-architecture/standards-based/data-standards/ipv6) address.
-   `CNAME`, points to another domain name (which results in a new `DNS request` for the name that is returned).
-   `MX`, for Email.
-   `NS`, to point to other `DNS Servers`

## Sources

-   https://en.wikipedia.org/wiki/List_of_DNS_record_types
