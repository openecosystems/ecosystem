---
title: Synchronous
aliases:
    - Synchronous
pcx_content_type: definition
summary: >-
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/latency
    - /fundamentals/glossary/performance
    - /fundamentals/design-and-architecture/standards-based/design-patterns/request-acknowledge-process
---

# Synchronous

`Synchronous` communication means that when a request is sent to execute an action, that the process is blocked until a confirmation is received that the execution is completed. This provides a guarantee that the request has really completed, but this comes at the cost of [Latency](/fundamentals/glossary/latency), and thus at the cost of [Performance](/fundamentals/glossary/performance).

The [Request Acknowledge Process Pattern](/fundamentals/design-and-architecture/standards-based/design-patterns/request-acknowledge-process) could reduce the introduced [Latency](/fundamentals/glossary/latency), where the receiver doesn't respond with a _confirmation_ that the request has been completed, but with a _promise_ that the request will be completed. But that means that the receiver will have to implement measures to guarantee that it can keep the promise.

Synchronous communication can be useful if the response is needed to continue. It can also be useful when the request is Time Critical and an [[Exception Handling|Exception]] Must be raised if the operation could not be completed within a certain amount of time.
