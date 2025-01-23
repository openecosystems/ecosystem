---
title: Queue
pcx_content_type: definition
summary: >-
  A Queue is a channel where `messages` are received. Multiple Applications can listen to that channel, but only one of them will receive the `message`. If nothing is listening, the `messages` on the channel will be persisted, until something starts listening to the channel.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/dead-letter-queue
  - /fundamentals/glossary/esb
  - /fundamentals/glossary/performance
  - /fundamentals/glossary/race-condition
---

# Queue

A Queue is a channel where `messages` are received. Multiple Applications can listen to that channel, but only one of them will receive the `message`. If nothing is listening, the `messages` on the channel will be persisted, until something starts listening to the channel.

## Queue limits

Because the [ESB](/fundamentals/glossary/esb) has a finite amount of resources: Queues cannot grow indefinitely. Therefore Queues can have a limit for the amount of messages, a limit for the total size of the messages, or both. Depending on the `ESB` implementation, when the limit is reached, either:

1. new messages are rejected, or
2. new messages are accepted but the oldest messages are discarded.

If messages are rejected, they could be configured to be redirected to a [Dead Letter Queue](/fundamentals/glossary/dead-letter-queue).

The second approach is rather dangerous as architects and developers might assume message are not getting lost, as they are using a Queue, when in fact, they are.

## TTL

Messages that are put on a Queue can be assigned a TTL. When that time expires, the message is removed from the Queue.

In practice this approach could result in a [Performance](/fundamentals/glossary/performance) issue, depending on the amount of messages that is on the Queue. The engine can only remove expired messages by traversing the messages on the Queue and checking whether the TTL has expired. If there is a large amount of messages in the Queue, or if this pattern is used in many places, the Performance impact can become non-negligible.

## Race Condition

With Queues there is a risk of a [Race Condition](/fundamentals/glossary/race-condition): if a message cannot be processed, it could be put back on the Queue, where it is picked up again, cannot be processed, and is put back on the Queue again. Introducing a [Dead Letter Queue](/fundamentals/glossary/dead-letter-queue) is one of the ways this problem can be mitigated.
