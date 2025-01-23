---
date_created: 2023-10-13T16:11:35
title: Request Acknowledge Process Pattern
aliases:
  - Request Acknowledge Process Pattern
  - Asynchronous Processing Pattern
pcx_content_type: definition
summary: >-
  The `Request Acknowledge Process Pattern` is particularly useful when you have long-running or resource-intensive operations that could potentially block the main Service from responding promptly to client requests.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/failure
  - /fundamentals/glossary/performance
  - /fundamentals/glossary/scaling
  - /fundamentals/glossary/high-availability
  - /fundamentals/design-and-architecture/standards-based/design-patterns/loose-coupling
  - /fundamentals/glossary/design-pattern
  - /fundamentals/glossary/complexity
  - /fundamentals/glossary/retry-mechanism
  - /fundamentals/glossary/exception-handling
  - /fundamentals/glossary/duplicate-detection
---

# Request Acknowledge Process Pattern

The `Request Acknowledge Process Pattern` is particularly useful when you have long-running or resource-intensive operations that could potentially block the main Service from responding promptly to client requests.

Here is how the `Request Acknowledge Process Pattern` typically works:

1. The client sends a request to a Service to perform a task that might take a long time to complete.
2. Instead of processing the request immediately, the Service persists the client's request is some form of persistent storage, so that the request won't be lost even in the event of a [Failure](/fundamentals/glossary/failure).
3. The Service responds with a "promise" that the request will be processed, for example by providing a token that corresponds to the specific request.
4. The Service processes the request in the background, allowing the main service to continue handling other requests.
5. Once the long-running process is finished, the service could notify the client that processing has completed.

This pattern has several benefits:

- [Performance](/fundamentals/glossary/performance): The client receives the response promptly, even if the task takes a long time to complete.
- [Scaling](/fundamentals/glossary/scaling): The background process can be scaled separately.
- [Fault-Tolerance](/fundamentals/glossary/high-availability): By persisting requests, the system can recover from [Failures](/fundamentals/glossary/failure), ensuring that no request is lost.
- [Loose Coupling](/fundamentals/design-and-architecture/standards-based/design-patterns/loose-coupling): Clients and Services are Loosely Coupled, since they don't have to wait for each other.
- [[Concurrent|Concurrency]]: The system can execute multiple long-running tasks in parallel, making efficient use of resources.

This [Design Pattern](/fundamentals/glossary/design-pattern) also introduces [Complexity](/fundamentals/glossary/complexity), such as managing the storage of requests, [Retry Mechanism](/fundamentals/glossary/retry-mechanism), [Exception Handling](/fundamentals/glossary/exception-handling), [Duplicate Detection](/fundamentals/glossary/duplicate-detection), etc.
