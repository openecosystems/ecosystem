---
title: Projection
pcx_content_type: definition
summary: >-
  A `Projection` is a declaration of a sub-model of a data document.
hidden: true
has_more: true
aliases:
  - Projection
  - Projections
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/json
---

# Projection

A Data Document is a multi-dimensional data object, similar to what you would find in a [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json) file. When working with a document store, one generally does not interact with entire Documents, but only with a number of fields from those Documents.

`Projection` refers to the ability to specify which fields or properties of Data Documents should be retrieved or updated.

The combination of fields in a `Projection` is also called a "sub-model" and it is documented in an Entity Diagram.
