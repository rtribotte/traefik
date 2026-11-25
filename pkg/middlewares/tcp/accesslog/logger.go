package accesslog

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"go.opentelemetry.io/contrib/bridges/otellogrus"
)

// Handler writes the TCP access logs, in the JSON format, and exports them with OTLP when configured.
// It is shared by the access log middlewares of all the TCP routers.
type Handler struct {
	logger *logrus.Logger
	file   io.WriteCloser
	otlp   bool
}

// NewHandler creates a new TCP access log Handler.
func NewHandler(ctx context.Context, config *otypes.TCPAccessLog) (*Handler, error) {
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

	if config.OTLP != nil {
		otelLoggerProvider, err := config.OTLP.NewLoggerProvider(ctx)
		if err != nil {
			return nil, fmt.Errorf("setting up OpenTelemetry logger provider: %w", err)
		}

		logger.Hooks.Add(otellogrus.NewHook("traefik", otellogrus.WithLoggerProvider(otelLoggerProvider)))
		if !config.DualOutput {
			logger.Out = io.Discard
		}
	}

	return &Handler{
		logger: logger,
		file:   file,
		otlp:   config.OTLP != nil,
	}, nil
}

func (h *Handler) log(fields logrus.Fields) {
	entry := h.logger.WithFields(fields)

	var message string
	if h.otlp {
		// As for the HTTP access logs, the OTLP log body is the formatted log line.
		// The formatter reads the entry level and time, which logrus only sets when the log method is called.
		// Setting them here avoids formatting a body with the zero values, and makes logrus reuse
		// this timestamp, so the OTLP body and the regular output carry the same level and time.
		entry.Level = logrus.InfoLevel
		entry.Time = time.Now()

		mBytes, err := h.logger.Formatter.Format(entry)
		if err != nil {
			message = fmt.Sprintf("Failed to format TCP access log entry: %v", err)
		} else {
			message = string(mBytes)
		}
	}

	entry.Println(message)
}
