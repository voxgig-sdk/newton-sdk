# Newton PHP SDK Reference

Complete API reference for the Newton PHP SDK.


## NewtonSDK

### Constructor

```php
require_once __DIR__ . '/newton_sdk.php';

$client = new NewtonSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NewtonSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = NewtonSDK::test();
```


### Instance Methods

#### `Abs($data = null)`

Create a new `AbsEntity` instance. Pass `null` for no initial data.

#### `Arcco($data = null)`

Create a new `ArccoEntity` instance. Pass `null` for no initial data.

#### `Arcsin($data = null)`

Create a new `ArcsinEntity` instance. Pass `null` for no initial data.

#### `Arctan($data = null)`

Create a new `ArctanEntity` instance. Pass `null` for no initial data.

#### `Area($data = null)`

Create a new `AreaEntity` instance. Pass `null` for no initial data.

#### `Cos($data = null)`

Create a new `CosEntity` instance. Pass `null` for no initial data.

#### `Derive($data = null)`

Create a new `DeriveEntity` instance. Pass `null` for no initial data.

#### `Factor($data = null)`

Create a new `FactorEntity` instance. Pass `null` for no initial data.

#### `Integrate($data = null)`

Create a new `IntegrateEntity` instance. Pass `null` for no initial data.

#### `Log($data = null)`

Create a new `LogEntity` instance. Pass `null` for no initial data.

#### `Simplify($data = null)`

Create a new `SimplifyEntity` instance. Pass `null` for no initial data.

#### `Sin($data = null)`

Create a new `SinEntity` instance. Pass `null` for no initial data.

#### `Tan($data = null)`

Create a new `TanEntity` instance. Pass `null` for no initial data.

#### `Tangent($data = null)`

Create a new `TangentEntity` instance. Pass `null` for no initial data.

#### `Zero($data = null)`

Create a new `ZeroEntity` instance. Pass `null` for no initial data.

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AbsEntity

```php
$abs = $client->abs();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->abs()->load(["id" => "abs_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): AbsEntity`

Create a new `AbsEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArccoEntity

```php
$arcco = $client->arcco();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->arcco()->load(["id" => "arcco_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArccoEntity`

Create a new `ArccoEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArcsinEntity

```php
$arcsin = $client->arcsin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->arcsin()->load(["id" => "arcsin_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArcsinEntity`

Create a new `ArcsinEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArctanEntity

```php
$arctan = $client->arctan();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->arctan()->load(["id" => "arctan_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArctanEntity`

Create a new `ArctanEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## AreaEntity

```php
$area = $client->area();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->area()->load(["id" => "area_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): AreaEntity`

Create a new `AreaEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## CosEntity

```php
$cos = $client->cos();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->cos()->load(["id" => "cos_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): CosEntity`

Create a new `CosEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## DeriveEntity

```php
$derive = $client->derive();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->derive()->load(["id" => "derive_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): DeriveEntity`

Create a new `DeriveEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## FactorEntity

```php
$factor = $client->factor();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->factor()->load(["id" => "factor_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): FactorEntity`

Create a new `FactorEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## IntegrateEntity

```php
$integrate = $client->integrate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->integrate()->load(["id" => "integrate_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): IntegrateEntity`

Create a new `IntegrateEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## LogEntity

```php
$log = $client->log();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->log()->load(["id" => "log_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): LogEntity`

Create a new `LogEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SimplifyEntity

```php
$simplify = $client->simplify();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->simplify()->load(["id" => "simplify_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SimplifyEntity`

Create a new `SimplifyEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SinEntity

```php
$sin = $client->sin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->sin()->load(["id" => "sin_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SinEntity`

Create a new `SinEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TanEntity

```php
$tan = $client->tan();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->tan()->load(["id" => "tan_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TanEntity`

Create a new `TanEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TangentEntity

```php
$tangent = $client->tangent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->tangent()->load(["id" => "tangent_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TangentEntity`

Create a new `TangentEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ZeroEntity

```php
$zero = $client->zero();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->zero()->load(["id" => "zero_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ZeroEntity`

Create a new `ZeroEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new NewtonSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

