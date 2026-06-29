# jankytext

Clean copied terminal text without opening a website.

## What It Fixes

jankytext cleans common copy/paste problems from terminal and agent output:

- color and control characters
- trailing whitespace
- progress lines that redraw in place
- prose that was wrapped by a narrow terminal
- common copied shell prompts and quote markers

It tries to preserve code, command output, JSON, YAML, diffs, logs, and test output.

## Install

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

Phones are not covered well by the CLI version. For Android and iOS, a small web or mobile version is the more accessible future path.

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
make dist
```

Cleanup behavior is tested with unit cases and before/after fixtures in `internal/cleaner/testdata`.
