---
date_created: 2022-12-11T16:58:36
title: REST
aliases:
  - REST
  - API
pcx_content_type: definition
summary: >-
  A REST API is (an interface definition of) a Web Application that can return data and execute actions on data.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/http
  - /fundamentals/design-and-architecture/standards-based/data-standards/json
  - /fundamentals/design-and-architecture/standards-based/data-standards/uri
  - /fundamentals/design-and-architecture/standards-based/data-standards/xml
  - /fundamentals/glossary/http-method
  - /fundamentals/glossary/http-return-code
---

# REST

A REST API is (an interface definition of) a Web Application that can return data and execute actions on data.

A REST API is a framework for consuming services, which heavily relies on [HTTP](/fundamentals/design-and-architecture/standards-based/data-standards/http). The [URI](/fundamentals/design-and-architecture/standards-based/data-standards/uri) of the `request` indicates **what** data that is operated on. The [HTTP Method](/fundamentals/glossary/http-method) specifies **how** is operated on that data.

The [HTTP Return code](/fundamentals/glossary/http-return-code) indicates whether the operation was successful, and, if not, gives an indication why not.

The `payload` is supposed to be [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json).

## Advantages and Disadvantages

The advantages of REST API is that it does NOT have the following things:

- Though [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) provides `XML Namespaces`, in practice one doesn't need it. After all, if you are calling a specific API the context (and thus the Namespace) is obvious.
- And though schema validation is a nice feature, one doesn't need that either, only in Design Time, and there are plenty of tools that can help in that area.
- JSON is both less bloaty than XML, and it is easier to parse.

The main disadvantage of REST APIs is that it still uses HTML, which is relatively slow with a relatively high overhead.
