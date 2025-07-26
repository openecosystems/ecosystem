---
date_created: 2022-12-11T16:58:55
title: FHIR
aliases:
    - FHIR
pcx_content_type: definition
summary: >-
    `FHIR` is an Industry Standard for describing Health Care related Artifacts.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/c-cda
    - /fundamentals/design-and-architecture/standards-based/data-standards/hl7v2
    - /fundamentals/design-and-architecture/standards-based/data-standards/hl7v3
    - /fundamentals/design-and-architecture/standards-based/data-standards/json
    - /fundamentals/design-and-architecture/standards-based/data-standards/rest
    - /fundamentals/design-and-architecture/standards-based/data-standards/soap
    - /fundamentals/design-and-architecture/standards-based/data-standards/xml
---

# Fast Healthcare Interoperability Resources (FHIR)

`FHIR` is an Industry Standard for describing Health Care related Artifacts.

It describes Artifacts, and also allows individual companies to extend that data model for their specific Application.

It also describes a number of Operations and specifies [REST API](/fundamentals/design-and-architecture/standards-based/data-standards/rest) end points for them.

The FHIR protocol describes different FHIR Resources for every Health Care artifact. The problem with how they defined the FHIR Resources, is that it still allows different ways of describing an artifact in the FHIR Resource. The result is that there are various FHIR "Dialects".

## FHIR Resources

FHIR has a lot of FHIR Resources, where every FHIR Resource describes a different artifact in relation to Health Care. Every FHIR Resource is a technical data object, which can be rendered in a number of Message Protocols (like [JSON](/fundamentals/design-and-architecture/standards-based/data-standards/json) or [XML](/fundamentals/design-and-architecture/standards-based/data-standards/xml)) using a number of Transport Protocols (like [REST](/fundamentals/design-and-architecture/standards-based/data-standards/rest) API or [SOAP](/fundamentals/design-and-architecture/standards-based/data-standards/soap)).

FHIR Resources can be placed into a number of categories:

-   Foundation
-   Base
-   Clinical
-   Financial
-   Specialized

In total there are about 150 different FHIR Resources.

Every FHIR Resource has a "header" that is common to all FHIR Resources. This allows for parsing the message without knowing in advance exactly what resource it is; the "header" will tell exactly what resource it is, which tells the Software how to proceed processing it.

## Specifications

There are two FHIR Standards:

-   [HL7v2](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v2)
-   [HL7v3](/fundamentals/design-and-architecture/standards-based/data-standards/hl7v3)

FHIR also provides a Clinical Document Architecture (CDA). [Consolidated CDA (C-CDA)](/fundamentals/design-and-architecture/standards-based/data-standards/c-cda) is another standard.

## Sources

-   https://www.hl7.org/fhir/
-   https://www.hl7.org/implement/standards/index.cfm?ref=nav
