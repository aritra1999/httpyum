# AGENTS.md - Guidelines for AI Coding Agents

This document provides essential information for AI coding agents working on the httpyum codebase.

## Project Overview

**httpyum** is a fast, interactive HTTP request runner for `.http` files built as a TUI application in Go using the Bubbletea framework (Elm architecture pattern).

## Build/Test/Lint Commands

```bash
make build          # Build the binary
make run            # Build and run with example.http
make dev            # Run in development mode (no build step)
make all            # Full cycle: clean, deps, fmt, vet, test, build

make test           # Run all tests
make test-verbose   # Run tests with verbose output
make test-race      # Run tests with race detector
make coverage       # Generate coverage report (coverage.html)

# Run a single test
go test -v -run TestFunctionName ./path/to/package

# Run tests in a specific package
go test -v ./internal/parser/...

make fmt            # Format code (required before committing)
make vet            # Run go vet
make lint           # Run all linters (go vet + staticcheck)
make tidy           # Tidy go modules
```

## Project Structure

```
httpyum/
├── cmd/httpyum/main.go      # CLI entry point
├── internal/
│   ├── parser/              # .http file parsing
│   │   ├── parser.go        # Core parsing logic
│   │   ├── types.go         # Data types (Request, Variable, ParsedFile)
│   │   ├── errors.go        # Custom ParseError type
│   │   └── dotenv.go        # Environment variable loading
│   ├── client/              # HTTP execution
│   │   ├── executor.go      # Request execution
│   │   ├── types.go         # Response types
│   │   └── errors.go        # Custom ExecutionError type
│   ├── ui/                  # TUI components (Bubbletea)
│   │   ├── tui.go           # Model, Init, Update, View
│   │   ├── views.go         # View rendering
│   │   ├── components.go    # Reusable components
│   │   ├── styles.go        # Lipgloss styles
│   │   └── list_delegate.go # Custom list delegate
│   └── config/              # CLI configuration
└── Makefile                 # Build automation
```

## Code Style Guidelines

### Naming Conventions

- **Packages**: lowercase, single word (`parser`, `client`, `ui`)
- **Files**: lowercase with underscores (`list_delegate.go`)
- **Exported**: PascalCase (`Request`, `NewExecutor`)
- **Unexported**: camelCase (`requestItem`, `checkText`)
- **Constants**: camelCase for internal, PascalCase for exported

### Import Organization

Group imports with blank line separators: standard library, internal packages, external packages.

```go
import (
    "fmt"
    "strings"

    "httpyum/internal/client"
    "httpyum/internal/parser"

    tea "github.com/charmbracelet/bubbletea"
)
```

### File Organization

- **types.go**: Define all data structures for the package
- **errors.go**: Define custom error types with constructors
- **Module-level vars**: Regex patterns, styles at package level

### Error Handling

- **Custom error types** with constructors in `errors.go`:
```go
type ParseError struct { Line int; Message string }
func NewParseError(line int, message string) *ParseError { ... }
func (e *ParseError) Error() string { ... }
```
- **Error wrapping**: `return nil, fmt.Errorf("reading file: %w", err)`
- **Early returns** on error - check and return immediately
- **Result structs** for complex operations with `Success bool` and `Error error` fields

### Type Patterns

- **Constructors**: `func NewExecutor(vars map[string]string) *Executor`
- **Method receivers**: `func (e *Executor) Execute(req *Request) *Result`

### TUI Architecture (Elm Pattern)

- **Model**: Application state struct
- **Init()**: Returns initial command
- **Update(msg)**: Handles messages, returns (Model, Cmd)
- **View()**: Renders model to string

View states: `ViewList`, `ViewResponse`, `ViewLoading`, `ViewError`

## Commit Message Convention

Uses Conventional Commits for automated releases:

- `feat:` - New feature (minor version bump)
- `fix:` - Bug fix (patch version bump)
- `perf:` / `refactor:` - Patch version bump
- `docs:` / `chore:` - No version bump
- `!` or `BREAKING CHANGE:` - Major version bump

## Dependencies

- **bubbletea**: TUI framework (Elm architecture)
- **bubbles**: UI components (list, viewport, etc.)
- **lipgloss**: Terminal styling

## Development Workflow

```bash
# Before committing
make fmt && make lint && make test

# Quick iteration
make dev

# Full validation
make all
```

## CI/CD

- Tests run on PR and push to main (Go 1.24, race detector)
- Coverage uploaded to Codecov
- Auto-release on conventional commit analysis
