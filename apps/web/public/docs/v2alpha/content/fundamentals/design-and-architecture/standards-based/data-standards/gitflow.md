---
date_created: 2022-12-11T17:16:55
title: GitFlow
aliases:
    - GitFlow
pcx_content_type: definition
summary: >-
    `GitFlow` is a [Branching Model](/fundamentals/glossary/branching-model) that does not depend on creating forks of central repositories. Instead, it creates separate timelines by agreeing on specific `branch names`.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/github-flow
    - /fundamentals/glossary/branching-model
    - /fundamentals/glossary/ci-cd
    - /fundamentals/glossary/production-environment
---

# GitFlow

`GitFlow` is not to be confused with [GitHub flow](/fundamentals/design-and-architecture/standards-based/data-standards/github-flow). They are completely different things.

`GitFlow` is a [Branching Model](/fundamentals/glossary/branching-model) that does not depend on creating forks of central repositories. Instead, it creates separate timelines by agreeing on specific `branch names`. It is "invented" by Vincent Driessen, who is Dutch.

## Application

`GitFlow` was documented back in 2010 when [CI/CD](/fundamentals/glossary/ci-cd) was not yet that big a thing. If you're heavily using [CI/CD](/fundamentals/glossary/ci-cd) a simpler [Branching Model](/fundamentals/glossary/branching-model) like [GitHub flow](/fundamentals/design-and-architecture/standards-based/data-standards/github-flow) might be more appropriate.

However, if your product is explicitly versioned and released, or if you need to support multiple versions of your product in the wild, then `GitFlow` is probably as relevant as it was in 2010, and probably more appropriate than [GitHub flow](/fundamentals/design-and-architecture/standards-based/data-standards/github-flow).

## Branching

The main branch is still called `master`. The idea is that every commit on `master` is something that at one point in time was actually running in the [Production Environment](/fundamentals/glossary/production-environment).

`master` is split to another branch that is called `develop`. Developers will create new features, and when they are finished they will merge them to the `develop` branch. The idea of this branch is that every commit is a **fully functional** version of the Application.

Now if a developer wants to build a new feature, he creates a `feature branch` on top of the current `develop` branch. So if multiple developers are working on the same Application, `develop` will have multiple `feature branches`. By convention `feature branches` are named "feature/feature_name", where "feature_name" is either a reference to a (Jira) ticket, a description of the feature, or both.When a developer finished a feature and tested it, he will merge that `feature branch` on top of `develop`. If that is not possible, because another developer merged his feature on top of `develop` first, then the developer should rebase his feature on top of what now is the `develop` branch, then merge with `develop` and push it to the central Git Repository. (It is good practice to check regularly of other developers have merged changes to `develop` and rebase your feature often, so that you can incorporate the changes others have made, so that you can be sure that your feature will still be working after you have merged it with all the features that the other developers have created).

If the new features are ready to be released, a `release branch` will be created. By convention the name is "release/x.y.z", where "x.y.z" is the version number of the new release. Meanwhile developers can continue creating new features and put them on the `develop` branch. The `release branch` will be used for QA. If bugs are discovered, they are fixed on the `release branch`. If everything checks out, the release branch can be "Finished". The following things will happen:

-   The `release branch` is merged with `master`, creating a new `master`. The new `master` contains all the features that were added to `develop` since the last release, as well as all the bug fixes from the `release branch`.
-   The new `master` will be tagged with the version that is in the `release branch`.
-   The `release branch` will be merged with `develop`, to make sure that the bug fixes will also end up in future releases. So the new `develop` will contain the bug fixes, as well as all the new features that did not end up on `master`.

Now suppose there is an issue in the Production Environment that requires a code change. What happens then is that a `hotfix branch` will be created on top of the current `master`. The name will be "hotfix/x.y.z", where "x.y.z" is slightly higher than the tag of the current `master`. The changes that will fix the bug will then be committed to the `hotfix branch` and tested. If everything checks out, the `hotfix branch` will be "Finished". The following things will happen.

-   The `hotfix branch` is merged with `master`, creating a new `master`, which is identical to the old one, but the hotfix. So no features that have been developed by developers will end up in the new `master`.
-   The new `master` is tagged with the version number that is in the `hotfix branch`.
-   The `hotfix branch` is merged with `develop`, so that the hotfix will also be applied in future releases. So the new `develop` will contain all the features that were not yet released, **and** the hotfix.

## Git Extension

`GitFlow` is an official `Git Extension`. It is automatically installed with most distributions, and if it is not, it can be added.

That it is an extension means that you can execute commands like

-   `git feature start featurename`, and
-   `git release finish 1.2.0`,

and then `git` will do all the things described above automatically.

## Sources

-   https://nvie.com/posts/a-successful-git-branching-model/
