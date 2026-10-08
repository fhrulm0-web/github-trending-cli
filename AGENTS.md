# AGENTS.md — github-trending-cli

> Operating guidelines and development rules for AI agents in this repository.
> Read this document before making any changes to the codebase.

---

## 1. Project Context

This project is a lightweight command-line interface (CLI) tool built in **Go** to display trending GitHub repositories.
The comprehensive specification resides in the internal specification documentation.

---

## 2. Core Rules (STRICTLY ENFORCED)

### FORBIDDEN
- **Do NOT use emojis or decorative icons** in README, markdown documentation, issue/PR templates, commit messages, or terminal UI. Maintain a clean, professional, and formal style without decorative symbols.
- **Do NOT add new dependencies** without explicit confirmation from the user. Currently authorized dependencies:
  - `github.com/spf13/cobra` (CLI parsing)
  - Packages from the Go standard library (`net/http`, `encoding/json`, `text/tabwriter`, `fmt`, `os`, `time`, etc.)
- **Do NOT implement features not defined in the project scope** (e.g., local caching, TUI, JSON output, authentication).
- **Do NOT delete or modify** `AGENTS.md`, `LICENSE`, or `CONTRIBUTING.md` without explicit user instruction.
- **Do NOT hardcode** values that can be defined as constants (API URLs, default values, limits, etc.).
- **Do NOT commit directly to `main`**. All changes must go through a dedicated feature branch.
- **Do NOT write documentation in languages other than English**. All documentation, code comments, and commit messages must be in English.

### MANDATORY
- Always verify that requested changes align with existing acceptance criteria before implementation.
- If requirements are ambiguous, **ask the user for clarification** before making assumptions.
- Maintain the modular directory structure defined in the architecture guidelines.

---

## 3. Go Coding Conventions

### Formatting
- Use **`gofmt`** / **`goimports`** for all `.go` source files.
- Indentation: **tabs** (standard Go convention, not spaces).
- Line length: Keep under 100 characters where practical.

### Naming
- Package names: **lowercase, singular** (e.g., `github`, `display`, `cmd`).
- Exported functions/types: **PascalCase** (e.g., `FetchRepositories`, `Repository`).
- Unexported functions/types: **camelCase** (e.g., `buildSearchQuery`, `formatNumber`).
- Constants: **PascalCase** or **ALL_CAPS** following Go conventions (e.g., `MaxLimit = 100`).

### Error Handling
- **Always handle errors explicitly** — do not use blank identifiers (`_`) to ignore errors without justification.
- User-facing error messages must be clear and actionable.
- Output error messages to `stderr` (`fmt.Fprintln(os.Stderr, ...)`), not `stdout`.
- Exit with status code `1` (`os.Exit(1)`) on fatal errors.

### Architecture and Separation of Concerns
- Maintain clean boundaries:
  - **`cmd/`**: CLI setup, flags definition, and input validation only. No business logic.
  - **`internal/`**: Core business logic, API communication, and formatting.
  - **`main.go`**: Minimal entry point that delegates directly to `cmd.Execute()`.

---

## 4. Recommended Workflow

Follow this sequence when implementing tasks:

1. **Review requirements** — Confirm the task scope.
2. **Create a feature branch** with a descriptive name (e.g., `feat/add-language-flag`).
3. **Develop with tests** (TDD where applicable): write tests first, then implementation.
4. **Verify quality**: Run `go vet ./...` and `go test ./...` — ensure zero warnings or failures.
5. **Commit using Conventional Commits**:
   - `feat(cmd): add language filter flag`
   - `fix(github): handle rate limit 403 response`
   - `docs(readme): update installation instructions`
6. **Update CHANGELOG.md** for new features or bug fixes.

---

## 5. Protected Files

| File | Rationale |
|---|---|
| `AGENTS.md` | Core agent governance file |
| `LICENSE` | Legal documentation |
| `go.mod` | Only modified when authorized dependencies change |

---

## 6. Quick References

- GitHub Search API Documentation: https://docs.github.com/en/rest/search/search#search-repositories
- Cobra CLI Library: https://cobra.dev/
- Go Standard Library: https://pkg.go.dev/std
- Conventional Commits: https://www.conventionalcommits.org/
- Effective Go: https://go.dev/doc/effective_go
