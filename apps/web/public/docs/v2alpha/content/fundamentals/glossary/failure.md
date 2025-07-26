---
title: Failure
aliases:
    - Failure
    - Failures
pcx_content_type: definition
summary: >-
    An event where a component becomes unavailable. Typically this does not result in loss, and minor actions are required to continue business, like restarting a server.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/design-patterns/design-for-failure
    - /fundamentals/glossary/availability-zone
---

# Failure

An event where a component becomes unavailable. Typically this does not result in loss, and minor actions are required to continue business, like restarting a server.

A good metaphor is a car getting a flat tire: you replace the tire and continue driving.

A way to make systems resistant to `Failure` is to deploy components redundantly: if one component has a `failure` the other one would still be working and the platform would still be available. This is also called [Design for Failure](/fundamentals/design-and-architecture/standards-based/design-patterns/design-for-failure).

Making use of [Availability Zones](/fundamentals/glossary/availability-zone) helps make a platform robust against an entire datacenter going down.
