package accesslog

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"go.opentelemetry.io/contrib/bridges/otellogrus"
)

// Handler writes the TCP access logs, in the JSON format, and exports them with OTLP when configured.
// It is shared by the access log middlewares of all the TCP routers.
type Handler struct {
	config *otypes.TCPAccessLog
	logger *logrus.Logger

	// The regular output is discarded when only OTLP is used, even after a rotation.
	discardOutput bool

	// mu protects the file, which is replaced on rotation.
	mu   sync.Mutex
	file io.WriteCloser

	// The log lines are only buffered when the buffering size is greater than zero.
	logChan chan logLine
	wg      sync.WaitGroup
}

type logLine struct {
	fields logrus.Fields
	time   time.Time
}

// NewHandler creates a new TCP access log Handler.
func NewHandler(ctx context.Context, config *otypes.TCPAccessLog) (*Handler, error) {
	var file io.WriteCloser = noopCloser{os.Stdout}
	if len(config.FilePath) > 0 {
		f, err := openFile(config.FilePath)
		if err != nil {
			return nil, err
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

	handler := &Handler{
		config: config,
		logger: logger,
		file:   file,
	}

	if config.OTLP != nil {
		otelLoggerProvider, err := config.OTLP.NewLoggerProvider(ctx)
		if err != nil {
			_ = file.Close()
			return nil, fmt.Errorf("setting up OpenTelemetry logger provider: %w", err)
		}

		logger.Hooks.Add(otellogrus.NewHook("traefik", otellogrus.WithLoggerProvider(otelLoggerProvider)))
		if !config.DualOutput {
			handler.discardOutput = true
			logger.Out = io.Discard
		}
	}

	if config.BufferingSize > 0 {
		handler.logChan = make(chan logLine, config.BufferingSize)
		handler.wg.Go(func() {
			for line := range handler.logChan {
				handler.write(line)
			}
		})
	}

	return handler, nil
}

// Close writes the buffered log lines, and closes the log file.
func (h *Handler) Close() error {
	if h.logChan != nil {
		close(h.logChan)
		h.wg.Wait()
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	return h.file.Close()
}

// Rotate closes and reopens the log file to allow for rotation by an external source.
func (h *Handler) Rotate() error {
	if h.config.FilePath == "" {
		return nil
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	file, err := openFile(h.config.FilePath)
	if err != nil {
		return err
	}

	if !h.discardOutput {
		h.logger.SetOutput(file)
	}

	_ = h.file.Close()
	h.file = file

	return nil
}

func (h *Handler) log(fields logrus.Fields) {
	// The line is timestamped when the connection ends, even when it is written later on.
	line := logLine{fields: fields, time: time.Now()}

	if h.logChan != nil {
		h.logChan <- line
		return
	}

	h.write(line)
}

func (h *Handler) write(line logLine) {
	entry := h.logger.WithFields(line.fields).WithTime(line.time)

	var message string
	if h.config.OTLP != nil {
		// As for the HTTP access logs, the OTLP log body is the formatted log line.
		// The formatter reads the entry level, which logrus only sets when the log method is called.
		entry.Level = logrus.InfoLevel

		mBytes, err := h.logger.Formatter.Format(entry)
		if err != nil {
			message = fmt.Sprintf("Failed to format TCP access log entry: %v", err)
		} else {
			message = string(mBytes)
		}
	}

	entry.Println(message)
}

func openFile(filePath string) (*os.File, error) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("creating TCP access log path %s: %w", dir, err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o664)
	if err != nil {
		return nil, fmt.Errorf("opening TCP access log file %s: %w", filePath, err)
	}

	return file, nil
}

// noopCloser keeps the standard output open when the handler is closed.
type noopCloser struct {
	*os.File
}

func (n noopCloser) Close() error {
	return nil
}
