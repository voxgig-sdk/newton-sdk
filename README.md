# Newton SDK

Tiny HTTP service for symbolic math: derivatives, integrals, simplification, factoring, trig and more

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Newton API

[Newton](https://newton.now.sh/api/v2) is a small HTTP micro-service for advanced math, written and hosted by [Gerald Nash (aunyks)](https://github.com/aunyks/newton-api). It exposes a single URL pattern that takes an operation name and a URL-encoded expression and returns a JSON result.

What you get from the API:
- Algebra: `simplify`, `factor`, `zeroes`
- Calculus: `derive`, `integrate`, `tangent`, `area`
- Trigonometry: `cos`, `sin`, `tan`, `arccos`, `arcsin`, `arctan`
- Misc: `abs`, `log`

All operations share the same shape: `GET /api/v2/:operation/:expression`, where `:expression` is URL-encoded (for example `x%5E2` for `x^2`). Tangent calls take `c|f(x)` and area calls take `c:d|f(x)`. Responses are JSON objects of the form `{"operation": ..., "expression": ..., "result": ...}`.

The service is open and unauthenticated, has CORS enabled, and is built on top of the metadelta library. There is no published rate-limit policy, so treat it as best-effort and cache results where you can.

## Try it

**TypeScript**
```bash
npm install newton
```

**Python**
```bash
pip install newton-sdk
```

**PHP**
```bash
composer require voxgig/newton-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/newton-sdk/go
```

**Ruby**
```bash
gem install newton-sdk
```

**Lua**
```bash
luarocks install newton-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { NewtonSDK } from 'newton'

const client = new NewtonSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o newton-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "newton": {
      "command": "/abs/path/to/newton-mcp"
    }
  }
}
```

## Entities

The API exposes 15 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Abs** | Absolute value of an expression via `GET /api/v2/abs/:expression`. | `/abs/{expression}` |
| **Arcco** | Inverse cosine (arccos) of an expression via `GET /api/v2/arccos/:expression`. | `/arccos/{expression}` |
| **Arcsin** | Inverse sine (arcsin) of an expression via `GET /api/v2/arcsin/:expression`. | `/arcsin/{expression}` |
| **Arctan** | Inverse tangent (arctan) of an expression via `GET /api/v2/arctan/:expression`. | `/arctan/{expression}` |
| **Area** | Definite integral / area under a curve via `GET /api/v2/area/:c:d|f(x)`. | `/area/{expression}` |
| **Cos** | Cosine of an expression via `GET /api/v2/cos/:expression`. | `/cos/{expression}` |
| **Derive** | Symbolic derivative of an expression via `GET /api/v2/derive/:expression`. | `/derive/{expression}` |
| **Factor** | Factored form of an expression via `GET /api/v2/factor/:expression`. | `/factor/{expression}` |
| **Integrate** | Symbolic indefinite integral via `GET /api/v2/integrate/:expression`. | `/integrate/{expression}` |
| **Log** | Logarithm of an expression via `GET /api/v2/log/:expression`. | `/log/{expression}` |
| **Simplify** | Algebraic simplification of an expression via `GET /api/v2/simplify/:expression`. | `/simplify/{expression}` |
| **Sin** | Sine of an expression via `GET /api/v2/sin/:expression`. | `/sin/{expression}` |
| **Tan** | Tangent of an expression via `GET /api/v2/tan/:expression`. | `/tan/{expression}` |
| **Tangent** | Tangent line to f(x) at point c via `GET /api/v2/tangent/:c|f(x)`. | `/tangent/{expression}` |
| **Zero** | Real zeroes (roots) of an expression via `GET /api/v2/zeroes/:expression`. | `/zeroes/{expression}` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from newton_sdk import NewtonSDK

client = NewtonSDK({})


# Load a specific abs
abs, err = client.Abs(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'newton_sdk.php';

$client = new NewtonSDK([]);


// Load a specific abs
[$abs, $err] = $client->Abs(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/newton-sdk/go"

client := sdk.NewNewtonSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "Newton_sdk"

client = NewtonSDK.new({})


# Load a specific abs
abs, err = client.Abs(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("newton_sdk")

local client = sdk.new({})


-- Load a specific abs
local abs, err = client:Abs(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = NewtonSDK.test()
const result = await client.Abs().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = NewtonSDK.test(None, None)
result, err = client.Abs(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = NewtonSDK::test(null, null);
[$result, $err] = $client->Abs(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Abs(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = NewtonSDK.test(nil, nil)
result, err = client.Abs(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Abs(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Newton API

- Upstream: [https://newton.now.sh/api/v2](https://newton.now.sh/api/v2)
- API docs: [https://github.com/aunyks/newton-api](https://github.com/aunyks/newton-api)

- Licensed under the **GNU General Public License v3.0** (2016-2020).
- Source is published on GitHub by the author, [Gerald Nash (aunyks)](https://github.com/aunyks/newton-api).
- Derivative works that redistribute the service must remain under a GPL-compatible licence.
- This SDK only wraps the public HTTP endpoints and does not relicense the upstream code.

---

Generated from the Newton API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
