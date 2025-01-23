---
title: Retry Mechanism
aliases:
  - Retry Mechanism
pcx_content_type: definition
summary: >-
  A Retry Mechanism is a mechanism that monitors a request, and on the detection of a Failure automatically fires a repeat of the request.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/dos-attack
  - /fundamentals/glossary/duplicate-detection
  - /fundamentals/glossary/idempotence
  - /fundamentals/glossary/rate-limiting
  - /fundamentals/glossary/ttl
---

# Retry Mechanism

A Retry Mechanism is a mechanism that monitors a request, and on the detection of a Failure automatically fires a repeat of the request.

A retry should only be considered if there is a chance of success. I.e.: in case of a HTTP-404 ("resource cannot be found") a retry is inadvisable as it wastes resources and could have unexpected results (like getting an ever-increasing number of requests that need to be retried, which will eventually crash the system).

Retries could result in a [DoS](/fundamentals/glossary/dos-attack) if the Retry Mechanism is poorly implemented. One way to prevent a DoS is [Rate Limiting](/fundamentals/glossary/rate-limiting). On the client-side this could be achieved by applying Exponential Backoff. If API Gateways are used, then these could also provide means to limit the rate.

Retries, by definition, can cause duplicates. I.e.: if processing takes longer than a [TTL](/fundamentals/glossary/ttl), then the request will be processed AND retried. Duplicates should be handled correctly. Two possible ways of handling duplicates:

- [Duplicate Detection](/fundamentals/glossary/duplicate-detection): If the receiver can determine that the message has been processed, it can skip processing duplicates.
- [Idempotence](/fundamentals/glossary/idempotence): The system can be designed in such a way that repeatedly processing the same message will always have the same result. I.e.: "Change the value to 2" is idempotent, whereas "Increase the value by 1" is not.

## Sources

- https://denalibalser.medium.com/best-practices-for-retry-685bf58de79
