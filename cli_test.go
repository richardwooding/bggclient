package main

import (
	"testing"

	"github.com/alecthomas/kong"
)

// The kong grammar is validated at build time; this guards against
// invalid struct tags in the CLI definition.
func TestCLIGrammar(t *testing.T) {
	parser, err := kong.New(&cli)
	if err != nil {
		t.Fatalf("invalid CLI grammar: %v", err)
	}
	for _, args := range [][]string{
		{"search", "Catan", "--exact"},
		{"boardgame", "13", "42", "--stats", "--comments"},
		{"collection", "someuser", "--own", "--min-rating=5"},
		{"geeklist", "11205", "--comments"},
		{"serve", "--http=:8080"},
	} {
		if _, err := parser.Parse(args); err != nil {
			t.Errorf("parsing %v: %v", args, err)
		}
	}
}
