# djm-cli

A small Go-first CLI for creating and working with server-rendered web applications built from the Daniel J. Manning reusable Go web modules.

The CLI keeps project setup intentionally simple: clone a working starter, install focused modules when needed, run the development server, and execute the standard project checks without introducing a separate framework runtime.

## Commands

```text
djm new <name> [--module <path>]
djm add <feature>
djm dev
djm check
```

## Create a project

```bash
djm new good-app
```

By default, the generated Go module path is:

```text
github.com/danieljmanningdev/good-app
```

Override it explicitly:

```bash
djm new good-app --module github.com/example/good-app
```

Or set a reusable module prefix:

```bash
export DJM_MODULE_PREFIX=github.com/example
```

`djm new` currently:

1. Clones [`go-starter-auth-app`](https://github.com/danieljmanningdev/go-starter-auth-app).
2. Removes the starter repository Git history.
3. Updates the Go module path.
4. Runs `go mod tidy`.
5. Initialises a fresh Git repository.

If setup fails after cloning, the partially generated project directory is removed.

## Add a module

```bash
djm add core
djm add auth
djm add security
djm add jsonld
djm add file-utils
```

Available modules map to:

| Feature | Module |
| --- | --- |
| `core` | `github.com/danieljmanningdev/go-web-core` |
| `auth` | `github.com/danieljmanningdev/go-web-auth` |
| `security` | `github.com/danieljmanningdev/go-web-security` |
| `jsonld` | `github.com/danieljmanningdev/go-jsonld-schema` |
| `file-utils` | `github.com/danieljmanningdev/go-file-utils` |

The command uses `go get` so the dependency remains under normal Go module management.

## Run development server

From a compatible generated application:

```bash
djm dev
```

This runs:

```bash
go run ./cmd/server
```

## Run project checks

```bash
djm check
```

The command runs:

```text
gofmt -w .
go vet ./...
go test ./...
govulncheck ./...   when installed
git diff --check
```

## Install locally

Clone the repository and install the `djm` binary:

```bash
git clone https://github.com/danieljmanningdev/djm-cli.git
cd djm-cli
go install ./cmd/djm
```

Ensure your Go binary directory is on `PATH`, then run:

```bash
djm help
```

Once releases are available, the command can also be installed directly from the module path:

```bash
go install github.com/danieljmanningdev/djm-cli/cmd/djm@latest
```

## Philosophy

`djm-cli` is an orchestration layer rather than a framework.

The reusable implementation stays in focused Go modules such as `go-web-core`, `go-web-security`, and `go-web-auth`. The CLI exists to make those modules and starter applications quicker to use without duplicating them into one monolithic codebase.

The wider stack remains server-rendered and Go-first, with HTML as the default, HTMX for progressive enhancement, and client-side JavaScript added only where browser-side behaviour genuinely requires it.

## Validate UX documentation

Projects using the DJM UX JSON documentation structure can validate
screens, components, flows, and design tokens against the schemas
embedded in the CLI.

```bash
djm ux validate
```

## Development

```bash
gofmt -w .
go vet ./...
go test ./...
git diff --check
```

## License

Licensed under the MIT License. See [`LICENSE`](LICENSE).
