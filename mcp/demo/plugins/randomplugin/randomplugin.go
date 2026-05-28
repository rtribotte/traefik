// Package randomplugin is a Traefik local (Yaegi) middleware plugin used by the
// trace-latency demo scenario. Its name is intentionally non-descriptive: it
// adds a random delay to every request, but nothing in the name or
// configuration reveals that, so the latency can only be found in the traces.
package randomplugin

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

// Config holds the middleware configuration. It is deliberately empty: there is
// no knob that hints at the behaviour.
type Config struct{}

// CreateConfig returns the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// RandomPlugin delays each request by a random duration before passing it on.
type RandomPlugin struct {
	next http.Handler
	name string
}

// New builds the middleware.
func New(_ context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	return &RandomPlugin{next: next, name: name}, nil
}

func (p *RandomPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Random delay between 500ms and 3s, imitating real, non-deterministic
	// slowness rather than a fixed, guessable value.
	delay := time.Duration(500+rand.Intn(2500)) * time.Millisecond
	time.Sleep(delay)
	p.next.ServeHTTP(rw, req)
}
