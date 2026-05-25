# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make build          # Build to bin/app
make test           # go test ./...
make test-integration # go test -tags=integration ./...
make lint           # golangci-lint run
make precommit      # lint + test (run before committing)
make clean          # remove bin/

go test ./internal/config      # test a single package
go test -v ./internal/database # verbose
```

## Architecture

`pgxcli` is an interactive PostgreSQL REPL CLI in Go. It follows a strict layered design:

```
cmd/pgxcli/main.go        → entry point; wires context, printer, cobra root command
internal/cli/             → cobra command, flags, connection setup, password prompting
internal/app/             → REPL loop, command routing, history, rendering
internal/app/ui/          → Bubble Tea REPL model + connection forms (huh) + UI components
internal/app/renderer/    → result formatting and table rendering helpers
internal/database/        → pgx connection, query/exec dispatch, special commands
internal/parser/          → SQL statement splitter
internal/config/          → TOML config load/validate (embeds defaults)
internal/completer/       → keyword autocompletion with schema metadata
internal/cliio/           → Printer interface (stdout/stderr abstraction)
internal/logger/          → slog-based file logger
```

### REPL data flow

1. `cli.NewRootCmd` loads config/logger, initializes pager, and connects `database.Client`.
2. `app.pgxCLI.Start()` builds the Bubble Tea model (`internal/app/ui`) with input, spinner, and status components.
3. Input is handled by Bubbline editline with history and SQL highlighting; multi-statement SQL is split by `parser.SplitSQLStatements`.
4. Builtins (e.g. `\clear`) are handled directly; special commands (`\d`, `\q`, `\c`, `\conninfo`) go through `pgxspecial`.
5. SQL statements execute via `database.Client.ExecuteQuery`; `OnErrorAction` controls stop/resume for multi-statement input.
6. `internal/app/renderer` formats output using table config; `cliio.Printer` emits output and uses a pager when needed.
7. History is persisted via editline (default `~/.pgxcli_history.jsonl`) on close.

### Special commands

`\d`, `\l`, `\dt`, `\q`, `\c`, `\conninfo` etc. are handled by the external `pgxspecial` package (`github.com/balajz/pgxspecial`). `database.executor.executeSpecial` wraps it; rows are materialized into `result.SpecialRow` before being handed back to the REPL.

### Configuration

- Config file auto-created at `~/.config/pgxcli/config.toml` on first run
- `"default"` is a sentinel value; `config.Load()` resolves it to OS-appropriate paths
- `config.validate()` blocks startup on invalid values
- `OnErrorAction`: `STOP` aborts multi-statement execution on error; `RESUME` continues

### Prompt placeholders

`\u` user, `\d` database, `\h` host (short), `\H` host (full), `\p` port, `\t` timestamp, `\n` newline — resolved in `database.Client.ParsePrompt`.

## Conventions

**Errors:** wrap with `fmt.Errorf("context: %w", err)`; log with key `"error"` (not `"err"`): `logger.Error("msg", "error", err)`.

**Logging:** `internal/logger` initializes a file-backed `slog.Logger`; pass the logger down; never use `log` package directly.

**Testing:** use `testify`; mock `conn` interface via `MockConn`/`MockRows` in `database/mocks_test.go`; config tests use `t.TempDir()` with `HOME`/`XDG_CONFIG_HOME` overridden.

**Lint:** golangci-lint is configured in `.golangci.yml` with `revive`, `misspell`, `gocyclo` (min 15), `goconst`, `unconvert`, `unparam`. Tests are excluded from linting. `internal/completer` is excluded from lint paths.

**Interfaces:** `Application`, `Printer`, `Connector` — prefer programming to interfaces for testability.
