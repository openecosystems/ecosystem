---
title: AWS
pcx_content_type: definition
summary: >-
  `Amazon Web Services (AWS)` is the biggest [CSP](/fundamentals/glossary/#csp) at the time of writing.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/caching
  - /fundamentals/glossary/cdn
  - /fundamentals/glossary/certificate
  - /fundamentals/glossary/csp
  - /fundamentals/glossary/ddos-attack
  - /fundamentals/glossary/iam
  - /fundamentals/glossary/infrastructure-as-code
  - /fundamentals/glossary/kubernetes
  - /fundamentals/glossary/s3
  - /fundamentals/glossary/service-mesh
  - /fundamentals/glossary/sqs
  - /fundamentals/glossary/waf
---

# AWS

`Amazon Web Services (AWS)` is the biggest [CSP](/fundamentals/glossary/csp) at the time of writing.

Below are some services that are provided by AWS.

## Athena

AWS's Serverless Interactive Query Service.

Helps parse CSV files or Parquet files.

Uses SQL dialect.

AWS Athena is expensive, though. Executing a query against `1 TB` of data will take `236 seconds` at a Cost of `$5.75`.[^1]

## Aurora

AWS Aurora is AWS's Serverless implementation of Relational Databases. It comes in the following flavors:

- Aurora MySQL
- Aurora Postgres

With AWS RDS you could provision both AWS Aurora (MySQL) and plain MySQL. AWS Aurora makes Scaling much easier for a Price. But if you don't need it, MySQL is cheaper. (The same goes for AWS Aurora (PostgreSQL) and PostgreSQL, obviously).

## Beanstalk

Orchestration Engine in AWS to deploy infrastructure and applications.

## CloudFormation

AWS CloudFormation is AWS's service for [Infrastructure as Code](/fundamentals/glossary/infrastructure-as-code). It allows one to declare the Desired State of an infrastructure, and AWS CloudFormation will then make all the necessary changes to make the Current State identical to the Desired State.

In AWS CloudFormation the Desired State is defined is an AWS CloudFormation Template.

## CloudFront

AWS CloudFront is AWS's implementation of a [CDN](/fundamentals/glossary/cdn).

## CloudWatch

AWS CloudWatch is AWS's Cloud Monitoring System. It can both used for capturing and storing Monitoring data, as well as using it to generate Alerts.

## Cognito

AWS Cognito lets you add user sign-up, sign-in, and access control to your web and mobile apps quickly and easily. Amazon Cognito scales to millions of users and supports sign-in with Social Identity Providers, such as Apple, Facebook, Google, and Amazon, and Enterprise Identity Providers via SAML 2.0 and OpenID Connect.

## DynamoDB

AWS DynamoDB is a Cloud NoSQL Database.

## EC2

Elastic Compute Cloud (EC2) allows users to rent Virtual Machines on which they can run their own computer applications.

## EKS

Amazon EKS is a managed [Kubernetes](/fundamentals/glossary/kubernetes) service to run Kubernetes in the AWS cloud and on-premises data centers. In the cloud, Amazon EKS automatically manages the availability and scalability of the Kubernetes control plane nodes responsible for scheduling containers, managing application availability, storing cluster data, and other key tasks. With Amazon EKS, you can take advantage of all the performance, scale, reliability, and availability of AWS infrastructure, as well as integrations with AWS networking and security services. On-premises, EKS provides a consistent, fully-supported Kubernetes solution with integrated tooling and simple deployment to AWS Outposts, virtual machines, or bare metal servers.

## ElastiCache

AWS ElastiCache can provision either a Memcached or Redis instance to provide [Caching](/fundamentals/glossary/caching) capabilities.

## EventBridge

AWS EventBridge is a serverless event bus that lets you receive, filter, transform, route, and deliver events.

## IAM

AWS's implementation of [IAM](/fundamentals/glossary/iam) to securely manage identities and access to AWS Services and Resources.

## KMS

AWS KMS is AWS's Key Management System, that allows one to create and control keys and [certificates](/fundamentals/glossary/certificate) that can be used to encrypt or digitally sign data.

## Lambda

Event-driven serverless computing platform on AWS.

## RDS

AWS Relational Database Service (AWS RDS) allows for provisioning Relational Databases:

- AWS Aurora
- MariaDB
- MySQL
- Oracle Database
- Postgres

## Route 53

Route 53 is AWS's implementation of DNS. It allows for registering hostnames in Private Networks as well as Public Networks.

Route 53 is not to be confused with a [Service Mesh](/fundamentals/glossary/service-mesh). Though both manage hostnames and IP Address, both solve a different problem. A Service Mesh does not use Route 53.

## S3

AWS Simple Storage Service (AWS S3). Files are stored in `Buckets`, which can have `Folders` which can have `Files`. For more information, see [S3](/fundamentals/glossary/s3).

## SageMaker

AWS SageMaker is a cloud machine-learning platform that was launched in November 2017. SageMaker enables developers to create, train, and deploy machine-learning models in the cloud. SageMaker also enables developers to deploy ML models on embedded systems and edge-devices.

## SES

AWS SES is a cost-effective, flexible, and Scalable email service that enables developers to send mail from within any application. You can configure Amazon SES quickly to support several email use cases, including transactional, marketing, or mass email communications. Amazon SES's flexible IP deployment and email authentication options help drive higher deliverability and protect sender reputation, while sending analytics measure the impact of each email. With Amazon SES, you can send email securely, globally, and at scale.

## Shield

AWS Shield provides protection against a [DDoS Attack](/fundamentals/glossary/ddos-attack). It's one of the most expensive AWS Services.

## SNS

AWS's `Notification` service, a.k.a. `Topics`.

Whenever a message is `posted` to a `topic`, any instance that is listening to that `topic` will receive a copy of the message. If nothing is `consuming` messages from that `topic`, the message will get lost.

This is the opposite of AWS SQS, where only **one** `consumer` will receive a message and messages **never** get lost.

## SQS

AWS SQS is AWS' Queue Service. This service allows for sending and receiving messages over a Queue. For more information, see [SQS](/fundamentals/glossary/sqs).

## WAF

AWS Web Application Firewall is AWS's implementation of a [WAF](/fundamentals/glossary/waf).

[^1]: https://aws.amazon.com/blogs/big-data/analyzing-data-in-s3-using-amazon-athena/
