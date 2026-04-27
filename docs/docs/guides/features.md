---
title: Features
description: An overview of what pgxcli brings to your PostgreSQL workflow.
sidebar_position: 6
---

pgxcli is an interactive PostgreSQL REPL built in Go. This page highlights the core features that make it fast, readable, and comfortable to use.

---

## Syntax Highlighting

pgxcli applies real-time syntax highlighting to your SQL as you type, using [Chroma](https://github.com/alecthomas/chroma) — the same engine that powers many static site generators and code tools.

![pgxcli syntax highlighting — monokai and onedark themes](/img/syntax.png)

Colors are applied to keywords, identifiers, string literals, operators, and comments, so long queries remain easy to read at a glance.

pgxcli automatically detects your terminal's color depth — TrueColor, 256-color, or 16-color — and picks the right formatter for your environment.

### Changing the Theme

The highlighting theme is controlled by the `style` setting in your config file:

```toml
[main]
style = "monokai"
```

Some popular choices:

| Theme | Description |
|-------|-------------|
| `monokai` | Default — warm, high-contrast dark theme |
| `dracula` | Cool purples and pinks |
| `nord` | Arctic, blue-toned dark theme |
| `onedark` | Atom One Dark port |
| `catppuccin-mocha` | Soft pastel dark theme |
| `github-dark` | GitHub's dark mode palette |
| `gruvbox` | Retro groove dark theme |
| `solarized-dark` | Classic Solarized dark |
| `solarized-light` | Classic Solarized light |

:::tip
Browse the full list of available themes at [xyproto.github.io/splash/docs](https://xyproto.github.io/splash/docs/index.html).
:::

See [Configuration → style](/docs/guides/configuration#style) for the complete reference.

---

## Autocompletion

pgxcli completes SQL keywords as you type. Press `Tab` to cycle through suggestions. Completion is context-aware and works across `SELECT`, `FROM`, `WHERE`, `JOIN`, and other clauses.

---

## Persistent History

Every query you run is saved to a history file (`.pgxcli_history.jsonl` by default). Use the arrow keys to navigate previous commands. Up to 1000 entries are kept.

The history file location is configurable — see [Configuration → history_file](/docs/guides/configuration#history_file).

---

## Table Output

Query results are rendered as formatted tables. pgxcli supports a wide range of border styles — from plain `ascii` to decorative themed styles like `rounded`, `markdown`, and `galaxy`.

:::tip
Use `table.style = "markdown"` to copy-paste query results directly into Markdown documents.
:::

See [Configuration → table.style](/docs/guides/configuration#tablestyle) for the full style list.

---

## Pager Support

When query output is long, pgxcli pipes it through a pager (`less` on Linux/macOS, `more` on Windows) so you can scroll through results without flooding your terminal.

Pager behavior is configurable:

| Value | Behavior |
|-------|----------|
| `auto` | Page only when output exceeds terminal height |
| `always` | Always page |
| `never` | Print directly to the terminal |

Set a custom pager with the `PAGER` environment variable, e.g. `PAGER="less -S"`.

---

## Multi-Statement Execution

Paste multiple SQL statements separated by semicolons and pgxcli will execute each one in order. If a statement fails, the `on_error` setting controls whether execution stops or continues.

See [Special Commands](/docs/guides/special-commands#sql-execution) for more detail.
