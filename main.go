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

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	kctx := kong.Parse(&cli,
		kong.Name("bggclient"),
		kong.Description("BoardGameGeek XML API client — CLI and MCP server."),
		kong.UsageOnError(),
		kong.Vars{"version": fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)},
		kong.BindTo(ctx, (*context.Context)(nil)),
	)
	kctx.FatalIfErrorf(kctx.Run(&cli.Globals))
}
