---
title: Point-to-Point
pcx_content_type: definition
summary: >-
  `Point-to-Point` is an [MEP](/fundamentals/glossary/#mep) where the publisher of a message is programmed to send the message to specific receivers.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/at-least-once
  - /fundamentals/glossary/at-most-once
  - /fundamentals/glossary/esb
  - /fundamentals/glossary/exactly-once
  - /fundamentals/glossary/mdp
  - /fundamentals/glossary/mep
  - /fundamentals/glossary/publish-subscribe
  - /fundamentals/glossary/queue
  - /fundamentals/glossary/time-critical
  - /fundamentals/glossary/topic
  - /fundamentals/glossary/ttl
---

# Point-to-Point

`Point-to-Point` is an [MEP](/fundamentals/glossary/mep) where the publisher of a message is programmed to send the message to specific receivers.

This pattern is mostly used to decouple logical systems when their dependency is not [Time Critical](/fundamentals/glossary/time-critical). Usually one would use the [MDP](/fundamentals/glossary/mdp) [At Least Once](/fundamentals/glossary/at-least-once) or [Exactly Once](/fundamentals/glossary/exactly-once) to make sure that messages from one system to the other don't get lost. With an [ESB](/fundamentals/glossary/esb) one would typically use a [Queue](/fundamentals/glossary/queue) for this [MEP](/fundamentals/glossary/mep).

In some scenarios [At Most Once](/fundamentals/glossary/at-most-once) is used, for example when a system sends messages that reflect the current state: as soon as a new message appears, the previous one is no longer relevant as it no longer reflects the _current_ state. In these scenarios with [ESB](/fundamentals/glossary/esb) one would typically use a [Topic](/fundamentals/glossary/topic), or a [Queue](/fundamentals/glossary/queue) where the messages [Expire](/fundamentals/glossary/ttl). One should seriously consider to use [Publish-Subscribe](/fundamentals/glossary/publish-subscribe), though.
