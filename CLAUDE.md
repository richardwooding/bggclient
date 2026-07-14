# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

BoardGameGeek XML API 1 client, shipped three ways: a kong-based CLI (`bggclient`), an MCP server (`bggclient serve`, official MCP Go SDK), and a Go library (`xml1` package). Module: `github.com/richardwooding/bggclient`, Go 1.26.

BGG requires registered applications to send a Bearer token (`Options.APIToken` / `--token` / `BGG_API_TOKEN`); unauthenticated live calls return 401. Tests never hit the network (govcr cassette), so no token is needed for development.

## Commands

- Build: `go build ./...`
- Test (all): `go test ./...`
- Run a single feature scenario: `go test ./xml1 -run TestFeatures/<scenario name with underscores>` (godog exposes scenarios as subtests)
- Modernize: `go fix ./...` — CI fails if `go fix -diff ./...` produces output
- Lint/format: `pre-commit run --all-files` (go-fmt, golangci-lint, go-mod-tidy, actionlint, etc.)
- Release dry-run: `goreleaser release --snapshot --clean --skip=ko` (ko needs a registry or Docker daemon)
- Run the CLI: `go run . search "Catan"` (needs `BGG_API_TOKEN` for live calls)
- Commits follow Conventional Commits (commitlint hook). CI runs on `main` and `develop`; `develop` is the default PR target.

## Architecture

- Root package (`main.go`, `cli.go`, `commands.go`, `serve.go`) — kong CLI. `Globals` holds shared flags (base URL, token, request interval, timeout, `--no-color`); commands map flags to `xml1` calls and print colourised JSON via a lipgloss token-stream colorizer in `cli.go` (colour auto-disables for non-TTY/`NO_COLOR`). `version`/`commit`/`date` vars in `main.go` are injected by goreleaser ldflags.
- `internal/bggopts/` — flag/input structs shared by CLI and MCP, tri-tagged (kong `help:` / `json:` / `jsonschema:`), with `Options()` converters to `xml1` functional options. Note naming traps in xml1: `Comments` (boardgame) vs `Comment` (collection) vs `GeeklistComments` (presence-only var).
- `internal/mcpserver/` — `New(api, version)` returns an `*mcp.Server` with four tools (`bgg_search`, `bgg_get_boardgames`, `bgg_get_collection`, `bgg_get_geeklist`). Handlers use the generic `mcp.AddTool` typed pattern; output structs are the `xml1/model` types directly (they carry json tags for this reason). Transports: stdio (default) and streamable HTTP (`serve --http`).
- `xml1/` — the API client. `NewAPI(Options{...})` wires a `rate.Limiter` (default one request per 5s). All endpoints funnel through `getInternal`: rate limit → Bearer token header → GET → retries (max 5; 429 always retryable, 202 also retryable for collection/geeklist since BGG builds those asynchronously). 404 → `customerrors.NotFoundError`.
- `xml1/model/` — response structs, dual xml+json tagged (`XMLName` fields are `json:"-"`). `model.Decode(io.Reader)` dispatches on the XML root element; an `errors` root becomes a typed Go error via `customerrors.New`.

## Testing

- `xml1` is tested with godog BDD: scenarios in `xml1/features/*.feature`, step definitions in `xml1/api_test.go`.
- `internal/mcpserver` is tested end-to-end via `mcp.NewInMemoryTransports()` + an MCP client, asserting on `StructuredContent`.
- All HTTP is replayed from the govcr cassette `xml1/fixtures/bgg.json` (~1.9MB recorded request/response file — not a schema; also loaded cross-package as `../../xml1/fixtures/bgg.json`). New endpoints or parameter combinations require recording new cassette interactions; the file is excluded from the `check-added-large-files` pre-commit hook. Tests pass `RequestInterval: time.Millisecond` to bypass rate-limit waits.

## Release

Tag `v*` on main → `.github/workflows/release.yml` runs goreleaser: GitHub release with archives, Homebrew cask pushed to `richardwooding/homebrew-tap` (needs `HOMEBREW_TAP_GITHUB_TOKEN` secret), and a ko-built OCI image at `ghcr.io/richardwooding/bggclient`. The gloam site in `docs/` deploys to GitHub Pages via `pages.yml`; gloam assets are vendored and refreshed by `gloam-sync.yml`.
