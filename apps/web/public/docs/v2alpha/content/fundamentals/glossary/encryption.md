---
title: Encryption
aliases:
    - Encryption
    - Encrypt
    - Encrypted
    - Encrypting
    - Decrypt
    - Decrypted
    - Decrypting
pcx_content_type: definition
summary: >-
    Encryption is the process of encoding information in an unreadable format. It can also be done in a way that only the intended recipient can decrypt the information.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/asymmetric-encryption
    - /fundamentals/glossary/authentication
    - /fundamentals/glossary/encryption-key
    - /fundamentals/glossary/envelope-encryption
    - /fundamentals/glossary/nist
    - /fundamentals/glossary/pfs
    - /fundamentals/glossary/private-key
    - /fundamentals/glossary/public-key
    - /fundamentals/glossary/symmetric-encryption
---

# Encryption

Encryption is the process of encoding information in an unreadable format. It can also be done in a way that only the intended recipient can decrypt the information.

Market standards are set by [NIST](/fundamentals/glossary/nist). Every system that is used in Health Care systems must meet these standards.

## Symmetric and Asymmetric Encryption

There are two forms of Encryption:

-   [Symmetric Encryption](/fundamentals/glossary/symmetric-encryption)
-   [Asymmetric Encryption](/fundamentals/glossary/asymmetric-encryption)

[Symmetric Encryption](/fundamentals/glossary/symmetric-encryption) means that the same [Encryption Key](/fundamentals/glossary/encryption-key) can be used both to Encrypt and to Decrypt the information. This is cheap in terms of computing resources, but the downside is that the Encryption Key has to be exchanged between the sender and receiver, which introduces a Security Risk.

[Asymmetric Encryption](/fundamentals/glossary/asymmetric-encryption) is a different algorithm where data can be Encrypted with one key and Decrypted by another. One is kept private, a.k.a. the [Private Key](/fundamentals/glossary/private-key), and the other is distributed freely, a.k.a. the [Public Key](/fundamentals/glossary/public-key). If data is Encrypted using the recipient's Public Key, then the recipient is the only one that can Decrypt the message using his Private Key. So the advantage of this form of Encryption, is that the `Decryption Key` does not have to be exchanged. The downside is that this form of Encryption requires a disproportionate amount of computing resources.

## Authentication

[Asymmetric Encryption](/fundamentals/glossary/asymmetric-encryption) can also be used for [Authentication](/fundamentals/glossary/authentication). I.e.: if the sender encrypts a piece of data using his Private Key, then anyone can use the Public Key to decrypt it, and assert that the sender is the only one that could have encrypted that piece of data.

## Perfect Forward Secrecy

Because [Symmetric Encryption](/fundamentals/glossary/symmetric-encryption) is cheap but Insecure, and [Asymmetric Encryption](/fundamentals/glossary/asymmetric-encryption) is Secure, but expensive, the two are often combined when sending Encrypted data. This is called Perfect Forward Secrecy ([PFS](/fundamentals/glossary/pfs)).

## Envelope Encryption

[Envelope Encryption](/fundamentals/glossary/envelope-encryption) is to _storing_ data what [PFS](/fundamentals/glossary/pfs) is to _transporting_ data.
