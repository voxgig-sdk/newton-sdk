# Newton PHP SDK



The PHP SDK for the Newton API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Abs()` — with named operations (`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/newton-sdk/releases](https://github.com/voxgig-sdk/newton-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'newton_sdk.php';

$client = new NewtonSDK();
```

### 3. Load an abs

```php
try {
    // load() returns the ENTITY — call data_get() for the Abs record (throws on error).
    $abs = $client->Abs()->load(["id" => "example_id"]);
    print_r($abs);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $arcco = $client->Arcco()->load(["id" => "example_id"]);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = NewtonSDK::test([
    "entity" => ["arcco" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$arcco = $client->Arcco()->load(["id" => "test01"]);
print_r($arcco);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new NewtonSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
NEWTON_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### NewtonSDK

```php
require_once 'newton_sdk.php';
$client = new NewtonSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = NewtonSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### NewtonSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Abs` | `($data): AbsEntity` | Create an Abs entity instance. |
| `Arcco` | `($data): ArccoEntity` | Create an Arcco entity instance. |
| `Arcsin` | `($data): ArcsinEntity` | Create an Arcsin entity instance. |
| `Arctan` | `($data): ArctanEntity` | Create an Arctan entity instance. |
| `Area` | `($data): AreaEntity` | Create an Area entity instance. |
| `Cos` | `($data): CosEntity` | Create a Cos entity instance. |
| `Derive` | `($data): DeriveEntity` | Create a Derive entity instance. |
| `Factor` | `($data): FactorEntity` | Create a Factor entity instance. |
| `Integrate` | `($data): IntegrateEntity` | Create an Integrate entity instance. |
| `Log` | `($data): LogEntity` | Create a Log entity instance. |
| `Simplify` | `($data): SimplifyEntity` | Create a Simplify entity instance. |
| `Sin` | `($data): SinEntity` | Create a Sin entity instance. |
| `Tan` | `($data): TanEntity` | Create a Tan entity instance. |
| `Tangent` | `($data): TangentEntity` | Create a Tangent entity instance. |
| `Zero` | `($data): ZeroEntity` | Create a Zero entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$abs = $client->Abs();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Abs record (throws on error).
$abs = $client->Abs()->load(["id" => "abs_id"]);
```


### Arcco

Create an instance: `$arcco = $client->Arcco();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Arcco record (throws on error).
$arcco = $client->Arcco()->load(["id" => "arcco_id"]);
```


### Arcsin

Create an instance: `$arcsin = $client->Arcsin();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Arcsin record (throws on error).
$arcsin = $client->Arcsin()->load(["id" => "arcsin_id"]);
```


### Arctan

Create an instance: `$arctan = $client->Arctan();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Arctan record (throws on error).
$arctan = $client->Arctan()->load(["id" => "arctan_id"]);
```


### Area

Create an instance: `$area = $client->Area();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Area record (throws on error).
$area = $client->Area()->load(["id" => "area_id"]);
```


### Cos

Create an instance: `$cos = $client->Cos();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Cos record (throws on error).
$cos = $client->Cos()->load(["id" => "cos_id"]);
```


### Derive

Create an instance: `$derive = $client->Derive();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Derive record (throws on error).
$derive = $client->Derive()->load(["id" => "derive_id"]);
```


### Factor

Create an instance: `$factor = $client->Factor();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Factor record (throws on error).
$factor = $client->Factor()->load(["id" => "factor_id"]);
```


### Integrate

Create an instance: `$integrate = $client->Integrate();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Integrate record (throws on error).
$integrate = $client->Integrate()->load(["id" => "integrate_id"]);
```


### Log

Create an instance: `$log = $client->Log();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Log record (throws on error).
$log = $client->Log()->load(["id" => "log_id"]);
```


### Simplify

Create an instance: `$simplify = $client->Simplify();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Simplify record (throws on error).
$simplify = $client->Simplify()->load(["id" => "simplify_id"]);
```


### Sin

Create an instance: `$sin = $client->Sin();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Sin record (throws on error).
$sin = $client->Sin()->load(["id" => "sin_id"]);
```


### Tan

Create an instance: `$tan = $client->Tan();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Tan record (throws on error).
$tan = $client->Tan()->load(["id" => "tan_id"]);
```


### Tangent

Create an instance: `$tangent = $client->Tangent();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Tangent record (throws on error).
$tangent = $client->Tangent()->load(["id" => "tangent_id"]);
```


### Zero

Create an instance: `$zero = $client->Zero();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `expression` | `string` |  |
| `operation` | `string` |  |
| `result` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Zero record (throws on error).
$zero = $client->Zero()->load(["id" => "zero_id"]);
```


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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── newton_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`newton_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$arcco = $client->Arcco();
$arcco->load(["id" => "example_id"]);

// $arcco->data_get() now returns the arcco data from the last load
// $arcco->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
