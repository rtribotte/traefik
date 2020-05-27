package metrics

import (
	"context"
	"crypto/tls"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/containous/alice"
	"github.com/containous/traefik/v2/pkg/log"
	"github.com/containous/traefik/v2/pkg/metrics"
	"github.com/containous/traefik/v2/pkg/middlewares"
	"github.com/containous/traefik/v2/pkg/middlewares/retry"
	traefiktls "github.com/containous/traefik/v2/pkg/tls"
	gokitmetrics "github.com/go-kit/kit/metrics"
)

const (
	protoHTTP      = "http"
	protoSSE       = "sse"
	protoWebsocket = "websocket"
	typeName       = "Metrics"
	nameEntrypoint = "metrics-entrypoint"
	nameService    = "metrics-service"
	nameBackend    = "metrics-backend"
)

type metricsMiddleware struct {
	next                     http.Handler
	reqsCounter              gokitmetrics.Counter
	reqsTLSCounter           gokitmetrics.Counter
	reqDurationHistogram     metrics.ScalableHistogram
	upstreamBytesHistogram   metrics.ScalableHistogram
	downstreamBytesHistogram metrics.ScalableHistogram
	openConnsGauge           gokitmetrics.Gauge
	baseLabels               []string
	backends                 bool
}

// NewEntryPointMiddleware creates a new metrics middleware for an Entrypoint.
func NewEntryPointMiddleware(ctx context.Context, next http.Handler, registry metrics.Registry, entryPointName string) http.Handler {
	log.FromContext(middlewares.GetLoggerCtx(ctx, nameEntrypoint, typeName)).Debug("Creating middleware")

	return &metricsMiddleware{
		next:                     next,
		reqsCounter:              registry.EntryPointReqsCounter(),
		reqsTLSCounter:           registry.EntryPointReqsTLSCounter(),
		reqDurationHistogram:     registry.EntryPointReqDurationHistogram(),
		upstreamBytesHistogram:   registry.EntryPointUpstreamBytesHistogram(),
		downstreamBytesHistogram: registry.EntryPointDownstreamBytesHistogram(),
		openConnsGauge:           registry.EntryPointOpenConnsGauge(),
		baseLabels:               []string{"entrypoint", entryPointName},
	}
}

// NewServiceMiddleware creates a new metrics middleware for a Service.
func NewServiceMiddleware(ctx context.Context, next http.Handler, registry metrics.Registry, serviceName string) http.Handler {
	log.FromContext(middlewares.GetLoggerCtx(ctx, nameService, typeName)).Debug("Creating middleware")

	return &metricsMiddleware{
		next:                     next,
		reqsCounter:              registry.ServiceReqsCounter(),
		reqsTLSCounter:           registry.ServiceReqsTLSCounter(),
		reqDurationHistogram:     registry.ServiceReqDurationHistogram(),
		upstreamBytesHistogram:   registry.ServiceUpstreamBytesHistogram(),
		downstreamBytesHistogram: registry.ServiceDownstreamBytesHistogram(),
		openConnsGauge:           registry.ServiceOpenConnsGauge(),
		baseLabels:               []string{"service", serviceName},
	}
}

// NewBackendsMiddleware creates a new metrics middleware for service's Backends.
func NewBackendsMiddleware(ctx context.Context, next http.Handler, registry metrics.Registry, serviceName string) http.Handler {
	log.FromContext(middlewares.GetLoggerCtx(ctx, nameBackend, typeName)).Debug("Creating middleware")

	return &metricsMiddleware{
		next:                     next,
		reqsCounter:              registry.BackendsReqsCounter(),
		reqsTLSCounter:           registry.BackendsReqsTLSCounter(),
		reqDurationHistogram:     registry.BackendsReqDurationHistogram(),
		upstreamBytesHistogram:   registry.BackendsUpstreamBytesHistogram(),
		downstreamBytesHistogram: registry.BackendsDownstreamBytesHistogram(),
		baseLabels:               []string{"service", serviceName},
	}
}

// WrapEntryPointHandler Wraps metrics entrypoint to alice.Constructor.
func WrapEntryPointHandler(ctx context.Context, registry metrics.Registry, entryPointName string) alice.Constructor {
	return func(next http.Handler) (http.Handler, error) {
		return NewEntryPointMiddleware(ctx, next, registry, entryPointName), nil
	}
}

// WrapServiceHandler Wraps metrics service to alice.Constructor.
func WrapServiceHandler(ctx context.Context, registry metrics.Registry, serviceName string) alice.Constructor {
	return func(next http.Handler) (http.Handler, error) {
		return NewServiceMiddleware(ctx, next, registry, serviceName), nil
	}
}

