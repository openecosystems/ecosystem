---
date_created: 2022-12-01T11:43:28
title: User Error
pcx_content_type: definition
summary: >-
  A `User Error` is an error made by the human User of a complex system, usually a computer system. Also known as `PEBMAC`, `ID-10-T`, `PICNIC`, or `IBM Error`. One should replace the User and try again.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/exception-handling
  - /fundamentals/glossary/happy-path
---

# User Error

A `User Error` is an error made by the human User of a complex system, usually a computer system. Also known as `PEBMAC`[^1], `ID-10-T`, `PICNIC`[^2], or `IBM Error`[^3]. One should replace the User and try again.

There is no such thing as a `User Error`. There is nothing that a User should be able to do that would make the system behave unexpectedly. What it really means is that the system doesn't have proper [Exception Handling](/fundamentals/glossary/exception-handling), or that the system was not designed well enough, only taking [Happy Paths](/fundamentals/glossary/happy-path) into account, for example.

[^1]: Problem Exists Between Monitor And Chair (PEBMAC)
[^2]: Problem In Chair, Not In Computer (PICNIC)
[^3]: Idiot Behind Machine Error (IBM Error)
