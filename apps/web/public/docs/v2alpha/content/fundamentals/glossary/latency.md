---
title: Latency
aliases:
  - Latency
pcx_content_type: definition
summary: >-
  A.k.a. "Delay". The time it takes for a request to result in a response. `Latency` is an important [Metric](/fundamentals/glossary/#metric) for measuring [Performance](/fundamentals/glossary/#performance).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/concurrent
  - /fundamentals/glossary/database
  - /fundamentals/glossary/metric
  - /fundamentals/glossary/performance
  - /fundamentals/glossary/sequential
  - /fundamentals/glossary/suicide-mechanism
  - /fundamentals/glossary/user-experience
---

# Latency

A.k.a. "Delay". The time it takes for a request to result in a response. `Latency` is an important [Metric](/fundamentals/glossary/metric) for measuring [Performance](/fundamentals/glossary/performance).

`Latency` is also an important [Metric](/fundamentals/glossary/metric) for [User Experience](/fundamentals/glossary/user-experience). I.e.: a User feels that the system is responding "instantaneous" if the system can respond to User input within 100ms. If something takes longer than 10 seconds, the User will want to go and do something different. (For more information, see [User Experience](/fundamentals/glossary/user-experience)).

`Latency` not often is taken into account when designing or implementing systems or components. But it's crucially important. Users may develop a negative association with the product, if the system is usually slow to respond. But more importantly, the chance of systems running out of resources increases exponentially as `Latency` increases beyond a certain threshold. This is also called a [Suicide Mechanism](/fundamentals/glossary/suicide-mechanism): if latency increases, the number of requests at any point in time increases, which requires more system resources and _decreases_ system [performance](/fundamentals/glossary/performance), which makes the problems worse.

Things one can do to improve `Latency`, and things that one Should take into account, when designing or developing systems, are:

- Limit the number of `hops`: a component that is invoked by another component, Should NOT, in turn, invoke yet another component. Likewise a component that interacts with a [Database](/fundamentals/glossary/database) Should have _one_ interaction with the `database`, not many.
- Execute operations [Concurrently](/fundamentals/glossary/concurrent) instead of [Sequentially](/fundamentals/glossary/sequential).
