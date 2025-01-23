---
date_created: 2022-12-11T17:15:47
title: TCP
aliases:
  - TCP
  - TCP/IP
pcx_content_type: definition
summary: >-
  Where [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/#ipv4) and [IPv6](/fundamentals/design-and-architecture/standards-based/data-standards/#ipv6) specify how packets are sent across a Network, TCP adds the concepts of "Connections" and TCP IP Ports (among others).
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/ipv4
  - /fundamentals/design-and-architecture/standards-based/data-standards/ipv6
  - /fundamentals/design-and-architecture/standards-based/data-standards/http
  - /fundamentals/design-and-architecture/standards-based/data-standards/smtp
---

# TCP

Where [IPv4](/fundamentals/design-and-architecture/standards-based/data-standards/ipv4) and [IPv6](/fundamentals/design-and-architecture/standards-based/data-standards/ipv6) merely specify how packets are sent across a Network, TCP adds the concepts of "Connections" and `TCP Ports` (among others).

A `TCP Port` provides multiplexing, allowing multiple networking applications to run on a single host by assigning `TCP Port` numbers to the various applications. A `TCP Port` can range anywhere between `0` and `65535`. Quite a few port numbers have been reserved, or have a predefined application. I.e.: port `80` is reserved for applications that talk [HTTP](/fundamentals/design-and-architecture/standards-based/data-standards/http), and port `25` is reserved for [SMTP](/fundamentals/design-and-architecture/standards-based/data-standards/smtp).

## Sources

- https://www.ietf.org/rfc/rfc793.txt
