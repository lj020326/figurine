// Copyright 2026 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.
// License can be found in the LICENSE file.

package main

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

const testFontName = "Poison.flf"

func TestParseArgsCombinesShortBooleanFlags(t *testing.T) {
	opts, input, err := parseArgs([]string{"-ls", "Sample", "Text"})
	if err != nil {
		t.Fatalf("parse args: %v", err)
	}
	if !opts.listFonts {
		t.Fatal("list flag was not set")
	}
	if !opts.showSample {
		t.Fatal("sample flag was not set")
	}
	if !slices.Equal(input, []string{"Sample", "Text"}) {
		t.Fatalf("input = %q, want Sample Text", input)
	}
}

func TestParseArgsReadsLongFontValue(t *testing.T) {
	opts, input, err := parseArgs([]string{"--font", testFontName, "hello"})
	if err != nil {
		t.Fatalf("parse args: %v", err)
	}
	if opts.fontName != testFontName {
		t.Fatalf("font name = %q, want %s", opts.fontName, testFontName)
	}
	if !slices.Equal(input, []string{"hello"}) {
		t.Fatalf("input = %q, want hello", input)
	}
}

func TestParseArgsStopsAtSeparator(t *testing.T) {
	_, input, err := parseArgs([]string{"--", "--not-a-flag"})
	if err != nil {
		t.Fatalf("parse args: %v", err)
	}
	if !slices.Equal(input, []string{"--not-a-flag"}) {
		t.Fatalf("input = %q, want --not-a-flag", input)
	}
}

func TestParseArgsReportsUnknownFlag(t *testing.T) {
	_, _, err := parseArgs([]string{"--bogus"})
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "unknown flag --bogus") {
		t.Fatalf("error = %q, want unknown flag", err)
	}
}

func TestRunHelpWritesUsefulUsage(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"--help"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
	output := stdout.String()
	for _, want := range []string{"Usage:", "Examples:", "-ls Sample Text", "--font name"} {
		if !strings.Contains(output, want) {
			t.Fatalf("help output does not contain %q", want)
		}
	}
}
