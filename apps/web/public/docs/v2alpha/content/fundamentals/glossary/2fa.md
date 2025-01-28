---
title: 2FA
pcx_content_type: definition
summary: >-
  Two-factor authentication (2FA) adds an additional layer of login security to Open Ecosystems accounts by requiring users to provide two mechanisms of authentication instead of just one.
hidden: true
has_more: true
aliases:
  - MFA
links_to:
  - /fundamentals/glossary/brute-force-attack
  - /fundamentals/glossary/dictionary-attack
  - /fundamentals/glossary/hashing
  - /fundamentals/glossary/otp
  - /fundamentals/glossary/password-manager
  - /fundamentals/glossary/social-engineering
  - /fundamentals/glossary/tpm-chip
---

# 2FA

There are three mechanisms for authentication:

| Mechanism  | Meaning                     | Example                                      |
| ---------- | --------------------------- | -------------------------------------------- |
| Inherence  | Something the user _is_.    | Fingerprint, Facial recognition.             |
| Knowledge  | Something the user _knows_. | Open Ecosystems password.                    |
| Possession | Something the user _has_.   | An authentication code from a mobile device. |

Two-factor authentication (2FA) adds an additional layer of login security to Open Ecosystems accounts by requiring users to provide two mechanisms of authentication instead of just one.

## Inherence

`Inherence` is a mechanism of authentication where a user is identified by what he is. The most common examples are Fingerprint Recognition and Facial Recognition, both of which can be achieved with relatively cheap hardware. High security facilities could be protected by Retinal Scans, but the hardware to make that possible is not available to the general public, which is why you will typically find this in movies.

Though Inherence is pretty hard to circumvent (after all, mimicking what a User _is_, is not easy), its vulnerability lies in that a User can be tricked in to Authenticating when he shouldn't. To give an example: Users are conditioned to put their finger on the Fingerprint Reader whenever a request to do so pops up. They are conditioned to do so, even if they are not informed of what they are Authenticating for, or without evaluating the information if they are informed.

### Fingerprint Recognition

Everyone's fingerprint is unique. Therefore, if one can recognize one's fingerprint, one can uniquely identify a person and that identification can be used to Authenticate an individual.

On Devices that have a Fingerprint Reader, the fingerprint can then be used to unlock the key that is stored in the [TPM Chip](/fundamentals/glossary/tpm-chip), which, in turn, can be used for Authentication. The model to recognize the fingerprint is stored locally on the Device.

### Facial Recognition

