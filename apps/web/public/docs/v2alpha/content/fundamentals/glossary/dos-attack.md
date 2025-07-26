---
title: DoS attack
pcx_content_type: definition
summary: >-
    A Denial of Service (DoS) happens when one client sends more requests than a system can handle, which also impacts other clients. A DoS can be an attack, but it doesn't have to be.<br><br>

    **Related terms:** [DDoS attack](/fundamentals/glossary/#ddos-attack)
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/ddos-attack
    - /fundamentals/glossary/exponential-backoff
    - /fundamentals/glossary/rate-limiting
    - /fundamentals/glossary/retry-mechanism
---

# DoS attack

A Denial of Service (DoS) happens when one client sends more requests than a system can handle, which also impacts other clients.

A DoS can be an attack, but it doesn't have to be. It could also be caused by a poorly implemented [Retry Mechanism](/fundamentals/glossary/retry-mechanism), for example (i.e., a Retry Mechanism without [Exponential Backoff](/fundamentals/glossary/exponential-backoff) . (A [DDoS Attack](/fundamentals/glossary/ddos-attack) also floods the system with requests so that it doesn't work anymore, but that is an attack).

By applying [Rate Limiting](/fundamentals/glossary/rate-limiting) one can prevent a DoS.

**Related terms:** [DDoS attack](/fundamentals/glossary/#ddos-attack)
