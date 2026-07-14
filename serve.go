package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/richardwooding/bggclient/internal/mcpserver"
)

type ServeCmd struct {
	HTTP string `help:"Serve MCP over streamable HTTP on this address (e.g. :8080) instead of stdio." placeholder:"ADDR"`
}

func (c *ServeCmd) Run(ctx context.Context, g *Globals) error {
	server := mcpserver.New(g.newAPI(), version)
	if c.HTTP == "" {
		return server.Run(ctx, &mcp.StdioTransport{})
	}

	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)
	httpServer := &http.Server{Addr: c.HTTP, Handler: handler}
	go func() {
		<-ctx.Done()
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("shutdown: %v", err)
		}
	}()
	log.Printf("serving MCP over HTTP on %s", c.HTTP)
	if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
