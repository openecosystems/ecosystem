---
title: Publish-Subscribe
pcx_content_type: definition
summary: >-
  `Publish-Subscribe` is a [MEP](/fundamentals/glossary/#mep) where publishers of messages are not programmed to send those messages to specific receivers. Instead, the message is published to a "channel" and zero or more receivers could subscribe to that "channel" and receive a copy of that message. If there are no subscribers, the message gets lost without anyone having seen it.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/at-least-once
  - /fundamentals/glossary/esb
  - /fundamentals/glossary/exactly-once
  - /fundamentals/glossary/mep
  - /fundamentals/glossary/topic
  - /fundamentals/glossary/topic-to-queue-bridge
---

# Publish-Subscribe

`Publish-Subscribe` is a [MEP](/fundamentals/glossary/mep) where publishers of messages are not programmed to send those messages to specific receivers. Instead, the message is published to a "channel" and zero or more receivers could subscribe to that "channel" and receive a copy of that message. If there are no subscribers, the message gets lost without anyone having seen it.

With an [ESB](/fundamentals/glossary/esb) `Publish-Subscribe` is implemented with a [Topic](/fundamentals/glossary/topic). If one also needs [At Least Once](/fundamentals/glossary/at-least-once) or [Exactly Once](/fundamentals/glossary/exactly-once), one would combine the [Topic](/fundamentals/glossary/topic) with a [Topic to Queue Bridge](/fundamentals/glossary/topic-to-queue-bridge).
