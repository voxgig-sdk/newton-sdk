# Newton Golang SDK

The Golang SDK for the Newton API. Provides an entity-oriented interface using standard Go conventions — no generics required, data flows as `map[string]any`.


## Install
```bash
go get github.com/voxgig-sdk/newton-sdk/go
```

If the module is not yet published to a registry, use a `replace` directive
in your `go.mod` to point to a local checkout:

```bash
go mod edit -replace github.com/voxgig-sdk/newton-sdk/go=../path/to/github.com/voxgig-sdk/newton-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```go
package main

import (
    "fmt"

    sdk "github.com/voxgig-sdk/newton-sdk/go"
    "github.com/voxgig-sdk/newton-sdk/go/core"
)

func main() {
    client := sdk.NewNewtonSDK(map[string]any{})
```

### 3. Load a abs

```go
    result, err = client.Abs(nil).Load(
        map[string]any{"id": "example_id"}, nil,
    )
    if err != nil {
        panic(err)
    }

    rm = core.ToMapAny(result)
    if rm["ok"] == true {
        fmt.Println(rm["data"])
    }
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.TestSDK(nil, nil)

result, err := client.Planet(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
// result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewNewtonSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NEWTON_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewNewtonSDK

```go
func NewNewtonSDK(options map[string]any) *NewtonSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *NewtonSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NewtonSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Abs` | `(data map[string]any) NewtonEntity` | Create a Abs entity instance. |
| `Arcco` | `(data map[string]any) NewtonEntity` | Create a Arcco entity instance. |
| `Arcsin` | `(data map[string]any) NewtonEntity` | Create a Arcsin entity instance. |
| `Arctan` | `(data map[string]any) NewtonEntity` | Create a Arctan entity instance. |
| `Area` | `(data map[string]any) NewtonEntity` | Create a Area entity instance. |
| `Cos` | `(data map[string]any) NewtonEntity` | Create a Cos entity instance. |
| `Derive` | `(data map[string]any) NewtonEntity` | Create a Derive entity instance. |
| `Factor` | `(data map[string]any) NewtonEntity` | Create a Factor entity instance. |
| `Integrate` | `(data map[string]any) NewtonEntity` | Create a Integrate entity instance. |
| `Log` | `(data map[string]any) NewtonEntity` | Create a Log entity instance. |
| `Simplify` | `(data map[string]any) NewtonEntity` | Create a Simplify entity instance. |
| `Sin` | `(data map[string]any) NewtonEntity` | Create a Sin entity instance. |
| `Tan` | `(data map[string]any) NewtonEntity` | Create a Tan entity instance. |
| `Tangent` | `(data map[string]any) NewtonEntity` | Create a Tangent entity instance. |
| `Zero` | `(data map[string]any) NewtonEntity` | Create a Zero entity instance. |

### Entity interface (NewtonEntity)

All entities implement the `NewtonEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(any, error)`. The `any` value is a
`map[string]any` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `"ok"` | `bool` | `true` if the HTTP status is 2xx. |
| `"status"` | `int` | HTTP status code. |
| `"headers"` | `map[string]any` | Response headers. |
| `"data"` | `any` | Parsed JSON response body. |

On error, `"ok"` is `false` and `"err"` contains the error value.

### Entities

#### Abs

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/abs/{expression}`

#### Arcco

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/arccos/{expression}`

#### Arcsin

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/arcsin/{expression}`

#### Arctan

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/arctan/{expression}`

#### Area

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/area/{expression}`

#### Cos

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/cos/{expression}`

#### Derive

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/derive/{expression}`

#### Factor

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/factor/{expression}`

#### Integrate

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/integrate/{expression}`

#### Log

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/log/{expression}`

#### Simplify

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/simplify/{expression}`

#### Sin

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/sin/{expression}`

#### Tan

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/tan/{expression}`

#### Tangent

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/tangent/{expression}`

#### Zero

| Field | Description |
| --- | --- |
| `"expression"` |  |
| `"operation"` |  |
| `"result"` |  |

Operations: Load.

API path: `/zeroes/{expression}`



## Entities


### Abs

Create an instance: `abs := client.Abs(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Abs(nil).Load(map[string]any{"id": "abs_id"}, nil)
```


### Arcco

Create an instance: `arcco := client.Arcco(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Arcco(nil).Load(map[string]any{"id": "arcco_id"}, nil)
```


### Arcsin

Create an instance: `arcsin := client.Arcsin(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Arcsin(nil).Load(map[string]any{"id": "arcsin_id"}, nil)
```


### Arctan

Create an instance: `arctan := client.Arctan(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Arctan(nil).Load(map[string]any{"id": "arctan_id"}, nil)
```


### Area

Create an instance: `area := client.Area(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Area(nil).Load(map[string]any{"id": "area_id"}, nil)
```


### Cos

Create an instance: `cos := client.Cos(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Cos(nil).Load(map[string]any{"id": "cos_id"}, nil)
```


### Derive

Create an instance: `derive := client.Derive(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Derive(nil).Load(map[string]any{"id": "derive_id"}, nil)
```


### Factor

Create an instance: `factor := client.Factor(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Factor(nil).Load(map[string]any{"id": "factor_id"}, nil)
```


### Integrate

Create an instance: `integrate := client.Integrate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Integrate(nil).Load(map[string]any{"id": "integrate_id"}, nil)
```


### Log

Create an instance: `log := client.Log(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Log(nil).Load(map[string]any{"id": "log_id"}, nil)
```


### Simplify

Create an instance: `simplify := client.Simplify(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Simplify(nil).Load(map[string]any{"id": "simplify_id"}, nil)
```


### Sin

Create an instance: `sin := client.Sin(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Sin(nil).Load(map[string]any{"id": "sin_id"}, nil)
```


### Tan

Create an instance: `tan := client.Tan(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Tan(nil).Load(map[string]any{"id": "tan_id"}, nil)
```


### Tangent

Create an instance: `tangent := client.Tangent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Tangent(nil).Load(map[string]any{"id": "tangent_id"}, nil)
```


### Zero

Create an instance: `zero := client.Zero(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```go
result, err := client.Zero(nil).Load(map[string]any{"id": "zero_id"}, nil)
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/newton-sdk/go/
├── newton.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/newton-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
moon := client.Moon(nil)
moon.Load(map[string]any{"planet_id": "earth", "id": "luna"}, nil)

// moon.Data() now returns the loaded moon data
// moon.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
