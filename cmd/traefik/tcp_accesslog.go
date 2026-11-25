package main

import (
	"github.com/rs/zerolog/log"
	tcpaccesslog "github.com/traefik/traefik/v3/pkg/middlewares/tcp/accesslog"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
)

// setupTCPAccessLog creates a TCP access log handler from the config, or returns nil if not enabled.
func setupTCPAccessLog(conf *otypes.TCPAccessLog) *tcpaccesslog.Handler {
	if conf == nil {
		return nil
	}

	handler, err := tcpaccesslog.NewHandler(conf)
	if err != nil {
		log.Warn().Err(err).Msg("Unable to create TCP access logger")
		return nil
	}

	return handler
}
