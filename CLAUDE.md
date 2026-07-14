# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Go client library for the BoardGameGeek XML API 1 (https://boardgamegeek.com/xmlapi). Library only — no main.go. Module: `github.com/richardwooding/bggclient`, Go 1.25.

## Commands

- Build: `go build ./...`
- Test (all): `go test ./...` — CI runs `go build -v ./...` and `go test -v ./...`
- Run a single feature scenario: `go test ./xml1 -run TestFeatures/<scenario name with underscores>` (godog exposes scenarios as subtests), or narrow via `godog.Options.Paths`/tags in `xml1/api_test.go`
- Lint/format: `pre-commit run --all-files` (go-fmt, golangci-lint, go-mod-tidy, actionlint, etc.)
- Commits must follow Conventional Commits (enforced by commitlint pre-commit hook). CI runs on `main` and `develop`; `develop` is the default PR target.

## Architecture

All code is in `xml1/` (one package per BGG API version; only XML API 1 is implemented):

- `xml1/api.go` — the client. `NewAPI(Options{HttpClient, BaseURL})` wires in a `rate.Limiter`. All endpoints funnel through `getInternal`, which applies rate limiting, builds the URL, and handles retries recursively (max 5, `MAX_ALLOWED_RETRIES`). HTTP 429 is retryable for all endpoints; collection and geeklist endpoints also retry on 202 (BGG returns 202 while building a response asynchronously). 404 becomes `customerrors.NotFoundError`.
- Endpoint options use the functional options pattern: each endpoint has its own option type (`SearchOption`, `BoardgameOption`, `CollectionOption`, `GeeklistOption` in `*options.go` files) mutating a query-param map.
- `xml1/model/` — XML response structs behind the sealed `XML1Model` interface. `model.Decode(io.Reader)` inspects the XML root element (`boardgames`, `items`, `geeklist`, `errors`) and dispatches to the right struct; an `errors` root is converted into a typed Go error.
- `xml1/customerrors/` — typed errors. `customerrors.New(message)` maps known BGG error message strings to specific error types.

## Testing

Tests are BDD (godog/Cucumber): `xml1/api_test.go` is the suite entry point (`TestFeatures`) with step definitions; scenarios live in `xml1/features/*.feature`. Step state is passed through `context.Context` with unexported key types.

Tests never hit the network: HTTP is replayed via a govcr cassette at `xml1/fixtures/bgg.json` (a ~1.9MB recorded request/response file — not a schema). To test a new endpoint or parameter combination, a new interaction must be recorded into the cassette; it is excluded from the `check-added-large-files` pre-commit hook.
