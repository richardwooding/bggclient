package main

import (
	"context"

	"github.com/richardwooding/shelfofshame/internal/gameopts"
	"github.com/richardwooding/shelfofshame/xml1"
)

type SearchCmd struct {
	Query string `arg:"" help:"Text to search game names for."`
	Exact bool   `help:"Only return exact name matches."`
}

func (c *SearchCmd) Run(ctx context.Context, g *Globals) error {
	var opts []xml1.SearchOption
	if c.Exact {
		opts = append(opts, xml1.ExactSearch)
	}
	result, err := g.newAPI().SearchBoardgames(ctx, c.Query, opts...)
	if err != nil {
		return err
	}
	return g.printJSON(result)
}

type BoardgameCmd struct {
	IDs []string `arg:"" name:"id" help:"BGG object ids of the games to fetch (max 20)."`
	gameopts.BoardgameFlags
}

func (c *BoardgameCmd) Run(ctx context.Context, g *Globals) error {
	opts, err := c.Options()
	if err != nil {
		return err
	}
	result, err := g.newAPI().GetBoardgamesById(ctx, c.IDs, opts...)
	if err != nil {
		return err
	}
	return g.printJSON(result)
}

type CollectionCmd struct {
	Username string `arg:"" help:"BGG username whose collection to fetch."`
	gameopts.CollectionFlags
}

func (c *CollectionCmd) Run(ctx context.Context, g *Globals) error {
	result, err := g.newAPI().GetCollection(ctx, c.Username, c.Options()...)
	if err != nil {
		return err
	}
	return g.printJSON(result)
}

type GeeklistCmd struct {
	ID       string `arg:"" help:"Numeric id of the geeklist."`
	Comments bool   `help:"Include comments on the geeklist."`
}

func (c *GeeklistCmd) Run(ctx context.Context, g *Globals) error {
	var opts []xml1.GeeklistOption
	if c.Comments {
		opts = append(opts, xml1.GeeklistComments)
	}
	result, err := g.newAPI().GetGeeklist(ctx, c.ID, opts...)
	if err != nil {
		return err
	}
	return g.printJSON(result)
}
