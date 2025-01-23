---
title: At Least Once
pcx_content_type: definition
summary: >-
  `At Least Once` is an [MDP](/fundamentals/glossary/#mdp). With this pattern the sender would send the message, and have a [Retry Mechanism](/fundamentals/glossary/#retry-mechanism) in place in case sending the message fails.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/mdp
  - /fundamentals/glossary/retry-mechanism
  - /fundamentals/glossary/idempotence
  - /fundamentals/glossary/duplicate-detection
  - /fundamentals/glossary/metric
---

# At Least Once

`At Least Once` is an [MDP](/fundamentals/glossary/mdp). With this pattern the sender would send the message, and have a [Retry Mechanism](/fundamentals/glossary/retry-mechanism) in place in case sending the message fails. Though there is a mechanism that ensures the message is being sent, there is **no mechanism** in place to prevent the message from being processed more than once.

One example would be where one sends an HTTP request to a service that implements a business process: some steps could succeed, but if one of the steps causes a timeout exception, the sender will resend the message and the message could be processed more than once.

One could use this pattern in scenarios where it doesn't matter if a message is processed more than once (because the operation is [Idempotent](/fundamentals/glossary/idempotence)), or where the receiver implements [Duplicate Detection](/fundamentals/glossary/duplicate-detection) (thus achieving Exactly Once in a different way). Again, a [Metric](/fundamentals/glossary/metric) could be a good example: if we store twice that the disk was used for 65% at a certain point in time, that doesn't really change the meaning of the data.
