---
title: HTTP-404
pcx_content_type: definition
summary: >-
  '`HTTP 404` "Not found" means that the resource that is indicated with the [URI](/fundamentals/design-and-architecture/standards-based/data-standards/uri) cannot be found.'
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/uri
  - /fundamentals/glossary/authorization
---

# HTTP-404

`HTTP 404` "Not found" means that the resource that is indicated with the [URI](/fundamentals/design-and-architecture/standards-based/data-standards/uri) cannot be found. This can either refer to a service, or to the data. The response may contain fields that help make the distinction.

I.e.:

- If there is no service that can handle "http://server/non-existent-service", the server will return a `HTTP 404.`
- But if there is a service that can look up books, for example, but the [URI](/fundamentals/design-and-architecture/standards-based/data-standards/uri) refers to a book that doesn't exist (as in "http://server/books/book-that-does-not-exist"), then, too, the server would return a `HTTP 404`. One of the HTTP Headers will probably indicate that the type is "data".

Servers may also return a `HTTP 404` instead of a `HTTP 403` ("forbidden") to obfuscate that a resource exists that the client is not [Authorized](/fundamentals/glossary/authorization) to access.
