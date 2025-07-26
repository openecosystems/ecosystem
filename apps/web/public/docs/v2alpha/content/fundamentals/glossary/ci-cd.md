---
title: CI/CD
aliases:
    - CI/CD
    - CI
    - CD
    - CI Pipeline
    - CI Pipelines
    - CD Pipeline
    - CD Pipelines
    - Continuous Integration / Continuous Deployment (CI/CD)
pcx_content_type: definition
summary: >-
    "`CI/CD` is the combined practices of continuously merging all developer's working copies into a shared mainline (including building, testing, validation, and [Versioning](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning)), and automatically deploying the build output."
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning
    - /fundamentals/glossary/pull-request
    - /fundamentals/glossary/artifact-repository
    - /fundamentals/glossary/dtap
---

# Continuous Integration / Continuous Deployment (CI/CD)

`CI/CD` is the combined practices of continuously merging all developer's working copies into a shared mainline (including building, testing, validation, and [Versioning](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning)), and automatically deploying the build output.

## Continuous Integration (CI)

With `CI` a developer creates a new [Pull Request](/fundamentals/glossary/pull-request) whenever he believes his code changes are ready to be merged to the main branch. This triggers the `CI Pipeline` to perform various activities, among which:

-   **Formatting**: Check that the code meets formatting standards.
-   **Linting**: Ensure that the code meets the standards of the programming language in question.
-   **Building**: Build the code into an executable artifact.
-   **Testing**: Execute all the unit tests and make sure they pass.

If any of these steps fail, then the developer is notified that some code changes are required, and he is required to update the [Pull Request](/fundamentals/glossary/pull-request).

If everything checks out, then optionally there is an extra step where another developer has to review the changes. After approval the `CI Pipeline` will generate a new [Version](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) and upload the built artifact to an [Artifact Repository](/fundamentals/glossary/artifact-repository).

## Continuous Deployment (CD)

`CD` is an approach in which software components are frequently deployed. For `CD` to provide value, the entire deployment process MUST be automated. It must be able to retrieve a specific [Version](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) of a software component from the [Artifact Repository](/fundamentals/glossary/artifact-repository), and deploy it in a specific [Environment](/fundamentals/glossary/dtap).

The `CD Pipeline` should be able to deploy different [Versions](/fundamentals/design-and-architecture/standards-based/data-standards/semantic-versioning) of a software component into different [Environments](/fundamentals/glossary/dtap).

Next to being able to deploy a software component, the `CD Pipeline` must also be able to deploy configuration that is [Environment](/fundamentals/glossary/dtap) specific.

It is also good practice that the `CD Pipeline` is able to execute integration tests to assert that the [Environment](/fundamentals/glossary/dtap) is still sound after the new software component is deployed, and is able to roll back if integration tests fail.