// WrapBackendsHandler Wraps metrics service to alice.Constructor.
func WrapBackendsHandler(ctx context.Context, registry metrics.Registry, serviceName string) alice.Constructor {
	return func(next http.Handler) (http.Handler, error) {
		return NewBackendsMiddleware(ctx, next, registry, serviceName), nil
	}
}

func (m *metricsMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var labels []string
	labels = append(labels, m.baseLabels...)
	labels = append(labels, "method", getMethod(req), "protocol", getRequestProtocol(req))

	m.openConnsGauge.With(labels...).Add(1)
	defer m.openConnsGauge.With(labels...).Add(-1)

	recorder := newResponseRecorder(rw)
	start := time.Now()

	m.next.ServeHTTP(recorder, req)

	if m.backends {
		labels = append(labels, "host", req.URL.Host)
	}

	if req.ContentLength >= 0 {
		m.upstreamBytesHistogram.With(labels...).Observe(float64(req.ContentLength))
	}

	contentLength := rw.Header().Get("Content-Length")
	if contentLength != "" {
		cl, err := strconv.Atoi(contentLength)
		if err == nil {
			m.downstreamBytesHistogram.With(labels...).Observe(float64(cl))
		}
	}
	// TODO test it
	//Hi @0xVox,
	//
	//	Thanks for your interest
	//Today, the only data we got for one server (backend) is its url. As you mentioned, the access log provides this url.
	//	A new metric can be introduced by adding a new label which will be this url.
	//	If it's what your are proposing, to do so, one can collect this url on the request, after the execution of the next `ServeHTTP` func in the metrics middleware.

	labels = append(labels, "code", strconv.Itoa(recorder.getCode()))

	histograms := m.reqDurationHistogram.With(labels...)
	histograms.ObserveFromStart(start)

	// TLS metrics
	if req.TLS != nil {
		var tlsLabels []string
		tlsLabels = append(tlsLabels, m.baseLabels...)
		tlsLabels = append(tlsLabels, "tls_version", getRequestTLSVersion(req), "tls_cipher", getRequestTLSCipher(req))

		m.reqsTLSCounter.With(tlsLabels...).Add(1)
	}

	m.reqsCounter.With(labels...).Add(1)
}

func getRequestProtocol(req *http.Request) string {
	switch {
	case isWebsocketRequest(req):
		return protoWebsocket
	case isSSERequest(req):
		return protoSSE
	default:
		return protoHTTP
	}
}

// isWebsocketRequest determines if the specified HTTP request is a websocket handshake request.
func isWebsocketRequest(req *http.Request) bool {
	return containsHeader(req, "Connection", "upgrade") && containsHeader(req, "Upgrade", "websocket")
}

// isSSERequest determines if the specified HTTP request is a request for an event subscription.
func isSSERequest(req *http.Request) bool {
	return containsHeader(req, "Accept", "text/event-stream")
}

func containsHeader(req *http.Request, name, value string) bool {
	items := strings.Split(req.Header.Get(name), ",")
	for _, item := range items {
		if value == strings.ToLower(strings.TrimSpace(item)) {
			return true
		}
	}
	return false
}

func getMethod(r *http.Request) string {
	if !utf8.ValidString(r.Method) {
		log.Warnf("Invalid HTTP method encoding: %s", r.Method)
		return "NON_UTF8_HTTP_METHOD"
	}
	return r.Method
}

func getRequestTLSVersion(req *http.Request) string {
	switch req.TLS.Version {
	case tls.VersionTLS10:
		return "1.0"
	case tls.VersionTLS11:
		return "1.1"
	case tls.VersionTLS12:
		return "1.2"
	case tls.VersionTLS13:
		return "1.3"
	default:
		return "unknown"
	}
}

func getRequestTLSCipher(req *http.Request) string {
	if version, ok := traefiktls.CipherSuitesReversed[req.TLS.CipherSuite]; ok {
		return version
	}

	return "unknown"
}

type retryMetrics interface {
	ServiceRetriesCounter() gokitmetrics.Counter
}

// NewRetryListener instantiates a MetricsRetryListener with the given retryMetrics.
func NewRetryListener(retryMetrics retryMetrics, serviceName string) retry.Listener {
	return &RetryListener{retryMetrics: retryMetrics, serviceName: serviceName}
}

// RetryListener is an implementation of the RetryListener interface to
// record RequestMetrics about retry attempts.
type RetryListener struct {
	retryMetrics retryMetrics
	serviceName  string
}

// Retried tracks the retry in the RequestMetrics implementation.
func (m *RetryListener) Retried(req *http.Request, attempt int) {
	m.retryMetrics.ServiceRetriesCounter().With("service", m.serviceName).Add(1)
}
