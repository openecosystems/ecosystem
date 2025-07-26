---
date_created: 2022-12-11T17:16:25
title: DNS
pcx_content_type: definition
summary: >-
    `Dynamic Name Service`, a network service that can resolve Domain Names to current [IP Addresses](/fundamentals/glossary/ip-address).
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv4
    - /fundamentals/design-and-architecture/standards-based/data-standards/ipv6
    - /fundamentals/glossary/caching
    - /fundamentals/glossary/ip-address
    - /fundamentals/glossary/ttl
---

# DNS

`Dynamic Name Service`, a network service that can resolve Domain Names to current [IP Addresses](/fundamentals/glossary/ip-address).

In order to prevent a `DNS Server` to be hit for each and every request, a `Domain Name` is typically resolved once using a `DNS Server`, and the result is then [Cached](/fundamentals/glossary/caching). For how long depends on the record's [TTL](/fundamentals/glossary/ttl).

There are a few different types of records, each of which have different meanings, among which:

-   `A`, returns a [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) address.
-   `AAAA`, returns an [IPv6](/fundamentals/design-and-architecture/standards-based/data-standards/ipv6) address.
-   `CNAME`, points to another domain name (which results in a new `DNS request` for the name that is returned).
-   `MX`, for Email.
-   `NS`, to point to other `DNS Servers`.

## Sources

-   https://www.ietf.org/rfc/rfc1035.txt
-   <https://en.wikipedia.org/wiki/List_of_DNS_record_types>
