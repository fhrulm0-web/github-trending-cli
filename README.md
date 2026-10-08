# ghtrend

A fast, lightweight command-line interface (CLI) tool built in Go to discover trending GitHub repositories directly from your terminal.

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/fhrulm0-web/github-trending-cli?include_prereleases&color=blue)](https://github.com/fhrulm0-web/github-trending-cli/releases)

---

## Features

- **Trending Repositories**: Discover top GitHub repositories sorted by star count.
- **Flexible Time Ranges**: Filter by `day`, `week`, `month`, or `year`.
- **Language Filter**: Filter repositories by programming language (e.g., `go`, `rust`, `python`, `typescript`).
- **Custom Limits**: Display anywhere between `1` and `100` repositories per query.
- **Clean Terminal Output**: Formatted tabular view using Go standard library (`text/tabwriter`) with formatted numbers and description truncation.
- **Zero External Bloat**: Single lightweight binary with minimal dependencies (`cobra` + standard library).

---

## Installation

### Option 1: Via go install (Recommended)

Ensure you have [Go](https://go.dev/dl/) installed:

```bash
go install github.com/fhrulm0-web/github-trending-cli@latest
```

The `ghtrend` binary will be installed to `$GOPATH/bin` (or `~/go/bin`). Ensure that directory is in your system's `PATH`.

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/fhrulm0-web/github-trending-cli.git
cd github-trending-cli

# Build binary
go build -o ghtrend .

# Run directly
./ghtrend --help
```

---

## Usage

### Quick Examples

```bash
# Default: Top 10 trending repositories from the past week
ghtrend

# Top 5 trending Go repositories created today
ghtrend --duration day --language go --limit 5

# Top 20 trending repositories this month
ghtrend --duration month --limit 20

# Top trending Rust projects this year
ghtrend --duration year --language rust
```

### Output Example

```text
Trending go repositories (month) — sorted by stars
────────────────────────────────────────────────────────────────────────────────
 #   Repository                          Stars     Language     Description
 ─   ────────────────────                ───────   ──────────   ────────────────────────────
 1   yetone/magpie                       6,625     Go           Every agent's model. One place. Codex on DeepSeek…
 2   linguo2625469/workbuddy2api-panel   2,226     Go           OpenAI compatible API gateway with Web UI…
 3   unreallabsai/unreal-agent           2,175     Go           Async-first agent harness
────────────────────────────────────────────────────────────────────────────────
Showing 3 results | Source: GitHub Search API | Duration: month
```

---

## Command Flags

| Flag | Shorthand | Type | Default | Description |
|---|---|---|---|---|
| `--duration` | `-d` | `string` | `week` | Time range: `day`, `week`, `month`, `year` |
| `--limit` | `-l` | `int` | `10` | Number of repositories to display (`1`–`100`) |
| `--language` | — | `string` | `""` | Filter by programming language (e.g. `go`, `rust`) |
| `--help` | `-h` | — | — | Show help menu |
| `--version` | `-v` | — | — | Print binary version |

---

## Project Structure

```text
github-trending-cli/
├── cmd/
│   ├── root.go             # Cobra CLI configuration and flag definitions
│   └── root_test.go        # CLI integration and validation tests
├── internal/
│   ├── display/
│   │   ├── table.go        # Terminal table formatting and text truncation
│   │   └── table_test.go   # Table rendering unit tests
│   └── github/
│       ├── client.go       # HTTP client and GitHub API request handling
│       ├── client_test.go  # Mocked HTTP server tests
│       ├── models.go       # API JSON response data models
│       ├── query.go        # Search query builder and date math
│       └── query_test.go   # Query builder unit tests
├── main.go                 # Application entry point
├── SPEC.md                 # Single source of truth specification
├── AGENTS.md               # Guidelines and rules for AI agents
├── CHANGELOG.md            # Version release history
├── CONTRIBUTING.md         # Contribution guidelines
├── LICENSE                 # MIT License
└── go.mod                  # Go module definition
```

---

## Running Tests

Run the test suite across all packages:

```bash
go test -v ./...
```

Run static analysis:

```bash
go vet ./...
```

---

## Contributing

Contributions are welcome. Please check out [CONTRIBUTING.md](CONTRIBUTING.md) and feel free to submit issues or pull requests.

---

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
