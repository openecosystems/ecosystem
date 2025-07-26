---
title: Password Manager
pcx_content_type: definition
summary: >-
    A Password Manager is a computer program or Cloud service that allows users to store, generate, and manage their [passwords](/fundamentals/glossary/2fa/#password).
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/2fa
---

# Password Manager

A Password Manager is a computer program or Cloud service that allows users to store, generate, and manage their [passwords](/fundamentals/glossary/2fa/#passwords).

Passwords have quite a few Security issues, many of which are physiological and psychological physiological in nature (most prominently man's limited capacity for remembering multiple strong Passwords for extensive periods of time, and his unwillingness to compensate for that limitation with effort). As a result Users recycle relatively weak Passwords. Password Managers provide a technical solution to complement man's limitations and allow for using unique strong passwords for various service that require Authentication.

Whenever a User registers for a service and has to provide a Password, the Password Manager will intervene and offer to generate a unique strong Password, which will be stored by the Password Manager. The next time the User wants to login to that same service, again the Password Manager will intervene, and offer to enter the credentials that were generated before.

Though Password Managers solve some problems, they create others:

-   The Password Manager becomes a Single Point of Failure: If the User loses access to the Password Manager, he loses access to all his Passwords. Also, if access to the Password Manager is compromised, then **all** the User's Passwords are compromised. See https://blog.lastpass.com/2022/08/notice-of-recent-security-incident/ for a recent example.
-   Because with a Password Manager entering Passwords is as simple as clicking a button, Users can be tricked into providing credentials when they shouldn't. (Also see Social Engineering and Fingerprint Recognition).
