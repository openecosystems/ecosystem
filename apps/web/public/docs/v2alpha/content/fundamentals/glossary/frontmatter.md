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
{{<raw>}}<pre class="CodeBlock CodeBlock-with-rows CodeBlock-scrolls-horizontally CodeBlock-is-light-in-light-theme CodeBlock--language-txt" language="txt"><code><span class="CodeBlock--rows"><span class="CodeBlock--rows-content"><span class="CodeBlock--row"><span class="CodeBlock--row-indicator"></span><div class="CodeBlock--row-content"><span class="CodeBlock--token-plain">txt</span></div></span></span></span></code></pre>{{</raw>}}

`Frontmatter` is a block of text at the top of the [Markdown](/fundamentals/design-and-architecture/standards-based/data-standards/markdown) file, and can be provided as:

-   [YAML](/fundamentals/design-and-architecture/standards-based/data-standards/yaml): enclosed within `---` and `---`.
-   [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json): enclosed within `{` and `}`.
-   [TOML](/fundamentals/design-and-architecture/standards-based/data-standards/toml): enclosed within `+++` and `+++`.
