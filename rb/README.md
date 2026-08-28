# Newton Ruby SDK



The Ruby SDK for the Newton API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Abs` — with named operations (`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/newton-sdk/releases](https://github.com/voxgig-sdk/newton-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Newton_sdk"

client = NewtonSDK.new
```

### 3. Load an abs

```ruby
begin
  # load returns the ENTITY — call data_get for the Abs record (raises on error).
  abs = client.Abs.load({ "id" => "example_id" })
  puts abs
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  arcco = client.Arcco.load({ "id" => "example_id" })
rescue => err
  warn "load failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = NewtonSDK.test({
  "entity" => { "arcco" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
arcco = client.Arcco.load({ "id" => "test01" })
puts arcco
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = NewtonSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### NewtonSDK

```ruby
require_relative "Newton_sdk"
client = NewtonSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = NewtonSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NewtonSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Abs` | `(data) -> AbsEntity` | Create an Abs entity instance. |
| `Arcco` | `(data) -> ArccoEntity` | Create an Arcco entity instance. |
| `Arcsin` | `(data) -> ArcsinEntity` | Create an Arcsin entity instance. |
| `Arctan` | `(data) -> ArctanEntity` | Create an Arctan entity instance. |
| `Area` | `(data) -> AreaEntity` | Create an Area entity instance. |
| `Cos` | `(data) -> CosEntity` | Create a Cos entity instance. |
| `Derive` | `(data) -> DeriveEntity` | Create a Derive entity instance. |
| `Factor` | `(data) -> FactorEntity` | Create a Factor entity instance. |
| `Integrate` | `(data) -> IntegrateEntity` | Create an Integrate entity instance. |
| `Log` | `(data) -> LogEntity` | Create a Log entity instance. |
| `Simplify` | `(data) -> SimplifyEntity` | Create a Simplify entity instance. |
| `Sin` | `(data) -> SinEntity` | Create a Sin entity instance. |
| `Tan` | `(data) -> TanEntity` | Create a Tan entity instance. |
| `Tangent` | `(data) -> TangentEntity` | Create a Tangent entity instance. |
| `Zero` | `(data) -> ZeroEntity` | Create a Zero entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `NewtonError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Abs

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/abs/{expression}`

#### Arcco

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/arccos/{expression}`

#### Arcsin

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/arcsin/{expression}`

#### Arctan

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/arctan/{expression}`

#### Area

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/area/{expression}`

#### Cos

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/cos/{expression}`

#### Derive

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/derive/{expression}`

#### Factor

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/factor/{expression}`

#### Integrate

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/integrate/{expression}`

#### Log

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/log/{expression}`

#### Simplify

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/simplify/{expression}`

#### Sin

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/sin/{expression}`

#### Tan

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/tan/{expression}`

#### Tangent

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/tangent/{expression}`

#### Zero

| Field | Description |
| --- | --- |
| `expression` | The mathematical expression that was processed |
| `id` |  |
| `operation` | The mathematical operation that was performed |
| `result` | The result of the mathematical operation |

Operations: Load.

API path: `/zeroes/{expression}`



## Entities


### Abs

Create an instance: `abs = client.Abs`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Abs record (raises on error).
abs = client.Abs.load({ "id" => "abs_id" })
```


### Arcco

Create an instance: `arcco = client.Arcco`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Arcco record (raises on error).
arcco = client.Arcco.load({ "id" => "arcco_id" })
```


### Arcsin

Create an instance: `arcsin = client.Arcsin`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Arcsin record (raises on error).
arcsin = client.Arcsin.load({ "id" => "arcsin_id" })
```


### Arctan

Create an instance: `arctan = client.Arctan`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Arctan record (raises on error).
arctan = client.Arctan.load({ "id" => "arctan_id" })
```


### Area

Create an instance: `area = client.Area`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Area record (raises on error).
area = client.Area.load({ "id" => "area_id" })
```


### Cos

Create an instance: `cos = client.Cos`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Cos record (raises on error).
cos = client.Cos.load({ "id" => "cos_id" })
```


### Derive

Create an instance: `derive = client.Derive`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Derive record (raises on error).
derive = client.Derive.load({ "id" => "derive_id" })
```


### Factor

Create an instance: `factor = client.Factor`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Factor record (raises on error).
factor = client.Factor.load({ "id" => "factor_id" })
```


### Integrate

Create an instance: `integrate = client.Integrate`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Integrate record (raises on error).
integrate = client.Integrate.load({ "id" => "integrate_id" })
```


### Log

Create an instance: `log = client.Log`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Log record (raises on error).
log = client.Log.load({ "id" => "log_id" })
```


### Simplify

Create an instance: `simplify = client.Simplify`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Simplify record (raises on error).
simplify = client.Simplify.load({ "id" => "simplify_id" })
```


### Sin

Create an instance: `sin = client.Sin`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Sin record (raises on error).
sin = client.Sin.load({ "id" => "sin_id" })
```


### Tan

Create an instance: `tan = client.Tan`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Tan record (raises on error).
tan = client.Tan.load({ "id" => "tan_id" })
```


### Tangent

Create an instance: `tangent = client.Tangent`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Tangent record (raises on error).
tangent = client.Tangent.load({ "id" => "tangent_id" })
```


### Zero

Create an instance: `zero = client.Zero`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `String` | The mathematical expression that was processed |
| `id` | `String` |  |
| `operation` | `String` | The mathematical operation that was performed |
| `result` | `String` | The result of the mathematical operation |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Zero record (raises on error).
zero = client.Zero.load({ "id" => "zero_id" })
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Newton_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Newton_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
arcco = client.Arcco
arcco.load({ "id" => "example_id" })

# arcco.data_get now returns the arcco data from the last load
# arcco.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
