---
date_created: 2022-12-11T17:16:47
title: GitHub Flow
aliases:
  - GitHub Flow
pcx_content_type: definition
summary: >-
  `GitHub Flow` is a [Branching Model](/fundamentals/glossary/branching-model) defined by `GitHub`.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/branching-model
  - /fundamentals/design-and-architecture/standards-based/data-standards/gitflow
  - /fundamentals/glossary/pull-request
---

# GitHub Flow

`GitHub Flow` is a [Branching Model](/fundamentals/glossary/branching-model) defined by `GitHub`. It is simpler than [GitFlow](/fundamentals/design-and-architecture/standards-based/data-standards/gitflow) in one sense, but more complex in another.

Instead of having `Feature branches` like in [GitFlow](/fundamentals/design-and-architecture/standards-based/data-standards/gitflow), one `forks` an existing Git Repository. You make the changes you want, which results in a new `master`, which you push to your own Git Repository.

You can then create a [Pull Request](/fundamentals/glossary/pull-request) on the original Git Repository, that basically says: "I made some changes in my own repo, would you care to have a look at it?" The maintainer of the original Git Repository can then pull in the changes from your repo, have a look at it, and if he likes it, he will squash all the commits on your master into one commit which references your fork, add it to his `master` and that is how your feature is added to the main repo.

## Sources

- https://docs.github.com/en/get-started/quickstart/github-flow
