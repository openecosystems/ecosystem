---
title: Region
pcx_content_type: definition
summary: >-
  [CSPs](/fundamentals/glossary/#csp) have [Data Centers](/fundamentals/glossary/#data-center) at various locations across the globe. Those various locations are called `Regions`.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/availability-zone
  - /fundamentals/glossary/csp
  - /fundamentals/glossary/data-center
  - /fundamentals/glossary/geo-redundancy
  - /fundamentals/glossary/latency
---

# Region

[CSPs](/fundamentals/glossary/csp) have [Data Centers](/fundamentals/glossary/data-center) at various locations across the globe. Those various locations are called `Regions`. This has a number of purposes, among which:

- To provide [Geo-Redundancy](/fundamentals/glossary/geo-redundancy).
- To reduce [Latency](/fundamentals/glossary/latency): systems respond faster if clients in Australia can use services in Australia instead of services in the US.
- To distribute load.

[CSPs](/fundamentals/glossary/csp) also tend to have more than one datacenter relatively close to one another within a `Region`. These datacenters within a `Region` are called [Availability Zones](/fundamentals/glossary/availability-zone). This serves a number of purposes, among which:

- Redundancy and low [Latency](/fundamentals/glossary/latency): If Redundancy would have to be achieved with `Regions`, then that would result in higher `Latencies`, because `Regions` are pretty far apart (i.e. west us, central us, east us). By having multiple datacenters within a `Region` Redundancy can be achieved with little latency, because the datacenters within a `Region` are relatively close to one another (10 miles, give or take).
