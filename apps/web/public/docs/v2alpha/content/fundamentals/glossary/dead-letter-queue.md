---
title: Dead Letter Queue
pcx_content_type: definition
summary: >-
    A `Dead Letter Queue` is a [Design Pattern](/fundamentals/glossary/#design-pattern) where one moves messages to a dedicated [Queue](/fundamentals/glossary/#queue) called "the `Dead Letter Queue`" if the message meets one or more [Exception](/fundamentals/glossary/#exception-handling) criteria.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/exception-handling
    - /fundamentals/glossary/application-layer
    - /fundamentals/glossary/aws
    - /fundamentals/glossary/design-pattern
    - /fundamentals/glossary/event-plane
    - /fundamentals/glossary/exactly-once
    - /fundamentals/glossary/exception-handling
    - /fundamentals/glossary/guaranteed-delivery
    - /fundamentals/glossary/incident
    - /fundamentals/glossary/middleware-layer
    - /fundamentals/glossary/monitoring
    - /fundamentals/glossary/publish-subscribe
    - /fundamentals/glossary/queue
    - /fundamentals/glossary/racetrack-problem
    - /fundamentals/glossary/request-response
    - /fundamentals/glossary/ttl
---

# Dead Letter Queue

A `Dead Letter Queue` is a [Design Pattern](/fundamentals/glossary/design-pattern) where one moves messages to a dedicated [Queue](/fundamentals/glossary/queue) called "the `Dead Letter Queue`" if the message meets one or more [Exception](/fundamentals/glossary/exception-handling) criteria.

## Exception criteria

### Queue does not exist

If the message is sent to a [Queue](/fundamentals/glossary/queue) that does not exist, it could be sent to the `Dead Letter Queue`.

This would have to be implemented on the [Event Plane](/fundamentals/glossary/event-plane).

### Queue length limit exceeded

Sometimes there is a limit to how many messages a [Queue](/fundamentals/glossary/queue) can hold. So if the [Queue](/fundamentals/glossary/queue) fills up because the messages are not being processed (or not being processed as fast as they are produced), then new messages cannot be placed in the [Queue](/fundamentals/glossary/queue). To not lose those messages, they could be placed in the `Dead Letter Queue` instead.

This would have to be implemented on the [Event Plane](/fundamentals/glossary/event-plane).

### Message or Queue length limit exceeded

Sometimes there are restrictions to the size a message can have, or the maximum size a [Queue](/fundamentals/glossary/queue) can have. In either of those cases, if that limit is exceeded then new messages cannot be placed on the [Queue](/fundamentals/glossary/queue). To not lose them, they could be placed in the `Dead Letter Queue` instead.

This would have to be implemented on the [Event Plane](/fundamentals/glossary/event-plane). Some implementations, though, like [AWS SQS](/fundamentals/glossary/aws/#sqs) for example, don't support messages that exceed a certain size, period (`256 kB` in the example of [AWS SQS](/fundamentals/glossary/aws/#sqs)). That means that if a message cannot be placed in the [Queue](/fundamentals/glossary/queue) because it's too big, that it cannot be placed in the `Dead Letter Queue` either. To not lose that message, the `Dead Letter Queue` cannot be a [Queue](/fundamentals/glossary/queue), but something else should serve as a `Dead Letter Queue`, like a file system. This would have to be implemented in the [Application Layer](/fundamentals/glossary/application-layer), though.

### Message is rejected by another Queue exchange

Some [Event Plane](/fundamentals/glossary/event-plane) implementations, like RabbitMQ, support explicitly rejecting a message (which could be done as part of [Exception Handling](/fundamentals/glossary/exception-handling)). The [Event Plane](/fundamentals/glossary/event-plane) could be configured to send those messages to the `Dead Letter Queue` instead.

### Message reaches a threshold read counter

A "Read Limit" is a feature that is provided by some [Event Plane](/fundamentals/glossary/event-plane) implementations to address the [Racetrack Problem](/fundamentals/glossary/racetrack-problem) in the [Middleware Layer](/fundamentals/glossary/middleware-layer) instead of the [Application Layer](/fundamentals/glossary/application-layer): if a messages was picked up from the [Queue](/fundamentals/glossary/queue) and then put back (because an [Exception](/fundamentals/glossary/exception-handling) occurred while processing it, for example), then it probably never will get processed, and by moving it to the `Dead Letter Queue` one will not run into the [Racetrack Problem](/fundamentals/glossary/racetrack-problem).

### The message expires

Many [Event Plane](/fundamentals/glossary/event-plane) implementations support a [TTL](/fundamentals/glossary/ttl). This should always be applied when [Request-Response](/fundamentals/glossary/request-response) is used, and could be applied when [Publish-Subscribe](/fundamentals/glossary/publish-subscribe) is used for notifications.

One could possibly configure the [Event Plane](/fundamentals/glossary/event-plane) to put expired messages to a `Dead Letter Queue`. The real question is whether you should. I think they answer is usually "No". After all, the [TTL](/fundamentals/glossary/ttl) was deliberately set because after that time the message is no longer relevant. So if it is rendered irrelevant by the [TTL](/fundamentals/glossary/ttl) expiring, why retain that message?

### The message is not processed correctly

If a message fails to process correctly, because of an [Exception](/fundamentals/glossary/exception-handling) for example, the default operation is to put the message back in the [Queue](/fundamentals/glossary/queue). This, however, could easily result in the [Racetrack Problem](/fundamentals/glossary/racetrack-problem), where the message gets processed over and over again and fails every time.

Instead of putting the message back in the [Queue](/fundamentals/glossary/queue), it could be put in the `Dead Letter Queue` instead. This, however, would have to be implemented in the [Application Layer](/fundamentals/glossary/application-layer), as that is the lowest level that is aware that the message fails to be processed. It is somewhat tricky to implement though:

-   You are already in an [Exception](/fundamentals/glossary/exception-handling) state.
-   If [Guaranteed Delivery](/fundamentals/glossary/guaranteed-delivery) is one of the requirements, meeting that requirement is not trivial when removing a message from one [Queue](/fundamentals/glossary/queue) and putting it in the `Dead Letter Queue`. It will be even less trivial if [Exactly Once](/fundamentals/glossary/exactly-once) is also a requirement.

The simpler solution might be [Message reaches a threshold read counter](#message-reaches-a-threshold-read-counter) in those cases.

## Handling the Dead Letter Queue

Redirecting messages to a `Dead Letter Queue` is only useful if one fixes the issues that got the messages in the `Dead Letter Queue` in the first place, and then reprocesses them, or deliberately decides to discard them.

Especially if there is an [Incident](/fundamentals/glossary/incident) that causes messages to be directed to the `Dead Letter Queue`, there could be a lot of messages there that need to be processed. One should think carefully about how to implement that.

There are two different approaches to handling the `Dead Letter Queue`.

### One Dead Letter Queue for all messages

The advantage of this approach is that one would have one place where failed messages end up. This means that one also only has to [Monitor](/fundamentals/glossary/monitoring) one [Queue](/fundamentals/glossary/queue) for failed messages.

Because every kind of failed message now ends up in one [Queue](/fundamentals/glossary/queue), and because there are a lot of them, Tooling is required to fix and reprocess messages efficiently. If one chooses for this approach one should pick an [Event Plane](/fundamentals/glossary/event-plane) implementation that provides those adequate tools.

### Multiple dedicated Dead Letter Queues

In this scenario every [Queue](/fundamentals/glossary/queue) that makes use of `Dead Letter Queue`s gets its own dedicated `Dead Letter Queue`.

The advantage of this approach is that you always know what kind of message you're dealing with, because it only contains one kind of messages. If the problem is fixed and the messages need to be reprocessed, that, too, is easy, because one knows what [Queue](/fundamentals/glossary/queue) the messages came from.

The disadvantage of this approach is that you will have as many `Dead Letter Queue`s as there are [Queues](/fundamentals/glossary/queue). This many [Queues](/fundamentals/glossary/queue) is harder to [Monitor](/fundamentals/glossary/monitoring) than just one [Queue](/fundamentals/glossary/queue). Also if there are many `Dead Letter Queue`s and many different kinds of failed messages, one has to jump back and forth between a lot of `Dead Letter Queue`s to resolve all issues.

Another thing to take into account is that not all [Event Plane](/fundamentals/glossary/event-plane) implementations allow for configuring which [Queue](/fundamentals/glossary/queue) to use for the `Dead Letter Queue` in different scenarios. This may result in that the `Dead Letter Queue` Pattern has to be implemented in the [Application Layer](/fundamentals/glossary/application-layer) instead of in the [Middleware Layer](/fundamentals/glossary/middleware-layer).
