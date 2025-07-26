---
title: gRPC
aliases:
    - gRPC
date_created: 2022-12-11T17:06:49
author: Guillaume Hanique
pcx_content_type: definition
summary: >-
    gRPC uses a binary format to encode data, which is *much* faster, cheaper, and compact than many other Message Protocols.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/html
    - /fundamentals/design-and-architecture/standards-based/data-standards/json
    - /fundamentals/design-and-architecture/standards-based/data-standards/xml
    - /fundamentals/glossary/message-protocol
    - /fundamentals/glossary/performance
---

# gRPC

`gRPC` uses a binary format to encode data, which is _much_ more [Performant](/fundamentals/glossary/performance) than many other [Message Protocols](/fundamentals/glossary/message-protocol).

The fact that it is binary also implies that it is much harder to consume, but this is not the case. The message format can be declared using Proto, and there are implementations for almost all conceivable Programming Languages to use that definition to generate code to serialize and deserialize the message.

That leaves the downside that one `gRPC` cannot be read messages in transit, and see its content, where both [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) and [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json) would be humanly readable to some extend.

Another reason why `gRPC` is much faster than most other [Message Protocols](/fundamentals/glossary/message-protocol), is that it is not based on [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html). [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html) adds overhead to communication that makes it easy to understand, but is overhead nevertheless.

## Sources

-   https://grpc.io/
-   https://grpc.github.io/grpc/core/index.html
