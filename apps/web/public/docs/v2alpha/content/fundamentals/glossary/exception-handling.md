---
title: Exception Handling
aliases:
    - Exception Handling
pcx_content_type: definition
summary: >-
    Sometimes when running code things happen that were not expected, i.e. when an object is expected to exist but doesn't, or when a response is expected within a certain amount of time, but isn't. The art to prevent Applications from crashing in these cases, is called `Exception Handling`.
hidden: true
has_more: true
links_to:
    - /fundamentals/glossary/error-detection
---

# Exception Handling

Sometimes when running code things happen that were not expected, i.e. when an object is expected to exist but doesn't, or when a response is expected within a certain amount of time, but isn't. The art to prevent Applications from crashing in these cases, is called `Exception Handling`.

In some scenarios `Exception Handling` is especially important, for example in `C` and `C++` where one must release memory and resources to prevent leaks.

Some programming languages like `C#` and `Java` make this easier because they implement a `Garbage Collector` to automatically release memory for example. Though this generally works well for memory (but not always!), it is not always enough for resources. I.e., for a `Database Connection` it may not be sufficient to just release allocated memory, but you may also need to call the `Database.Close()` method, for example. There sometimes are also solutions for this. `.NET` has an interface called `IDisposable`, for example.

`Exception Handling` should not be confused with [SpecError Detection](/fundamentals/glossary/error-detection); SpecError Detection has a broader scope than a Programming Language, Application, or a `System`.
