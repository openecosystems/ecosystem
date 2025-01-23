---
title: Request-Response
pcx_content_type: definition
summary: >-
  With `Request-Response` a message is sent from one component to another with the expectation to receive a response, because the response is required for further processing. `Request-Response` is an [MEP](/fundamentals/glossary/#mep).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/esb
  - /fundamentals/glossary/exception-handling
  - /fundamentals/glossary/mep
  - /fundamentals/glossary/queue
  - /fundamentals/glossary/topic
  - /fundamentals/glossary/ttl
---

# Request-Response

`Request-Response` is an [MEP](/fundamentals/glossary/mep).

With `Request-Response` a message is sent from one component to another with the expectation to receive a response, because the response is required for further processing.

Applications implementing the `Request-Response` pattern should have proper [Exception Handling](/fundamentals/glossary/exception-handling) in place to handle [timeouts](/fundamentals/glossary/ttl), where it takes too long to receive a response.

For [ESB](/fundamentals/glossary/esb) implementations `Request-Response` could either be implemented with [Queues](/fundamentals/glossary/queue) or with [Topics](/fundamentals/glossary/topic). If Queues are used, messages should _always_ have a [TTL](/fundamentals/glossary/ttl) that is approximately equal to the timeout value, because otherwise the Queue could gradually fill up over time.
