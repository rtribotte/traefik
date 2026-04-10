package clientip

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientIP(t *testing.T) {
	tests := []struct {
		desc       string
		remoteAddr string
		ctxValue   string // empty = no context value
		want       string
	}{
		{
			desc:       "context resolved IP wins",
			remoteAddr: "8.8.8.8:1234",
			ctxValue:   "1.2.3.4",
			want:       "1.2.3.4",
		},
		{
			desc:       "no context, RemoteAddr with port is stripped",
			remoteAddr: "8.8.8.8:1234",
			want:       "8.8.8.8",
		},
		{
			desc:       "no context, RemoteAddr without port returned as-is",
			remoteAddr: "8.8.8.8",
			want:       "8.8.8.8",
		},
		{
			desc:       "no context, IPv6 with port",
			remoteAddr: "[::1]:1234",
			want:       "::1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req.RemoteAddr = tt.remoteAddr
			if tt.ctxValue != "" {
				req = req.WithContext(WithClientIP(req.Context(), tt.ctxValue))
			}
			assert.Equal(t, tt.want, ClientIP(req))
		})
	}
}

func TestWithClientIPAndFromContext(t *testing.T) {
	ctx := context.Background()
	_, ok := fromContext(ctx)
	assert.False(t, ok)

	ctx = WithClientIP(ctx, "1.2.3.4")
	got, ok := fromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, "1.2.3.4", got)
}

func TestResolvedAddrStrategy_GetIP(t *testing.T) {
	s := &ResolvedAddrStrategy{}

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.RemoteAddr = "8.8.8.8:1234"
	assert.Equal(t, "8.8.8.8", s.GetIP(req))

	req = req.WithContext(WithClientIP(req.Context(), "1.2.3.4"))
	assert.Equal(t, "1.2.3.4", s.GetIP(req))
}
