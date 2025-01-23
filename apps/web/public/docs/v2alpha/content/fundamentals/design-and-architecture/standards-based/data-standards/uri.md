---
title: URI
aliases:
  - URI
  - Uniform Resource Identifier (URI)
  - Uniform Resource Identifier
pcx_content_type: definition
date_created: 2022-12-11T17:01:51
author: Guillaume Hanique
summary: >-
  A Uniform Resource Identifier (URI) is a unique sequence of characters that identifies a logical or physical resource used by web technologies. URIs may be used to identify anything, including real-world objects, such as people and places, concepts, or information resources such as web pages and books.
hidden: true
has_more: false
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/url
  - /fundamentals/design-and-architecture/standards-based/data-standards/urn
  - /fundamentals/design-and-architecture/standards-based/data-standards/rdf
  - /fundamentals/design-and-architecture/standards-based/data-standards/owl
  - /fundamentals/design-and-architecture/standards-based/data-standards/iri
---

# URI

A Uniform Resource Identifier (URI) is a unique sequence of characters that identifies a logical or physical resource used by web technologies. URIs may be used to identify anything, including real-world objects, such as people and places, concepts, or information resources such as web pages and books.

Some URIs provide a means of locating and retrieving information resources on a network; these are Uniform Resource Locators ([URLs](/fundamentals/design-and-architecture/standards-based/data-standards/url)). A URL provides the location of the resource. A `URI` identifies the resource by name at the specified location or URL.

Other `URI`s provide only a unique name, without a means of locating or retrieving the resource or information about it, these are Uniform Resource Names ([URNs](/fundamentals/design-and-architecture/standards-based/data-standards/urn)).

`URI`s are not only used by Browsers. They are also used to identify anything in [RDF](/fundamentals/design-and-architecture/standards-based/data-standards/rdf) and [OWL](/fundamentals/design-and-architecture/standards-based/data-standards/owl).

The `URI` syntax is:

```text
URI = scheme ":" ["//" authority] path ["?" query] ["#" fragment]
```

Some examples:

```
          userinfo       host      port
          ┌──┴───┐ ┌──────┴──────┐ ┌┴┐
  https://john.doe@www.example.com:123/forum/questions/?tag=networking&order=newest#top
  └─┬─┘   └───────────┬──────────────┘└───────┬───────┘ └───────────┬─────────────┘ └┬┘
  scheme          authority                  path                 query           fragment

  ldap://[2001:db8::7]/c=GB?objectClass?one
  └┬─┘   └─────┬─────┘└─┬─┘ └──────┬──────┘
  scheme   authority   path      query

  mailto:John.Doe@example.com
  └─┬──┘ └────┬─────────────┘
  scheme     path

  news:comp.infosystems.www.servers.unix
  └┬─┘ └─────────────┬─────────────────┘
  scheme            path

  tel:+1-816-555-1212
  └┬┘ └──────┬──────┘
  scheme    path

  telnet://192.0.2.16:80/
  └─┬──┘   └─────┬─────┘│
  scheme     authority  path

  urn:oasis:names:specification:docbook:dtd:xml:4.1.2
  └┬┘ └──────────────────────┬──────────────────────┘
  scheme                    path
```

[[URI|URIs]] can only use a subset of characters. All other characters must be URL Encoded. [IRI](/fundamentals/design-and-architecture/standards-based/data-standards/iri) extends `URIs` by also allowing international characters.

## Sources

- https://en.wikipedia.org/wiki/Uniform_Resource_Identifier
- https://www.rfc-editor.org/rfc/rfc3986
