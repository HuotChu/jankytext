# jankytext

Clean messy terminal and agent text without copying it twice.

## What It Fixes

jankytext cleans common copy/paste problems from terminal and agent output:

- color and control characters
- trailing whitespace
- progress lines that redraw in place
- prose that was wrapped by a narrow terminal
- common copied shell prompts and quote markers

It tries to preserve code, command output, JSON, YAML, diffs, logs, stack traces, and test output.

## Use In An Agent Chat

This is the main mobile-friendly workflow.

No CLI install is needed for this mode. Configure the agent with one of the prompt or skill files below, then use `jankytext` inside the conversation.

When an agent has already produced messy terminal output, reply:

```text
jankytext
```

The agent should clean the most recent relevant output already in the conversation and return only the cleaned text.

Useful variants:

```text
jankytext that
jankytext the last output
jankytext the error above
/jankytext
```

You should not need to copy messy text out of the chat, paste it back into the chat, and copy it again.

Reusable agent instructions are in:

- `prompts/generic-agent.md`
- `prompts/chatgpt.md`
- `prompts/claude.md`
- `skills/jankytext/SKILL.md`

The shared behavior contract is in `docs/agent-behavior.md`.

## Use On Desktop

The desktop CLI cleans your clipboard in place.

1. Copy messy text from a terminal, agent, log, or transcript.
2. Run this from any terminal:

```sh
jankytext
```

3. Paste. Your clipboard now contains the cleaned text.

The CLI should not require copying cleaned output back out of the terminal.

## Install The CLI

Download the latest release for your system:

[github.com/HuotChu/jankytext/releases/latest](https://github.com/HuotChu/jankytext/releases/latest)

Available builds:

- Windows: `jankytext_..._windows_amd64.zip`
- macOS Apple Silicon: `jankytext_..._darwin_arm64.tar.gz`
- macOS Intel: `jankytext_..._darwin_amd64.tar.gz`
- Linux x64: `jankytext_..._linux_amd64.tar.gz`
- Linux ARM64: `jankytext_..._linux_arm64.tar.gz`

Unzip or untar the download, then put `jankytext` somewhere your terminal or agent can run it. On Windows, the program is `jankytext.exe`.

Check that it works:

```sh
jankytext --version
```

If you already have Go installed, this also works:

```sh
go install github.com/HuotChu/jankytext/cmd/jankytext@latest
```

## Less Common CLI Uses

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
make dist
```

Cleanup behavior is tested with unit cases and before/after fixtures in `internal/cleaner/testdata`.
