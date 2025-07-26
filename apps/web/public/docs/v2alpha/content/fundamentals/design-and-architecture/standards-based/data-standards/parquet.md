---
date_created: 2022-12-11T17:07:53
title: Parquet
pcx_content_type: definition
summary: >-
    `Parquet` is an open source binary column-oriented data format that is very efficient.
hidden: true
has_more: true
links_to:
    - /fundamentals/design-and-architecture/standards-based/data-standards/csv
    - /fundamentals/design-and-architecture/standards-based/data-standards/gzip
    - /fundamentals/design-and-architecture/standards-based/data-standards/lzo
    - /fundamentals/design-and-architecture/standards-based/data-standards/snappy
    - /fundamentals/glossary/compression
    - /fundamentals/glossary/one-to-many
---

# Parquet

`Parquet` is an open source binary column-oriented data format that is very efficient.

## Advantages

-   Store Big Data.
-   Can be parsed by [AWS Athena](/fundamentals/glossary/aws/#athena) when stored in [AWS S3](/fundamentals/glossary/aws/#s3).
-   Compact
-   Data is fully typed

## Disadvantages

-   No schema evolution

## Under the Hood

A `Parquet` file is not as simple as a [CSV](/fundamentals/design-and-architecture/standards-based/data-standards/csv) file, for example. Where a [CSV](/fundamentals/design-and-architecture/standards-based/data-standards/csv) file has one header and many records, a `Parquet` file has:

-   File Metadata
    -   Row Groups
        -   Column Groups
            -   Pages

... where Plural indicates a [One-to-Many](/fundamentals/glossary/one-to-many) relationship.

## Encoding

`Parquet` supports Encoding.

## Compression

`Parquet`s support [Compression](/fundamentals/glossary/compression).

Below is a table where a set of data is stored in different file formats:

| File Format                                                                     | Compression                                                                           | Query time (s) | Size (GB) |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | -------------: | --------: |
| [CSV](/fundamentals/design-and-architecture/standards-based/data-standards/csv) | None                                                                                  |         2892.3 |    437.46 |
| `Parquet`                                                                       | [Snappy](/fundamentals/design-and-architecture/standards-based/data-standards/snappy) |           28.9 |     54.83 |
| `Parquet`                                                                       | [GZIP](/fundamentals/design-and-architecture/standards-based/data-standards/gzip)     |           40.3 |     36.78 |
| `Parquet`                                                                       | None                                                                                  |           43.4 |    138.54 |
| `Parquet`                                                                       | [LZO](/fundamentals/design-and-architecture/standards-based/data-standards/lzo)       |           50.6 |      55.6 |

As you can see:

-   `Parquet` outperforms CSV in every way.
-   Enabling Compression improves both the File Size as well as Performance.

| Category                | [LZO](/fundamentals/design-and-architecture/standards-based/data-standards/lzo) | [GZIP](/fundamentals/design-and-architecture/standards-based/data-standards/gzip) | [Snappy](/fundamentals/design-and-architecture/standards-based/data-standards/snappy) |
| ----------------------- | ------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| Compression Size        | Medium                                                                          | Smallest                                                                          | Medium                                                                                |
| Compression Speed       | Fast                                                                            | Slow                                                                              | Fastest                                                                               |
| Decompression Speed     | Fastest                                                                         | Slow                                                                              | Fast                                                                                  |
| Frequency of Data Usage | Hot                                                                             | Cold                                                                              | Hot                                                                                   |
| Splitable               | Yes                                                                             | No                                                                                | Yes                                                                                   |

## Schemas and Data Types

`Parquet` supports Schemas and Data Types.

There is one caveat, though: Schemas are not easily changed, if it is at all possible. It may be possible to make non-breakable changes (like adding a column)[^1], or it may not. Removing columns, renaming them, or changing their data type is a breaking change by definition and requires regenerating the `Parquet` files (which can be immensely expensive if one has lots of data).

## References

-   [What is Apache Parquet file](https://www.youtube.com/watch?v=PaDUxrI6ThA)

[^1]: The video [Avro vs Parquet](https://www.youtube.com/watch?v=UrWthx8T3UY](https://youtu.be/UrWthx8T3UY?t=292) 4:52 suggests that `Parquets` only supports schema _appends_.
