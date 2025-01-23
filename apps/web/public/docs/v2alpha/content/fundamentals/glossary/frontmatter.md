---
title: Frontmatter
aliases:
  - Frontmatter
pcx_content_type: definition
summary: >-
  `Frontmatter` is an extension to [Markdown](/fundamentals/design-and-architecture/standards-based/data-standards/markdown) for adding metadata to documents.
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/markdown
  - /fundamentals/design-and-architecture/standards-based/data-standards/yaml
  - /fundamentals/design-and-architecture/standards-based/data-standards/json
  - /fundamentals/design-and-architecture/standards-based/data-standards/toml
---

# Frontmatter

`Frontmatter` is an extension to [Markdown](/fundamentals/design-and-architecture/standards-based/data-standards/markdown) for adding metadata to documents.

Example:

```txt
---
title: Frontmatter
description: This article about Frontmatter
---

# Frontmatter

`Frontmatter` is an extension to Markdown for adding metadata to documents.
```

`Frontmatter` is a block of text at the top of the [Markdown](/fundamentals/design-and-architecture/standards-based/data-standards/markdown) file, and can be provided as:

- [YAML](/fundamentals/design-and-architecture/standards-based/data-standards/yaml): enclosed within `---` and `---`.
- [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json): enclosed within `{` and `}`.
- [TOML](/fundamentals/design-and-architecture/standards-based/data-standards/toml): enclosed within `+++` and `+++`.
