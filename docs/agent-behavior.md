# jankytext Agent Behavior

jankytext has two primary meanings, depending on where it is used.

## In An Agent Chat

When the user says `jankytext`, clean the most recent relevant messy text that is already present in the conversation. Prefer terminal output, command output, logs, stack traces, copied agent transcripts, or wrapped prose from recent assistant messages.

Do not ask the user to paste the same text again unless there is no relevant messy text in the conversation.

Return only the cleaned text unless the user asks for explanation, comparison, or diagnostics.

Also treat these as equivalent requests:

- `jankytext that`
- `jankytext this`
- `jankytext the last output`
- `jankytext the error above`
- `/jankytext`

If more than one recent block could be intended, clean the most recent likely block. Ask a short clarification only when choosing the wrong block would be materially confusing.

## In A Terminal

When run as a CLI command, `jankytext` cleans the system clipboard in place. The user should not need to copy text back out of terminal output.

Piped input and file input are supported for scripting and automation, but they are not the primary human workflow.

## Cleanup Rules

- Remove ANSI color/control sequences and non-useful control characters.
- Resolve terminal redraw artifacts such as carriage-return progress updates.
- Trim trailing whitespace.
- Reflow prose wrapped by a narrow terminal.
- Strip common shell prompts and quote markers when doing so does not damage the content.
- Preserve code, command output, JSON, YAML, diffs, logs, stack traces, and test output.
- Preserve paragraph boundaries and intentional blank lines.

## Output Rules

- Return clean text only.
- Do not wrap the answer in commentary like "Here is the cleaned text."
- Do not use Markdown fences unless the cleaned content itself needs fences to remain clear.
- Do not summarize, explain, or modify meaning.
- If there is nothing relevant to clean, say: `No recent janky text found.`
