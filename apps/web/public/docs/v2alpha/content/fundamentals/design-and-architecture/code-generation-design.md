---
pcx_content_type: reference
title: Code Generation Design
weight: 13
---

# Code Generation Design

-   Safety: The protobuf code-generation system avoids the errors that are common with hand-built serialization code.
-   Correctness: SwiftProtobuf passes both its own extensive test suite and Google's full conformance test for protobuf correctness.
-   Schema-driven: Defining your data structures in a separate .proto schema file clearly documents your communications conventions.
-   Idiomatic: SwiftProtobuf takes full advantage of the Swift language. In particular, all generated types provide full Swift copy-on-write value semantics.
-   Efficient binary serialization: The .serializedBytes() method returns a bag of bytes with a compact binary form of your data. You can deserialize the data using the init(serializedBytes:) initializer.
-   Standard JSON serialization: The .jsonUTF8Data() method returns a JSON form of your data that can be parsed with the init(jsonUTF8Bytes:) initializer.
-   Hashable, Equatable: The generated struct can be put into a Set<> or Dictionary<>.
-   Performant: The binary and JSON serializers have been extensively optimized.
-   Extensible: You can add your own Swift extensions to any of the generated types.
