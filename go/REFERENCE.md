# Newton Golang SDK Reference

Complete API reference for the Newton Golang SDK.


## NewtonSDK

### Constructor

```go
func NewNewtonSDK(options map[string]any) *NewtonSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `TestSDK(testopts, sdkopts map[string]any) *NewtonSDK`

Create a test client with mock features active. Both arguments may be `nil`.

```go
client := sdk.TestSDK(nil, nil)
```


### Instance Methods

#### `Abs(data map[string]any) NewtonEntity`

Create a new `Abs` entity instance. Pass `nil` for no initial data.

#### `Arcco(data map[string]any) NewtonEntity`

Create a new `Arcco` entity instance. Pass `nil` for no initial data.

#### `Arcsin(data map[string]any) NewtonEntity`

Create a new `Arcsin` entity instance. Pass `nil` for no initial data.

#### `Arctan(data map[string]any) NewtonEntity`

Create a new `Arctan` entity instance. Pass `nil` for no initial data.

#### `Area(data map[string]any) NewtonEntity`

Create a new `Area` entity instance. Pass `nil` for no initial data.

#### `Cos(data map[string]any) NewtonEntity`

Create a new `Cos` entity instance. Pass `nil` for no initial data.

#### `Derive(data map[string]any) NewtonEntity`

Create a new `Derive` entity instance. Pass `nil` for no initial data.

#### `Factor(data map[string]any) NewtonEntity`

Create a new `Factor` entity instance. Pass `nil` for no initial data.

#### `Integrate(data map[string]any) NewtonEntity`

Create a new `Integrate` entity instance. Pass `nil` for no initial data.

#### `Log(data map[string]any) NewtonEntity`

Create a new `Log` entity instance. Pass `nil` for no initial data.

#### `Simplify(data map[string]any) NewtonEntity`

Create a new `Simplify` entity instance. Pass `nil` for no initial data.

#### `Sin(data map[string]any) NewtonEntity`

Create a new `Sin` entity instance. Pass `nil` for no initial data.

#### `Tan(data map[string]any) NewtonEntity`

Create a new `Tan` entity instance. Pass `nil` for no initial data.

#### `Tangent(data map[string]any) NewtonEntity`

Create a new `Tangent` entity instance. Pass `nil` for no initial data.

#### `Zero(data map[string]any) NewtonEntity`

Create a new `Zero` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AbsEntity

```go
abs := client.Abs(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Abs(nil).Load(map[string]any{"id": "abs_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AbsEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArccoEntity

```go
arcco := client.Arcco(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Arcco(nil).Load(map[string]any{"id": "arcco_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArccoEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArcsinEntity

```go
arcsin := client.Arcsin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Arcsin(nil).Load(map[string]any{"id": "arcsin_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArcsinEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArctanEntity

```go
arctan := client.Arctan(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Arctan(nil).Load(map[string]any{"id": "arctan_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArctanEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AreaEntity

```go
area := client.Area(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Area(nil).Load(map[string]any{"id": "area_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AreaEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CosEntity

```go
cos := client.Cos(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Cos(nil).Load(map[string]any{"id": "cos_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CosEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DeriveEntity

```go
derive := client.Derive(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Derive(nil).Load(map[string]any{"id": "derive_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DeriveEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FactorEntity

```go
factor := client.Factor(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Factor(nil).Load(map[string]any{"id": "factor_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FactorEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IntegrateEntity

```go
integrate := client.Integrate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Integrate(nil).Load(map[string]any{"id": "integrate_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IntegrateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LogEntity

```go
log := client.Log(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Log(nil).Load(map[string]any{"id": "log_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LogEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SimplifyEntity

```go
simplify := client.Simplify(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Simplify(nil).Load(map[string]any{"id": "simplify_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SimplifyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SinEntity

```go
sin := client.Sin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Sin(nil).Load(map[string]any{"id": "sin_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SinEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TanEntity

```go
tan := client.Tan(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Tan(nil).Load(map[string]any{"id": "tan_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TanEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TangentEntity

```go
tangent := client.Tangent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Tangent(nil).Load(map[string]any{"id": "tangent_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TangentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ZeroEntity

```go
zero := client.Zero(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Zero(nil).Load(map[string]any{"id": "zero_id"}, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ZeroEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewNewtonSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

