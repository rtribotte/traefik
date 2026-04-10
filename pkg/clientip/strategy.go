package clientip

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strings"
)

const (
	xForwardedFor = "X-Forwarded-For"
)

// Strategy a strategy for IP selection.
type Strategy interface {
	GetIP(req *http.Request) string
}

// Get an IP selection strategy.
// If nil it returns a strategy that honors the client IP resolved at the
// entrypoint (see ip.ResolvedAddrStrategy) and falls back to the remote
// address; otherwise it returns a strategy based on the configuration using
// the X-Forwarded-For Header. Depth override the ExcludedIPs.
func (s *IPStrategy) Get() (Strategy, error) {
	if s == nil {
		return &ResolvedAddrStrategy{}, nil
	}

	if s.Depth > 0 {
		if s.IPv6Subnet != nil && (*s.IPv6Subnet <= 0 || *s.IPv6Subnet > 128) {
			return nil, fmt.Errorf("invalid IPv6 subnet %d value, should be greater to 0 and lower or equal to 128", *s.IPv6Subnet)
		}

		return &DepthStrategy{
			Depth:      s.Depth,
			IPv6Subnet: s.IPv6Subnet,
		}, nil
	}

	if len(s.ExcludedIPs) > 0 {
		checker, err := NewChecker(s.ExcludedIPs)
		if err != nil {
			return nil, err
		}
		return &PoolStrategy{
			Checker: checker,
		}, nil
	}

	if s.IPv6Subnet != nil && (*s.IPv6Subnet <= 0 || *s.IPv6Subnet > 128) {
		return nil, fmt.Errorf("invalid IPv6 subnet %d value, should be greater to 0 and lower or equal to 128", *s.IPv6Subnet)
	}

	return &ResolvedAddrStrategy{
		IPv6Subnet: s.IPv6Subnet,
	}, nil
}

// RemoteAddrStrategy a strategy that always return the remote address.
type RemoteAddrStrategy struct {
	// IPv6Subnet instructs the strategy to return the first IP of the subnet where IP belongs.
	IPv6Subnet *int
}

// GetIP returns the selected IP.
func (s *RemoteAddrStrategy) GetIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}

	if s.IPv6Subnet != nil {
		return getIPv6SubnetIP(ip, *s.IPv6Subnet)
	}

	return ip
}

// ResolvedAddrStrategy returns the client IP resolved by the entrypoint
// forwarded headers middleware (via the ip.ClientIP context helper) when
// available, and otherwise falls back to the remote address.
//
// It is used as the default strategy when no explicit IPStrategy has been
// configured on a middleware, so that IPAllowList, IPWhiteList, rate limiter,
// in-flight request counter, and any other strategy-aware middleware
// automatically honor the real client IP resolved once per request at the
// entrypoint, while preserving the original behavior when the entrypoint has
// not been configured to resolve it.
type ResolvedAddrStrategy struct {
	// IPv6Subnet instructs the strategy to return the first IP of the subnet where IP belongs.
	IPv6Subnet *int
}

// GetIP returns the selected IP.
func (s *ResolvedAddrStrategy) GetIP(req *http.Request) string {
	addr, ok := FromContext(req.Context())
	if !ok {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			return req.RemoteAddr
		}
		addr = ip
	}

	if s.IPv6Subnet != nil {
		return getIPv6SubnetIP(addr, *s.IPv6Subnet)
	}

	return addr
}

// DepthStrategy a strategy based on the depth inside the X-Forwarded-For from right to left.
type DepthStrategy struct {
	Depth int
	// IPv6Subnet instructs the strategy to return the first IP of the subnet where IP belongs.
	IPv6Subnet *int
}

// GetIP returns the selected IP.
func (s *DepthStrategy) GetIP(req *http.Request) string {
	xff := req.Header.Get(xForwardedFor)
	xffs := strings.Split(xff, ",")

	if len(xffs) < s.Depth {
		return ""
	}

	ip := strings.TrimSpace(xffs[len(xffs)-s.Depth])

	if s.IPv6Subnet != nil {
		return getIPv6SubnetIP(ip, *s.IPv6Subnet)
	}

	return ip
}

// PoolStrategy is a strategy based on an IP Checker.
// It allows to check whether addresses are in a given pool of IPs.
type PoolStrategy struct {
	Checker *Checker
}

// GetIP checks the list of Forwarded IPs (most recent first) against the
// Checker pool of IPs. It returns the first IP that is not in the pool, or the
// empty string otherwise.
func (s *PoolStrategy) GetIP(req *http.Request) string {
	if s.Checker == nil {
		return ""
	}

	xff := req.Header.Get(xForwardedFor)
	xffs := strings.Split(xff, ",")

	for i := len(xffs) - 1; i >= 0; i-- {
		xffTrimmed := strings.TrimSpace(xffs[i])
		if len(xffTrimmed) == 0 {
			continue
		}
		if contain, _ := s.Checker.Contains(xffTrimmed); !contain {
			return xffTrimmed
		}
	}

	return ""
}

// getIPv6SubnetIP returns the IPv6 subnet IP.
// It returns the original IP when it is not an IPv6, or if parsing the IP has failed with an error.
func getIPv6SubnetIP(ip string, ipv6Subnet int) string {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return ip
	}

	if !addr.Is6() {
		return ip
	}

	prefix, err := addr.Prefix(ipv6Subnet)
	if err != nil {
		return ip
	}

	return prefix.Addr().String()
}
