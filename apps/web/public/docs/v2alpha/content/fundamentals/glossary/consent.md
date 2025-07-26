---
date_created: 2023-10-05T11:03:26
title: Consent
aliases:
    - Consent
pcx_content_type: definition
summary: >-
    `Consent` is a record where an individual authorizes a legal entity to access, use, or disclose [PII](/fundamentals/glossary/pii) and under what conditions that is allowed.
hidden: true
has_more: true
has_links: false
links_to:
    - /fundamentals/glossary/pii
    - /fundamentals/glossary/hipaa
    - /fundamentals/glossary/zero-trust-architecture
    - /fundamentals/glossary/authentication
    - /fundamentals/glossary/authorization
---

# Consent

`Consent` is a record where an individual authorizes a legal entity to access, use, or disclose [PII](/fundamentals/glossary/pii) and under what conditions that is allowed.

`Consent` can also be implicit. [HIPAA](/fundamentals/glossary/hipaa) includes several scenarios. If `Consent` is implicit, there won't be a Consent Record.

Also in scenarios where an explicit `Consent` is required, there is the possibility that the Consent Record doesn't exist yet because it is delayed, for example in case of emergencies.

InOpen Ecosystem `Consent` is integrated with [ZTA](/fundamentals/glossary/zero-trust-architecture) in the foundations of the platform: not only is every request [Authenticated](/fundamentals/glossary/authentication) and [Authorized](/fundamentals/glossary/authorization), if the request is to access data [ZTA](/fundamentals/glossary/zero-trust-architecture) will also check if a `Consent Record` exists that allows access to that data.
