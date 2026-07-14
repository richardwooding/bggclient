package main

import (
	"context"
	"errors"
	"testing"
)

// A cancelled ctx bound at parser construction must reach the command's
// Run method and short-circuit before any network I/O.
func TestRunHonoursCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	parser, err := newParser(ctx)
	if err != nil {
		t.Fatalf("newParser: %v", err)
	}
	kctx, err := parser.Parse([]string{"search", "anything"})
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	err = kctx.Run(&cli.Globals)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Run err = %v, want context.Canceled", err)
	}
}
