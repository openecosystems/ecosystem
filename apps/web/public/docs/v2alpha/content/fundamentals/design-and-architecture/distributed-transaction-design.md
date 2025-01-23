---
pcx_content_type: reference
title: Distributed Transaction Design
weight: 11
---

# Distributed Transaction Design

Our transaction management is built on five pillars:

- Edge
- Discovery
- Routines
- Connectors
- Encryption

One of our design goals is: [Never Lose Data](/fundamentals/design-and-architecture/design-goals/)

## Edge

- Edge Jurisdiction Routing (where we guarantee Data Sovereignty, Data Residency and Data Classification). Your data lives in the correct jurisdiction based on your tenant.
- Ensures all transactions are rate limited, observed, governed, access controlled, and consented

## Discovery

- Dynamically register Connected Tests, Connected Devices, Health services, Clinical Protocols, Systems, Services and Connectors
- Allows tenants to turn on and off systems

## Routines

- Routines

## Connectors

- Connectors are our extensibility framework. All of our services that manage resources can elect to extend their transaction model with a connector. This allows tenants to dynamically change how a transaction is routed. For example, you can configure your tenant to route logistics request to DoorDash or UPS. Or you can configure your tenant to use our Machine Learning algorithm or roll out your own.

## Encryption

- We have designed the platform to support encryption everywhere.

All transactions are tightly observed, governed, audited, rate-limited, inspected, access-controlled, and checked for consent

## Transaction Concepts

### Canceling a transaction

### Transaction rollback

### Transaction Compression

### Transaction Deadlines

### Transaction Error Handling
