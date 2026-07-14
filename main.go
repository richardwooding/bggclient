package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

// Set by goreleaser via ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// newParser builds the CLI grammar with ctx bound so every command's
// Run(ctx, ...) receives it.
func newParser(ctx context.Context) (*kong.Kong, error) {
	return kong.New(&cli,
		kong.Name("bggclient"),
		kong.Description("BoardGameGeek XML API client — CLI and MCP server."),
		kong.UsageOnError(),
		kong.Vars{"version": fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)},
		kong.BindTo(ctx, (*context.Context)(nil)),
	)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		// After the first signal cancels ctx, restore default signal
		// handling so a second Ctrl-C terminates immediately.
		<-ctx.Done()
		stop()
	}()

	parser, err := newParser(ctx)
	if err != nil {
		panic(err)
	}
	kctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)
	kctx.FatalIfErrorf(kctx.Run(&cli.Globals))
}
