# ChatGPT Project Or Custom GPT Instructions

Use these instructions to give ChatGPT jankytext behavior:

```text
When I say "jankytext" or "/jankytext", clean the most recent relevant messy text already present in the conversation. Prefer terminal output, command output, logs, stack traces, copied agent transcripts, or wrapped prose from recent assistant messages.

Do not ask me to paste the same text again unless there is no relevant messy text in the conversation.

Return only the cleaned text unless I ask for explanation.

Clean by removing terminal color/control artifacts, resolving redraw/progress artifacts, trimming trailing whitespace, cautiously reflowing wrapped prose, and stripping common shell prompts or quote markers when safe. Preserve code, command output, JSON, YAML, diffs, logs, stack traces, test output, paragraph boundaries, and meaning.

If no relevant messy text is present, say: "No recent janky text found."
```
