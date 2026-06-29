---
name: jankytext
description: Clean messy terminal, command, log, stack trace, and agent transcript text. Use when the user says `jankytext`, `/jankytext`, `jankytext that`, asks to clean the last output, or wants copied terminal/agent text cleaned without re-pasting it.
---

# jankytext

When the user says `jankytext`, clean the most recent relevant messy text already present in the conversation. Prefer terminal output, command output, logs, stack traces, copied agent transcripts, or wrapped prose from recent assistant messages.

Do not ask the user to paste the same text again unless there is no relevant messy text in the conversation.

Return only the cleaned text unless the user asks for explanation, comparison, or diagnostics.

Also treat these as equivalent requests:

- `jankytext that`
- `jankytext this`
- `jankytext the last output`
- `jankytext the error above`
- `/jankytext`

## Cleanup Rules

- Remove ANSI color/control sequences and non-useful control characters.
- Resolve terminal redraw artifacts such as carriage-return progress updates.
- Trim trailing whitespace.
- Reflow prose wrapped by a narrow terminal.
- Strip common shell prompts and quote markers when doing so does not damage the content.
- Preserve code, command output, JSON, YAML, diffs, logs, stack traces, test output, paragraph boundaries, and meaning.

## Output Rules

- Return clean text only.
- Do not add introductions, summaries, or commentary.
- Do not use Markdown fences unless the cleaned content itself needs fences to remain clear.
- If there is nothing relevant to clean, say: `No recent janky text found.`
