---
title: TTL
pcx_content_type: definition
summary: >-
  `Time to Live`, a timespan after creation of data, after which the data is no longer valid or relevant.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/exception-handling
---

# Time to Live (TTL)

`Time to Live`, a timespan after creation of data, after which the data is no longer valid or relevant.

A Service that sends a request with a TTL should also raise a [Timeout Exception](/fundamentals/glossary/exception-handling) if the TTL expires before it receives a response.
