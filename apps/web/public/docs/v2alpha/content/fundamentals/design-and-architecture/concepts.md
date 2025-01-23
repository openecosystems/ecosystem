---
pcx_content_type: reference
title: Concepts
weight: 4
---

# Concepts

## Model

Data is modeled within the Spec as strongly-typed structured data with strongly enforced semantic rules, which allow for portable and type-safe APIs written in any modern programmatic language, such as Golang, Java, Python, CSharp, Typescript, and Rust. The Spec data model extends Google’s protocol buffers technology, which is a binary based protocol that optimizes for small data sizes and efficient network transport across long wire distances. The Spec extends Google’s protocol buffers with semantic rules that allow architects and data owners to enforce invariants (what must always be true) on a particular field, method or service definition. For example, a FHIR resource field can be constrained to be no shorter than 3 characters and no longer than 15 characters. Another resource method can be constrained to prevent data from being decrypted except for explicit consent of the owner.

## Serialize

Once modeled, data can be instantiated and serialized. All data “within” the system gets serialized as binary protocol buffer byte arrays. While data that exits the system can be serialized as either protocol buffers, JSON, XML, CSV, or any text based protocol. This serialization technique supports cross-language communication and storage.

## Encrypt and Decrypt

NIST requires that we encrypt data in transit and at rest, however, once a filesystem is mounted to a piece of software, that software must have the keys to decrypt that data. This poses a major security issue for highly sensitive information where we do not want even system administrators to access. The Spec takes things even further with consent based encryption and decryption.

By default, data is always encrypted at rest and in transit. The Spec approach to consent based encryption and decryption uses a combination of both asymmetric and symmetric encryption, using a technique known as envelope based encryption. Encryption type one is fast, while the second is strong. We need both speed at scale and security at scale. By combining these two approaches, we achieve the best of both worlds. When data is modeled, the architect can elect to define which portions of the model must use consent based decryption using a semantic rule.

## Transact

The network is one of a systems worst enemies. It is like the wind, it does what it wants and introduces failures and unexplained latencies at varying and unpredictable points in the system. Transaction management is one of the most important yet once of the least understood by engineers because modern frameworks “handle” this out of the box. The trouble is that while those frameworks are fantastic at managing a transaction within a process boundary, they do nothing to manage it across boundaries, context or machines.

The Spec places a lot of emphasis on transaction management as it is critical to preventing data loss and acheiveing ACIDity across a globally distributed system. The Spec design pattern optimizes for security, simplicity, stability and speed through an event-driven architecture that supports both strongly and eventually consistent ACID transactions. To achieve eventually consistent ACID Transactions using an event-driven architecture, we use EventPlane Sourcing to persist multi-service transactions, and Command Query Responsibility Segregation (CQRS) to mutate and query the data.

## Transport

Now that we have: (1) modeled our data using strong types and semantic rules; (2) serialized it for speed and portability; (3) encrypted and decrypted using consent; (4) ensured ACIDity across a distributed network; we now will focus on how we transport the data between processes and services, and clients. The Spec supports the following transports by default: GraphQL, REST, GRPC, Server Sent Events, text-based TCP, and binary TCP.

## Store

Explain how we store data in different types of stores: SQL, NoSQL, Warehouse, Networked Cache, Unnetworked Cache,

## Secure

## Audit

Explain through our event plane we achieve an immutable log that is not optional or
controlled by engineers

## Capture changes immutably on a block chain

Explain how using our event plane, we can selectively send data to any block chain for an immutatble and public record

## Exchange

Explain the Spec Connector and how it is used to integrate with any third part system across any transport system

## Observe
