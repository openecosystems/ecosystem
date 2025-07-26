---
date_created: 2022-12-11T17:09:35
title: URN
pcx_content_type: definition
summary: >-
    A URN is a [URI](/fundamentals/design-and-architecture/standards-based/data-standards/#uri) that uses the `urn` scheme to identify a logical or physical resource used by web technology, but it does not provide information to locate the object.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/uri
    - /fundamentals/design-and-architecture/standards-based/data-standards/url
---

# URN

A URN is a [URI](/fundamentals/design-and-architecture/standards-based/data-standards/uri) that uses the `urn` scheme to identify a logical or physical resource used by web technology, but it does not provide information to locate the object. (That would be specified in a [URL](/fundamentals/design-and-architecture/standards-based/data-standards/url)). `URNs` are global unique persistent identifiers assigned within defined namespaces so that they will be available for long periods of time, even after the resource that is identified by it ceased to exist.

Every `URN` starts with `urn:` followed by the _namespace identifier_ (or "path" as it is called for an `URI`), followed by the _namespace specific string_, i.e.: `urn:uuid:123e4567-e89b-12d3-a456-426614174000`.

The _namespace specific string_ can be expanded with separators to pass parameters to the object that is indicated with the `URN`, but I have yet to see this used.

AWS uses `URNs` extensively. Every object that is created also is assigned a `URN`.

## Sources

-   https://en.wikipedia.org/wiki/Uniform_Resource_Name
-   https://www.rfc-editor.org/rfc/rfc8141.html
