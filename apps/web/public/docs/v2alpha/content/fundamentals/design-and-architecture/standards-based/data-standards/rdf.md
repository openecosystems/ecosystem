---
date_created: 2022-12-11T16:57:27
author: Guillaume Hanique
title: RDF
aliases:
    - RDF
    - Resource Description Framework
    - Resource Description Framework (RDF)
pcx_content_type: definition
summary: >-
    Resource Description Framework (RDF) is a specification for expressing information about resources in a [Data Graph](/fundamentals/glossary/data-graph).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/data-graph
    - /fundamentals/design-and-architecture/standards-based/data-standards/iri
    - /fundamentals/design-and-architecture/standards-based/data-standards/xml
    - /fundamentals/glossary/graph-database
    - /fundamentals/design-and-architecture/standards-based/data-standards/rdf-xml
    - /fundamentals/design-and-architecture/standards-based/data-standards/n-triples
    - /fundamentals/design-and-architecture/standards-based/data-standards/turtle
    - /fundamentals/design-and-architecture/standards-based/data-standards/trig
    - /fundamentals/design-and-architecture/standards-based/data-standards/json-ld
    - /fundamentals/design-and-architecture/standards-based/data-standards/rdfa
---

# Resource Description Framework (RDF)

Resource Description Framework (RDF) is a specification for expressing information about resources in a [Data Graph](/fundamentals/glossary/data-graph).

With RDF all data is modeled with `Triples`. A `Triple` is the atomic data entity in the RDF data model. It consists of - you guessed it - three fields:

-   `Subject`: An [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri) identifying an entity.
-   `Predicate`: An [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri) specifying how the subject relates to the `Object`.
-   `Object`: A [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri) to an object (which could be another `Tripple`).

RDF heavily depends on schemas, similar to how every XML Element MUST be of a specific XML Namespace, that is defined in an XSD.

So suppose there is a Person whose first name is John and whose last name is Doe, then this is more or less what it would look like in RDF:

-   First we need a `Subject`. The `Subject` MUST be an [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri). So let's assume we created our own schema that defines that there is a thing called a "Person". Our John Doe also must have an ID, which MUST be part of the URI. Then this could be a valid `Subject`: `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e`.
-   The next thing we need is a `Predicate`. Let's start with the first name. We cannot just say that the `Predicate` is "firstname", because the `Predicate` MUST be a [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri). We could either add predicates to our own schema, or we could use one of the standard ones. Let's use "firstName" of the "foaf" namespace (Friend of a Friend). Then the `Predicate` would be `http://xmlns.com/foaf/0.1/firstName`.
-   The last thing we need is the `Object`. In this case it's the value "John", which is a "Literal" in the `rdfs` namespace. So the `Object` would be: "John". ("John" must still be of a specific type).

Then we need to repeat this process for the last name. There is no `Predicate` called "lastname", though, so let's use "family_name".

This is John Doe:

| Subject                                                                 | Predicate                               | Object |
| ----------------------------------------------------------------------- | --------------------------------------- | ------ |
| `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` | `http://xmlns.com/foaf/0.1/firstName`   | John   |
| `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` | `http://xmlns.com/foaf/0.1/family_name` | Doe    |

Now let's say that John Doe is married to Jane Doe. We'd have to create another two `Triples` for `Jane`, and then we can create another `Triple` to marry the two. There is no existing [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri) for "married to", but there is one for "knows". So here is John marrying Jane:

| Subject                                                                 | Predicate                               | Object                                                                  |
| ----------------------------------------------------------------------- | --------------------------------------- | ----------------------------------------------------------------------- |
| `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` | `http://xmlns.com/foaf/0.1/firstName`   | John                                                                    |
| `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` | `http://xmlns.com/foaf/0.1/family_name` | Doe                                                                     |
| `http://openecosystems.com/person#7b9c0cad-a13e-4e42-850d-d7d0d0c50160` | `http://xmlns.com/foaf/0.1/firstName`   | Jane                                                                    |
| `http://openecosystems.com/person#7b9c0cad-a13e-4e42-850d-d7d0d0c50160` | `http://xmlns.com/foaf/0.1/family_name` | Doe                                                                     |
| `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` | `http://xmlns.com/foaf/0.1/knows`       | `http://openecosystems.com/person#22761879-a876-4e38-a9a0-9e44fe498e3e` |

## XML Implications

-   **Validation**: Just like one can unambiguously validate an XML document if one has all the XSD, one can also unambiguously validate an RDF if one has all the RDF schema's.
-   **Strongly typed**: Though the data that is stored in an RDF has no structure, all the `Objects` are Strongly Typed. Just like one could define an "Enumeration" in an XSD, one could also define those in an RDF Schema.
-   **Bloaty**: The official RDF specification is even more bloaty than [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml) itself.
-   **Prefixes**: Because an RDF is a specification on top of XSL, one can create aliases for various namespaces. I.e.: By adding the XML attribute `xmlns:foaf="http://xmlns.com/foaf/0.1/"`, "`http://xmlns.com/foaf/0.1/family_name`" can be shortened to "foaf:family_name".

## Performance

RDF has the same problem as XML: it is very heavy on computing resources. For that reason the official RDF specification is not practical at scale.[^1]

The concept of a `Triple` is still valid though. So RDF based [Graph Databases](/fundamentals/glossary/graph-database) could not adhere to the RDF standard, but only to the concept of `Triples`. They will also apply much smarter ways of storing those `Triples`, so that data can be queried far more efficiently.

## Advantages and Disadvantages

One advantage of RDF is that the Data Model is very Simple: it only consists of `Triples`. This Simplicity also comes with a disadvantage: the Data Model has no structure, so the Data Model gives no indication whether something is a `property` or an `entity`.

## Serialization

There are several ways to serialize RDF data:

-   [RDF/XML](/fundamentals/design-and-architecture/standards-based/data-standards/rdf-xml)
-   [N-Triples](/fundamentals/design-and-architecture/standards-based/data-standards/n-triples)
-   [Turtle](/fundamentals/design-and-architecture/standards-based/data-standards/turtle)
-   [TriG](/fundamentals/design-and-architecture/standards-based/data-standards/trig)
-   [JSON-LD](/fundamentals/design-and-architecture/standards-based/data-standards/json-ld)
-   [RDFa](/fundamentals/design-and-architecture/standards-based/data-standards/rdfa)

## References

-   https://www.w3.org/TR/rdf11-primer/
-   https://www.w3.org/RDF/
-   https://www.w3.org/1999/02/22-rdf-syntax-ns
-   https://www.w3.org/2001/sw/RDFCore/Schema/20010913/
-   http://xmlns.com/foaf/0.1/

[^1]: https://en.wikipedia.org/wiki/Semantic_triple

## Sources

-   https://www.w3.org/RDF/
