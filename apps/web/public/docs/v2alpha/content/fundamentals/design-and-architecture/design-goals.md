---
pcx_content_type: reference
title: Design Goals
weight: 3
---

# Enterprise Design Goals

There are a few critical design goals we must address when making tradeoff decisions between one integration approach versus another. Below are those goals:

## Secure. Speed-to-Market

Achieving fast speed-to-market in a secure way is our primary goal. There are many business pressures that drive this, not least of which is the competitive landscape. And design must primarily achieve this goal using proven technology.

## Interoperability

We want the compatibility of an EMR without the complexity of an EMR. One of the most challenging aspects of healthcare is the difficulty of exchanging data between the different parties: Personal Wellness Devices, Hospital/Urgent Cares, Specialist, Pharmacies, Laboratories and Primary Care Providers.

## Trusted audit trail

Our most prized data possession is our Profile data. Refer to the Data Enrichment Model for a deeper understanding of why this is the case. To ensure this data is protected and trustworthy as a source of truth, we must audit every action an employee, a system, or the user themselves takes on a profile record. We must capture details such as: (1) system emailed participant; (2) employee Joe Smith opened this profile record; (3) and user changed their preference.

## Data Provenance

Track dataflow from beginning to end

## Observability

To optimize the technical processes that undergird the business model and our ability to delivery at scale, we must be able to observe what works, what works but works inefficiently, what fails, and what fails unexplainably. Since each team works rather independently, the systems they build end up being independent as well. This poses additional challenges when one wants to observe the entire business process soup to nuts.

## Systems of Engagement vs Systems of Record

This concept is detailed in the Model but distinguishing between these two types of systems is important to achieving secure speed-to-market, while providing great experiences for our customers, users and employees. The design must account and optimize and secure these systems separately.

## Variable yet consistent data schemas

Because of our business approach and need to adapt quickly to market conditions, we must allow for variable schemas. However, this variability must not compromise consistency, normalcy and exactness.

## ACID Transactions

The design must achieve transaction acidity.

·      Atomicity – Changes are made atomically

·      Consistency – The state of the varying data stores is always consistent

·      Isolation – Even though transactions are executed concurrently it appears they are executed serially

·      Durability – Once a transaction has committed it cannot be undone without committing a new transaction, both of which must be audited

## Dynamic prioritiztion

Some request must be treated with a higher priority. We need the ability to dynamically prioritize such request

## Compliance and Security

System to system and system to user interactions must be secure, trusted, and accountable

## Extensible Architecture

Extendable

## Team Independence

Different teams have different timelines, delivery cadences and business pressures. We must reduce the amount of inter-team dependence as much as possible. Allow each team to work without reasoning about other team’s needs or data structures.

## Separation of Concerns

Seperate the network details from the business details
