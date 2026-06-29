# jankytext

Clean copied terminal text without opening a website.

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

## Use From An Agent

jankytext works with Codex, Claude Code, Devin, VS Code agent plugins, and other terminal-based tools because it is just a command.

1. Copy messy text from a terminal, agent, log, or transcript.
2. Ask your agent:

```text
run jankytext
```

or just:

```text
jankytext
```

3. Paste. Your clipboard now contains the cleaned text.

If you later add a jankytext skill or slash command in an agent, use whatever shortcut that agent exposes, such as `/jankytext`.

## Use Directly

From any terminal:

```sh
jankytext
```

That command reads your clipboard, cleans the text, and puts the cleaned text back on your clipboard.

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
