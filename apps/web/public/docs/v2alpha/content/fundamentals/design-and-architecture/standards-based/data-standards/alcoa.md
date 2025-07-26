---
date_created: 2022-12-11T17:20:09
title: ALCOA
pcx_content_type: definition
summary: >-
    ALCOA is a set of principles to ensure data quality. It is an implementation of [GDocP](/fundamentals/design-and-architecture/standards-based/data-standards/#gdocp).
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/gdocp
---

# Attributable, Legible, Contemporaneous, Original, Accurate (ALCOA)

`ALCOA` is a set of principles to ensure data quality. It is an implementation of [GDocP](/fundamentals/design-and-architecture/standards-based/data-standards/gdocp).

`ALCOA` is an acronym of:

-   **Attributable**: Who, When, What and Why.
-   **Legible**: Readable and Permanent.
-   **Contemporaneous**: Timestamped.
-   **Original**: Original, Immutable.
-   **Accurate**: Data must be unambiguous.

## Attributable

The raw data must be traceable to an _authorized_ person recording the data.

When a change is made, the following data must be added to a document:

-   **When** it was changed (a date or timestamp)
-   **Who** changed it (i.e. Initials)
-   **Why** it was changed (description)

For digital records an audit trail should describe:

-   When the record was created, edited, signed or viewed
-   Who performed the action
-   Why the action was performed

## Legible

Everything that is written must be legible and unambiguous.

"0" and "O" could be ambiguous, as could various ways for writing down a date.

Signatures must be identifiable.

All data must be in an identifiable format.

## Contemporaneous

Data must be recorded at the time the event occurs. Using another date than the current one is prohibited.

For electronic systems timestamps must be automatically set by systems, and not be posted.

## Original

Original records MUST NEVER be destroyed.

Data can only be modified if there is a record of:

-   what it used to be,
-   when it was changed,
-   who changed it, and
-   why it was changed.

In any case it MUST be possible to retrieve the original record.

## Accurate

A record must completely reflect the true observation. Either the document is complete, consistent and correct, or an explanation has to be attached why it is not.

Systems that process records must have mechanisms is place that prevent recording of inaccurate data.

Data must be correct, truthful, complete, valid, and reliable.

## Sources

-   https://blink.ucsd.edu/research/_files/ALCOA-Standards-210304.pdf
