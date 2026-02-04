package timeout

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/traefik/traefik/v3/pkg/middlewares"
)

const (
	typeName       = "Timeouts"
	middlewareName = "traefik-internal-timeouts"
)

type timeout struct {
	readTimeout  time.Duration
	writeTimeout time.Duration

	next http.Handler
}

// New creates a timeout middleware.
func New(ctx context.Context, next http.Handler) (http.Handler, error) {
	middlewares.GetLogger(ctx, middlewareName, typeName).Debug().Msg("Creating middleware")

	return &timeout{
		next: next,
	}, nil
}

func (t *timeout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rc := http.NewResponseController(w)

	tw := &timeoutWriter{
		ResponseWriter: w,
		rc:             rc,
		timeout:        t.writeTimeout,
	}

	tr := &timeoutReader{
		ReadCloser: r.Body,
		rc:         rc,
		timeout:    t.readTimeout,
	}

	// Create new request with wrapped body.
	r2 := new(http.Request)
	*r2 = *r
	r2.Body = tr

	t.next.ServeHTTP(tw, r2)
}

type timeoutReader struct {
	io.ReadCloser
	rc      *http.ResponseController
	timeout time.Duration
}

// Read reads data from the connection, and resets the read deadline on success.
func (tr *timeoutReader) Read(b []byte) (int, error) {
	read, err := tr.ReadCloser.Read(b)
	if err != nil {
		return read, err
	}

	// Reset deadline before after each successful read.
	_ = tr.rc.SetReadDeadline(time.Now().Add(tr.timeout))
	return read, nil
}

type timeoutWriter struct {
	http.ResponseWriter
	rc      *http.ResponseController
	timeout time.Duration

	headersSent bool
}

// Write writes data to the connection, and resets the write deadline on success.
func (tw *timeoutWriter) Write(b []byte) (int, error) {
	written, err := tw.ResponseWriter.Write(b)
	if err != nil {
		return written, err
	}

	// Reset deadline before after each successful write.
	_ = tw.rc.SetWriteDeadline(time.Now().Add(tw.timeout))
	return written, err
}
