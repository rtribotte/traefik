package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolver_Resolve(t *testing.T) {
	trusted, err := NewChecker([]string{"10.0.0.0/8", "192.168.1.5"})
	require.NoError(t, err)

	tests := []struct {
		desc       string
		header     string
		remoteAddr string
		headers    http.Header
		want       string
	}{
		{
			desc:       "untrusted peer returns empty",
			header:     "X-Forwarded-For",
			remoteAddr: "8.8.8.8:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4"}},
			want:       "",
		},
		{
			desc:       "trusted peer, XFF single value",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4"}},
			want:       "1.2.3.4",
		},
		{
			desc:       "trusted peer, XFF chain, walks past trusted hops",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4, 10.0.0.2, 10.0.0.3"}},
			want:       "1.2.3.4",
		},
		{
			desc:       "trusted peer, XFF chain with untrusted in middle stops at untrusted",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4, 5.6.7.8, 10.0.0.3"}},
			want:       "5.6.7.8",
		},
		{
			desc:       "trusted peer, repeated XFF headers are folded",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4", "10.0.0.2"}},
			want:       "1.2.3.4",
		},
		{
			desc:       "trusted peer, all XFF entries are trusted -> empty",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"10.0.0.2, 10.0.0.3"}},
			want:       "",
		},
		{
			desc:       "trusted peer, missing header -> empty",
			header:     "X-Forwarded-For",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{},
			want:       "",
		},
		{
			desc:       "trusted peer, single-value CF-Connecting-IP",
			header:     "CF-Connecting-IP",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"Cf-Connecting-Ip": []string{"1.2.3.4"}},
			want:       "1.2.3.4",
		},
		{
			desc:       "trusted peer, single-value header is not split on comma",
			header:     "CF-Connecting-IP",
			remoteAddr: "10.0.0.1:1234",
			headers:    http.Header{"Cf-Connecting-Ip": []string{"1.2.3.4, 5.6.7.8"}},
			want:       "1.2.3.4, 5.6.7.8",
		},
		{
			desc:       "trusted CIDR exact IP",
			header:     "X-Forwarded-For",
			remoteAddr: "192.168.1.5:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4"}},
			want:       "1.2.3.4",
		},
		{
			desc:       "IPv6 trusted peer, IPv4 resolved",
			header:     "X-Forwarded-For",
			remoteAddr: "[::1]:1234",
			headers:    http.Header{"X-Forwarded-For": []string{"1.2.3.4"}},
			want:       "", // ::1 is not in the trust list
		},
	}

	r := &Resolver{Header: "", Trusted: trusted}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			r.Header = tt.header
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req.RemoteAddr = tt.remoteAddr
			req.Header = tt.headers
			assert.Equal(t, tt.want, r.Resolve(req))
		})
	}
}

func TestResolver_Resolve_NilSafety(t *testing.T) {
	var r *Resolver
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	assert.Equal(t, "", r.Resolve(req))

	r = &Resolver{}
	assert.Equal(t, "", r.Resolve(req))
}
