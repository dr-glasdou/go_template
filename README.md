# Go Template

Go starter template with hot reload, linting, and git hooks.

## Quick Start

```bash
make setup   # Install deps + git hooks
make dev    # Run with hot reload
make build  # Build binary to ./bin/gotemplate
```

## Commands

| Command | Description |
|---------|-------------|
| `make setup` | Install deps + git hooks |
| `make dev` | Hot reload (air) |
| `make test` | Run tests (-v -race) |
| `make lint` | Run golangci-lint |
| `make build` | Build binary |
| `make clean` | Clean build artifacts |

## Project Structure

```
.
├── main.go          # Entry point
├── Makefile         # Build commands
├── go.mod           # Go module
├── .air.toml        # Hot reload config
├── .golangci.yml    # Linter config
└── lefthook.yml     # Git hooks config
```

## Requirements

- Go 1.26+
- golangci-lint
- air (for hot reload)
- lefthook (for git hooks)
