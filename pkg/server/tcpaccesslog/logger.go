package tcpaccesslog

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	otypes "github.com/traefik/traefik/v3/pkg/observability/types"
	"go.opentelemetry.io/contrib/bridges/otellogrus"
)

// Logger writes the TCP access logs, in the JSON format, and exports them with OTLP when configured.
// It is shared by the access log middlewares of all the TCP routers.
type Logger struct {
	config    *otypes.TCPAccessLog
	formatter logrus.Formatter

	// fileLogger writes the log lines to the file.
	// It is nil when only OTLP is used.
	fileLogger *logrus.Logger

	// otlpLogger exports the log lines with OTLP, with the formatted log line as body.
	// It is nil when OTLP is not used.
	// It is distinct from the file one, so that the body does not end up in the message of the written log line.
	otlpLogger *logrus.Logger

	// muFile protects the file, which is replaced on rotation.
	muFile sync.Mutex
	file   io.WriteCloser

	// The log lines are only buffered when the buffering size is greater than zero.
	logChan chan logLine
	wg      sync.WaitGroup
}

type logLine struct {
	fields logrus.Fields
	time   time.Time
}

// NewLogger creates a new TCP access Logger.
func NewLogger(ctx context.Context, config *otypes.TCPAccessLog) (*Logger, error) {
	var file io.WriteCloser = noopCloser{os.Stdout}
	if len(config.FilePath) > 0 {
		f, err := openFile(config.FilePath)
		if err != nil {
			return nil, err
		}
		file = f
	}

	logger := &Logger{
		config:    config,
		formatter: &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano},
		file:      file,
	}

	// The zero value of the logrus logger mutex makes the writes thread-safe.
	if config.OTLP == nil || config.DualOutput {
		logger.fileLogger = &logrus.Logger{
			Out:       file,
			Formatter: logger.formatter,
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.InfoLevel,
		}
	}

	if config.OTLP != nil {
		otelLoggerProvider, err := config.OTLP.NewLoggerProvider(ctx)
		if err != nil {
			_ = file.Close()
			return nil, fmt.Errorf("setting up OpenTelemetry logger provider: %w", err)
		}

		// The OTLP logger only exports the log lines through its hook,
		// so it does not format them for its output.
		logger.otlpLogger = &logrus.Logger{
			Out:       io.Discard,
			Formatter: noopFormatter{},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.InfoLevel,
		}
		logger.otlpLogger.Hooks.Add(otellogrus.NewHook("traefik", otellogrus.WithLoggerProvider(otelLoggerProvider)))
	}

	if config.BufferingSize > 0 {
		logger.logChan = make(chan logLine, config.BufferingSize)
		logger.wg.Go(func() {
			for line := range logger.logChan {
				logger.write(line)
			}
		})
	}

	return logger, nil
}

// Close writes the buffered log lines, and closes the log file.
func (l *Logger) Close() error {
	if l.logChan != nil {
		close(l.logChan)
		l.wg.Wait()
	}

	l.muFile.Lock()
	defer l.muFile.Unlock()

	return l.file.Close()
}

// Rotate closes and reopens the log file to allow for rotation by an external source.
func (l *Logger) Rotate() error {
	if l.config.FilePath == "" {
		return nil
	}

	l.muFile.Lock()
	defer l.muFile.Unlock()

	file, err := openFile(l.config.FilePath)
	if err != nil {
		return err
	}

	if l.fileLogger != nil {
		l.fileLogger.SetOutput(file)
	}

	_ = l.file.Close()
	l.file = file

	return nil
}

func (l *Logger) log(fields logrus.Fields) {
	// The line is timestamped when the connection ends, even when it is written later on.
	line := logLine{fields: fields, time: time.Now()}

	if l.logChan != nil {
		l.logChan <- line
		return
	}

	l.write(line)
}

func (l *Logger) write(line logLine) {
	if l.fileLogger != nil {
		l.fileLogger.WithFields(line.fields).WithTime(line.time).Println()
	}

	if l.otlpLogger != nil {
		l.export(line)
	}
}

// export exports the given log line with OTLP.
// As for the HTTP access logs, the body is the formatted log line,
// as some observability platforms parse it instead of the attributes.
func (l *Logger) export(line logLine) {
	entry := l.otlpLogger.WithFields(line.fields).WithTime(line.time)

	// The formatter reads the entry level, which logrus only sets when the log method is called.
	formatted := entry.Dup()
	formatted.Level = logrus.InfoLevel

	// The empty JSON object keeps the body parseable as the other ones when the formatting fails.
	body := "{}"
	if b, err := l.formatter.Format(formatted); err != nil {
		// The log line is still exported with its attributes,
		// but the formatting error is not sent as if it was the log line.
		log.Error().Err(err).Msg("Unable to format the TCP access log body")
	} else {
		body = string(b)
	}

	entry.Println(body)
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

// noopCloser keeps the standard output open when the logger is closed.
type noopCloser struct {
	*os.File
}

func (n noopCloser) Close() error {
	return nil
}

// noopFormatter does not format the log lines,
// for the loggers which only export them through their hooks.
type noopFormatter struct{}

func (noopFormatter) Format(*logrus.Entry) ([]byte, error) {
	return nil, nil
}
