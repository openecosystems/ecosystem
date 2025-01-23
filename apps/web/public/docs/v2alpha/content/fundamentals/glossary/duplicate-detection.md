---
title: Duplicate Detection
aliases:
  - Duplicate Detection
pcx_content_type: definition
summary: >-
  For the scenarios where [Idempotence](/fundamentals/glossary/#idempotence) is not possible, but processing each request must be guaranteed nevertheless, one should implement `Duplicate Detection` to prevent requests being processed more than once if they are received more than once.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/at-least-once
  - /fundamentals/glossary/caching
  - /fundamentals/glossary/exception-handling
  - /fundamentals/glossary/idempotence
  - /fundamentals/glossary/mep
  - /fundamentals/glossary/message-id
  - /fundamentals/glossary/request-response
  - /fundamentals/glossary/retry-mechanism
---

# Duplicate Detection

Depending on the situation you may need to make sure that requests are processed [At Least Once](/fundamentals/glossary/at-least-once). To achieve that you may need to resend requests in case of [Exceptions](/fundamentals/glossary/exception-handling). The problem with resending requests, is that they may end up being received more than once.

One should attempt to design a system in such a way that it is [Idempotent](/fundamentals/glossary/idempotence): a system that is Idempotent will reach the same state regardless of how often the same request is processed. This way one can maintain a consistent state simply by implementing a [Retry Mechanism](/fundamentals/glossary/retry-mechanism).

For the scenarios where [Idempotence](/fundamentals/glossary/idempotence) is not possible, but processing each request must be guaranteed nevertheless, one should implement `Duplicate Detection` to prevent requests being processed more than once if they are received more than once.

One way to implement `Duplicate Detection`, is by identifying a key that makes a request unique, and storing that in a [Cache](/fundamentals/glossary/caching). (A good candidate for the unique key could be a [Message ID](/fundamentals/glossary/message-id)). Before processing a request, the system checks that that request's unique key is in the Cache. If it is, it skips processing the request. If the request is not in the Cache, then it should be processed.

If the request follows the [Request-Response](/fundamentals/glossary/request-response) [MEP](/fundamentals/glossary/mep), one should not only [Cache](/fundamentals/glossary/caching) the request, but also the response. In this scenario the Cache is written twice: once before processing the request (to make sure that if the request is received again while being processed, that processing can be prevented), and once when the response has been determined.
