---
date_created: 2022-12-11T17:14:48
title: WSDL
aliases:
    - WSDL
pcx_content_type: definition
summary: >-
    Web Service Description Language (WSDL) is a language to describe [SOAP](/fundamentals/design-and-architecture/standards-based/data-standards/#soap) Web Services.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/soap
    - /fundamentals/design-and-architecture/standards-based/data-standards/xsd
---

# Web Service Description Language (WSDL)

Web Service Description Language (WSDL) is a language to describe [SOAP](/fundamentals/design-and-architecture/standards-based/data-standards/soap) Web Service. A `WSDL` defines a number of things:

-   A list of `types` (see [XSD](/fundamentals/design-and-architecture/standards-based/data-standards/xsd)) that describe the `XML Elements` that can be used.
-   A list of `interfaces` that describe:
    -   The `XML Element` that should be used for a `Fault`
    -   The `Operations` that that interface implements, where each `Operation` has a:
        -   `input`, which references the `XML Element` that is used for sending messages to the service for that operation,
        -   `output`, which references the `XML Element` that is used for when the service responds
-   A list of `bindings`, which describes another layer of abstraction over `Operations`, and what `Interface` implements it. (This is where things are getting vague).
-   A `service` that has multiple `endpoint`s, that link to a specific `binding`.

## Sources

-   https://www.w3.org/TR/2007/REC-wsdl20-20070626/
