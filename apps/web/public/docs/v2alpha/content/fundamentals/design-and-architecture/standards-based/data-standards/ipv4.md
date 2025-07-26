---
title: IPv4
aliases:
    - IPv4
pcx_content_type: definition
summary: >-
    Internet Protocol v4 (IPv4) is a full specification of how communication should work across networks.
hidden: true
has_more: true
todo: Find sources
links_to:
    - /fundamentals/glossary/cidr-block
---

# IPv4

Internet Protocol v4 (IPv4) is a full specification of how communication should work across networks.

## Addressing

A subset of this specification is `Addressing`. Specific to `IPv4` is that an address consists of 32 bits or 4 bytes. It's usually written down like `x.y.z.a`, where each letter can have a value between 0 and 255.

The `Address` is also used for routing packets across the network. That means you cannot simply choose any IP address (similar to that you cannot have a house with a zip code `91000` in New York).

Some network ranges are reserved. Network ranges can be specified using the [CIDR Block](/fundamentals/glossary/cidr-block) notation.

Some special ranges (there are more):

-   0.0.0.0/8: The current network.
-   10.0.0.0/8: Private network.
-   127.0.0.0/8: Local host.
-   172.16.0.0/12: Private network.
-   192.0.0.0/24: Private network.
-   192.168.0.0/16: Private network. Usually home networks.
-   255.255.255.255/32: Broadcast address.
