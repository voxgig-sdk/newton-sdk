# Newton Ruby SDK Reference

Complete API reference for the Newton Ruby SDK.


## NewtonSDK

### Constructor

```ruby
require_relative 'Newton_sdk'

client = NewtonSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NewtonSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = NewtonSDK.test
```


### Instance Methods

#### `Abs(data = nil)`

Create a new `Abs` entity instance. Pass `nil` for no initial data.

#### `Arcco(data = nil)`

Create a new `Arcco` entity instance. Pass `nil` for no initial data.

#### `Arcsin(data = nil)`

Create a new `Arcsin` entity instance. Pass `nil` for no initial data.

#### `Arctan(data = nil)`

Create a new `Arctan` entity instance. Pass `nil` for no initial data.

#### `Area(data = nil)`

Create a new `Area` entity instance. Pass `nil` for no initial data.

#### `Cos(data = nil)`

Create a new `Cos` entity instance. Pass `nil` for no initial data.

#### `Derive(data = nil)`

Create a new `Derive` entity instance. Pass `nil` for no initial data.

#### `Factor(data = nil)`

Create a new `Factor` entity instance. Pass `nil` for no initial data.

#### `Integrate(data = nil)`

Create a new `Integrate` entity instance. Pass `nil` for no initial data.

#### `Log(data = nil)`

Create a new `Log` entity instance. Pass `nil` for no initial data.

#### `Simplify(data = nil)`

Create a new `Simplify` entity instance. Pass `nil` for no initial data.

#### `Sin(data = nil)`

Create a new `Sin` entity instance. Pass `nil` for no initial data.

#### `Tan(data = nil)`

Create a new `Tan` entity instance. Pass `nil` for no initial data.

#### `Tangent(data = nil)`

Create a new `Tangent` entity instance. Pass `nil` for no initial data.

#### `Zero(data = nil)`

Create a new `Zero` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AbsEntity

```ruby
abs = client.Abs
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Abs.load({ "id" => "abs_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AbsEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArccoEntity

```ruby
arcco = client.Arcco
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Arcco.load({ "id" => "arcco_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArccoEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArcsinEntity

```ruby
arcsin = client.Arcsin
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Arcsin.load({ "id" => "arcsin_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArcsinEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArctanEntity

```ruby
arctan = client.Arctan
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Arctan.load({ "id" => "arctan_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArctanEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AreaEntity

```ruby
area = client.Area
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Area.load({ "id" => "area_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AreaEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CosEntity

```ruby
cos = client.Cos
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Cos.load({ "id" => "cos_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CosEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DeriveEntity

```ruby
derive = client.Derive
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Derive.load({ "id" => "derive_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DeriveEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FactorEntity

```ruby
factor = client.Factor
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Factor.load({ "id" => "factor_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FactorEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IntegrateEntity

```ruby
integrate = client.Integrate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Integrate.load({ "id" => "integrate_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IntegrateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LogEntity

```ruby
log = client.Log
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Log.load({ "id" => "log_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LogEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SimplifyEntity

```ruby
simplify = client.Simplify
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Simplify.load({ "id" => "simplify_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SimplifyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SinEntity

```ruby
sin = client.Sin
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Sin.load({ "id" => "sin_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SinEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TanEntity

```ruby
tan = client.Tan
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Tan.load({ "id" => "tan_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TanEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TangentEntity

```ruby
tangent = client.Tangent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Tangent.load({ "id" => "tangent_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TangentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ZeroEntity

```ruby
zero = client.Zero
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `String` | Yes | The mathematical expression that was processed |
| `id` | `String` | No |  |
| `operation` | `String` | Yes | The mathematical operation that was performed |
| `result` | `String` | Yes | The result of the mathematical operation |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Zero.load({ "id" => "zero_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ZeroEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = NewtonSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

