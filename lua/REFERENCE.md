# Newton Lua SDK Reference

Complete API reference for the Newton Lua SDK.


## NewtonSDK

### Constructor

```lua
local sdk = require("newton_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Abs(data)`

Create a new `Abs` entity instance. Pass `nil` for no initial data.

#### `Arcco(data)`

Create a new `Arcco` entity instance. Pass `nil` for no initial data.

#### `Arcsin(data)`

Create a new `Arcsin` entity instance. Pass `nil` for no initial data.

#### `Arctan(data)`

Create a new `Arctan` entity instance. Pass `nil` for no initial data.

#### `Area(data)`

Create a new `Area` entity instance. Pass `nil` for no initial data.

#### `Cos(data)`

Create a new `Cos` entity instance. Pass `nil` for no initial data.

#### `Derive(data)`

Create a new `Derive` entity instance. Pass `nil` for no initial data.

#### `Factor(data)`

Create a new `Factor` entity instance. Pass `nil` for no initial data.

#### `Integrate(data)`

Create a new `Integrate` entity instance. Pass `nil` for no initial data.

#### `Log(data)`

Create a new `Log` entity instance. Pass `nil` for no initial data.

#### `Simplify(data)`

Create a new `Simplify` entity instance. Pass `nil` for no initial data.

#### `Sin(data)`

Create a new `Sin` entity instance. Pass `nil` for no initial data.

#### `Tan(data)`

Create a new `Tan` entity instance. Pass `nil` for no initial data.

#### `Tangent(data)`

Create a new `Tangent` entity instance. Pass `nil` for no initial data.

#### `Zero(data)`

Create a new `Zero` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AbsEntity

```lua
local abs = client:Abs(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Abs():load({ id = "abs_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AbsEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArccoEntity

```lua
local arcco = client:Arcco(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Arcco():load({ id = "arcco_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArccoEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArcsinEntity

```lua
local arcsin = client:Arcsin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Arcsin():load({ id = "arcsin_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArcsinEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArctanEntity

```lua
local arctan = client:Arctan(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Arctan():load({ id = "arctan_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArctanEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AreaEntity

```lua
local area = client:Area(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Area():load({ id = "area_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AreaEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CosEntity

```lua
local cos = client:Cos(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Cos():load({ id = "cos_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CosEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DeriveEntity

```lua
local derive = client:Derive(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Derive():load({ id = "derive_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DeriveEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FactorEntity

```lua
local factor = client:Factor(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Factor():load({ id = "factor_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FactorEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## IntegrateEntity

```lua
local integrate = client:Integrate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Integrate():load({ id = "integrate_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IntegrateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## LogEntity

```lua
local log = client:Log(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Log():load({ id = "log_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LogEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SimplifyEntity

```lua
local simplify = client:Simplify(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Simplify():load({ id = "simplify_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SimplifyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SinEntity

```lua
local sin = client:Sin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Sin():load({ id = "sin_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SinEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TanEntity

```lua
local tan = client:Tan(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Tan():load({ id = "tan_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TanEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TangentEntity

```lua
local tangent = client:Tangent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Tangent():load({ id = "tangent_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TangentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ZeroEntity

```lua
local zero = client:Zero(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | `string` | Yes |  |
| `operation` | `string` | Yes |  |
| `result` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Zero():load({ id = "zero_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ZeroEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

