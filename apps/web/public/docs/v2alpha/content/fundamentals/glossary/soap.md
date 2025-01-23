---
title: SOAP
aliases:
  - SOAP
pcx_content_type: definition
summary: >-
  "`Simple Object Access Protocol`. But it's not simple. It allows for describing services, similar to [Swagger](/fundamentals/design-and-architecture/standards-based/data-standards/swagger), but then in [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml). It is described in a [WSDL](/fundamentals/design-and-architecture/standards-based/data-standards/wsdl)."
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/swagger
  - /fundamentals/design-and-architecture/standards-based/data-standards/xml
  - /fundamentals/design-and-architecture/standards-based/data-standards/wsdl
  - /fundamentals/design-and-architecture/standards-based/data-standards/xsd
  - /fundamentals/glossary/debugging
---

# Simple Object Access Protocol (SOAP)

`Simple Object Access Protocol`. But it's not simple.

It allows for describing services, similar to [Swagger](/fundamentals/design-and-architecture/standards-based/data-standards/swagger), but then in [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml). It is described in a [WSDL](/fundamentals/design-and-architecture/standards-based/data-standards/wsdl).

`Messages` are transmitted in a `SOAP Envelope`, which is an `XML Schema` (see [XSD](/fundamentals/design-and-architecture/standards-based/data-standards/xsd) that basically has a root object with two properties: `Header` and `Body`. When sending a `SOAP Message` one must also specify a `SOAP Operation`. There is a number of ways to transmit one.

## Advantages and Disadvantages

The advantage of `SOAP` and [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) is that it allows for unambiguous data formats: there is no discussion whether or not a message meets its schema. Another nice feature of [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) is the concept of XML Namespaces, which allows for different definitions of an object with the same name depending on the context.

The advantage is also a disadvantage: it makes everything all the more complex, as a result of which one finds oneself [Debugging](/fundamentals/glossary/debugging) too often why a specific message is not getting processed. Another downside is that [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) is pretty bloaty, which results in requiring more bandwidth, and more computing power and memory to process [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml).
