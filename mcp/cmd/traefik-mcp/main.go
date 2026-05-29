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
	"github.com/traefik/traefik-mcp/internal/configschema"
	"github.com/traefik/traefik-mcp/internal/loki"
	"github.com/traefik/traefik-mcp/internal/prom"
	"github.com/traefik/traefik-mcp/internal/rag"
	"github.com/traefik/traefik-mcp/internal/server"
	"github.com/traefik/traefik-mcp/internal/staticconf"
	"github.com/traefik/traefik-mcp/internal/tempo"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

const version = "0.1.0"

func main() {
	apiURL := flag.String("traefik.api-url", "http://localhost:8080", "Base URL of the Traefik API.")
	name := flag.String("traefik.name", "primary", "Name identifying this Traefik instance in tool output.")
	timeout := flag.Duration("traefik.timeout", 5*time.Second, "Timeout for Traefik API requests.")
	accessLog := flag.String("traefik.access-log", "", "Path to Traefik's JSON access log file (enables the access-log tool).")
	appLog := flag.String("traefik.app-log", "", "Path to Traefik's JSON application log file (enables the app-log tool).")
	tempoURL := flag.String("tempo.url", "", "Base URL of a Tempo (otel-lgtm) instance (enables the trace tools), e.g. http://localhost:3200.")
	lokiURL := flag.String("loki.url", "", "Base URL of a Loki (otel-lgtm) instance (enables querying OTLP-shipped access logs), e.g. http://localhost:3100.")
	promURL := flag.String("prometheus.url", "", "Base URL of a Prometheus (otel-lgtm) instance (enables querying OTLP-shipped metrics), e.g. http://localhost:9090.")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	target := traefik.NewHTTPTarget(*name, *apiURL, &http.Client{Timeout: *timeout})

	var tempoClient *tempo.Client
	if *tempoURL != "" {
		tempoClient = tempo.New(*tempoURL, &http.Client{Timeout: *timeout})
	}

	var lokiClient *loki.Client
	if *lokiURL != "" {
		lokiClient = loki.New(*lokiURL, &http.Client{Timeout: *timeout})
	}

	var promClient *prom.Client
	if *promURL != "" {
		promClient = prom.New(*promURL, &http.Client{Timeout: *timeout})
	}

	caps := detectCapabilities(*appLog)

	validator, err := configschema.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: config validation disabled: %v\n", err)
	}

	var retriever rag.Retriever
	if r, err := rag.NewEmbedded(); err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: documentation search disabled: %v\n", err)
	} else {
		retriever = r
	}

	srv := server.New("traefik-mcp", version, server.Deps{
		Target:        target,
		AccessLogPath: *accessLog,
		AppLogPath:    *appLog,
		Tempo:         tempoClient,
		Loki:          lokiClient,
		Prom:          promClient,
		Caps:          caps,
		Validator:     validator,
		Retriever:     retriever,
	})

	if err := srv.Run(ctx, &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: %v\n", err)
		os.Exit(1)
	}
}

// detectCapabilities recovers Traefik's static configuration from its app log
// and derives which data-source tools to expose. A nil result means the static
// configuration is unknown (no app log, or Traefik not at debug level), in which
// case the server falls back to registering every tool.
func detectCapabilities(appLog string) *staticconf.Capabilities {
	if appLog == "" {
		return nil
	}

	f, err := os.Open(appLog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: cannot read app log for static config: %v\n", err)
		return nil
	}
	defer f.Close()

	cfg, err := staticconf.FromLog(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "traefik-mcp: %v; registering all tools\n", err)
		return nil
	}

	caps := staticconf.Detect(cfg)
	fmt.Fprintf(os.Stderr, "traefik-mcp: detected capabilities from static config: %+v\n", caps)
	return &caps
}
