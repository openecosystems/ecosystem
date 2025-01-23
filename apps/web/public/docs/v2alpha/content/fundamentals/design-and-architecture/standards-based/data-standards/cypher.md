---
date_created: 2022-12-11T17:17:49
title: Cypher
pcx_content_type: definition
summary: >-
  `Cypher` is a Query Language for [Labeled Property Graphs](/fundamentals/glossary/labeled-property-graph). It is aimed to be easily readable by both humans and machines. It's also designed to look familiar to people that know [SQL](/fundamentals/design-and-architecture/standards-based/data-standards/#sql).
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/sql
  - /fundamentals/glossary/labeled-property-graph
  - /fundamentals/glossary/query
---

# Cypher

`Cypher` is a Query Language for [Labeled Property Graphs](/fundamentals/glossary/labeled-property-graph). It is aimed to be easily readable by both humans and machines. It's also designed to look familiar to people that know [SQL](/fundamentals/design-and-architecture/standards-based/data-standards/sql).

Consider the following `Cypher` [Query](/fundamentals/glossary/query) to find out who (or what) John Doe is in love with:

```cypher
MATCH (Person { name: "John Doe" })-[:LOVES]->(whom) RETURN whom
```

The different [Labeled Property Graph's](/fundamentals/glossary/labeled-property-graph) aspects can easily be recognized in this [query](/fundamentals/glossary/query):

- The `Node` that has the `Label` "Person" and a `Property` "name" that equal "John Doe"
- That has a `Relationship` "LOVES" to another node,
- that other `Node` we assign to a variable "whom",
- which we return.

## Sources

- https://github.com/opencypher/openCypher
- https://s3.amazonaws.com/artifacts.opencypher.org/openCypher9.pdf
