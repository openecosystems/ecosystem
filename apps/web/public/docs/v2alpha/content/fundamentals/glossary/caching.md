---
title: Caching
aliases:
    - Cache
    - Caching
pcx_content_type: definition
summary: >-
    The art of remembering answers to questions for a period of time, so that if the same question is asked within that time frame, the answer can be provided without hitting the Back End. This concept is quite often used to increase performance, for example with Web Services or Database access.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/back-end
    - /fundamentals/glossary/performance
    - /fundamentals/glossary/database
    - /fundamentals/glossary/latency
    - /fundamentals/glossary/dns
    - /fundamentals/glossary/cdn
    - /fundamentals/glossary/risk
    - /fundamentals/glossary/ttl
---

# Caching

The art of remembering answers to questions for a period of time, so that if the same question is asked within that time frame, the answer can be provided without hitting the [Back End](/fundamentals/glossary/back-end).

This concept is quite often used to increase [performance](/fundamentals/glossary/performance), for example with Web Services or [Database](/fundamentals/glossary/database) access.

## Cache mechanism

Instead of resolving a request using the [Back End](/fundamentals/glossary/back-end) the Application will first attempt to retrieve the result from the Cache. If the result is in the Cache, which is called a Cache Hit, the result can be returned immediately. If the result is not in the Cache, which is called a Cache Miss, the Application will have to use the [Back End](/fundamentals/glossary/back-end) to determine the result and then have the discipline to put the result in the Cache so that the next time the result _can_ be returned from the Cache.

A Cache is a Key-Value Store. Results can be retrieved from the Cache by providing the `key`; Results can be stored in the Cache by providing a `key` and the Result.

## Performance

Because Caching implementations are typically In-Memory, they can look up values with sub-millisecond [Latency](/fundamentals/glossary/latency), something that can never be achieved with most [Back End](/fundamentals/glossary/back-end) systems like [Databases](/fundamentals/glossary/database). The ability to reduce [Latency](/fundamentals/glossary/latency) to such small values

## Common applications

If Caching is used, a Cache Miss is relatively expensive: instead of one call to the [Back End](/fundamentals/glossary/back-end) you will have two calls to the Cache (one Cache Miss and one to store the result in the Cache) _and_ a call to the [Back End](/fundamentals/glossary/back-end). That means Caching only makes sense if the number of Cache Hits is bigger than the number of Cache Misses. Or, in other words: don't use Caching if every request is expected to be unique.

Caching typically works well with the following use cases:

-   Database Lookups
-   [DNS](/fundamentals/glossary/dns)
-   Static content ([CDN](/fundamentals/glossary/cdn))

## Risk obsolete data

By definition a Cache returns a snapshot of a result that was previously generated. That means that if the source data has changed, that by retrieving the result from the Cache you may get **obsolete data**. This is a generic [Risk](/fundamentals/glossary/risk) that must be taken into account when working with Caches.

One could assume that if source data is very volatile, Caching should not be used. But this assumption may not be accurate. At a high data rate even a Cache with a [TTL](/fundamentals/glossary/ttl) of only 1 second can result in a tremendous [Performance](/fundamentals/glossary/performance) boost.

## Cache expiration

In order to mitigate the Risk of returning obsolete data, many Caching implementations allow for setting **Cache Expiration**, where data is automatically removed after a certain amount of time (see [TTL](/fundamentals/glossary/ttl)). What the [TTL](/fundamentals/glossary/ttl) should be greatly depends on the use case. It could range anywhere from a second up to years.

Also data should be explicitly be removed from the Cache, if one can reasonably assume that source data has changed in a way that it would result in a different Result.
