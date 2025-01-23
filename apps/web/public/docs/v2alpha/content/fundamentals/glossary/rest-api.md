---
title: REST API
pcx_content_type: definition
summary: >-
  A REST API is a Web Application that can return data and execute actions on data.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/html
  - /fundamentals/design-and-architecture/standards-based/data-standards/http
  - /fundamentals/design-and-architecture/standards-based/data-standards/json
  - /fundamentals/design-and-architecture/standards-based/data-standards/xml
  - /fundamentals/glossary/http-method
  - /fundamentals/glossary/http-return-code
---

# REST API

A `REST API` is a Web Application that can return data and execute actions on data. A `REST API` is a framework for consuming services, which heavily relies on [HTTP](/fundamentals/design-and-architecture/standards-based/data-standards/http). The `URI` of the `request` indicates the `kind of data` that is operated on. The [HTTP Method](/fundamentals/glossary/http-method) specifies how is operated on that data.

The [HTTP Return code](/fundamentals/glossary/http-return-code) indicates whether the operation was successful, and, if not, gives an indication why not.

The `payload` is supposed to be [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json).

## Advantages and Disadvantages

The advantages of `REST API` is that it does NOT have the following things:

- Though [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) provides XML Namespaces, in practice one doesn't need it. After all, if you are calling a specific API the context (and thus the Namespace) is obvious.
- And though schema validation is a nice feature, one doesn't need that either, only in Design Time, and there are plenty of tools that can help in that area.
- [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json) is both less bloaty than [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml), and it is easier to parse.

The main disadvantage of `REST API`s is that it still uses [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html), which is relatively slow with a relatively high overhead.
