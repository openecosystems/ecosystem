---
title: Integrity
pcx_content_type: definition
summary: >-
    Data Integrity is the maintenance of, and the assurance of, data accuracy and consistency over its entire life-cycle.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/blockchain
    - /fundamentals/glossary/encryption
    - /fundamentals/glossary/hashing
    - /fundamentals/glossary/private-key
---

# Integrity

Data Integrity is the maintenance of, and the assurance of, data accuracy and consistency over its entire life-cycle.

## Cryptography

The way it works is that the `sender` generates a `Hash` (see [Hashing](/fundamentals/glossary/hashing)) of the data that it is sending. The sender then [encrypts](/fundamentals/glossary/encryption) the `Hash` using his [Private Key](/fundamentals/glossary/private-key) and sends both the data as well as the encrypted `hash`.

The `recipient` can calculate his own `Hash` from the data he received, decrypt the received `Hash` using the `sender`'s Public Key and compare both `Hashes`. If the `Hashes` are different, the data has been modified during transport. If they are not, there is a guarantee that it has not been modified.

## Quorum

A `Quorum` can be used to provide Data Integrity: if there are multiple instances of the data, then what the majority agrees on, is considered the Truth. This principle is applied in airplanes, clustered file systems, [Blockchain](/fundamentals/glossary/blockchain), etc.

In order to prevent Split-Brain, there should always be an odd number of instances. (Otherwise, if 50% thinks it's A, and 50% thinks it's B, then you still don't know, whereas if 67% thinks it's A and 33% thinks it's B, you know it's A.)
