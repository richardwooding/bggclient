// Package mcpserver exposes the BGG XML API 1 client as MCP tools.
package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/richardwooding/bggclient/internal/bggopts"
	"github.com/richardwooding/bggclient/xml1"
	"github.com/richardwooding/bggclient/xml1/model"
)

type server struct {
	api *xml1.API
}

// New returns an MCP server exposing search, boardgame, collection, and
// geeklist tools backed by the given API client. All tool calls share the
// client's rate limiter, so calls may block while BGG's request interval
// elapses.
func New(api *xml1.API, version string) *mcp.Server {
	s := &server{api: api}
	srv := mcp.NewServer(&mcp.Implementation{Name: "bggclient", Version: version}, nil)
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "bgg_search",
		Description: "Search BoardGameGeek for board games by name. Returns matching games with their BGG ids and publication years.",
	}, s.search)
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "bgg_get_boardgames",
		Description: "Get full details for up to 20 board games by BGG id, optionally including user comments, rating statistics, and historical data.",
	}, s.getBoardgames)
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "bgg_get_collection",
		Description: "Get a BoardGameGeek user's game collection, filterable by ownership, ratings, plays, wishlist, and trade status.",
	}, s.getCollection)
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "bgg_get_geeklist",
		Description: "Get a BoardGameGeek geeklist by id, optionally including comments.",
	}, s.getGeeklist)
	return srv
}

type SearchIn struct {
	Query string `json:"query" jsonschema:"text to search game names for"`
	Exact bool   `json:"exact,omitempty" jsonschema:"only return exact name matches"`
}

func (s *server) search(ctx context.Context, req *mcp.CallToolRequest, in SearchIn) (*mcp.CallToolResult, model.Boardgames, error) {
	var opts []xml1.SearchOption
	if in.Exact {
		opts = append(opts, xml1.ExactSearch)
	}
	result, err := s.api.SearchBoardgames(ctx, in.Query, opts...)
	if err != nil {
		return nil, model.Boardgames{}, err
	}
	return nil, *result, nil
}

type BoardgamesIn struct {
	IDs []string `json:"ids" jsonschema:"BGG numeric object ids of the games to fetch, maximum 20"`
	bggopts.BoardgameFlags
}

func (s *server) getBoardgames(ctx context.Context, req *mcp.CallToolRequest, in BoardgamesIn) (*mcp.CallToolResult, model.Boardgames, error) {
	opts, err := in.Options()
	if err != nil {
		return nil, model.Boardgames{}, err
	}
	result, err := s.api.GetBoardgamesById(ctx, in.IDs, opts...)
	if err != nil {
		return nil, model.Boardgames{}, err
	}
	return nil, *result, nil
}

type CollectionIn struct {
	Username string `json:"username" jsonschema:"BGG username whose collection to fetch"`
	bggopts.CollectionFlags
}

func (s *server) getCollection(ctx context.Context, req *mcp.CallToolRequest, in CollectionIn) (*mcp.CallToolResult, model.Items, error) {
	result, err := s.api.GetCollection(ctx, in.Username, in.Options()...)
	if err != nil {
		return nil, model.Items{}, err
	}
	return nil, *result, nil
}

type GeeklistIn struct {
	ID       string `json:"id" jsonschema:"numeric id of the geeklist"`
	Comments bool   `json:"comments,omitempty" jsonschema:"include comments on the geeklist"`
}

func (s *server) getGeeklist(ctx context.Context, req *mcp.CallToolRequest, in GeeklistIn) (*mcp.CallToolResult, model.Geeklist, error) {
	var opts []xml1.GeeklistOption
	if in.Comments {
		opts = append(opts, xml1.GeeklistComments)
	}
	result, err := s.api.GetGeeklist(ctx, in.ID, opts...)
	if err != nil {
		return nil, model.Geeklist{}, err
	}
	return nil, *result, nil
}
