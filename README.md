# jankytext

Clean copied terminal text without opening a website.

## Use It

1. Copy messy text from a terminal, coding agent, log, or transcript.
2. Run:

```sh
jankytext
```

3. Paste. Your clipboard now contains the cleaned text.

That is the main workflow.

## Install

With Go installed:

```sh
go install github.com/HuotChu/jankytext/cmd/jankytext@latest
```

Make sure your Go binary directory is on `PATH`. By default this is usually `~/go/bin`.

Check that it works:

```sh
jankytext --version
```

## Use It From An Agent

jankytext works with Codex, Claude Code, Devin, VS Code agent plugins, and other terminal-based tools because it is just a command.

Copy the messy text, then ask the agent:

```text
Run jankytext to clean my clipboard.
```

If the agent needs the exact command, it is:

```sh
jankytext
```

## What It Fixes

jankytext cleans common copy/paste problems from terminal and agent output:

- color and control characters
- trailing whitespace
- progress lines that redraw in place
- prose that was wrapped by a narrow terminal
- common copied shell prompts and quote markers

It tries to preserve code, command output, JSON, YAML, diffs, logs, and test output.

## Less Common Uses

Preview the cleaned clipboard without changing it:

```sh
jankytext clip --preview
```

Clean a file or piped text instead of the clipboard:

```sh
jankytext < copied.txt
cat copied.txt | jankytext
```

Use a different cleanup level:

```sh
jankytext --mode conservative < copied.txt
jankytext --mode aggressive < copied.txt
```

Most users should not need these commands.

## Development

```sh
make test
make build
```

Cleanup behavior is tested with unit cases and before/after fixtures in `internal/cleaner/testdata`.
