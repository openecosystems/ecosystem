---
title: Exactly Once
pcx_content_type: definition
summary: >-
  This [MDP](/fundamentals/glossary/#mdp) pattern provides [Guaranteed Delivery](/fundamentals/glossary/guaranteed-delivery), but it also prevents sending (or receiving) duplicates. It is very hard to implement. Perhaps it's better to find a way to make the system [Idempotent](/fundamentals/glossary/#idempotenc) and use [At Least Once](/fundamentals/glossary/#at-least-once) instead.
hidden: true
has_more: false
links_to:
  - /fundamentals/glossary/at-least-once
  - /fundamentals/glossary/guaranteed-delivery
  - /fundamentals/glossary/idempotenc
  - /fundamentals/glossary/idempotence
  - /fundamentals/glossary/mdp
---

# Exactly Once

This [MDP](/fundamentals/glossary/mdp) pattern provides [Guaranteed Delivery](/fundamentals/glossary/guaranteed-delivery), but it also prevents sending (or receiving) duplicates. It is very hard to implement. Perhaps it's better to find a way to make the system [Idempotent](/fundamentals/glossary/idempotence) and use [At Least Once](/fundamentals/glossary/at-least-once) instead.
