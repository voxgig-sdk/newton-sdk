# Newton Python SDK Reference

Complete API reference for the Newton Python SDK.


## NewtonSDK

### Constructor

```python
from newton_sdk import NewtonSDK

client = NewtonSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NewtonSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = NewtonSDK.test()
```


### Instance Methods

#### `Abs(data=None)`

Create a new `AbsEntity` instance. Pass `None` for no initial data.

#### `Arcco(data=None)`

Create a new `ArccoEntity` instance. Pass `None` for no initial data.

#### `Arcsin(data=None)`

Create a new `ArcsinEntity` instance. Pass `None` for no initial data.

#### `Arctan(data=None)`

Create a new `ArctanEntity` instance. Pass `None` for no initial data.

#### `Area(data=None)`

Create a new `AreaEntity` instance. Pass `None` for no initial data.

#### `Cos(data=None)`

Create a new `CosEntity` instance. Pass `None` for no initial data.

#### `Derive(data=None)`

Create a new `DeriveEntity` instance. Pass `None` for no initial data.

#### `Factor(data=None)`

Create a new `FactorEntity` instance. Pass `None` for no initial data.

#### `Integrate(data=None)`

Create a new `IntegrateEntity` instance. Pass `None` for no initial data.

#### `Log(data=None)`

Create a new `LogEntity` instance. Pass `None` for no initial data.

#### `Simplify(data=None)`

Create a new `SimplifyEntity` instance. Pass `None` for no initial data.

#### `Sin(data=None)`

Create a new `SinEntity` instance. Pass `None` for no initial data.

#### `Tan(data=None)`

Create a new `TanEntity` instance. Pass `None` for no initial data.

#### `Tangent(data=None)`

Create a new `TangentEntity` instance. Pass `None` for no initial data.

#### `Zero(data=None)`

Create a new `ZeroEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> tuple`

Make a direct HTTP request to any API endpoint. Returns `(result, err)`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `(result_dict, err)`

#### `prepare(fetchargs=None) -> tuple`

Prepare a fetch definition without sending. Returns `(fetchdef, err)`.


---

## AbsEntity

```python
abs = client.Abs()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Abs().load({"id": "abs_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AbsEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArccoEntity

```python
arcco = client.Arcco()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Arcco().load({"id": "arcco_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArccoEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArcsinEntity

```python
arcsin = client.Arcsin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Arcsin().load({"id": "arcsin_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArcsinEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArctanEntity

```python
arctan = client.Arctan()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Arctan().load({"id": "arctan_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArctanEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AreaEntity

```python
area = client.Area()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Area().load({"id": "area_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AreaEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CosEntity

```python
cos = client.Cos()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Cos().load({"id": "cos_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CosEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DeriveEntity

```python
derive = client.Derive()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Derive().load({"id": "derive_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DeriveEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FactorEntity

```python
factor = client.Factor()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Factor().load({"id": "factor_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FactorEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IntegrateEntity

```python
integrate = client.Integrate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Integrate().load({"id": "integrate_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IntegrateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LogEntity

```python
log = client.Log()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Log().load({"id": "log_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LogEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SimplifyEntity

```python
simplify = client.Simplify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Simplify().load({"id": "simplify_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SimplifyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SinEntity

```python
sin = client.Sin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Sin().load({"id": "sin_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SinEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TanEntity

```python
tan = client.Tan()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Tan().load({"id": "tan_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TanEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TangentEntity

```python
tangent = client.Tangent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Tangent().load({"id": "tangent_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TangentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ZeroEntity

```python
zero = client.Zero()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> tuple`

Load a single entity matching the given criteria.

```python
result, err = client.Zero().load({"id": "zero_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ZeroEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = NewtonSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

