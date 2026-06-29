package clipboard

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func Read() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		return output("pbpaste")
	case "windows":
		return output("powershell", "-NoProfile", "-Command", "Get-Clipboard")
	case "linux":
		if text, err := output("wl-paste", "--no-newline"); err == nil {
			return text, nil
		}
		return output("xclip", "-selection", "clipboard", "-out")
	default:
		return "", fmt.Errorf("clipboard read is not supported on %s", runtime.GOOS)
	}
}

func Write(text string) error {
	switch runtime.GOOS {
	case "darwin":
		return input(text, "pbcopy")
	case "windows":
		return input(text, "powershell", "-NoProfile", "-Command", "Set-Clipboard")
	case "linux":
		if err := input(text, "wl-copy"); err == nil {
			return nil
		}
		return input(text, "xclip", "-selection", "clipboard", "-in")
	default:
		return fmt.Errorf("clipboard write is not supported on %s", runtime.GOOS)
	}
}

func output(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", commandError(name, err, stderr.String())
	}
	return string(out), nil
}

func input(text string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = bytes.NewBufferString(text)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return commandError(name, err, stderr.String())
	}
	return nil
}

func commandError(name string, err error, stderr string) error {
	if errors.Is(err, exec.ErrNotFound) {
		return fmt.Errorf("%s is not installed or not on PATH", name)
	}
	if stderr != "" {
		return fmt.Errorf("%s failed: %w: %s", name, err, stderr)
	}
	return fmt.Errorf("%s failed: %w", name, err)
}
