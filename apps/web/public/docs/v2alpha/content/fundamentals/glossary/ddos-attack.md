---
title: DDoS attack
pcx_content_type: definition
summary: >-
    A Distributed Denial of Service (DDoS) attack is a malicious attempt to disrupt normal traffic of a targeted server, service, or network by overwhelming the target or its surrounding infrastructure with a flood of Internet traffic.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/dos-attack
    - /fundamentals/glossary/csp
    - /fundamentals/glossary/aws
---

<!-- This document is an original CloudFlare Document from which the cloudflare links are removed. -->

# DDoS attack

A Distributed Denial of Service (DDoS) attack is a malicious attempt to disrupt normal traffic of a targeted server, service, or network by overwhelming the target or its surrounding infrastructure with a flood of Internet traffic.

A DDoS Attack happens when a Hacker has hijacked a number of computers (a.k.a. a "Botnet") and uses them to send many requests a single end-point. The intent is to send so many requests to a platform that it cannot serve normal requests anymore, thus disabling the service that platform provides.

It is disproportionally hard to protect against this kind of attack. With a normal [Denial of Service Attack](/fundamentals/glossary/dos-attack) one could block all requests from certain IP Addresses, but because with a DDoS Attack requests are coming from everywhere, this is not feasible.

Some [CSPs](/fundamentals/glossary/csp) have services that can be used to protect against a DDoS Attack. One example is [AWS Shield](/fundamentals/glossary/aws#shield).
