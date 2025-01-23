# libs-partner-typescript-protobuf

This library was generated with [Nx](https://nx.dev).

## Building

Run `nx build libs-partner-typescript-protobuf` to build the library.

## Running unit tests

Run `nx test libs-partner-typescript-protobuf` to execute the unit tests via [Jest](https://jestjs.io).


"exports": {
      ".": {
        "import": "./dist/esm/index.js",
        "require": "./dist/cjs/index.js"
      },
      "./v1alpha": {
        "import": "./src/index.js",
        "require": "./src/index.js"
      },
      "./v1beta": {
        "import": "./src/index.js",
        "require": "./src/index.js"
      },
      "./v2alpha": {
        "import": "./src/index.js",
        "require": "./src/index.js"
      },
      "./v2beta": {
        "import": "./src/index.js",
        "require": "./src/index.js"
      }
    },
    "typesVersions": {
      "*": {
        "v1alpha": [
          "./dist/cjs/codegenv1/index.d.ts"
        ],
        "v1beta": [
          "./dist/cjs/reflect/index.d.ts"
        ],
        "v2alpha": [
          "./dist/cjs/wkt/index.d.ts"
        ],
        "v2beta": [
          "./dist/cjs/wire/index.d.ts"
        ]
      }
    }
