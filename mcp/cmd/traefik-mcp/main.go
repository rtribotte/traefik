// Command traefik-mcp exposes a running Traefik instance to MCP clients over
// stdio.
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/server"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

const version = "0.1.0"

func main() {
	apiURL := flag.String("traefik.api-url", "http://localhost:8080", "Base URL of the Traefik API.")
	name := flag.String("traefik.name", "primary", "Name identifying this Traefik instance in tool output.")
	timeout := flag.Duration("traefik.timeout", 5*time.Second, "Timeout for Traefik API requests.")
	accessLog := flag.String("traefik.access-log", "", "Path to Traefik's JSON access log file (enables the access-log tool).")
	appLog := flag.String("traefik.app-log", "", "Path to Traefik's JSON application log file (enables the app-log tool).")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	target := traefik.NewHTTPTarget(*name, *apiURL, &http.Client{Timeout: *timeout})

	srv := server.New("traefik-mcp", version, server.Deps{Target: target, AccessLogPath: *accessLog, AppLogPath: *appLog})

	if err := srv.Run(ctx, &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: %v\n", err)
		os.Exit(1)
	}
}
