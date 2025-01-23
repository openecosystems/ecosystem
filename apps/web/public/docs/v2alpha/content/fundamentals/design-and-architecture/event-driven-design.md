---
pcx_content_type: reference
title: Event-Driven Design
value_plane_context: eventplane
platform_context: eventplane
weight: 7
---

# Event-Driven Design

To meet these competing [Design Goals](/fundamentals/design-and-architecture/design-goals/),
we are introducing an event-driven architecture that will allow the different teams and systems to communicate over a well-defined signal specification.
We refer to this as the Event Specification.

### Strong Consistency

Some of our services will require strong consistency across all our data centers.

### Eventual Consistency

As described in our design goals, we must achieve ACIDity with our transactions.
Because we have introduced an event-driven architecture, it has the consequence of transactions being “eventually” consistent.
This is the idea that consistency isn’t achieved immediately, but rather eventually, where eventually is usually within seconds.

#### Eventual Consistency ACID-ITY

To achieve eventually consistent ACID Transactions using an event-driven architecture,
we use Event Sourcing to persist multi-service transactions, and Command Query Responsibility Segregation (CQRS) to query the data.

#### Command Query Responsibility Separation (CQRS)


#### Event Sourcing



### Polyglot Persistence


