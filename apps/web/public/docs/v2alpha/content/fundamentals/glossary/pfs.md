---
title: PFS
pcx_content_type: definition
summary: >-
    Perfect Forward Secrecy (PFS) combines [Symmetric Encryption](/fundamentals/glossary/#symmetric-encryption) and [Asymmetric Encryption](/fundamentals/glossary/#asymmetric-encryption) in a way that makes it both Secure *and* Cheap. PFS is applied in various Transport Protocols and Message Protocols like HTTPS and [SOAP](/fundamentals/glossary/#soap).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/asymmetric-encryption
    - /fundamentals/glossary/encryption
    - /fundamentals/glossary/encryption-key
    - /fundamentals/glossary/soap
    - /fundamentals/glossary/symmetric-encryption
---

# Perfect Forward Secrecy (PFS)

Both [Symmetric Encryption](/fundamentals/glossary/symmetric-encryption) and [Asymmetric Encryption](/fundamentals/glossary/asymmetric-encryption) have problems: one is _Cheap_ but not _Secure_, the other one is _Secure_ but not _Cheap_.

`PFS` combines Symmetric Encryption and Asymmetric Encryption in a way that makes it both Secure _and_ Cheap. `PFS` is applied in various Transport Protocols and Message Protocols like HTTPS and [SOAP](/fundamentals/glossary/soap).

-   A new [Symmetric Key](/fundamentals/glossary/encryption-key) is created, which is used to [Encrypt](/fundamentals/glossary/encryption) the data, which is Cheap.
-   The `Symmetric Key` is Encrypted with [Asymmetric Encryption](/fundamentals/glossary/encryption-key), which is Secure, and still Cheap because the `Symmetric Key` is quite small.
-   Both the [Symmetric Encrypted](/fundamentals/glossary/symmetric-encryption) data and the [Asymmetrically Encrypted](/fundamentals/glossary/asymmetric-encryption) `Symmetric Key` are sent to the recipient.
-   The recipient will first decrypt the `Symmetric Key` key using `Asymmetric Decryption` and then use the `Symmetric Key` to decrypt the data.
