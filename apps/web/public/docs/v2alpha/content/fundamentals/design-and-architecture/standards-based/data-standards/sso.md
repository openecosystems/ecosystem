---
date_created: 2022-12-11T17:14:29
title: SSO
pcx_content_type: definition
summary: >-
  `SSO` is an Authentication scheme that allows a User to login with a single ID to any of several related, yet independent, software systems.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/message-protocol
  - /fundamentals/design-and-architecture/standards-based/data-standards/saml
  - /fundamentals/glossary/2fa/#passwords
  - /fundamentals/glossary/authentication
  - /fundamentals/glossary/single-point-of-failure
  - /fundamentals/glossary/xss
---

# SSO

`SSO` is an Authentication scheme that allows a User to login with a single ID to any of several related, yet independent, software systems. I.e.: by logging into one's Google Account one can also log in to Trello.

The [Message Protocol](/fundamentals/glossary/message-protocol) that is used for implementing `SSO` is [SAML 2.0](/fundamentals/design-and-architecture/standards-based/data-standards/saml).

## Advantages

- Risk mitigation: a website no longer has to manage credentials for its user, another party is doing that.
- No need for one more [password](/fundamentals/glossary/2fa/#passwords)
- Simpler administration
- Better network security

## Disadvantages

- `SSO` turns the system that provides [Authentication](/fundamentals/glossary/authentication) into a [Single Point of Failure](/fundamentals/glossary/single-point-of-failure): if that system would become unavailable, the User would no longer be able to login into **any** site that uses that `SSO` provider.
- If `SSO` credentials or tokens are compromised, not just one service gets compromised, but many.
- `SSO` can be subject to web filtering. I.e.: many schools block Facebook, which inherently blocks access to any website that allows `SSO` with Facebook.

## Security

`SSO` is relatively Secure. Some vulnerabilities were reported in 2012 and 2014, but none have been discovered since. Guarding against [XSS](/fundamentals/glossary/xss) is crucial, though: in 2020 not having the proper [XSS](/fundamentals/glossary/xss) protections in place allowed for hijacking the `SSO` token, which caused a breach to several federal websites.

`SSO` should also be combined with Single Log-Out (SLO), which makes sure that if a User logged out from the `SSO` provider, that he is also automatically logged out from all the websites that use that `SSO` provider.

## Privacy

Technically `SSO` can work without having to reveal identifying information like an Email Address. But often the User isn't given the choice what information he wants to share.

## Sources

- https://en.wikipedia.org/wiki/Single_sign-on
- https://en.wikipedia.org/wiki/SAML_2.0
