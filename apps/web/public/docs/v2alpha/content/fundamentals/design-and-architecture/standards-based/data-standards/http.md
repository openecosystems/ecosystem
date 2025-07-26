---
date_created: 2022-12-11T17:01:28
title: HTTP
pcx_content_type: definition
summary: >-
    `HTTP` is a [Transport Protocol](/fundamentals/glossary/#transport-protocol) for Web Applications.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/rest
    - /fundamentals/design-and-architecture/standards-based/data-standards/url
    - /fundamentals/glossary/get-parameters
    - /fundamentals/glossary/http-method
    - /fundamentals/glossary/http-return-code
    - /fundamentals/glossary/transport-protocol
---

# HTTP

`HTTP` is a [Transport Protocol](/fundamentals/glossary/transport-protocol) for Web Applications.

The `HTTP` protocol has a number of "components":

-   `Request`:
    -   The [URL](/fundamentals/design-and-architecture/standards-based/data-standards/url), which indicates what resource is being accessed.
    -   [Get Parameters](/fundamentals/glossary/get-parameters). These are name/value pairs. They are separated from the `url` with a `?`. There can be multiple name/value pairs after the question mark, which are separated by a `&` character.
    -   `Message Body`, which could be of any Message Protocol.
    -   [HTTP Method](/fundamentals/glossary/http-method), which indicates how the message should be processed. Typical values are:
        -   `GET`: retrieve info,
        -   `PUT`: create something new,
        -   `POST`: create something new, change something, or perform an action,
        -   `DELETE`: delete something,
        -   and there are a few more (See [REST](/fundamentals/design-and-architecture/standards-based/data-standards/rest) API for example).
    -   `Headers`, which can be used to send meta-data. (One that is used quite commonly is `Content-Type`, which indicates what Message Protocol is being used).
-   `Response`:
    -   Same as `Request`, except:
        -   `URL` does not apply,
        -   `Get Parameters` do not apply,
        -   `Method` does not apply
    -   [HTTP Return code](/fundamentals/glossary/http-return-code)

## Sources

-   https://httpwg.org/specs/
-   https://www.w3.org/Protocols/
