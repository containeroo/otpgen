package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestRunRootGeneratesToken(t *testing.T) {
	originalNow := timeNow
	originalGenerate := generateCode
	t.Cleanup(func() {
		timeNow = originalNow
		generateCode = originalGenerate
	})

	fixedTime := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	timeNow = func() time.Time {
		return fixedTime
	}

	var gotSecret string
	var gotTime time.Time
	generateCode = func(secret string, now time.Time) (string, error) {
		gotSecret = secret
		gotTime = now
		return "123456", nil
	}

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"  supersecret  "})

	if err := root.Execute(); err != nil {
		t.Fatalf("expected command to succeed, got error: %v", err)
	}

	if gotSecret != "supersecret" {
		t.Fatalf("expected trimmed secret %q, got %q", "supersecret", gotSecret)
	}

	if !gotTime.Equal(fixedTime) {
		t.Fatalf("expected generation time %v, got %v", fixedTime, gotTime)
	}

	if got := out.String(); got != "123456\n" {
		t.Fatalf("expected token output with trailing newline, got %q", got)
	}
}

func TestRunRootRejectsEmptySecret(t *testing.T) {
	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"   "})

	err := root.Execute()
	if err == nil {
		t.Fatal("expected empty secret validation error")
	}

	if !strings.Contains(err.Error(), "secret must not be empty") {
		t.Fatalf("expected empty secret error, got %v", err)
	}
}

func TestRunRootWrapsGenerationErrors(t *testing.T) {
	originalGenerate := generateCode
	generateCode = func(secret string, now time.Time) (string, error) {
		return "", errors.New("invalid secret")
	}
	t.Cleanup(func() {
		generateCode = originalGenerate
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"secret"})

	err := root.Execute()
	if err == nil {
		t.Fatal("expected generation error")
	}

	if !strings.Contains(err.Error(), "generate TOTP: invalid secret") {
		t.Fatalf("expected wrapped generation error, got %v", err)
	}
}
