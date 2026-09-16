# Newton API

The Newton API is a micro-service that provides advanced math functionalities such as derivatives, integrals, simplifications, and more. Users can send requests with math expressions and receive results in JSON format, making it a powerful tool for mathematical computations.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 15 entities and 15 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### Abs

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Arcco

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Arcsin

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Arctan

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Area

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Cos

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Derive

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Factor

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Integrate

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Log

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Simplify

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Sin

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Tan

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Tangent

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Zero

Results: Successful operation.

SDK operations: `load`.

Key fields to recognise:

- `expression`: The mathematical expression that was processed
- `operation`: The mathematical operation that was performed
- `result`: The result of the mathematical operation

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| Abs | `load` | `GET /abs/{expression}` | See reference |
| Arcco | `load` | `GET /arccos/{expression}` | See reference |
| Arcsin | `load` | `GET /arcsin/{expression}` | See reference |
| Arctan | `load` | `GET /arctan/{expression}` | See reference |
| Area | `load` | `GET /area/{expression}` | See reference |
| Cos | `load` | `GET /cos/{expression}` | See reference |
| Derive | `load` | `GET /derive/{expression}` | See reference |
| Factor | `load` | `GET /factor/{expression}` | See reference |
| Integrate | `load` | `GET /integrate/{expression}` | See reference |
| Log | `load` | `GET /log/{expression}` | See reference |
| Simplify | `load` | `GET /simplify/{expression}` | See reference |
| Sin | `load` | `GET /sin/{expression}` | See reference |
| Tan | `load` | `GET /tan/{expression}` | See reference |
| Tangent | `load` | `GET /tangent/{expression}` | See reference |
| Zero | `load` | `GET /zeroes/{expression}` | See reference |

## Connect to the API

- Production server: `https://newton.now.sh/api/v2`

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Lua | `lua/` | Build from source |
| PHP | `php/` | Build from source |
| Python | `py/` | Build from source |
| Ruby | `rb/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `newton_list`: List records for an entity. No active entity supports this operation.
- `newton_load`: Load one record for an entity. Supported entities: `abs`, `arcco`, `arcsin`, `arctan`, `area`, `cos`, `derive`, `factor`, `integrate`, `log`, `simplify`, `sin`, `tan`, `tangent`, `zero`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

