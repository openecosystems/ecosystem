---
title: SQS
pcx_content_type: definition
summary: >-
    AWS SQS is AWS' Queue Service. This service allows for sending and receiving messages over a Queue.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/queue
---

# Simple Queue Service (SQS)

AWS SQS is AWS' Queue Service. This service allows for sending and receiving messages over a [Queue](/fundamentals/glossary/queue).

By definition a message that is persisted on a Queue is guaranteed to never get lost. The only way to lose a message, is to `consume` it (i.e.: get the message from the queue), and then failing to process it. (As opposed to AWS SNS, where a message gets lost if there are no `consumers`).

There can be many instances `posting` messages to a specific Queue. There can also be many instances `consuming` messages from the Queue. However, only one of the `consumers` will get a specific message. (As opposed to AWS SNS, where a message that is `posted` to a Topic will be processed by _all_ `consumers`). This is a convenient way to implement Load Balancing.

## Multiple Consumers

I believe it is possible to pipe an AWS SNS Topic to an AWS SQS Queue, to have the message be consumed by multiple `consumers` as well to have `guaranteed delivery`.

If one would have multiple Consumers on one AWS Queue and they use different "group names", then each of those Consumers should get their own copy of the Message.

## Large Messages

Messages that can be posted on a Queue **cannot be larger than 256kB**. There is a public library that could be used to post larger messages to a Queue: if a message is too large it will store the payload in an AWS S3 `bucket` and `post` a message to the Queue that contains a reference to the file in AWS S3. On the `consumer` end it will receive the message with a reference to the payload, dereference it and provide that to the payload. It should be obvious that this will increase Latency **tremendously**.

## Delays

An AWS SQS Queue can be configured as a "Delay Queue". (I currently cannot think of a use case).

It is also possible to set a delay on a message explicitly when putting it on the AWS Queue.
