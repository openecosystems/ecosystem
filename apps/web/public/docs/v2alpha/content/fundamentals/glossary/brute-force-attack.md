---
title: Brute Force Attack
pcx_content_type: definition
summary: >-
  A Brute Force attack consists of an attacker submitting many passwords or passphrases with the hope of eventually guessing correctly, in order to gain illegal access to confidential data.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/2fa
  - /fundamentals/glossary/api-gateway
  - /fundamentals/glossary/aws/#shield
  - /fundamentals/glossary/back-end
  - /fundamentals/glossary/http-return-code
  - /fundamentals/glossary/ip-address
  - /fundamentals/glossary/metric
  - /fundamentals/glossary/rate-limiting
  - /fundamentals/glossary/waf
---

# Brute Force Attack

A Brute Force attack consists of an attacker submitting many [passwords](/fundamentals/glossary/2fa/#passwords) or passphrases with the hope of eventually guessing correctly, in order to gain illegal access to confidential data.

## Protecting against Brute Force Attacks

### Hide cause of error

Hackers buy lists of usernames, and they will use those lists to see if they can be used to gain access to a system.

A Brute Force Attack is often used to guess a User's username / [Password](/fundamentals/glossary/2fa/#passwords) combination. Some systems help the hacker by telling him whether the username was wrong ("No user can be found with this email address"), or whether the [Password](/fundamentals/glossary/2fa/#passwords) was wrong ("Incorrect password"). By giving a hacker this information, the system is basically telling the hacker whether he should try a Brute Force Attack for this username, or whether he should continue searching for usernames that exist on this system.

Systems should never expose this kind of information. Instead they should merely state that the username / Password combination is incorrect. (One should also make sure that none if this information is exposed by the [HTTP Return code](/fundamentals/glossary/http-return-code) of the [Back End](/fundamentals/glossary/back-end) services).

### Account locking

If multiple login attempts have failed for a given username, then the corresponding User account should be locked. This prevents that a hacker gains access to the account, even if he guesses the [password](/fundamentals/glossary/2fa/#passwords) correctly.

### Metric Monitoring

[Metrics](/fundamentals/glossary/metric) could be generated from a login-service that report how often what [IP Address](/fundamentals/glossary/ip-address) did a login attempt. If that [Metric](/fundamentals/glossary/metric) exceeds a certain threshold, that [IP Address](/fundamentals/glossary/ip-address) should be blocked from accessing the system.

This could be implemented with advanced solutions like [WAFs](/fundamentals/glossary/waf) or [AWS Shield](/fundamentals/glossary/aws/#shield), but enforcing a [Rate Limit](/fundamentals/glossary/rate-limiting) using an [API Gateway](/fundamentals/glossary/api-gateway) could also do the trick.
