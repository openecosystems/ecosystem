---
date_created: 2022-12-11T16:58:06
author: Guillaume Hanique
title: OWL
pcx_content_type: definition
summary: >-
  The Web Ontology Language (OWL) is a semantic web language design to represent rich and complex knowledge about things, a group of things, and relations between things. The current standard is OWS2. It was published in 2009. It is based on [RDF](/fundamentals/design-and-architecture/standards-based/data-standards/rdf).
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/rdf
  - /fundamentals/glossary/complexity
  - /fundamentals/glossary/labeled-property-graph
---

# OWL

The Web Ontology Language (`OWL`) is a semantic web language design to represent rich and complex knowledge about things, a group of things, and relations between things. The current standard is `OWL2`. It was published in 2009. It is based on [RDF](/fundamentals/design-and-architecture/standards-based/data-standards/rdf).

Its intent is to reduce [Complexity](/fundamentals/glossary/complexity) and augment data with meta data.

There are three main `OWL` artifacts:

- **Concepts** represent a set of classes, entities, or things within a domain, which are used to classify `Instances` or other `Concepts`.
- **Instances** are used to refer to the things represented by the `Concept`. These may include concrete objects, or abstracts like numbers or words.
- **Relationships** specify how the objects relate to one another.

`Object Properties` are used to link `Instances` to other `Instances`.

`Data Properties` link `Instances` to Values. I.e.: The `Instance` "Petrol" has a "type" `Property` with the value "Diesel", and a "Liters" `Property` for how much.

`OWL` is implemented on top of [RDF](/fundamentals/design-and-architecture/standards-based/data-standards/rdf), but looks remarkably like a [Labeled Property Graph](/fundamentals/glossary/labeled-property-graph) implemented with RDF. The main difference is that `OWL` has restrictions to the `Labels` and `Relationships` that can be used. These `Labels` and `Relationships` are defined in the [RDF](/fundamentals/design-and-architecture/standards-based/data-standards/rdf)-`OWL` schema.

## Sources

- https://www.w3.org/TR/owl2-syntax/
- https://www.w3.org/TR/2012/REC-owl2-overview-20121211/
