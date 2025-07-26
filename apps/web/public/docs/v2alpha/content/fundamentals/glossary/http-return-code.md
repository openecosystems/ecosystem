---
title: HTTP Return Code
pcx_content_type: definition
summary: >-
    When a `HTTP Request` is processed by a Web Application, a number of things can happen while processing the request. Depending on what happens, the `Web Application` will return a different `HTTP Code`.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/http-200
    - /fundamentals/glossary/http-404
    - /fundamentals/glossary/http-429
---

# HTTP Return Code

When a `HTTP Request` is processed by a Web Application, a number of things can happen while processing the request. Depending on what happens, the `Web Application` will return a different `HTTP Code`.

-   `200` means [HTTP-OK](/fundamentals/glossary/http-200), which indicates a happy flow when processing the request.
-   The 300-range contains various redirect messages.
-   Errors in the 400-range, `4XX`, mean that the request was understood, but there was something wrong with the request. I.e., perhaps the resource that was identified by the URI did not exist ([404](/fundamentals/glossary/http-404)), or the requester does not have permission to access the resource (`403`), or perhaps there were so many requests in too short a time that this request could not be processed ([429](/fundamentals/glossary/http-429)).
-   Errors in the 500-range, `5xx`, mean that something went wrong while executing the request. I.e., it could be that a backend service that is required for processing the request is offline. (`502`).

See these pages for more information on a few specific return codes:

-   [HTTP-200](/fundamentals/glossary/http-200)
-   [HTTP-404](/fundamentals/glossary/http-404)
-   [HTTP-429](/fundamentals/glossary/http-429)
