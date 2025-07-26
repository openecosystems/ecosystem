---
title: XDS Affinity Domain
aliases:
    - XDS
    - XDS Domain
    - XDS Affinity Domain
pcx_content_type: definition
summary: >-
    A Cross-Enterprise Document Sharing Domain (XDS Affinity Domain) is a concept used in healthcare. It defines a group of [Providers](/fundamentals/actors/provider) that agree to share [PHI](/fundamentals/glossary/phi) using standard protocols and specifications.
hidden: true
has_more: true
has_links: false
links_to:
    - /fundamentals/actors/provider
    - /fundamentals/glossary/phi
    - /fundamentals/actors/patient
    - /fundamentals/glossary/xds-affinity-domain
    - /fundamentals/glossary/audit-logging
---

# Cross-Enterprise Document Sharing Domain (XDS Domain)

A Cross-Enterprise Document Sharing Domain (XDS Affinity Domain) is a concept used in healthcare. It defines a group of [Providers](/fundamentals/actors/provider) that agree to share [PHI](/fundamentals/glossary/phi) using standard protocols and specifications.

The `Affinity Domain` is the community of [Providers](/fundamentals/actors/provider) that have a common interest in sharing [PHI](/fundamentals/glossary/phi) within a region or network. These [Providers](/fundamentals/actors/provider) typically collaborate to establish a common set of policies and procedures for [PHI](/fundamentals/glossary/phi) exchange.

One would expect the following components to be set up:

-   **Document Sources**: the individual [Providers](/fundamentals/actors/provider) that generate or store [PHI](/fundamentals/glossary/phi).
-   **Document Repositories**: The systems where the shared [PHI](/fundamentals/glossary/phi) is stored.
-   **Document Registry**: Central index or catalog to find shared records.
-   **Patient Indentifier Cross-Reference Manager (PIX Manager)**: Associates the different [Patient](/fundamentals/actors/patient) identifiers used by the [Providers](/fundamentals/actors/provider) within the [XDS Domain](/fundamentals/glossary/xds-affinity-domain).
-   **Audit Trail and Node Authentication (ATNA)**: For [audit-logging](/fundamentals/glossary/audit-logging) and security.

By establishing an [XDS Affinity Domain](/fundamentals/glossary/xds-affinity-domain) [Providers](/fundamentals/actors/provider) can streamline the exchange of [PHI](/fundamentals/glossary/phi), improve care coordination, and enhance overall quality of healthcare services. It's a crucial concept in achieving interoperability and securely sharing [PHI](/fundamentals/glossary/phi) across [Providers](/fundamentals/actors/provider) within a specific region or network.
