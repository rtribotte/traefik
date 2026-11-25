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

// Logger writes the TCP access logs, in the JSON format, and exports them with OTLP when configured.
// It is shared by the access log middlewares of all the TCP routers.
type Logger struct {
	config       *otypes.TCPAccessLog
	logrusLogger *logrus.Logger

	// The regular output is discarded when only OTLP is used, even after a rotation.
	discardOutput bool

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

	// The zero value of the logrus logger mutex makes the writes thread-safe.
	logrusLogger := &logrus.Logger{
		Out:       file,
		Formatter: &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	logger := &Logger{
		config:       config,
		logrusLogger: logrusLogger,
		file:         file,
	}

	if config.OTLP != nil {
		otelLoggerProvider, err := config.OTLP.NewLoggerProvider(ctx)
		if err != nil {
			_ = file.Close()
			return nil, fmt.Errorf("setting up OpenTelemetry logger provider: %w", err)
		}

		logrusLogger.Hooks.Add(otellogrus.NewHook("traefik", otellogrus.WithLoggerProvider(otelLoggerProvider)))
		if !config.DualOutput {
			logger.discardOutput = true
			logrusLogger.Out = io.Discard
		}
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

	if !l.discardOutput {
		l.logrusLogger.SetOutput(file)
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
	entry := l.logrusLogger.WithFields(line.fields).WithTime(line.time)

	var message string
	if l.config.OTLP != nil {
		// As for the HTTP access logs, the OTLP log body is the formatted log line.
		// The formatter reads the entry level, which logrus only sets when the log method is called.
		entry.Level = logrus.InfoLevel

		mBytes, err := l.logrusLogger.Formatter.Format(entry)
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

// noopCloser keeps the standard output open when the logger is closed.
type noopCloser struct {
	*os.File
}

func (n noopCloser) Close() error {
	return nil
}
