---
title: At Most Once
pcx_content_type: definition
summary: >-
    This is the simplest [MDP](/fundamentals/glossary/#mdp) pattern to implement. With this pattern a message is sent to another component, without there being any mechanisms in place to guarantee that the message actually arrives at its destination.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/mdp
    - /fundamentals/glossary/metric
---

# At Most Once

This is the simplest [MDP](/fundamentals/glossary/mdp) pattern to implement. With this pattern a message is sent to another component, without there being any mechanisms in place to guarantee that the message actually arrives at its destination.

One could use this pattern if a single message getting lost doesn't really have an impact, which is usually the case for volatile data (data that only means something now, but quickly becomes irrelevant). One example of this would be a message indicating some [Metric](/fundamentals/glossary/metric) like the amount of disk space used: if we miss this one, we'll use the next one.
