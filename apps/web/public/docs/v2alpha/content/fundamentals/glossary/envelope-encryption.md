---
title: Envelope Encryption
pcx_content_type: definition
summary: >-
    `Envelope Encryption` is similar to [PFS](/fundamentals/glossary/#pfs), but then for *storing* data (as opposed to transmitting data).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/encryption
    - /fundamentals/glossary/encryption-key
    - /fundamentals/glossary/key-ring
    - /fundamentals/glossary/pfs
---

# Envelope Encryption

If one [Encryption Key](/fundamentals/glossary/encryption-key) is used to [Encrypt](/fundamentals/glossary/encryption) and Decrypt data, if that Key ever gets compromised _all_ data that was Encrypted with it gets compromised.

For data transports this generally isn't as big of a problem, because once the data has been transported, the transport is gone and there is nothing left to Decrypt. To further reduce the Risk, [PFS](/fundamentals/glossary/pfs) is typically applied.

But for **storing** Encrypted data the problem is huge, because the Encrypted data is there to stay. So what one would want is something similar to [PFS](/fundamentals/glossary/pfs), but then for data that **stored** instead of transported. That is where `Envelope Encryption` comes in.

With `Envelope Encryption` the [Encryption Key](/fundamentals/glossary/encryption-key) is regularly rotated. Every time a new Encryption Key is generated, it is added to the [Key Ring](/fundamentals/glossary/key-ring).

Now, when data must be [Encrypted](/fundamentals/glossary/encryption), it is Encrypted with the most recent [Encryption Key](/fundamentals/glossary/encryption-key) (or a new one); instead of just storing the Encrypted data, an `Encryption Envelope` is stored that contains both the Encrypted data and a reference to the [Encryption Key](/fundamentals/glossary/encryption-key) that is stored in the [Key Ring](/fundamentals/glossary/key-ring).

If data must be Decrypted, the [Decryption Key](/fundamentals/glossary/encryption-key) is retrieved from the [Key Ring](/fundamentals/glossary/key-ring) using the reference in the `Encryption Envelope`, which is then used to Decrypt the data.
