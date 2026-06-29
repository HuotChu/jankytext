# AGENTS.md — jankytext

## Project Overview
jankytext is an agent-agnostic text cleanup product for terminal output, command output, logs, stack traces, and agent transcripts. In agent chats, `jankytext` means cleaning the most recent relevant messy text already present in conversation. On desktop, the Go CLI cleans the system clipboard in place.

## Tech Stack
> **Note to AI agents:** Check `go.mod` for the current Go directive and module path.
- **Language:** Go.
- **Package style:** Standard-library-first CLI with internal packages for cleaner logic and clipboard integration.
- **Build entry point:** `cmd/jankytext`.

## Project Structure
> **Note to AI agents:** List the project root to discover current layout.
- **cmd/jankytext/** — CLI argument parsing, command dispatch, and user-facing command behavior.
- **docs/** — Shared product behavior specs for agent and CLI surfaces.
- **internal/cleaner/** — Deterministic text cleanup logic and tests.
- **internal/cleaner/testdata/** — Before/after fixture cases for cleanup behavior.
- **internal/clipboard/** — Platform clipboard adapters used by `jankytext clip`.
- **prompts/** — Copyable instructions for generic agents, ChatGPT, and Claude.
- **skills/jankytext/** — Codex-style skill package for agent-native jankytext behavior.

## Code Conventions
- Keep cleanup behavior deterministic and fixture-backed.
- Prefer conservative heuristics that preserve code, command output, JSON, YAML, diffs, logs, and test output.
- Keep agent-native behavior aligned with `docs/agent-behavior.md`.
- Keep desktop CLI workflows clipboard-in-place; do not require users to copy cleaned text back out of terminal output.
- Keep integrations thin; they should follow the shared behavior contract or call the CLI/shared cleaner code instead of duplicating cleanup rules.
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
