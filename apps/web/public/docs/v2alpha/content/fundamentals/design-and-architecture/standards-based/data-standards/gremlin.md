---
date_created: 2022-12-12T15:34:44
author: Guillaume Hanique
title: Gremlin
pcx_content_type: definition
summary: >-
  `Gremlin` is a query language for [Graph Databases](/fundamentals/glossary/#graph-database).
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/graph-database
  - /fundamentals/glossary/cli
  - /fundamentals/design-and-architecture/standards-based/design-patterns/method-chaining
---

# Gremlin

`Gremlin` is a query language for [Graph Databases](/fundamentals/glossary/graph-database). It was created by the same group that created Apache TinkerPop.

`Gremlin` also comes with a [CLI](/fundamentals/glossary/cli) that one can use to execute queries.

`Gremlin` uses [Method Chaining](/fundamentals/design-and-architecture/standards-based/design-patterns/method-chaining) for its queries.

Example:

```
g.V()
  .hasLabel('category')
  .as('a', 'b')
  .select('a', 'b')
  .by('name')
  .by(inE('category').count());
```

## Sources

- https://tinkerpop.apache.org/gremlin.html
- https://en.wikipedia.org/wiki/Gremlin_(query_language)
