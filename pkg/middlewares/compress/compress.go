package compress

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/containous/traefik/v2/pkg/config/dynamic"
	"github.com/containous/traefik/v2/pkg/log"
	"github.com/containous/traefik/v2/pkg/middlewares"
	"github.com/containous/traefik/v2/pkg/tracing"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	typeName = "Compress"
)

// Compress is a middleware that allows to compress the response.
type compress struct {
	next     http.Handler
	name     string
	excludes []string
}

// New creates a new compress middleware.
func New(ctx context.Context, next http.Handler, conf dynamic.Compress, name string) (http.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName)).Debug("Creating middleware")

	excludes := []string{"application/grpc"}
	for _, v := range conf.ExcludedContentTypes {
		mediaType, _, err := mime.ParseMediaType(v)
		if err != nil {
			return nil, err
		}

		excludes = append(excludes, mediaType)
	}

	return &compress{next: next, name: name, excludes: excludes}, nil
}

const (
	brEncoding      = "br"
	gzipEncoding    = "gzip"
	deflateEncoding = "deflate"
)

func (w *compressResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *compressResponseWriter) writeHeader() {
	if w.code != 0 {
		w.ResponseWriter.WriteHeader(w.code)
		w.code = 0
	}
}

func (w *compressResponseWriter) WriteHeader(c int) {
	//w.ResponseWriter.Header().Del("Content-Length")
	if w.code == 0 {
		w.code = c
	}
}

func (w *compressResponseWriter) Write(b []byte) (int, error) {
	h := w.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}
	h.Del("Content-Length")

	if contains(w.exlusions, h.Get("Content-Type")) {
		w.WriteHeader(http.StatusOK)
		return w.ResponseWriter.Write(b)
	}

	e := h.Get("Content-Encoding")
	// Don't compress short pieces of data since it won't improve performance. Ignore compressed data too
	if len(b) < 1400 || e == brEncoding || e == gzipEncoding || e == deflateEncoding {
		w.WriteHeader(http.StatusOK)
		return w.ResponseWriter.Write(b)
	}

	h.Set("Content-Encoding", w.encoding)
	w.WriteHeader(http.StatusOK)

	if w.encoding == brEncoding {
		if w.level < brotli.BestSpeed || w.level > brotli.BestCompression {
			w.level = brotli.DefaultCompression
		}

		w.Writer = brotli.NewWriterLevel(w, w.level)
	} else {
		if w.level < gzip.HuffmanOnly || w.level > gzip.BestCompression {
			w.level = gzip.DefaultCompression
		}

		w.Writer, _ = gzip.NewWriterLevel(w, w.level)
	}
	defer w.Writer.Close()

	return w.Writer.Write(b)
}

// CompressHandler gzip/brotli compresses HTTP responses for clients that support it
// via the 'Accept-Encoding' header.
//
// Compressing TLS traffic may leak the page contents to an attacker if the
// page contains user input: http://security.stackexchange.com/a/102015/12208
func compressHandler(h http.Handler) http.Handler {
	return compressHandlerLevel(h, 6, []string{})
}

// CompressHandlerLevel gzip/brotli compresses HTTP responses with specified compression level
// for clients that support it via the 'Accept-Encoding' header.
//
// The compression level should be valid for encodings you use
func compressHandlerLevel(h http.Handler, level int, exclusions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// detect what encoding to use
		var encoding string
		for _, curEnc := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
			curEnc = strings.TrimSpace(curEnc)
			if curEnc == brEncoding || curEnc == gzipEncoding {
				encoding = curEnc
				//if curEnc == brEncoding {
				break
				//}
			}
		}

		// if we weren't able to identify an encoding we're familiar with, pass on the
		// request to the handler and return
		if encoding == "" {
			h.ServeHTTP(w, r)
			return
		}

		r.Header.Del("Accept-Encoding")
		w.Header().Add("Vary", "Accept-Encoding")

		w = &compressResponseWriter{
			ResponseWriter: w,
			encoding:       encoding,
			level:          level,
			exlusions:      exclusions,
		}

		h.ServeHTTP(w, r)
	})
}

type compressResponseWriter struct {
	exlusions []string
	Writer    io.WriteCloser
	http.ResponseWriter
	encoding string
	level    int
	code     int
}

func (w *compressResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}

	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func (w *compressResponseWriter) Flush() {
	if fw, ok := w.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}

//func (w *compressResponseWriter) WriteHeader(code int) {
//	if !w.exclusionComputed {
//		w.exclusionComputed = true
//		contentType := w.Header().Get("Content-Type")
//
//		if !contains(w.exlusions(), contentType) {
//			w.originalWriter = nil
//		} else {
//			// Copy headers to original response writer fallback
//			for key, values := range w.compressWriter.Header() {
//				for _, value := range values {
//					w.originalWriter.Header().Add(key, value)
//				}
//			}
//		}
//	}
//
//	if w.originalWriter != nil {
//		w.originalWriter.WriteHeader(code)
//		return
//	}
//
//	w.compressWriter.WriteHeader(code)
//}

func (c *compress) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	mediaType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		log.FromContext(middlewares.GetLoggerCtx(context.Background(), c.name, typeName)).Debug(err)
	}

	if contains(c.excludes, mediaType) {
		c.next.ServeHTTP(rw, req)
	} else {
		//ctx := middlewares.GetLoggerCtx(req.Context(), c.name, typeName)
		compressHandlerLevel(c.next, 6, c.excludes).ServeHTTP(rw, req)
	}
}

func (c *compress) GetTracingInformation() (string, ext.SpanKindEnum) {
	return c.name, tracing.SpanKindNoneEnum
}

//func gzipHandler(ctx context.Context, h http.Handler) http.Handler {
//	wrapper, err := gziphandler.GzipHandlerWithOpts(
//		gziphandler.CompressionLevel(gzip.DefaultCompression),
//		gziphandler.MinSize(gziphandler.DefaultMinSize))
//	if err != nil {
//		log.FromContext(ctx).Error(err)
//	}
//
//	return wrapper(h)
//}

func contains(values []string, val string) bool {
	for _, v := range values {
		if strings.Contains(val, v) {
			return true
		}
	}
	return false
}
