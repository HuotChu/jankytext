# AGENTS.md — jankytext

## Project Overview
jankytext is an agent-agnostic Go CLI for cleaning text copied from terminals, coding agents, command output, and transcript panes. The core product is browser-free and deterministic, with future editor, launcher, and agent integrations expected to wrap the same executable rather than reimplement cleanup logic.

## Tech Stack
> **Note to AI agents:** Check `go.mod` for the current Go directive and module path.
- **Language:** Go.
- **Package style:** Standard-library-first CLI with internal packages for cleaner logic and clipboard integration.
- **Build entry point:** `cmd/jankytext`.

## Project Structure
> **Note to AI agents:** List the project root to discover current layout.
- **cmd/jankytext/** — CLI argument parsing, command dispatch, and user-facing command behavior.
- **internal/cleaner/** — Deterministic text cleanup logic and tests.
- **internal/cleaner/testdata/** — Before/after fixture cases for cleanup behavior.
- **internal/clipboard/** — Platform clipboard adapters used by `jankytext clip`.

## Code Conventions
- Keep cleanup behavior deterministic and fixture-backed.
- Prefer conservative heuristics that preserve code, command output, JSON, YAML, diffs, logs, and test output.
- Keep agent-specific behavior out of core cleanup logic unless it generalizes to copied terminal or transcript text.
- Keep integrations thin; they should call the CLI or shared cleaner code instead of duplicating cleanup rules.
- Use the standard library unless a dependency clearly improves cross-platform reliability.

## Formatting & Linting
- Go source is formatted with `gofmt`.
- Project commands are exposed through `Makefile`.
- Generated binaries belong in `bin/`, which is ignored by git.

## Testing
- Cleanup changes should include focused unit tests or before/after fixtures.
- When a real copied terminal sample cleans badly, add an `.in` and `.want` fixture under `internal/cleaner/testdata`.
- `make test` is the canonical test command.
- `make build` is the canonical local build command.

## Commit Conventions
No project-specific commit message format is currently documented. Use concise, imperative commit subjects that describe the user-visible or maintainer-visible change.
