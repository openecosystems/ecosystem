---
title: Endpoint Validation
pcx_content_type: definition
summary: >-
    `Endpoint Validation` is the process to assert that the endpoint of a communication is the intended recipient and not a hostile recipient pretending to be the intended one.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/ssl
    - /fundamentals/glossary/private-key
    - /fundamentals/glossary/certificate-authority
---

# Endpoint Validation

`Endpoint Validation` is the process to assert that the endpoint of a communication is the intended recipient and not a hostile recipient pretending to be the intended one.

`Endpoint Validation` is provided by [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl) and HTTPS. I.e.: When Browsing to the website of your bank, their server identifies itself with a [Private Key](/fundamentals/glossary/private-key) that is signed by a [Certificate Authority](/fundamentals/glossary/certificate-authority), who we trust to have asserted the physical identity of the party that requested the [Private Key](/fundamentals/glossary/private-key).
