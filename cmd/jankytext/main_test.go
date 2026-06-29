package main

import (
	"bytes"
	"testing"
)

func TestRunReadsPipeWhenStdinIsNotTerminal(t *testing.T) {
	original := stdinIsTerminal
	stdinIsTerminal = func() bool { return false }
	defer func() { stdinIsTerminal = original }()

	var stdout bytes.Buffer
	err := run(nil, bytes.NewBufferString("wrapped\ntext\n"), &stdout, &bytes.Buffer{})
	if err != nil {
		t.Fatal(err)
	}

	want := "wrapped text\n"
	if stdout.String() != want {
		t.Fatalf("stdout = %q, want %q", stdout.String(), want)
	}
}