Though Facial Recognition is often associated to movies where people are traced because their faces show up on street cameras, Facial Recognition can also be used for Authentication to uniquely identify an individual (as everyone's face is unique).

If a Device has a camera, a complex algorithm is used to process an image that is captured by that camera to match specific features of one's face against a previously stored model. If that succeeds, that outcome can be used to unlock the key that is stored in the [TPM Chip](/fundamentals/glossary/tpm-chip), which, in turn, can be used for Authentication. The model to recognize the fingerprint is stored locally on the device.

The strength of this form of Authentication depends on how advanced the algorithm is used process the facial features. Poor ones can be tweaked with a photograph of the person. Because those algorithms are usually built by white people, they intend to work better on white people than on black people.

## Knowledge

`Knowledge` is an Authentication Mechanism that is based on what the User knows. Typical examples are:

- A password
- Security Questions
- PIN Code
- Security Pattern

### Passwords

Passwords are a way of Authentication where the User Authenticates using a series of characters that he only knows. Even though this form of Authentication has been used the longest, and is still the most widely used form of Authentication today, a Password has several security vulnerabilities:

- Because Passwords are hard to remember, people tend to use the same Password everywhere. As a consequence, once the password has been compromised, a hacker not only has access to one Web Application the User uses, but to many.
- Passwords should be [Hashed](/fundamentals/glossary/hashing) and then stored. This would make it impossible to recover the Password from the data that was stored in the system. However, not every Website applies this pattern. These Websites are also typically easy to hack (if they had sufficient Security mechanisms in place, they would probably also have hashed the passwords before they stored them). A compromise on one of these websites reveals a User's Password, which would also give access to other sites where the User used the same Password.
- Because Passwords are hard to remember, people tend to use weak Passwords. Weak Passwords are easier to remember, but consequently also easier to guess. Websites tried to mitigate this problem by requiring people to use lowercase, uppercase, numbers, and "special characters". This, however, did not have the expected result. See the video below. This makes many passwords susceptible to [Dictionary Attacks](/fundamentals/glossary/dictionary-attack) and [Brute Force Attacks](/fundamentals/glossary/brute-force-attack).
- [Password Managers](/fundamentals/glossary/password-manager) are another approach to addressing weak passwords and reuse. The idea is that the Password Manager generates strong Passwords, where it remembers what website uses what passwords. This approach is slightly better than using the same Password on multiple websites: if the Password Manager's password is compromised, so are all Passwords that are stored in it. Another disadvantage is that if the User would lose access to the Password Manager, he would lose access to all the websites he used it for.

{{%youtube id="aHaBH4LqGsI"%}}

### Security Questions

Security Questions are a form of Authentication where the User, during the registration process, provides a number of Questions of which only he knows the answer. If, at a later time, the system needs to Authenticate the User, it can ask one or more (usually more) Security Questions, and compare the answers to the one that were given before.

This form of Authentication has one major security flaw: it is very vulnerable to [Social Engineering](/fundamentals/glossary/social-engineering). Through skillful manipulation of conversations hackers, or simply by scanning Social Media, hackers can retrieve or deduce answers to many Security Questions, potentially enough to get through the Authentication process.

One form of Social Engineering that many people are familiar with (without realizing it), are "games" on Facebook, for example: "Answer a few simple questions and we will tell you what kind of person you will marry!". What many people don't realize is that by logging in to that "game" and by answering those questions, they are actually volunteering information that is associated with their email address, which can be used to break Authentication using Security Questions.

### PIN Code

Using a PIN Code for Authentication may seem unintuitive: how can a (typically) 4-digit number be more secure than a strong Password? The answer lies in how it is implemented.

A PIN Code cannot be used to login to a remote system. It can only be used to login to a local system, like a mobile device or a specific Laptop. After a User successfully logged in using other Authentication Mechanisms, he receives a validated Identity Token. The Identity Token can subsequently be stored on the device, and protected by a PIN Code. If, at a later point in time, the User enters the correct PIN Code, it will unlock the Identity Token, which can then be used to interact with the system.

Another reason why a PIN Code is pretty secure, is because most of the time it is not the only form of protection: the devices themselves are usually also protected using a PIN Code, Security Pattern, Facial Recognition, or Fingerprint Recognition.

Secure implementations will use the [TPM Chip](/fundamentals/glossary/tpm-chip) to protect the Identity Token (which adds device specific metrics to the token, so that the same token cannot be used on another device). Less Secure implementations simply store the Identity Token in a store that is local to the application and manually check if the PIN Code matches the one that has been entered before.

Either way this method of Authentication is generally more secure than a Password, simply because of the fact that the PIN Code does not have to be transmitted to a remote server for verification.

Even storing the PIN Code in an application specific store is generally secure, because both iOS and Android have nearly unbreakable isolation of applications. Only if a device is Rooted (which is only possible on Android) can this Authentication Mechanism be compromised, and then only if the [TPM Chip](/fundamentals/glossary/tpm-chip) is not used. (This is also the reason why Google Pay disables itself on Rooted devices).

### Security Pattern

Instead of entering a Password or PIN Code, one could also draw a Security Pattern. Security Patterns are easier to remember and enter. At the time of writing, Authentication using a Security Pattern can only be found on Android.

Just like with a PIN Code a Security Pattern can only be used to Authenticate in a local device. If the pattern is entered correctly, it releases the validated Identity Token, which can then be used to interact with the system. If one uses the mechanism provided by Android then one also uses the [TPM Chip](/fundamentals/glossary/tpm-chip).

On one side a Security Pattern is more secure than a PIN Code, as there are more possible patterns then there combinations of (your) digits.

But a Security Pattern also has some Security Risks. Unfortunately "easier to remember" and "easier to draw", also implies "easier to steal". According to Why iPhones Don't Have Pattern Unlock[^1], someone looking over your shoulder can recognize and memorize the pattern 64% of the time by just looking once, as opposed to 11% when a (6 digit) PIN Code is entered (probably because a pattern is just one thing to remember, whereas a 6 digit PIN are 6 things to remember).

## Possession

Possession is an Authentication Mechanism that is based on something the User _has_. There are many examples:

- OTPs
- Smart Cards
- Security Keys
- Authenticator Apps

### OTP

An OTP is a one-time password. See [OTP](/fundamentals/glossary/otp) for more information.

### Smart Card

A Smart Card is a device that produce a rotating [OTP](/fundamentals/glossary/otp) or Challenge Code using Cryptography. A remote system has the same Cryptographic keys, which allows it to validate the Challenge Codes that are submitted by a User.

### Security Key

A Security Key depends upon a primary device such as a computer. A Security Key may be inserted into a USB port for Authentication.

### Authenticator Apps

An Authenticator App is Mobile Application that can produce time-based OTPs, that can subsequently be used for Authentication. During setup the service provider generates a secret key that is specific for that User. This is sent to the Authenticator App either as a Base32 string, or as a QR-Code. After registering the key, the Authenticator App uses that key to Encrypt time, of which a portion is extracted and displayed to the user as a six-digit code.

Since the server produced the key, it can use the same key, encrypt the same time, and get the same six-digits. This way the six-digits serve as an [OTP](/fundamentals/glossary/otp).

## Sources

- https://docs.microsoft.com/en-us/windows/security/information-protection/tpm/how-windows-uses-the-tpm
  [^1]: https://www.youtube.com/watch?v=OPZMNtAW4MM
