# bggclient

[![Go](https://github.com/richardwooding/bggclient/actions/workflows/go.yml/badge.svg)](https://github.com/richardwooding/bggclient/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/richardwooding/bggclient.svg)](https://pkg.go.dev/github.com/richardwooding/bggclient)

BoardGameGeek from your terminal and your AI assistant: a colourful CLI and an
[MCP](https://modelcontextprotocol.io) server for the
[BoardGameGeek XML API](https://boardgamegeek.com/wiki/page/BGG_XML_API), with the
underlying Go client available as a library.

**Website:** https://richardwooding.github.io/bggclient/

[![Powered by BGG](docs/powered-by-bgg-badge.png)](https://boardgamegeek.com)

All game data is sourced from [BoardGameGeek](https://boardgamegeek.com) and used
under the [XML API Terms of Use](https://boardgamegeek.com/wiki/page/XML_API_Terms_of_Use).

## You need your own BGG API token

BGG restricts the XML API to **registered applications** — without a token every
request returns `401 Unauthorized`. Registration is free for non-commercial use:

1. Read [Using the XML API](https://boardgamegeek.com/using_the_xml_api) and
   register your application there (requires a BGG account).
2. Create an API token for your registered application on the same page.
3. Provide it to bggclient via the `BGG_API_TOKEN` environment variable or the
   `--token` flag:

   ```sh
   export BGG_API_TOKEN=your-token
   ```

By using the API you agree to the
[XML API Terms of Use](https://boardgamegeek.com/wiki/page/XML_API_Terms_of_Use)
— note in particular that the data is licensed for **non-commercial** use and
may not be used to train AI/LLM systems.

## Install

**Homebrew**

```sh
brew install richardwooding/tap/bggclient
```

**Go**

```sh
go install github.com/richardwooding/bggclient@latest
```

**Container (ghcr.io)**

```sh
docker run --rm -e BGG_API_TOKEN ghcr.io/richardwooding/bggclient search "Catan"
```

## CLI

```sh
export BGG_API_TOKEN=your-token

# Search for games
bggclient search "Catan"
bggclient search "Brass Birmingham" --exact

# Fetch games by BGG id (up to 20), with statistics and comments
bggclient boardgame 13 224517 --stats --comments

# A user's collection, filtered
bggclient collection richardwooding --own --min-rating=7

# Geeklists
bggclient geeklist 11205 --comments
```

Output is pretty-printed, syntax-highlighted JSON. Colour is disabled
automatically when output is piped, when [`NO_COLOR`](https://no-color.org) is
set, or with `--no-color`.

Run `bggclient --help` for all flags, including `--request-interval` (default
5s, matching BGG's rate-limit guidance) and `--timeout`.

## MCP server

`bggclient serve` runs an MCP server over stdio exposing four tools:

| Tool | Description |
|------|-------------|
| `bgg_search` | Search board games by name |
| `bgg_get_boardgames` | Full details for up to 20 games, with stats/comments/history |
| `bgg_get_collection` | A user's collection with ownership/rating/plays filters |
| `bgg_get_geeklist` | Fetch a geeklist, optionally with comments |

**Claude Code**

```sh
claude mcp add bggclient --env BGG_API_TOKEN=your-token -- bggclient serve
```

**Claude Desktop** (`claude_desktop_config.json`)

```json
{
  "mcpServers": {
    "bggclient": {
      "command": "bggclient",
      "args": ["serve"],
      "env": { "BGG_API_TOKEN": "your-token" }
    }
  }
}
```

**Streamable HTTP** for remote hosts:

```sh
bggclient serve --http :8080
```

## Library

```sh
go get github.com/richardwooding/bggclient@latest
```

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/richardwooding/bggclient/xml1"
)

func main() {
	api := xml1.NewAPI(xml1.Options{
		BaseURL:  "https://boardgamegeek.com/xmlapi",
		APIToken: os.Getenv("BGG_API_TOKEN"),
	})
	ctx := context.Background()

	// Search for board games
	boardgames, err := api.SearchBoardgames(ctx, "Catan")
	if err != nil {
		log.Fatal(err)
	}
	for _, bg := range boardgames.Boardgames {
		fmt.Printf("%s https://boardgamegeek.com/boardgame/%s\n", bg.Name.Value, bg.ObjectID)
	}

	// A user's owned collection
	items, err := api.GetCollection(ctx, "richardwooding", xml1.Own(true))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d items owned\n", items.TotalItems)
}
```

The client rate-limits itself (one request per 5 seconds by default,
configurable via `Options.RequestInterval`) and retries on 429 and on BGG's
202 "still preparing" responses for collections and geeklists.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
