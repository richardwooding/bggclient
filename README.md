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
	"net/http"

	"github.com/richardwooding/bggclient/xml1"
)

func main() {
	api := xml1.NewAPI(xml1.Options{
		HttpClient: http.DefaultClient,
		BaseURL:    "https://boardgamegeek.com/xmlapi",
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
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
