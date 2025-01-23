---
date_created: 2022-12-11T17:09:14
author: Guillaume Hanique
title: URL
pcx_content_type: definition
summary: >-
  A Uniform Resource Locator (URL), colloquially termed a web address, is reference to a web resource that specifies its location on a computer network, and a mechanism for retrieving it.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/uri
  - /fundamentals/glossary/fqdn
  - /fundamentals/glossary/get-parameters
  - /fundamentals/glossary/ip-address
  - /fundamentals/glossary/transport-protocol
---

# URL

A Uniform Resource Locator (URL), colloquially termed a web address, is reference to a web resource that specifies its location on a computer network, and a mechanism for retrieving it. A URL is a specific type of URI that contains both the **location** of the item and the **identifier** of the item.

A `url`, i.e., "http://servername:1234/resourcename?param1=value1&param2=value2" consists of the following components:

- [Transport Protocol](/fundamentals/glossary/transport-protocol), i.e. "http", "https", "file", etc, followed by "://"
- The `servername`. This can be an [IP Address](/fundamentals/glossary/ip-address), hostname, or [FQDN](/fundamentals/glossary/fqdn).
- The `port` that the "Web Application" that serves the resource is listening on.
- The `resourcename`, which points to the resource on that server.
- A `?` followed by a list of [Get Parameters](/fundamentals/glossary/get-parameters), which are separated with a `&`.

Just like [URIs](/fundamentals/design-and-architecture/standards-based/data-standards/uri) `URLs` can only contain certain characters.

## Sources

- https://en.wikipedia.org/wiki/URL
- https://datatracker.ietf.org/doc/html/rfc1738
