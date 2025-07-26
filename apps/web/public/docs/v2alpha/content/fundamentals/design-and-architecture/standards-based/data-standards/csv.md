---
date_created: 2022-12-11T17:06:33
title: CSV
pcx_content_type: definition
summary: >-
    Comma Separated Values (CSV) is a file format for tabular data.
hidden: true
has_more: true
has_links: false
---

# CSV

In a Comma Separated Values (CSV)-file `records` (or: `rows`) are separated by a `newline` character. `Values` are separated by a separator character like ";", "|" (or any other character). (Even if a separator is used that is not a comma, one would still call it a `csv-file`).

Every `record` has the _same_ amount of `fields` in the _same_ order.

`Values` can include `new-line` characters, provided the `Values` are enclosed between "double-quotes". Not every tool is able to parse it, though. Excel cannot. Google Sheets can.

## Advantages

-   Easy to parse
-   Easy to read
-   Easy to understand

## Disadvantages

-   Data types have to be inferred
-   Parsing becomes hard when the data contains commas
-   Column names may or may not be there

## Sources

-   https://www.rfc-editor.org/rfc/rfc4180
