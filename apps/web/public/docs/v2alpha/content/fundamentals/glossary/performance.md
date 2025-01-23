---
title: Performance
aliases:
  - Performance
  - Performant
pcx_content_type: definition
summary: >-
  `Performance` is a vague term that describes how fast a system is, but it can be expressed with concrete [Metrics](/fundamentals/glossary/#metric).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/latency
  - /fundamentals/glossary/metric
  - /fundamentals/glossary/performance-monitoring
  - /fundamentals/glossary/performance-testing
  - /fundamentals/glossary/production-environment
  - /fundamentals/glossary/stress-testing
  - /fundamentals/glossary/suicide-mechanism
  - /fundamentals/glossary/throughput
---

# Performance

`Performance` is a vague term that describes how fast a system is, but it can be expressed with concrete [Metrics](/fundamentals/glossary/metric) like:

- [Latency](/fundamentals/glossary/latency)
- [Throughput](/fundamentals/glossary/throughput)
- `Consumption of compute resources`
- etc.

There is a [Suicide Mechanism](/fundamentals/glossary/suicide-mechanism) where lower performance (due to higher [Latency](/fundamentals/glossary/latency) for example) can result in more concurrent transactions, which require more `compute resources`, which degrades performance.

There are three perspectives on `Performance`: [Performance Testing](/fundamentals/glossary/performance-testing), [Performance Monitoring](/fundamentals/glossary/performance-monitoring), and [Stress Testing](/fundamentals/glossary/stress-testing). They are related, but they solve different problems:

- [Performance Testing](/fundamentals/glossary/performance-testing) should be done when developing a system or component. The idea is that when a component is first created, a performance test is executed on it where [Metrics](/fundamentals/glossary/metric) like [Latency](/fundamentals/glossary/latency) and [Throughput](/fundamentals/glossary/throughput) are captured, which will become that component's baseline. If then later the component is changed, its performance can be compared to its baseline, to make sure that a new version doesn't degrade the system.
- [Stress Testing](/fundamentals/glossary/stress-testing) is the "art" of stressing the system to the point where it breaks. It could be somewhat useful to find how a component behaves at some edge cases. What a breaking point means, could be somewhat obscure for various reasons. `Stess Testing` is most valuable if it is used to determine thresholds for [Performance Monitoring](/fundamentals/glossary/performance-monitoring).
- [Performance Monitoring](/fundamentals/glossary/performance-monitoring) is about how the system is really behaving in a [Production Environment](/fundamentals/glossary/production-environment). The best part is that you are looking at a realistic load. By considering how performance changes over time, one can detect when whole system is going to break before it actually breaks. [Stress Testing](/fundamentals/glossary/stress-testing) can provide valuable info here, because it can provide thresholds when we know things are going to break, so that we can intervene before it does.
