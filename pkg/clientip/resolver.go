package clientip

import (
	"net"
	"net/http"
	"strings"
)

// Resolver resolves the real client IP of a request by walking a configured
// source header against a pool of trusted proxy IPs.
//
// When the immediate peer (req.RemoteAddr) is in the trusted pool,
// the configured header is read and walked to extract the real client IP.
// When the immediate peer is NOT trusted
type Resolver struct {
	// Header is the source header to read.
	Header string
	// Trusted is the pool of trusted proxy IPs/CIDRs.
	Trusted *Checker
}

// Resolve returns the resolved client IP for the given request.
func (r *Resolver) Resolve(req *http.Request) string {
	if r == nil || r.Trusted == nil || r.Header == "" {
		return ""
	}

	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		host = req.RemoteAddr
	}
	if contain, _ := r.Trusted.Contains(host); !contain {
		return host
	}

	values := req.Header.Values(r.Header)
	if len(values) == 0 {
		return ""
	}

	if !isListHeader(r.Header) {
		for _, v := range values {
			if trimmed := strings.TrimSpace(v); trimmed != "" {
				return trimmed
			}
		}
		return ""
	}

	// List-valued header: walk right-to-left across all values (headers may be
	// repeated), skipping entries that are themselves trusted, and return the
	// first untrusted entry found.
	var parts []string
	for _, v := range values {
		for p := range strings.SplitSeq(v, ",") {
			if trimmed := strings.TrimSpace(p); trimmed != "" {
				parts = append(parts, trimmed)
			}
		}
	}
	for i := len(parts) - 1; i >= 0; i-- {
		if contain, _ := r.Trusted.Contains(parts[i]); !contain {
			return parts[i]
		}
	}
	return ""
}

// isListHeader reports whether the given header name is an HTTP header that
// carries a comma-separated list of values and should be walked as a chain.
// Everything that is not X-Forwarded-For is treated as a single-valued header
// containing the real client IP directly (which matches how CDNs like
// Cloudflare, Akamai, and Fastly set their real-IP headers).
func isListHeader(name string) bool {
	return strings.EqualFold(name, "X-Forwarded-For")
}
