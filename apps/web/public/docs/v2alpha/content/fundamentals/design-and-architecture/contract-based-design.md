---
pcx_content_type: reference
title: Contract-Based Design
weight: 10
links_to:
  - /fundamentals/platform/data/spec-entity
---

# Contract-Based Design

Because of the vast number of implementations across systems, services, connectors and jurisdictions,

we use Contract-Based Design to define precise and verifiable interface specifications that are semantically equivalent to a [Hoare Triple](https://en.wikipedia.org/wiki/Hoare_logic#Hoare_triple).

All of our contracts must answer the following questions:

- What does the contract expect from a Spec Actor?
- What does the contract guarantee?
- What [Spec Entity](/fundamentals/platform/data/spec-entity) does the contract maintain?

A [Hoare Triple](https://en.wikipedia.org/wiki/Hoare_logic#Hoare_triple) describes how the execution of Service or Connector changes the state of a [Spec Entity](/fundamentals/platform/data/spec-entity).

A Hoare triple is of the form

```text
{ P } C { Q }
```

where P and Q are assertions and C is a command

P is named the precondition and Q the postcondition: when the precondition is met, executing the command establishes the postcondition.

## Contracts

- Each contract specifies the Service Calls its supports
- The Input each service call expects
- The Return value each service call guarantees
- The possible errors the service call may return and their meanings

## Pre Conditions

A precondition is a condition that must always be true before the Platform will allow the service to execute a Service Call.

- For each field, the System will define acceptable and unacceptable field values or types

## Post Conditions

## Invariants

## Side Effects
