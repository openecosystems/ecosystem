---
title: Exponential Backoff
pcx_content_type: definition
summary: >-
  Exponential Backoff is an algorithm that uses feedback to multiplicatively decrease the rate of some process, in order to gradually find an acceptable rate. It is also applied to [Retry Mechanisms](/fundamentals/glossary/#retry-mechanism).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/dos-attack
  - /fundamentals/glossary/http-429
  - /fundamentals/glossary/rate-limiting
  - /fundamentals/glossary/retry-mechanism
---

# Exponential Backoff

Relates to:

- [Rate Limiting](/fundamentals/glossary/rate-limiting)

Exponential Backoff is an algorithm that uses feedback to multiplicatively decrease the rate of some process, in order to gradually find an acceptable rate. In the strict sense of the word it is a `closed-loop control system`.

In computer systems often a simplified algorithm is used to implement a [Retry Mechanism](/fundamentals/glossary/retry-mechanism) with [Rate Limiting](/fundamentals/glossary/rate-limiting) to prevent a [DoS](/fundamentals/glossary/dos-attack). (This is still considered "Exponential Backoff", even though this is not a `closed-loop control system`). An example: If there is an Exception when sending a request, the system will wait 1, 2, 4, 8, 16, ..., seconds respectively before trying again. To prevent very long delays after the situation has been restored, the maximum delay is often capped. I.e.: 1, 2, 4, 8, 8, 8, ... seconds between retries.

Sometimes there is a situation that a request doesn't result in an Exception, but that the service that has to handle the request is too busy. [HTTP 429](/fundamentals/glossary/http-429) is a mechanism to handle this scenario. (Note that if this happens that it may be an indication of a DoS in progress).
