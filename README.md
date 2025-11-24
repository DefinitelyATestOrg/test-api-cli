# Bruce Test API CLI

The official CLI for the Bruce Test API REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/bruce-test-api-cli/cmd/bruce-test-api@latest'
```

### Running Locally

```sh
go run cmd/bruce-test-api/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
bruce-test-api [resource] [command] [flags]
```

```sh
bruce-test-api foo \
  --version 1 \
  --user-id abc123 \
  --limit 20 \
  --tag red \
  --tag large \
  --x-flag fast \
  --x-flag debug \
  --x-flag verbose \
  --x-trace-id trace-9f82b1 \
<<JSON
{
  "preferences": {
    "theme": "dark",
    "alerts": true
  }
}
JSON
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
