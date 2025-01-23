---
title: Metric Monitoring
pcx_content_type: definition
summary: >-
  ``Metric Monitoring`` uses [Metrics](/fundamentals/glossary/#metric) to determine if the system is still running within normal parameters.
hidden: true
has_more: true
links_to:
  - /fundamentals/glossary/alerting
  - /fundamentals/glossary/elk-stack
  - /fundamentals/glossary/event-monitoring
  - /fundamentals/glossary/exception-handling
  - /fundamentals/glossary/machine-learning
  - /fundamentals/glossary/metric
  - /fundamentals/glossary/monitoring-maturity-level
---

# Metric Monitoring

`Metric Monitoring` uses [Metrics](/fundamentals/glossary/metric) to determine if the system is still running within normal parameters.

## The difference between Events and Metrics

Where an `Event` (see [Event Monitoring](/fundamentals/glossary/event-monitoring)) tells what happened, [Metrics](/fundamentals/glossary/metric) gives an aggregate view of _how much_ or _how often_ something happened.  
Where an `Event` says something about a specific point in time, a [Metric](/fundamentals/glossary/metric) says something about what happened _around_ that time.  
Where an `Event` can tell you that something _did_ go wrong, `Metric Monitoring` can tell you when something is _about_ to go wrong.

## Events to Metrics conversion

[Event Monitoring](/fundamentals/glossary/event-monitoring) provides insight on what the Applications are doing and allow one to respond to [Exceptions](/fundamentals/glossary/exception-handling). But there is much more information that could be derived from [Event Monitoring](/fundamentals/glossary/event-monitoring), like how often a User uses a specific function.

`Events` can be converted to [Metrics](/fundamentals/glossary/metric) to provide more insight in how the Application is behaving and how it is used.

## Rationale

`Metric Monitoring` is crucial for getting from `Reactive Monitoring` to `Proactive Monitoring`. (See [Monitoring Maturity Levels](/fundamentals/glossary/monitoring-maturity-level)).

## Examples

Examples:

- **CPU usage / Memory Usage**. If CPU usage is generally around 40%, but now it continuously clips at 100%, something is wrong, even though the system may not log an `Error Event`.
- **Disk usage**. If the amount of storage used on a disk gradually increases, there will be a moment when the disk is full and the system will crash.
- **Queue Size**. If Queues are used for messaging, then a Queue that starts to retain messages is an indication that the service instances that are reading from that Queue are either no longer working, or simply not fast enough.
- **Nr of Login Attempts**. If there is typically one login per second, but all of a sudden it increases to 20 logins per second, someone is probably trying to hack the system.

## Alerting

`Metric Monitoring` can also be used to generate `Alerts` (see [Alerting](/fundamentals/glossary/alerting)) when [Metrics](/fundamentals/glossary/metric) exceed certain thresholds. I.e.: If the disk usage exceeds 80%, send a notification to an Operator to have a look at it, to prevent a system crash.

Now, deciding on alerting `thresholds` is sometimes easy, like, for instance, a disk filling up. But for other [Metrics](/fundamentals/glossary/metric) deciding on good values may even be impossible. I.e.: if during daytime generally 50 people log in per minute and at night 2, then no threshold exists for detecting that someone is hacking the system. Say a hacker does 20 login attempts per minute, then we could set a threshold at 65 login, but that wouldn't trigger anything at night, when the number of logins would rise from 2 to 22.

For these scenarios [Machine Learning](/fundamentals/glossary/machine-learning) could provide a solution. One could teach a system what normal platform behavior is, and have it send an alert when current values exceed expected values for this time of the day, for example. One of the solutions that can do that is an [ELK Stack](/fundamentals/glossary/elk-stack).
