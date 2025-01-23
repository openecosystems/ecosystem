---
title: Disaster
pcx_content_type: definition
summary: >-
  A catastrophic event that results in long downtime (days or even weeks).<br><br>

  **Related terms:** [Failure](/fundamentals/glossary/#failure)
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/failure
  - /fundamentals/glossary/region
---

# Disaster

A catastrophic event that results in long downtime (days or even weeks). Examples:

- War
- Fire
- Floods
- Power grid disruption

Typically this results in loss of infrastructure and data.

When designing a system to survive a `Disaster` it is assumed that only **one** `Disaster` takes place at any point in time. I.e.: one can plan for a [Region](/fundamentals/glossary/region) to go down, but one does not also take measure to handle the scenario where the `backup region` also goes down.

If a [Failure](/fundamentals/glossary/failure) is a flat tire, then a `Disaster` is running the car into a brick wall where both the car and everything in it is permanently lost.

The measures to make systems resistant to a `Disaster` are vastly different from measures to make a system resistant to [Failure](/fundamentals/glossary/failure). For that reason it is advisable to distinguish between the two when determining Non-Functional Requirements.

**Related terms:** [Failure](/fundamentals/glossary/failure)
