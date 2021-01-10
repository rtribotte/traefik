package chain

import (
	"context"

	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tcp"
)

const (
	typeName = "ChainTCP"
)

type chainBuilder interface {
	BuildChain(ctx context.Context, middlewares []string) *tcp.Chain
}

// New creates a chain middleware.
func New(ctx context.Context, next tcp.Handler, config dynamic.TCPChain, builder chainBuilder, name string) (tcp.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName)).Debug("Creating middleware")

	middlewareChain := builder.BuildChain(ctx, config.Middlewares)
	return middlewareChain.Then(next)
}
