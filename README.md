# bggclient
BoardGameGeek Client Library.

Client for the [BoardGameGeek XML API](https://boardgamegeek.com/wiki/page/BGG_XML_API).

When using this library please abide by the [BoardGameGeek API Terms of Use](https://boardgamegeek.com/wiki/page/BGG_XML_API_Terms_of_Use).

## Installation

To install the library, use `go get`:

```sh
go get github.com/richardwooding/bggclient@latest
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/richardwooding/bggclient"
)

func main() {
    client := bggclient.NewClient(nil)
    ctx := context.Background()

    // Search for board games
    boardgames, err := client.SearchBoardgames(ctx, "Catan")
    if err != nil {
        log.Fatal(err)
    }
	for _, bg := range boardgames.Boardgames {
		fmt.Println(bg.Name)
	}
    fmt.Println(boardgames)

    // Get board game by ID
    boardgame, err := client.GetBoardgameById(ctx, "13")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(boardgame)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
