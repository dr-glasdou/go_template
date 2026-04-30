# AGENTS.md - gotemplate

## Developer Commands

```bash
make dev        # Hot reload (air)
make setup      # Install deps + git hooks
make test       # Run tests (-v -race)
make lint       # Run golangci-lint
make build      # Build binary to ./bin/gotemplate
make clean      # Clean build artifacts
```

## Test Files

Test files use `*.test.go` pattern (not `*_test.go`). Update `.golangci.yml` exclusion when modifying.

## Git Hooks (lefthook)

- `pre-commit`: golangci-lint + gofumpt formatting
- `pre-push`: go test ./...
- `commit-msg`: validates conventional commits

Conventional commit types: `feat|fix|chore|docs|refactor|test|style|ci`

Format: `type(scope): message` (e.g., `feat: add hot reload`)

## Entry Point

`main.go` → builds to `./bin/gotemplate`
