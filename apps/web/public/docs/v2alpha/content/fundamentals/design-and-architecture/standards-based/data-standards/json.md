---
date_created: 2022-12-11T17:00:28
title: JSON
aliases:
  - JSON
pcx_content_type: definition
summary: >-
  JSON is a way of describing any kind of artifact in a human-readable format, that is also easily parsed by JavaScript. Because of the convenience of using this format, many programming languages now have implementations to parse and render objects in `JSON`.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/yaml
---

# JSON

JSON is a way of describing any kind of artifact in a human-readable format, that is also easily parsed by JavaScript. Because of the convenience of using this format, many programming languages now have implementations to parse and render objects in `JSON`.

There are a few variants on `JSON`, one that has the same concept, but encodes it as a binary string (which is more efficient) and one that does something with cross-domains.

A format that addresses the same problem as `JSON` is [YAML](/fundamentals/design-and-architecture/standards-based/data-standards/yaml). [YAML](/fundamentals/design-and-architecture/standards-based/data-standards/yaml) is typically easier to read by humans, you can describe the same objects, and you can have comments inside a [YAML](/fundamentals/design-and-architecture/standards-based/data-standards/yaml)-file that are not part of the object declaration. This is at the cost of that this format depends on the indentation of each individual line, which could make it hard to read and modify if the object has many levels.

## Advantages

- Data can take any form (including arrays, and nested elements)
- Widely accepted
- Supported by pretty much any Programming Language

## Disadvantages

- No schema enforcing
- File can be big because of repeated key names
- Doesn't support comments

## Sources

- https://www.json.org/json-en.html
- https://en.wikipedia.org/wiki/JSON
