package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

type failingWriter struct{}

func (f failingWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("write failed")
}

func TestCompletionBashOutputsScript(t *testing.T) {
	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"completion", "bash"})

	if err := root.Execute(); err != nil {
		t.Fatalf("expected completion command to succeed, got error: %v", err)
	}

	if !strings.Contains(out.String(), "otpgen") {
		t.Fatalf("expected completion output to contain command name, got %q", out.String())
	}
}

func TestCompletionReturnsWriterErrors(t *testing.T) {
	root := newRootCmd()
	root.SetOut(failingWriter{})
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"completion", "zsh"})

	err := root.Execute()
	if err == nil {
		t.Fatal("expected completion command to return writer error")
	}

	if !strings.Contains(err.Error(), "write failed") {
		t.Fatalf("expected write error, got %v", err)
	}
}
