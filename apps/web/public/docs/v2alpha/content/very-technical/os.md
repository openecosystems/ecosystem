---
title: Operating System
pcx_content_type: overview
weight: 5
---

## Operating System

For local development and quick proof of concepts, you can run a Connector outside a Trusted Execution Environment (TEE)
on your local machine directly. However, we strongly recommend that all production connectors use Ecosystem Containers.

We have chosen to partner with Google's gVisor team to intercept all syscalls to your underlying operating system's kernel.

We support running Connector Enclave Applications as part of a special OCI runtime bundle specially designed for
Trusted Execution Environments with a full Linux kernel and a Linux distribution.

## Developer Experience

Using the OCI runtime bundle as a compatibility layer provides a more familiar development experience and more flexibility.
Ecosystem Containers supports multiple CPUs in the TEE and is therefore suitable for cases with high performance requirements.

## Performance

By using an OCI runtime with gVisor, we add additional layers of defense to our Containers, however, this comes with some performance overhead.
GVisor offers some guidance on how to reduce the overhead based on your goals.

## Linux

We require Linux because it gVisor requires direct access to the Linux kernel to function as a secure user-space kernel.
