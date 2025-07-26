---
date_created: 2022-12-11T17:17:37
title: Protobuf
aliases:
    - Protobuf
    - Protocol Buffer
    - Protocol Buffers
pcx_content_type: definition
summary: >-
    `Protocol Buffers`, or Protobuf, is a free and [Open Source](/fundamentals/glossary/open-source) Cross-Platform data format used to serialize structured data. It is useful in developing programs to communicate with each other over a Network or for storing data. Protobuf is more compact and [Performant](/fundamentals/glossary/performance) than [REST](/fundamentals/design-and-architecture/standards-based/data-standards/rest).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/open-source
    - /fundamentals/glossary/performance
    - /fundamentals/design-and-architecture/standards-based/data-standards/rest
    - /fundamentals/glossary/dsl
    - /fundamentals/glossary/code-generation
    - /fundamentals/glossary/streaming
    - /fundamentals/glossary/api
    - /fundamentals/glossary/messaging
---

# Protobuf

`Protocol Buffers`, or `Protobuf`, is a free and Open Source Cross-Platform data format used to serialize structured data. It is useful in developing programs to communicate with each other over a Network or for storing data. `Protobuf` is more compact and Performant than REST.

## Proto Files

`Protocol Buffers` comes with its own [DSL](/fundamentals/glossary/dsl) to describe the structure of data, called `Proto` files.

There are two `Proto` versions: `proto2` and `proto3`. These two versions are wire-compatible, but the [DSLs](/fundamentals/glossary/dsl) are only compatible to some extend.

## Code Generation

`Protobuf` also comes with its own source [code generator](/fundamentals/glossary/code-generation) (`Protoc`) which can generate source code from `Proto` files into several languages, that can subsequently be used to serialize, deserialize, and [stream](/fundamentals/glossary/streaming) the structured data.

The [code generator](/fundamentals/glossary/code-generation) also comes with an [API](/fundamentals/glossary/api) that can be used to create your own [code generator](/fundamentals/glossary/code-generation). The [API](/fundamentals/glossary/api) provides an interface to every aspect of the [DSL](/fundamentals/glossary/dsl) up to the comments. There is basically no limit to how this [API](/fundamentals/glossary/api) can be used. One could:

-   generate source code for Programming Languages that are not natively supported,
-   generate your own code, differently from how `Protoc` does it,
-   use the declarations in the [DSL](/fundamentals/glossary/dsl) to configure a [Messaging](/fundamentals/glossary/messaging) platform, or
-   use the [DSL](/fundamentals/glossary/dsl) to generate documentation.

## Sources

-   https://github.com/protocolbuffers/protobuf
-   https://en.wikipedia.org/wiki/Protocol_Buffers
