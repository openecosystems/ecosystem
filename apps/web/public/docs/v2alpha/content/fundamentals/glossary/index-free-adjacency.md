---
date_created: 2022-12-12T19:35:46
author: Guillaume Hanique
title: Index-free Adjacency
pcx_content_type: definition
summary: >-
  Index-free Adjacency is a key element of Graph Technology, referring to how it stores and queries [Data Graphs](/fundamentals/glossary/#data-graph).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/data-graph
  - /fundamentals/glossary/database-index
  - /fundamentals/glossary/graph-database
---

# Index-free Adjacency

`Index-free Adjacency` is a key element of Graph Technology, referring to how it stores and queries [Data Graphs](/fundamentals/glossary/data-graph).

At read-time `Index-free Adjacency` ensures extremely fast retrieval, without reliance on indexes. (Non-[Graph Databases](/fundamentals/glossary/graph-database) or non-native [Graph Databases](/fundamentals/glossary/graph-database) use a large number of [Indexes](/fundamentals/glossary/database-index), slowing down both read and write Transactions significantly).

At write-time `Index-free Adjacency` speeds up processing by ensuring that each `Node` is stored _directly_ to its adjacent `Nodes` and `Relationships`.

## Sources

- https://neo4j.com/blog/native-vs-non-native-graph-technology/
