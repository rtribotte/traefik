// Package slowdown is a Traefik local (Yaegi) middleware plugin that delays
// every request by a random duration. The delay is intentionally not
// configurable: the slowness cannot be inferred from the middleware
// configuration, so it can only be found by looking at the request traces.
package slowdown

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

// Config holds the middleware configuration. It is deliberately empty: there is
// no knob that hints at the delay.
type Config struct{}

// CreateConfig returns the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// SlowDown delays each request by a random duration before passing it on.
type SlowDown struct {
	next http.Handler
	name string
}

// New builds the middleware.
func New(_ context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	return &SlowDown{next: next, name: name}, nil
}

func (s *SlowDown) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Random delay between 500ms and 3s, imitating real, non-deterministic
	// slowness rather than a fixed, guessable value.
	delay := time.Duration(500+rand.Intn(2500)) * time.Millisecond
	time.Sleep(delay)
	s.next.ServeHTTP(rw, req)
}
