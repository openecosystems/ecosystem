---
title: Message ID
pcx_content_type: definition
summary: >-
    A `Message ID` is a unique number that Applications should assign to a message that it sends to another Application.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/tracing
---

# Message ID

A `Message ID` is a unique number that Applications should assign to a message that it sends to another Application, and it Should log that it is sending a message with that ID. It allows for some [Tracing](/fundamentals/glossary/tracing) of how messages travel to the system. I.e.: If Application A sends a message to Application B, then A's Application log will contain an entry that it send a message with a specific ID, and in B's Application log you will be able to find an entry that it received a message with the same ID.
