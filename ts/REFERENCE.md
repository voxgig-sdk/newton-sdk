# Newton TypeScript SDK Reference

Complete API reference for the Newton TypeScript SDK.


## NewtonSDK

### Constructor

```ts
new NewtonSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NewtonSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NewtonSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NewtonSDK` instance in test mode.


### Instance Methods

#### `Abs(data?: object)`

Create a new `Abs` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AbsEntity` instance.

#### `Arcco(data?: object)`

Create a new `Arcco` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArccoEntity` instance.

#### `Arcsin(data?: object)`

Create a new `Arcsin` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArcsinEntity` instance.

#### `Arctan(data?: object)`

Create a new `Arctan` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArctanEntity` instance.

#### `Area(data?: object)`

Create a new `Area` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AreaEntity` instance.

#### `Cos(data?: object)`

Create a new `Cos` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CosEntity` instance.

#### `Derive(data?: object)`

Create a new `Derive` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DeriveEntity` instance.

#### `Factor(data?: object)`

Create a new `Factor` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FactorEntity` instance.

#### `Integrate(data?: object)`

Create a new `Integrate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IntegrateEntity` instance.

#### `Log(data?: object)`

Create a new `Log` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LogEntity` instance.

#### `Simplify(data?: object)`

Create a new `Simplify` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SimplifyEntity` instance.

#### `Sin(data?: object)`

Create a new `Sin` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SinEntity` instance.

#### `Tan(data?: object)`

Create a new `Tan` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TanEntity` instance.

#### `Tangent(data?: object)`

Create a new `Tangent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TangentEntity` instance.

#### `Zero(data?: object)`

Create a new `Zero` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ZeroEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NewtonSDK.test()`.

**Returns:** `NewtonSDK` instance in test mode.


---

## AbsEntity

```ts
const abs = client.Abs()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Abs().load({ id: 'abs_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AbsEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArccoEntity

```ts
const arcco = client.Arcco()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Arcco().load({ id: 'arcco_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArccoEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArcsinEntity

```ts
const arcsin = client.Arcsin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Arcsin().load({ id: 'arcsin_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArcsinEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArctanEntity

```ts
const arctan = client.Arctan()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Arctan().load({ id: 'arctan_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArctanEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AreaEntity

```ts
const area = client.Area()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Area().load({ id: 'area_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AreaEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CosEntity

```ts
const cos = client.Cos()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Cos().load({ id: 'cos_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CosEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DeriveEntity

```ts
const derive = client.Derive()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Derive().load({ id: 'derive_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DeriveEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FactorEntity

```ts
const factor = client.Factor()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Factor().load({ id: 'factor_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FactorEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IntegrateEntity

```ts
const integrate = client.Integrate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Integrate().load({ id: 'integrate_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IntegrateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LogEntity

```ts
const log = client.Log()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Log().load({ id: 'log_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LogEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SimplifyEntity

```ts
const simplify = client.Simplify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Simplify().load({ id: 'simplify_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SimplifyEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SinEntity

```ts
const sin = client.Sin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Sin().load({ id: 'sin_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SinEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TanEntity

```ts
const tan = client.Tan()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Tan().load({ id: 'tan_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TanEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TangentEntity

```ts
const tangent = client.Tangent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Tangent().load({ id: 'tangent_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TangentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ZeroEntity

```ts
const zero = client.Zero()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Zero().load({ id: 'zero_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ZeroEntity` instance with the same client and
options.

#### `client()`

Return the parent `NewtonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NewtonSDK({
  feature: {
    test: { active: true },
  }
})
```

