---
title: Hostile Network
pcx_content_type: definition
summary: >-
    A Hostile Network is any Network where you have no control on who is using it and who can see information you're transmitting. The most well known example is the [internet](/fundamentals/glossary/#internet).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/encryption
    - /fundamentals/glossary/endpoint-validation
    - /fundamentals/glossary/integrity
    - /fundamentals/glossary/internet
    - /fundamentals/glossary/man-in-the-middle-attack
---

# Hostile Network

A `Hostile Network` is any Network where you have no control on who is using it and who can see information you're transmitting. The most well known example is the [internet](/fundamentals/glossary/#internet).

Using a `Hostile Network` has plenty of advantages, among which the ability to communicate with third parties. But in order to be able to use it one must address the disadvantages of such a network. This can include:

-   [Endpoint Validation](/fundamentals/glossary/endpoint-validation), to make sure that we are talking to the intended recipient, and not a [Man in the Middle](/fundamentals/glossary/man-in-the-middle-attack) pretending to be the intended recipient.
-   [Encryption](/fundamentals/glossary/encryption), to make sure that only the intended recipient can read the transmitted information.
-   [Integrity](/fundamentals/glossary/integrity), to enable the recipient to assert that it was us who sent the message (after all, anyone can Encrypt a message for the recipient, and potentially make it appear as if we sent the message).
