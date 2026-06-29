# jankytext

`jankytext` cleans text copied from terminals, coding agents, command output, and transcript panes so it can be pasted into prompts, docs, issues, and chat without visual junk or broken wrapping.

It is agent-agnostic by design. The core tool is a CLI; editor, launcher, and agent integrations can wrap the same executable later.

## Product stance

- Cross-agent: useful with Codex, Claude Code, Devin, VS Code agent plugins, terminals, and issue trackers.
- Browser-free: cleanup should work without opening a website or re-copying from a web page.
- Deterministic first: cleanup behavior should be covered by fixtures before it grows more clever.
- Integration-friendly: v1 is a CLI; later wrappers should call the same executable instead of reimplementing cleanup logic.

## Usage

```sh
jankytext < copied.txt
jankytext --mode conservative < copied.txt
jankytext --mode aggressive < copied.txt
jankytext clip
jankytext clip --preview
```

## Install

With Go installed:

```sh
go install github.com/HuotChu/jankytext/cmd/jankytext@latest
```

Make sure your Go binary directory is on `PATH`. By default this is usually `~/go/bin`.

After installing:

```sh
jankytext --version
jankytext clip --preview
```

For a local checkout:

```sh
make build
./bin/jankytext --help
```

## Modes

- `conservative`: normalize newlines, strip ANSI/control characters, resolve carriage-return redraws, and trim trailing whitespace.
- `standard`: conservative cleanup plus cautious prose unwrapping.
- `aggressive`: standard cleanup plus common quote-marker and shell-prompt stripping.

## Development

```sh
make test
make build
make install
go run ./cmd/jankytext --help
```

## Fixture tests

Cleanup behavior lives in `internal/cleaner` and is tested with plain unit cases plus before/after fixtures in `internal/cleaner/testdata`.

When a copied terminal sample cleans badly, add a focused `.in` and `.want` fixture before changing the heuristic.
