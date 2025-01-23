---
title: HTTP 429
pcx_content_type: definition
summary: >-
  HTTP 429 "Too many requests" is an [HTTP Return code](/fundamentals/glossary/#http-return-code) that can be returned by a service if it doesn't have enough resources to fulfill this request, or if a [Rate Limit](/fundamentals/glossary/#rate-limiting) has been imposed on the consumer, which has been exceeded. An HTTP 429 may be an indication that a [DoS Attack](/fundamentals/glossary/dos-attack) is in progress.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/dos-attack
  - /fundamentals/glossary/http-return-code
  - /fundamentals/glossary/rate-limiting
---

# HTTP 429

{{%Aside type="warning" header="Important"%}}
An HTTP 429 may be an indication that a [DoS Attack](/fundamentals/glossary/dos-attack) is in progress.
{{%/Aside%}}

HTTP 429 "Too many requests" is an [HTTP Return code](/fundamentals/glossary/http-return-code) that can be returned by a service if it doesn't have enough resources to fulfill this request, or if a [Rate Limit](/fundamentals/glossary/rate-limiting) has been imposed on the consumer, which has been exceeded.

The response may contain a `Retry-After` header, which can either contain a timestamp or a delay in seconds when a new request can be attempted. Clients should honor the request to wait before resending the request.

If the `Retry-After` header is absent, then the client should apply Exponential Backoff when resending the request.
