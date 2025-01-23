---
title: Terraform
pcx_content_type: definition
summary: >-
  [Terraform](https://www.terraform.io/) is a tool for building, changing, and versioning infrastructure, and provides components and documentation for building Open Ecosystems resources.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/api
  - /fundamentals/glossary/current-state
  - /fundamentals/glossary/desired-state
  - /fundamentals/glossary/iaas
  - /fundamentals/glossary/infrastructure-as-code
---

# Terraform

[Terraform](https://www.terraform.io/) is a tool for building, changing, and versioning infrastructure, and provides components and documentation for building Open Ecosystems resources.

`Terraform` is HashiCorp's [Infrastructure as Code](/fundamentals/glossary/infrastructure-as-code). It allows one to describe a [Desired State](/fundamentals/glossary/desired-state). By deploying it Terraform will make the necessary changes to make sure that the [Current State](/fundamentals/glossary/current-state) matches the Desired State.

## Desired State

With `Terraform` the `Desired State` is described using HCL[^1].

[Former2](https://former2.com/) is a great tool to create such a file from an existing infrastructure.

## State File

`Terraform` makes use of a `State File`. Whenever it deployed a `Desired State`, it will save it in the `State File`. If the `Desired State` is (changed and) deployed `Terraform` will not actually compare the `Desired State` to the `Current State`, but to the `State File`.

### Advantages

Having a `State File` has some advantages:

- It is much easier to compare the `Desired State` to a local file (the `State File`, than it is to the real `Current State`, because the latter will involve a lot of [API](/fundamentals/glossary/api) calls to the [IaaS](/fundamentals/glossary/iaas) or DCaaS layer.
- The `State File` can also be parsed by other tools, which gives them an easy to read model of what the infrastructure currently looks like.

### Disadvantages

Having a `State File` also has some disadvantages:

- The `State File` must be stored on Shared Storage, and it can never get lost, or it would break [Infrastructure as Code](/fundamentals/glossary/infrastructure-as-code). These requirements are not easily met.
- Nothing and nobody else can make changes to the infrastructure outside of `Terraform`. If changes are made, they will NOT be detected by `Terraform`, which results in an infrastructure that does NOT match the `Desired State`, although everything seems to be saying it is.

[^1]: HashiCorp Configuration Language
