package traefik

import (
	"context"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

// ServiceInfo mirrors the Traefik API representation of a service: the runtime
// ServiceInfo plus the per-server health map, which the API exposes as
// "serverStatus" but runtime.ServiceInfo keeps unexported.
type ServiceInfo struct {
	*runtime.ServiceInfo
	ServerStatus map[string]string `json:"serverStatus,omitempty"`
}

// TCPServiceInfo mirrors the API representation of a TCP service.
type TCPServiceInfo struct {
	*runtime.TCPServiceInfo
	ServerStatus map[string]string `json:"serverStatus,omitempty"`
}

// RawData is the decoded GET /api/rawdata payload. It reuses Traefik's runtime
// types so router/service/middleware shapes never drift from upstream, while
// exposing the serverStatus the API wraps around services.
type RawData struct {
	Routers        map[string]*runtime.RouterInfo        `json:"routers,omitempty"`
	Middlewares    map[string]*runtime.MiddlewareInfo    `json:"middlewares,omitempty"`
	Services       map[string]*ServiceInfo               `json:"services,omitempty"`
	TCPRouters     map[string]*runtime.TCPRouterInfo     `json:"tcpRouters,omitempty"`
	TCPMiddlewares map[string]*runtime.TCPMiddlewareInfo `json:"tcpMiddlewares,omitempty"`
	TCPServices    map[string]*TCPServiceInfo            `json:"tcpServices,omitempty"`
	UDPRouters     map[string]*runtime.UDPRouterInfo     `json:"udpRouters,omitempty"`
	UDPServices    map[string]*runtime.UDPServiceInfo    `json:"udpServices,omitempty"`
}

// FetchRawData retrieves and decodes the full runtime configuration snapshot.
func FetchRawData(ctx context.Context, target Target) (*RawData, error) {
	var raw RawData
	if err := target.Get(ctx, "/api/rawdata", &raw); err != nil {
		return nil, err
	}
	return &raw, nil
}
