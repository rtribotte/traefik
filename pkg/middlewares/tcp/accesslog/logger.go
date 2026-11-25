package accesslog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
)

// Handler writes the TCP access logs, in the JSON format.
// It is shared by the access log middlewares of all the TCP routers.
type Handler struct {
	logger *logrus.Logger
	file   io.WriteCloser
}

// NewHandler creates a new TCP access log Handler.
func NewHandler(config *otypes.TCPAccessLog) (*Handler, error) {
	var file io.WriteCloser = os.Stdout
	if len(config.FilePath) > 0 {
		f, err := os.OpenFile(config.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o664)
		if err != nil {
			return nil, fmt.Errorf("opening TCP access log file: %w", err)
		}
		file = f
	}

	// The zero value of the logger mutex makes the writes thread-safe.
	logger := &logrus.Logger{
		Out:       file,
		Formatter: &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	return &Handler{
		logger: logger,
		file:   file,
	}, nil
}

func (h *Handler) log(fields logrus.Fields) {
	h.logger.WithFields(fields).Println()
}
