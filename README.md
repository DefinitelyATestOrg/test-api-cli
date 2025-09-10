# Bruce Test API CLI

The official CLI for the Bruce Test API REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/bruce-hill/bruce-test-api-cli/cmd/bruce-test-api@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/bruce-test-api/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
bruce-test-api [resource] [command] [flags]
```

```sh
bruce-test-api foos retrieve
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
