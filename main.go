package main

import (
	"context"
	"fmt"
	"os"

	"github.com/emmanuelay/borsdata-mcp/client"
	"github.com/emmanuelay/borsdata-mcp/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// These will be set at build time
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	apiKey := os.Getenv("BORSDATA_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "BORSDATA_API_KEY environment variable is required")
		os.Exit(1)
	}

	versionInfo := fmt.Sprintf("v%s (%s) %s", version, commit, date)
	apiClient := client.New(apiKey, version)
	s := server.New(apiClient, versionInfo)

	if err := s.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
