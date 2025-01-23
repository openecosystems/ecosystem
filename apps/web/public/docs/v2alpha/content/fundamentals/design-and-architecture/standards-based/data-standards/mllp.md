---
date_created: 2022-12-11T17:00:28
title: MLLP
aliases:
  - MLLP
  - MLLPv2
  - Minimum Lower Layer Protocol
  - Minimum Lower Layer Protocol (MLLP)
pcx_content_type: definition
summary: >-
  `MLLP` is a protocol used to transfer [HL7v3](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v3) messages via [TCP/IP](/fundamentals/design-and-architecture/standards-based/data-standards/tcp).
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/hl7v3
  - /fundamentals/design-and-architecture/standards-based/data-standards/tcp
  - /fundamentals/design-and-architecture/standards-based/data-standards/synchronous
---

# Minimum Lower Layer Protocol (MLLP)

`MLLP` is a protocol used to transfer [HL7v3](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v3) messages via [TCP/IP](/fundamentals/design-and-architecture/standards-based/data-standards/tcp). It defines leading and trailing delimiters that can help the receiving applications determine the start and end of a [HL7](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v3) message. `MLLP` is inherently [Synchronous](/fundamentals/design-and-architecture/standards-based/data-standards/synchronous), because external systems almost always require the order of messages to be maintained.

`MLLPv2` is a requirement for transmitting [HL7v3](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v3) content. It adds the concept of Commit Acknowledgements to make the protocol more reliable.

The `MLLP` Adapter is a [TCP/IP](/fundamentals/design-and-architecture/standards-based/data-standards/tcp) socket adapter that uses the `MLLP` protocol.
