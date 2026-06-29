package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/HuotChu/jankytext/internal/cleaner"
	"github.com/HuotChu/jankytext/internal/clipboard"
)

const version = "0.1.0"

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "jankytext:", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	if len(args) > 0 && args[0] == "clip" {
		return runClip(args[1:], stdout)
	}
	return runClean(args, stdin, stdout)
}

func runClean(args []string, stdin io.Reader, stdout io.Writer) error {
	flags := flag.NewFlagSet("jankytext", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	mode := flags.String("mode", string(cleaner.ModeStandard), "cleanup mode: conservative, standard, aggressive")
	stripPrompts := flags.Bool("strip-prompts", false, "strip common shell prompts")
	showVersion := flags.Bool("version", false, "print version")
	help := flags.Bool("help", false, "print help")

	if err := flags.Parse(args); err != nil {
		return usageError(err)
	}
	if *help {
		printUsage(stdout)
		return nil
	}
	if *showVersion {
		fmt.Fprintln(stdout, version)
		return nil
	}

	options, err := cleanerOptions(*mode, *stripPrompts)
	if err != nil {
		return err
	}

	input, err := readInputs(stdin, flags.Args())
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, cleaner.Clean(input, options))
	return err
}

func runClip(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("jankytext clip", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	mode := flags.String("mode", string(cleaner.ModeStandard), "cleanup mode: conservative, standard, aggressive")
	stripPrompts := flags.Bool("strip-prompts", false, "strip common shell prompts")
	preview := flags.Bool("preview", false, "write cleaned text to stdout without changing the clipboard")
	help := flags.Bool("help", false, "print help")

	if err := flags.Parse(args); err != nil {
		return usageError(err)
	}
	if *help {
		printClipUsage(stdout)
		return nil
	}
	if flags.NArg() != 0 {
		return usageError(fmt.Errorf("clip does not accept file arguments"))
	}

	options, err := cleanerOptions(*mode, *stripPrompts)
	if err != nil {
		return err
	}
	input, err := clipboard.Read()
	if err != nil {
		return err
	}
	output := cleaner.Clean(input, options)
	if *preview {
		_, err = io.WriteString(stdout, output)
		return err
	}
	return clipboard.Write(output)
}

func cleanerOptions(modeName string, stripPrompts bool) (cleaner.Options, error) {
	mode := cleaner.Mode(modeName)
	switch mode {
	case cleaner.ModeConservative, cleaner.ModeStandard, cleaner.ModeAggressive:
		return cleaner.Options{Mode: mode, StripPrompts: stripPrompts}, nil
	default:
		return cleaner.Options{}, fmt.Errorf("unknown mode %q", modeName)
	}
}

func readInputs(stdin io.Reader, paths []string) (string, error) {
	if len(paths) == 0 {
		data, err := io.ReadAll(stdin)
		return string(data), err
	}

	var out []byte
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		out = append(out, data...)
		if len(data) > 0 && data[len(data)-1] != '\n' {
			out = append(out, '\n')
		}
	}
	return string(out), nil
}

func usageError(err error) error {
	return fmt.Errorf("%w\n\nrun `jankytext --help` for usage", err)
}

func printUsage(w io.Writer) {
	fmt.Fprint(w, `Usage:
  jankytext [flags] [file ...]
  jankytext clip [flags]

Flags:
  --mode conservative|standard|aggressive
  --strip-prompts
  --version
  --help
`)
}

func printClipUsage(w io.Writer) {
	fmt.Fprint(w, `Usage:
  jankytext clip [flags]

Flags:
  --mode conservative|standard|aggressive
  --strip-prompts
  --preview
  --help
`)
}
