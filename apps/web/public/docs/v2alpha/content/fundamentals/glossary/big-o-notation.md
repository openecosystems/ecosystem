---
date_created: 2022-12-01T21:54:00
title: Big O Notation
aliases:
  - Big O
  - Big O Notation
pcx_content_type: definition
summary: >-
  "In short the `Big O Notation` is a mathemathical notation that describes the limiting behavior of a function if the argument tends towards a particular value or infinitiy. In layman's terms it means: 'if the load doubles, how much more compute do I need?'"
hidden: true
has_more: true
has_links: false
links_to:
  - /fundamentals/glossary/performance
  - /fundamentals/glossary/scaling
---

# Big O Notation

In short the `Big O Notation` is a mathemathical notation that describes the limiting behavior of a function if the argument tends towards a particular value or infinitiy. In layman's terms it means: if the load doubles, how much more compute do I need?

The `Big O Notation` can be used to describe how the [Performance](/fundamentals/glossary/performance) of software change if the load increases. If the load doubles, do I need twice the amount of resources? Less than twice? More than twice? If [Scaling](/fundamentals/glossary/scaling) is to be possible at all, it must be less than twice.

Though theoretically you could write any formula, there are a few common ones:

- `O(1)`: Regardless of the size of the input, processing will always take the same time.
- `O(log n)`: If the data size doubles, then the time it takes to process it increases with a constant value.
- `O(n)`: If the dataset is twice as large, processing it will take twice as long.
- `O(m log n)`: If the data size doubles, then the time it takes to process it doubles AND increases with a constant value.
- `O(n2)`: If the data size doubles, the time it takes to process it quadruples.

Any process that behaves worse than `O(n)` cannot be [Scaled](/fundamentals/glossary/scaling). [Scaling](/fundamentals/glossary/scaling) an `O(n)` is very expensive. One should aim for implementations that are `O(log n)` (which is usually possible) or `O(1)` (which is hard to achieve).
