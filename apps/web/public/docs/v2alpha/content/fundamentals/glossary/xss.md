---
title: XSS
pcx_content_type: definition
summary: >-
  Cross-site scripting (XSS) is a type of security vulnerability that can be found in some web applications. XSS attacks enable attackers to inject client-side scripts into web pages viewed by other users.
hidden: true
has_more: true
has_links: false
---

# Cross-Site Scripting (XSS)

Cross-site scripting (XSS) is a type of security vulnerability that can be found in some web applications. XSS attacks enable attackers to inject client-side scripts into web pages viewed by other users. A cross-site scripting vulnerability may be used by attackers to bypass access controls such as the same-origin policy.

One of the oldest examples in the book is where a hacker posts a rich message in a bulletin board, where part of the "rich" message is some JavaScript that sends all the cookies to the hacker's server, which includes the login-tokens, when his message is viewed by other users of that bulletin board. The hacker can then use those login tokens to impersonate that user.

Nowadays several protections are in place to prevent XSS. For those scenarios where there is a use case to allow XSS it can be enabled explicitly.
