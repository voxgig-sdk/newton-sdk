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
| `$options["apikey"]` | `string` | API key for authentication. |
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

Make a direct HTTP request to any API endpoint. Returns `[$result, $err]`.

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

**Returns:** `array [$result, $err]`

#### `prepare(array $fetchargs = []): array`

Prepare a fetch definition without sending the request. Returns `[$fetchdef, $err]`.


---

## AbsEntity

```php
$abs = $client->Abs();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Abs()->load(["id" => "abs_id"]);
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
$arcco = $client->Arcco();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Arcco()->load(["id" => "arcco_id"]);
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
$arcsin = $client->Arcsin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Arcsin()->load(["id" => "arcsin_id"]);
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
$arctan = $client->Arctan();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Arctan()->load(["id" => "arctan_id"]);
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
$area = $client->Area();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Area()->load(["id" => "area_id"]);
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
$cos = $client->Cos();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Cos()->load(["id" => "cos_id"]);
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
$derive = $client->Derive();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Derive()->load(["id" => "derive_id"]);
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
$factor = $client->Factor();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Factor()->load(["id" => "factor_id"]);
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
$integrate = $client->Integrate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Integrate()->load(["id" => "integrate_id"]);
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
$log = $client->Log();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Log()->load(["id" => "log_id"]);
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
$simplify = $client->Simplify();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Simplify()->load(["id" => "simplify_id"]);
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
$sin = $client->Sin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Sin()->load(["id" => "sin_id"]);
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
$tan = $client->Tan();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Tan()->load(["id" => "tan_id"]);
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
$tangent = $client->Tangent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Tangent()->load(["id" => "tangent_id"]);
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
$zero = $client->Zero();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `expression` | ``$STRING`` | Yes |  |
| `operation` | ``$STRING`` | Yes |  |
| `result` | ``$STRING`` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Zero()->load(["id" => "zero_id"]);
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

