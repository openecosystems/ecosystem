---
title: Tracing
pcx_content_type: definition
summary: >-
  `Tracing` is a form of [Monitoring](/fundamentals/glossary/#monitoring) where messages or events are tracked throughout the system. At every step it is recorded how long that step took.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/latency
  - /fundamentals/glossary/monitoring
  - /fundamentals/glossary/performance
---

# Tracing

`Tracing` is a form of [Monitoring](/fundamentals/glossary/monitoring) where messages or events are tracked throughout the system. At every step it is recorded how long that step took.

This provides two kinds of information:

- If there is a [performance](/fundamentals/glossary/performance) issue when executing some kind of action, the `Trace Log` will indicate what component that processed the message or event took the longest. One could then examine that individual component to go and see what can be done to reduce the [Latency](/fundamentals/glossary/latency) in that component.
- `Errors` do not have to originate in the component where the error occurs. Sometimes there is a bug in the component that invokes the one where the error occurs. The `Trace Log` will show what other components processed the message or event before the current one did, which resulted in an Error.
