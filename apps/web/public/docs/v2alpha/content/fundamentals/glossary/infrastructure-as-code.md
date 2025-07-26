---
title: Infrastructure as Code
pcx_content_type: definition
summary: >-
    Infrastructure as Code is the process of managing and provisioning computer Data Centers through machine-readable declaration files, rather than physical hardware configuration or interactive configuration.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/aws
    - /fundamentals/glossary/terraform
---

# Infrastructure as Code

Infrastructure as Code is the process of managing and provisioning computer Data Centers through machine-readable declaration files, rather than physical hardware configuration or interactive configuration.

## Achieving Infrastructure as Code

There are two ways to achieve Infrastructure as Code:

-   Have a Desired State, a tool that can determine the Current State, and a tool that can detect the differences and determine and execute actions to transform the Current State to became identical to the Desired State.
-   Declare Idempotent actions that, regardless of how many times they are executed, will always result in the same (Desired) State.

## Implementations

Examples of Infrastructure as Code:

-   Ansible
-   [AWS CloudFormation](/fundamentals/glossary/aws#CloudFormation)
-   CDK Terraform
-   CDK v1
-   CDK v2
-   CDK
-   Pulumi
-   [Terraform](/fundamentals/glossary/terraform)
-   Troposphere
