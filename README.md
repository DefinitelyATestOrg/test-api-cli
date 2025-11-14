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
bruce-test-api client post-fnord \
  --first-pos A \
  --second-pos B \
  --array-items 1 \
  --+array-item \
  --array-items 2 \
  --name.full_name 'Abraham Lincoln' \
  --name.nickname 'Honest Abe' \
  --job President
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
