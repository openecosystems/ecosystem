---
title: CIDR Block
aliases:
  - CIDR Block
pcx_content_type: definition
summary: >-
  `Classless Inter-Domain Routing` is a method for allocating [IP Addresses](/fundamentals/glossary/ip-address) and routing. Its goal was to reduce the size of routing tables across the internet, and to slow down the exhaustion of [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) addresses. A `CIDR Block` specifies a subnet.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/ip-address
  - /fundamentals/design-and-architecture/standards-based/data-standards/ipv4
---

# CIDR Block

`Classless Inter-Domain Routing` is a method for allocating [IP Addresses](/fundamentals/glossary/ip-address) and routing. Its goal was to reduce the size of routing tables across the internet, and to slow down the exhaustion of [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) addresses. A `CIDR Block` specifies a subnet.

`CIDR Block` is based on variable-length subnet masking, which gives a finer control of the size of subnets.

`CIDR Blocks` are expressed with the "CIDR Notation", which is an [IP Address](/fundamentals/glossary/ip-address) followed by a suffix indicating the number of bits that are part of the subnet. I.e.: `192.0.2.0/24` means that the first `24` out of the `32` bits of that [IP Address](/fundamentals/glossary/ip-address) are part of the subnet.
