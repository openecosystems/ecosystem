---
title: Symmetric Encryption
pcx_content_type: definition
summary: >-
  A cryptographic algorithm to [Encrypt](/fundamentals/glossary/#encryption) data using a `key`, where the data can be Decrypted using the same `key`. The most commonly used algorithm is AES256.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/encryption
  - /fundamentals/glossary/encryption-key
---

# Symmetric Encryption

A cryptographic algorithm to Encrypt data using a `key`, where the data can be Decrypted using the same `key`. The most commonly used algorithm is AES256.

`Symmetric Encryption` is strong and requires little computing resources. The downside of `Symmetric Encryption` is that both the sender and receiver must have the same `key`, which implies that the key must be exchanged, which is a Security Risk.

Another downside of `Symmetric Encryption` is that one could theoretically calculate the Symmetric [Encryption Key](/fundamentals/glossary/encryption-key)[^1] if one has enough data and enough compute. (For AES256 it is not yet possible to have enough data and compute).

[^1]: This is what makes cracking WEP WiFi Networks possible. See https://www.wikihow.com/Break-WEP-Encryption.
