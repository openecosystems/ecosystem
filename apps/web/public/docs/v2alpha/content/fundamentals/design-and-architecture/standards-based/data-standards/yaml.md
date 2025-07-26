---
date_created: 2022-12-11T17:08:10
author: Guillaume Hanique
title: YAML
aliases:
    - YAML
pcx_content_type: definition
summary: >-
    `YAML` addresses the same domain as [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/#json). One can describe the same objects as one could with [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/#json). `YAML` is easier to read by humans, though, and `YAML` supports comments, which [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/#json) does not.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/json
---

# YAML Ain't Markup Language (YAML)

`YAML` addresses the same domain as[ JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json). One can describe the same objects as one could with [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json). `YAML` is easier to read by humans, though, and `YAML` supports comments, which [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json) does not.

A downside of `YAML` could be that its meaning depends on the indentation of each individual line. This could make `YAML` files somewhat hard to read and maintain if there are many levels.

Some rules, to get an idea:

-   `YAML` files should start with `---`.
-   `a: b` indicates a field "a" with the value "b".
-   `- item 1` indicates an item of an Array.
-   `a:` is the beginning of a Dictionary or Array. On the next line increase the indentation, and either declare a field or an item.
-   Anything after a `#` is a comment.
-   Values usually don't have to be quoted. It is only necessary if it contains special characters like a `:`.
-   The data type of a value is determined by its content (but could be set explicitly, which rarely happens). I.e. "John" is a string, "true" and "yes" are booleans, "3.5" is a floating point, etc.
-   `YAML` supports multi-line strings.
-   `YAML` supports IDs and References.

{{%Aside type="warning"%}}

`YAML` files often contain version number. Be careful to quote them, otherwise they will be interpreted as a number (and "3.10" will become "3.1").

{{%/Aside%}}
{{<raw>}}<pre class="CodeBlock CodeBlock-with-rows CodeBlock-scrolls-horizontally CodeBlock-is-light-in-light-theme CodeBlock--language-txt" language="txt"><code><span class="CodeBlock--rows"><span class="CodeBlock--rows-content"><span class="CodeBlock--row"><span class="CodeBlock--row-indicator"></span><div class="CodeBlock--row-content"><span class="CodeBlock--token-plain">yaml</span></div></span></span></span></code></pre>{{</raw>}}

## References

-   https://yaml.org/spec/1.2.2/
