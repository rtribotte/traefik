package routeback

import (
	"context"
	"net/http"

	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
)

const (
	typeName = "RouteBack"
)

// RouteBack is a middleware used to forward the request to the entryPoint muxer.
type routeBack struct {
	name string
	next http.Handler
}

// New creates a new handler.
func New(ctx context.Context, next http.Handler, name string) (http.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName)).Debug("Creating middleware")
	return &routeBack{next: next, name: name}, nil
}

func (a *routeBack) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	passedBy := req.Context().Value(typeName)

	if passedBy != nil {
		rw.WriteHeader(http.StatusMisdirectedRequest)
		if _, err := rw.Write([]byte(http.StatusText(http.StatusMisdirectedRequest))); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	ctx := context.WithValue(req.Context(), typeName, a.name)
	a.next.ServeHTTP(rw, req.WithContext(ctx))
}
