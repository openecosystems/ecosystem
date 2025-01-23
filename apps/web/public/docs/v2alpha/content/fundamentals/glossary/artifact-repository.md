---
title: Artifact Repository
aliases:
  - Artifact Repository
pcx_content_type: definition
summary: >-
  "An `Artifact Repository` stores [Versioned](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) artifacts (or: 'packages)."'
hidden: true
has_more: true
links_to:
  - /fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning
  - /fundamentals/glossary/ci-cd
  - /fundamentals/glossary/immutable
  - /fundamentals/design-and-architecture/standards-based/design-patterns/convention-over-configuration
---

# Artifact Repository

An `Artifact Repository` stores [Versioned](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) artifacts (or: packages), that are generally the output of a [CI](/fundamentals/glossary/ci-cd) built. [CD](/fundamentals/glossary/ci-cd) retrieves these artifacts to deploy them.

`Artifact Repositories` should be [Immutable](/fundamentals/glossary/immutable), which means that once a specific [Version](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) of an artifact has been stored, that it cannot be replaced by another artifact with the same [Version](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning). After all, that would defeat the purpose of [Versioning](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning).

By [Convention](/fundamentals/design-and-architecture/standards-based/design-patterns/convention-over-configuration) the structure of an `Artifact Repository` is:

- `Groups`
  - `Artifacts`
    - [Versions](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning)
      - (Binary)

Or in words:

- The `Group` generally refers to the organization that created the artifact, for example "com.oracle", or "org.apache".
- A `Group` can have multiple `Artifacts`, which refers to a specific component. To illustrate, for the group "org.apache" the name of the artifact could be "log4j".
- Every `Artifact` can have a number of versions, that should follow the [Semantic Versioning](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) guidelines.
- Lastly, every `Version` contains the binary / package.

From the `Group`, `Artifact`, and `Version` one can construct the [[URL]] to download the artifact, for example:

- `https://{RepositoryRoot}/{group.name1}/{artifact}/{version}/{group.name2}-{artifact}-{version}.zip

... where some implementations `{group.name1}` use dots (`.`), while others replace the dots with slashes. `{group.name2}` always uses dots.
