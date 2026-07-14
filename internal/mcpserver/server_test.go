package mcpserver

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/richardwooding/bggclient/xml1"
	"github.com/richardwooding/bggclient/xml1/model"
	"github.com/seborama/govcr/v15"
)

// newSession wires the MCP server to an in-memory client, backed by the
// shared govcr cassette so no network access happens.
func newSession(t *testing.T) *mcp.ClientSession {
	t.Helper()
	vcr := govcr.NewVCR(govcr.NewCassetteLoader("../../xml1/fixtures/bgg.json"))
	api := xml1.NewAPI(xml1.Options{
		HttpClient: vcr.HTTPClient(),
		BaseURL:    "https://boardgamegeek.com/xmlapi",
		// Responses are replayed from the cassette, so no need to rate limit.
		RequestInterval: time.Millisecond,
	})
	server := New(api, "test")

	ctx := t.Context()
	serverTransport, clientTransport := mcp.NewInMemoryTransports()
	if _, err := server.Connect(ctx, serverTransport, nil); err != nil {
		t.Fatalf("connecting server: %v", err)
	}
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client"}, nil)
	session, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatalf("connecting client: %v", err)
	}
	t.Cleanup(func() { _ = session.Close() })
	return session
}

func callTool[T any](t *testing.T, session *mcp.ClientSession, name string, args map[string]any) T {
	t.Helper()
	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{Name: name, Arguments: args})
	if err != nil {
		t.Fatalf("calling %s: %v", name, err)
	}
	if res.IsError {
		t.Fatalf("%s returned tool error: %+v", name, res.Content)
	}
	data, err := json.Marshal(res.StructuredContent)
	if err != nil {
		t.Fatalf("marshaling structured content: %v", err)
	}
	var out T
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshaling structured content: %v", err)
	}
	return out
}

func TestListTools(t *testing.T) {
	session := newSession(t)
	res, err := session.ListTools(context.Background(), nil)
	if err != nil {
		t.Fatalf("listing tools: %v", err)
	}
	want := map[string]bool{
		"bgg_search":         false,
		"bgg_get_boardgames": false,
		"bgg_get_collection": false,
		"bgg_get_geeklist":   false,
	}
	for _, tool := range res.Tools {
		if _, ok := want[tool.Name]; ok {
			want[tool.Name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("tool %s not registered", name)
		}
	}
}

func TestSearch(t *testing.T) {
	session := newSession(t)
	got := callTool[model.Boardgames](t, session, "bgg_search", map[string]any{"query": "Catan"})
	if len(got.Boardgames) == 0 {
		t.Fatal("expected search results, got none")
	}
	found := false
	for _, bg := range got.Boardgames {
		if bg.Name.Value == "Catan" || bg.Name.Value == "CATAN" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a boardgame named Catan in %d results", len(got.Boardgames))
	}
}

func TestGetCollection(t *testing.T) {
	session := newSession(t)
	got := callTool[model.Items](t, session, "bgg_get_collection", map[string]any{
		"username": "richardwooding",
		"own":      true,
	})
	if len(got.Items) == 0 {
		t.Fatal("expected collection items, got none")
	}
}

func TestGetGeeklist(t *testing.T) {
	session := newSession(t)
	got := callTool[model.Geeklist](t, session, "bgg_get_geeklist", map[string]any{"id": "11205"})
	if got.ID != 11205 {
		t.Errorf("geeklist ID = %d, want 11205", got.ID)
	}
}

func TestSearchError(t *testing.T) {
	session := newSession(t)
	res, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "bgg_get_boardgames",
		Arguments: map[string]any{"ids": []string{"not-a-number"}},
	})
	if err != nil {
		t.Fatalf("calling tool: %v", err)
	}
	if !res.IsError {
		t.Error("expected IsError for invalid id, got success")
	}
}
