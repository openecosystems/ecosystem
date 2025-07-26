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
{{<raw>}}<pre class="CodeBlock CodeBlock-with-rows CodeBlock-scrolls-horizontally CodeBlock-is-light-in-light-theme CodeBlock--language-txt" language="txt"><code><span class="CodeBlock--rows"><span class="CodeBlock--rows-content"><span class="CodeBlock--row"><span class="CodeBlock--row-indicator"></span><div class="CodeBlock--row-content"><span class="CodeBlock--token-plain">cypher</span></div></span></span></span></code></pre>{{</raw>}}

The different [Labeled Property Graph's](/fundamentals/glossary/labeled-property-graph) aspects can easily be recognized in this [query](/fundamentals/glossary/query):

-   The `Node` that has the `Label` "Person" and a `Property` "name" that equal "John Doe"
-   That has a `Relationship` "LOVES" to another node,
-   that other `Node` we assign to a variable "whom",
-   which we return.

## Sources

-   https://github.com/opencypher/openCypher
-   https://s3.amazonaws.com/artifacts.opencypher.org/openCypher9.pdf
