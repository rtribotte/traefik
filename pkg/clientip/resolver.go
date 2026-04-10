package clientip

import (
	"net"
	"net/http"
	"strings"
)

// Resolver resolves the real client IP of a request by walking a configured
// source header against a pool of trusted proxy IPs, matching nginx's
// ngx_http_realip_module (set_real_ip_from + real_ip_recursive) semantics.
//
// If the immediate peer (req.RemoteAddr) is not in the trusted pool, the
// Resolver returns an empty string: untrusted peers cannot influence real-IP
// resolution, mirroring nginx behavior and providing anti-spoofing by design.
type Resolver struct {
	// Header is the source header to read (e.g. X-Forwarded-For,
	// CF-Connecting-IP, True-Client-IP).
	Header string
	// Trusted is the pool of trusted proxy IPs/CIDRs. Resolution only runs
	// when the immediate peer is in this pool.
	Trusted *Checker
}

// Resolve returns the resolved client IP for the given request, or an empty
// string if the immediate peer is not trusted or no usable value is found.
//
// For list-valued headers (X-Forwarded-For) the header is walked right-to-left,
// skipping entries that are themselves in the trusted pool, and the first
// untrusted entry is returned. For single-valued headers (CF-Connecting-IP,
// X-Real-IP, True-Client-IP, ...) the first non-empty value is returned as-is.
func (r *Resolver) Resolve(req *http.Request) string {
	if r == nil || r.Trusted == nil || r.Header == "" {
		return ""
	}

	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		host = req.RemoteAddr
	}
	if contain, _ := r.Trusted.Contains(host); !contain {
		return ""
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
		for _, p := range strings.Split(v, ",") {
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
