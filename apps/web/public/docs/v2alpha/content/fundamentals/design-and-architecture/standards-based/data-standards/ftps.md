---
date_created: 2022-12-11T17:09:46
title: FTPS
pcx_content_type: definition
summary: >-
    File Transfer Protocol, Secure (FTPS) = [FTP](#ftp) + [SSL](#ssl). Or in other words: `FTPS` adds Security to File Transfer capabilities.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/ftp
    - /fundamentals/design-and-architecture/standards-based/data-standards/ssl
---

# FTPS

`FTPS` = [FTP](/fundamentals/design-and-architecture/standards-based/data-standards/ftp) + [SSL](/fundamentals/design-and-architecture/standards-based/data-standards/ssl). Or in other words: `FTPS` adds Security to File Transfer capabilities.

`FTPS` has the same disadvantage as [FTP](/fundamentals/design-and-architecture/standards-based/data-standards/ftp): it uses two ports, one for commands, one for the data, which makes it hard to implement through Firewalls.
