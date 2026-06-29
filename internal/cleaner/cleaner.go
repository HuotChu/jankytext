package cleaner

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Mode string

const (
	ModeConservative Mode = "conservative"
	ModeStandard     Mode = "standard"
	ModeAggressive   Mode = "aggressive"
)

type Options struct {
	Mode         Mode
	StripPrompts bool
}

type lineState struct {
	text      string
	protected bool
}

var (
	csiPattern       = regexp.MustCompile(`\x1b\[[0-?]*[ -/]*[@-~]`)
	oscPattern       = regexp.MustCompile(`\x1b\][^\a]*(?:\a|\x1b\\)`)
	remainingEscapes = regexp.MustCompile(`\x1b.`)
	promptPattern    = regexp.MustCompile(`^\s*(?:\[[^\]]+\]\s*)?(?:(?:[A-Za-z0-9_.-]+@[A-Za-z0-9_.-]+:[^$#%>]{0,80})|(?:[~/A-Za-z0-9_.-][^$#%>]{0,80}))[$#%>]\s+`)
)

func Clean(input string, options Options) string {
	mode := options.Mode
	if mode == "" {
		mode = ModeStandard
	}

	text := normalizeCRLF(input)
	text = stripANSI(text)
	text = resolveBackspaces(text)
	text = resolveCarriageReturns(text)
	text = normalizeNewlines(text)
	text = trimLineEnds(text)

	if mode == ModeConservative {
		return finalNewline(text, input)
	}

	lines := splitLineStates(text)
	if mode == ModeAggressive || options.StripPrompts {
		lines = stripPromptsAndQuotes(lines)
	}
	text = unwrapProse(lines)

	return finalNewline(text, input)
}

func normalizeCRLF(input string) string {
	return strings.ReplaceAll(input, "\r\n", "\n")
}

func normalizeNewlines(input string) string {
	return strings.ReplaceAll(input, "\r", "\n")
}

func stripANSI(input string) string {
	text := oscPattern.ReplaceAllString(input, "")
	text = csiPattern.ReplaceAllString(text, "")
	return remainingEscapes.ReplaceAllString(text, "")
}

func resolveBackspaces(input string) string {
	var out []rune
	for _, r := range input {
		if r == '\b' {
			if len(out) > 0 {
				out = out[:len(out)-1]
			}
			continue
		}
		if isAllowedControl(r) {
			out = append(out, r)
			continue
		}
		if unicode.IsControl(r) {
			continue
		}
		out = append(out, r)
	}
	return string(out)
}

func resolveCarriageReturns(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if idx := strings.LastIndexByte(line, '\r'); idx >= 0 {
			lines[i] = line[idx+1:]
		}
	}
	return strings.Join(lines, "\n")
}

func trimLineEnds(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	return strings.Join(lines, "\n")
}

func splitLineStates(text string) []lineState {
	raw := strings.Split(text, "\n")
	lines := make([]lineState, len(raw))
	for i, line := range raw {
		lines[i] = lineState{text: line}
	}
	return lines
}

func stripPromptsAndQuotes(lines []lineState) []lineState {
	out := make([]lineState, len(lines))
	for i, line := range lines {
		text := stripQuoteMarker(line.text)
		strippedPrompt := promptPattern.MatchString(text)
		text = promptPattern.ReplaceAllString(text, "")
		out[i] = lineState{
			text:      text,
			protected: line.protected || strippedPrompt,
		}
	}
	return out
}

func stripQuoteMarker(line string) string {
	trimmed := strings.TrimLeft(line, " \t")
	if !strings.HasPrefix(trimmed, ">") {
		return line
	}
	for strings.HasPrefix(trimmed, "> ") {
		trimmed = strings.TrimPrefix(trimmed, "> ")
	}
	if trimmed == ">" {
		return ""
	}
	return trimmed
}

func unwrapProse(lines []lineState) string {
	var out []string
	var paragraph []string
	inFence := false

	flush := func() {
		if len(paragraph) == 0 {
			return
		}
		out = append(out, joinProse(paragraph))
		paragraph = nil
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line.text)
		if isFence(trimmed) {
			flush()
			inFence = !inFence
			out = append(out, line.text)
			continue
		}
		if inFence || trimmed == "" || line.protected || isProtectedLine(line.text) {
			flush()
			out = append(out, line.text)
			continue
		}
		paragraph = append(paragraph, trimmed)
	}
	flush()

	return strings.Join(out, "\n")
}

func joinProse(lines []string) string {
	if len(lines) == 1 {
		return lines[0]
	}
	var b strings.Builder
	for i, line := range lines {
		if i == 0 {
			b.WriteString(line)
			continue
		}
		prev := b.String()
		if strings.HasSuffix(prev, "-") && len(prev) > 1 {
			b.WriteString(line)
		} else {
			b.WriteByte(' ')
			b.WriteString(line)
		}
	}
	return b.String()
}

func isProtectedLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return true
	}
	if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
		return true
	}
	prefixes := []string{
		"- ", "* ", "+ ", "1. ", "2. ", "3. ", "4. ", "5. ",
		"#", "|", ">", "```", "~~~", "diff ", "commit ", "@@",
		"+++", "---", "INFO", "WARN", "ERROR", "DEBUG", "TRACE",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}
	if looksLikeJSONOrYAML(trimmed) || looksLikePrompt(trimmed) || startsWithTimestamp(trimmed) || looksLikeGoTestLine(trimmed) {
		return true
	}
	return false
}

func looksLikeJSONOrYAML(line string) bool {
	if strings.HasPrefix(line, "{") || strings.HasPrefix(line, "}") || strings.HasPrefix(line, "[") || strings.HasPrefix(line, "]") {
		return true
	}
	if strings.Contains(line, ": ") {
		before, _, found := strings.Cut(line, ": ")
		if found && before != "" && !strings.Contains(before, " ") && utf8.RuneCountInString(before) <= 40 {
			return true
		}
	}
	return false
}

func looksLikePrompt(line string) bool {
	return promptPattern.MatchString(line)
}

func looksLikeGoTestLine(line string) bool {
	return strings.HasPrefix(line, "?") ||
		strings.HasPrefix(line, "ok ") ||
		strings.HasPrefix(line, "FAIL") ||
		strings.HasPrefix(line, "--- FAIL:")
}

func startsWithTimestamp(line string) bool {
	if len(line) < 10 {
		return false
	}
	for i, r := range line[:10] {
		if i == 4 || i == 7 {
			if r != '-' {
				return false
			}
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func isFence(line string) bool {
	return strings.HasPrefix(line, "```") || strings.HasPrefix(line, "~~~")
}

func isAllowedControl(r rune) bool {
	return r == '\n' || r == '\t' || r == '\r'
}

func finalNewline(output string, original string) string {
	if original == "" {
		return ""
	}
	if strings.HasSuffix(original, "\n") || strings.HasSuffix(original, "\r") {
		return output
	}
	return strings.TrimRight(output, "\n")
}
