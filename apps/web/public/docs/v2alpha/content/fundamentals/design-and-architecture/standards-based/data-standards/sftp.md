---
date_created: 2022-12-24T07:15:53
author: Guillaume Hanique
title: SFTP
pcx_content_type: definition
summary: >-
    `Secure Shell File Transfer Protocol (SFTP)` = [SSH](#ssh) + [FTP](#ftp). Or in other words: `SFTP` adds File Transfer capabilities to something that is already Secure.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/ssh
    - /fundamentals/design-and-architecture/standards-based/data-standards/ftp
    - /fundamentals/design-and-architecture/standards-based/data-standards/ftps
---

# Secure Shell File Transfer Protocol (SFTP)

`SFTP` = [SSH](/fundamentals/design-and-architecture/standards-based/data-standards/ssh) + [FTP](/fundamentals/design-and-architecture/standards-based/data-standards/ftp). Or in other words: `SFTP` adds File Transfer capabilities to something that is already Secure.

One advantage of `SFTP` is that it only needs one port to enable file transfer (while [FTPS](/fundamentals/design-and-architecture/standards-based/data-standards/ftps) requires two), which makes it easier to implement through Firewalls.
