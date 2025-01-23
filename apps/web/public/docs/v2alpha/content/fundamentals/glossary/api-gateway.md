---
title: API Gateway
pcx_content_type: definition
summary: >-
  An `API Gateway` is an API Management tool that sits between a client and a collection of [Back End](/fundamentals/glossary/#back-end) services. It acts as a reverse proxy to accept all [API](/fundamentals/glossary/#api) calls, and forwards them to internal services that can fulfill them.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/api
  - /fundamentals/glossary/back-end
  - /fundamentals/glossary/billing
  - /fundamentals/glossary/dos-attack
  - /fundamentals/glossary/lcm
  - /fundamentals/glossary/mediation
  - /fundamentals/glossary/metering
  - /fundamentals/glossary/throttling
---

# API Gateway

An `API Gateway` is an API Management tool that sits between a client and a collection of [Back End](/fundamentals/glossary/back-end) services. It acts as a reverse proxy to accept all [API](/fundamentals/glossary/api) calls, and forwards them to internal services that can fulfill them.

Exposing internal Microservices externally introduces new problems that will have to be addressed:

- Isolating consumers: No single external party should be able to cause a [DoS](/fundamentals/glossary/dos-attack) for other external parties.
- [LCM](/fundamentals/glossary/lcm) decoupling: It should be possible to [LCM](/fundamentals/glossary/lcm) internal services, without requiring external parties to switch to new versions at the same time.
- [Billing](/fundamentals/glossary/billing): External parties that use the system more, should pay more. (Pay-per-use).

An `API Gateway` addresses these problems with:

- [Throttling](/fundamentals/glossary/throttling)
- [Metering](/fundamentals/glossary/metering)
- [Mediation](/fundamentals/glossary/mediation)
