---
date_created: 2022-12-11T17:16:05
title: Semantic Versioning
aliases:
    - Semantic Versioning
    - Versioning
    - Version
    - Versioned
    - Versions
pcx_content_type: definition
summary: >-
    `Semantic Versioning` is an approach to address [Dependency Hell](/fundamentals/glossary/dependency-hell) by making incompatible dependencies more predictable.
hidden: true
has_more: false
links_to:
    - /fundamentals/glossary/dependency-hell
---

# Semantic Versioning

`Semantic Versioning` is an approach to address [Dependency Hell](/fundamentals/glossary/dependency-hell) by making incompatible dependencies more predictable. It basically says that a version number consists of three parts:

-   `Major Number`: Increased when a new version has breaking changes.
-   `Minor Number`: Increased when a new version only has non-breaking changes. A higher Minor Number often means that code has to be recompiled, but code should not have to change.
-   `Build Number`: Increases when fixing issues that do not introduce features.

## Sources

-   https://semver.org/
