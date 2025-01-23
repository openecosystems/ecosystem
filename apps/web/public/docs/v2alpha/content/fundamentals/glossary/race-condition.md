---
title: Race Condition
pcx_content_type: definition
summary: >-
  A Race Condition is one where a system's substantive behavior is dependent on the sequence or timing of other controllable events. It becomes a bug when one or more of those behaviors is undesirable.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/racetrack-problem
---

# Race Condition

> A Race Condition is one where a system's substantive behavior is dependent on the sequence or timing of other controllable events. It becomes a bug when one or more of those behaviors is undesirable.

A `Race Condition` is usually seen when multiple threads operate on a shared state without proper locking.

A good example is where one thread increases a value and the other decreases it. If one starts with `0`, and both threads run once, you expect `0` as the output. But in a race condition that doesn't have to be the case:

|      Thread 1 |       Thread 2 | Value |
| ------------: | -------------: | ----: |
|               |                |     0 |
|     read <- 0 |                |     0 |
|               |      read <- 0 |     0 |
| increase -> 1 |                |     0 |
|               | decrease -> -1 |     0 |
|         write |                |     1 |
|               |          write |    -1 |

A special kind of `Race Condition` is the [Racetrack Problem](/fundamentals/glossary/racetrack-problem).
