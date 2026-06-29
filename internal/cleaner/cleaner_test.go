package cleaner

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestConservativeStripsANSIAndTrailingWhitespace(t *testing.T) {
	input := "\x1b[32mhello\x1b[0m   \nworld\t \n"
	got := Clean(input, Options{Mode: ModeConservative})
	want := "hello\nworld\n"
	if got != want {
		t.Fatalf("Clean() = %q, want %q", got, want)
	}
}

func TestConservativeResolvesCarriageReturnRedraw(t *testing.T) {
	input := "Downloading 10%\rDownloading 100%\nDone\n"
	got := Clean(input, Options{Mode: ModeConservative})
	want := "Downloading 100%\nDone\n"
	if got != want {
		t.Fatalf("Clean() = %q, want %q", got, want)
	}
}

func TestStandardUnwrapsProseButPreservesBulletsAndCode(t *testing.T) {
	input := "This is a copied paragraph that\nwrapped in the terminal.\n\n- keep\n- bullets\n\n```go\nfmt.Println(\"keep\")\n```\n"
	got := Clean(input, Options{Mode: ModeStandard})
	want := "This is a copied paragraph that wrapped in the terminal.\n\n- keep\n- bullets\n\n```go\nfmt.Println(\"keep\")\n```\n"
	if got != want {
		t.Fatalf("Clean() = %q, want %q", got, want)
	}
}

func TestAggressiveStripsPromptsAndQuoteMarkers(t *testing.T) {
	input := "> ~/src/project $ go test ./...\n> ok github.com/HuotChu/jankytext\n"
	got := Clean(input, Options{Mode: ModeAggressive})
	want := "go test ./...\nok github.com/HuotChu/jankytext\n"
	if got != want {
		t.Fatalf("Clean() = %q, want %q", got, want)
	}
}

func TestPreservesJSONLikeLines(t *testing.T) {
	input := "{\n  \"name\": \"jankytext\",\n  \"ok\": true\n}\n"
	got := Clean(input, Options{Mode: ModeStandard})
	if got != input {
		t.Fatalf("Clean() = %q, want %q", got, input)
	}
}

func TestFixtureCases(t *testing.T) {
	cases := []struct {
		name string
		mode Mode
	}{
		{name: "wrapped_prose", mode: ModeStandard},
		{name: "prompt_transcript", mode: ModeAggressive},
		{name: "markdown_mixed_content", mode: ModeStandard},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := readFixture(t, tc.name+".in")
			want := readFixture(t, tc.name+".want")
			got := Clean(input, Options{Mode: tc.mode})
			if got != want {
				t.Fatalf("Clean fixture %s =\n%s\nwant\n%s", tc.name, quoteLines(got), quoteLines(want))
			}
		})
	}
}

func readFixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func quoteLines(text string) string {
	lines := strings.SplitAfter(text, "\n")
	var b strings.Builder
	for _, line := range lines {
		if line == "" {
			continue
		}
		b.WriteString("  ")
		b.WriteString(strconv.Quote(line))
		b.WriteByte('\n')
	}
	return b.String()
}
