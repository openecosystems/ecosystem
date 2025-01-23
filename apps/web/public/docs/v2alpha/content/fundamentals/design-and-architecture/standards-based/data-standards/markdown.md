---
date_created: 2023-10-11T16:56:19
title: Markdown
aliases:
  - Markdown
pcx_content_type: definition
summary: >-
  `Markdown` is a very lightweight plain text markup language for writing documentation. It is easy to read, easy to write, and easy to render to other markup languages like [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html).
hidden: true
has_more: true
todo: Find more details and sources
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/html
  - /fundamentals/design-and-architecture/standards-based/data-standards/pdf
  - /fundamentals/design-and-architecture/standards-based/design-patterns/convention-over-configuration
  - /fundamentals/glossary/frontmatter
  - https://url/of/the/link
---

# Markdown

`Markdown` is a very lightweight plain text markup language for writing documentation. It is easy to read, easy to write, and easy to render to other markup languages like [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html).

Every document needs a few headings, numbered or unnumbered lists, bold, italic, and underline, links, images, tables, and the like. With `Markdown` these can be specified with simple common characters.

`Markdown` can easily be rendered to [HTML](/fundamentals/design-and-architecture/standards-based/data-standards/html), for example, and there is also tooling that can render it to Word documents or [PDF](/fundamentals/design-and-architecture/standards-based/data-standards/pdf) files.

By [Convention](/fundamentals/design-and-architecture/standards-based/design-patterns/convention-over-configuration) the file extension for `Markdown` files is `.md`.

[Frontmatter](/fundamentals/glossary/frontmatter) is an "extension" of `Markdown` that allows for adding meta data.

By default links in `Markdown` use the following syntax:

```md
[Text of the link](https://url/of/the/link)
```

More modern tools support the `Wiki Syntax`:

```md
[[Name of the document]]
```

Here the text of the link and the target of the link are identical to the name of the document that is referenced. This is much easier to read and to type.

For more information on the `Markdown` syntax, see https://www.markdownguide.org/basic-syntax/.
