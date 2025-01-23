---
title: IAM
pcx_content_type: definition
summary: >-
  Identity and Access Management (IAM) is a framework of policies and technologies to ensure that the right users (that are part of the ecosystem connected to or within an enterprise) have the appropriate access to technology resources.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/audit-logging
  - /fundamentals/glossary/aws
  - /fundamentals/glossary/rbac
---

# Identity and Access Management (IAM)

Identity and Access Management (IAM) is a framework of policies and technologies to ensure that the right users (that are part of the ecosystem connected to or within an enterprise) have the appropriate access to technology resources.

[AWS](/fundamentals/glossary/aws) has its own implementation: [AWS IAM](/fundamentals/glossary/aws/#iam). This can be used to implement fine grained control to what permissions principals have within the AWS infrastructure. By itself it can only provide [RBAC](/fundamentals/glossary/rbac) for everything and everyone that has an identity within AWS. But if it is combined with [AWS Cognito](/fundamentals/glossary/aws/#cognito) then also external Users can be authenticated and associated with an AWS Principal, which could grant or deny permissions to use an AWS resource.

I.e.: Instead of using a Service Principal to get permissions to get data from a database, an end-User authenticated could authenticate himself with AWS Cognito using his Google Account, which is associated with an AWS Principal, which has a Policy that grants read permissions to a database table, which allows the Service to read the records from that table on behalf of the end-User. (This could also provide some pretty detailed [Audit Logging](/fundamentals/glossary/audit-logging)).

## Sources

- https://en.wikipedia.org/wiki/Identity_management
