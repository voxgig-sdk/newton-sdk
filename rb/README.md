# Newton Ruby SDK

The Ruby SDK for the Newton API. Provides an entity-oriented interface using idiomatic Ruby conventions.


## Install
```bash
gem install newton-sdk
```

Or add to your `Gemfile`:

```ruby
gem "newton-sdk"
```

Then run:

```bash
bundle install
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Newton_sdk"

client = NewtonSDK.new({})
```

### 3. Load a abs

```ruby
result, err = client.Abs(nil).load({ "id" => "example_id" }, nil)
raise err if err
puts result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
raise err if err

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
end
```

### Prepare a request without sending it

```ruby
fetchdef, err = client.prepare({
  "path" => "/api/resource/{id}",
  "method" => "DELETE",
  "params" => { "id" => "example" },
})
raise err if err

puts fetchdef["url"]
puts fetchdef["method"]
puts fetchdef["headers"]
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = NewtonSDK.test(nil, nil)

result, err = client.Newton(nil).load(
  { "id" => "test01" }, nil
)
# result contains mock response data
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
| `prepare` | `(fetchargs) -> [Hash, err]` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> [Hash, err]` | Build and send an HTTP request. |
| `Abs` | `(data) -> AbsEntity` | Create a Abs entity instance. |
| `Arcco` | `(data) -> ArccoEntity` | Create a Arcco entity instance. |
| `Arcsin` | `(data) -> ArcsinEntity` | Create a Arcsin entity instance. |
| `Arctan` | `(data) -> ArctanEntity` | Create a Arctan entity instance. |
| `Area` | `(data) -> AreaEntity` | Create a Area entity instance. |
| `Cos` | `(data) -> CosEntity` | Create a Cos entity instance. |
| `Derive` | `(data) -> DeriveEntity` | Create a Derive entity instance. |
| `Factor` | `(data) -> FactorEntity` | Create a Factor entity instance. |
| `Integrate` | `(data) -> IntegrateEntity` | Create a Integrate entity instance. |
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
| `load` | `(reqmatch, ctrl) -> [any, err]` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> [any, err]` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> [any, err]` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> [any, err]` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> [any, err]` | Remove an entity. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return `[any, err]`. The first value is a
`Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `false` and `err` contains the error value.

### Entities

#### Abs

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/abs/{expression}`

#### Arcco

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/arccos/{expression}`

#### Arcsin

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/arcsin/{expression}`

#### Arctan

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/arctan/{expression}`

#### Area

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/area/{expression}`

#### Cos

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/cos/{expression}`

#### Derive

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/derive/{expression}`

#### Factor

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/factor/{expression}`

#### Integrate

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/integrate/{expression}`

#### Log

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/log/{expression}`

#### Simplify

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/simplify/{expression}`

#### Sin

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/sin/{expression}`

#### Tan

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/tan/{expression}`

#### Tangent

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/tangent/{expression}`

#### Zero

| Field | Description |
| --- | --- |
| `expression` |  |
| `operation` |  |
| `result` |  |

Operations: Load.

API path: `/zeroes/{expression}`



## Entities


### Abs

Create an instance: `const abs = client.Abs()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const abs = await client.Abs().load({ id: 'abs_id' })
```


### Arcco

Create an instance: `const arcco = client.Arcco()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const arcco = await client.Arcco().load({ id: 'arcco_id' })
```


### Arcsin

Create an instance: `const arcsin = client.Arcsin()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const arcsin = await client.Arcsin().load({ id: 'arcsin_id' })
```


### Arctan

Create an instance: `const arctan = client.Arctan()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const arctan = await client.Arctan().load({ id: 'arctan_id' })
```


### Area

Create an instance: `const area = client.Area()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const area = await client.Area().load({ id: 'area_id' })
```


### Cos

Create an instance: `const cos = client.Cos()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const cos = await client.Cos().load({ id: 'cos_id' })
```


### Derive

Create an instance: `const derive = client.Derive()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const derive = await client.Derive().load({ id: 'derive_id' })
```


### Factor

Create an instance: `const factor = client.Factor()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const factor = await client.Factor().load({ id: 'factor_id' })
```


### Integrate

Create an instance: `const integrate = client.Integrate()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const integrate = await client.Integrate().load({ id: 'integrate_id' })
```


### Log

Create an instance: `const log = client.Log()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const log = await client.Log().load({ id: 'log_id' })
```


### Simplify

Create an instance: `const simplify = client.Simplify()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const simplify = await client.Simplify().load({ id: 'simplify_id' })
```


### Sin

Create an instance: `const sin = client.Sin()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const sin = await client.Sin().load({ id: 'sin_id' })
```


### Tan

Create an instance: `const tan = client.Tan()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const tan = await client.Tan().load({ id: 'tan_id' })
```


### Tangent

Create an instance: `const tangent = client.Tangent()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const tangent = await client.Tangent().load({ id: 'tangent_id' })
```


### Zero

Create an instance: `const zero = client.Zero()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | ``$STRING`` |  |
| `operation` | ``$STRING`` |  |
| `result` | ``$STRING`` |  |

#### Example: Load

```ts
const zero = await client.Zero().load({ id: 'zero_id' })
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
error is returned to the caller as a second return value.

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
moon = client.Moon
moon.load({ "planet_id" => "earth", "id" => "luna" })

# moon.data_get now returns the loaded moon data
# moon.match_get returns the last match criteria
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
