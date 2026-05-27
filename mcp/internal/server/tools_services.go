package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/traefik/traefik-mcp/internal/traefik"
)

// ServiceSummary is the structured projection of a service returned to the model.
type ServiceSummary struct {
	Name         string            `json:"name"`
	Status       string            `json:"status,omitempty"`
	Types        []string          `json:"types,omitempty"`
	Servers      []string          `json:"servers,omitempty"`
	ServerStatus map[string]string `json:"serverStatus,omitempty"`
	UsedBy       []string          `json:"usedBy,omitempty"`
	Errors       []string          `json:"errors,omitempty"`
}

type listServicesInput struct{}

type listServicesOutput struct {
	Services []ServiceSummary `json:"services"`
}

func listServices(target traefik.Target) mcp.ToolHandlerFor[listServicesInput, listServicesOutput] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listServicesInput) (*mcp.CallToolResult, listServicesOutput, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, listServicesOutput{}, err
		}

		out := listServicesOutput{Services: make([]ServiceSummary, 0, len(raw.Services))}
		for name, info := range raw.Services {
			out.Services = append(out.Services, serviceSummary(name, info))
		}
		return nil, out, nil
	}
}

type getServiceInput struct {
	Name string `json:"name" jsonschema:"fully qualified service name, e.g. whoami@docker"`
}

func getService(target traefik.Target) mcp.ToolHandlerFor[getServiceInput, ServiceSummary] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getServiceInput) (*mcp.CallToolResult, ServiceSummary, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, ServiceSummary{}, err
		}

		info, ok := raw.Services[in.Name]
		if !ok {
			return nil, ServiceSummary{}, fmt.Errorf("service %q not found", in.Name)
		}
		return nil, serviceSummary(in.Name, info), nil
	}
}

// ServiceHealth aggregates the per-server health of a service.
type ServiceHealth struct {
	Name    string            `json:"name"`
	Healthy bool              `json:"healthy"`
	Up      int               `json:"up"`
	Down    int               `json:"down"`
	Servers map[string]string `json:"servers,omitempty"`
}

type getServiceHealthInput struct {
	Name string `json:"name" jsonschema:"fully qualified service name, e.g. whoami@docker"`
}

func getServiceHealth(target traefik.Target) mcp.ToolHandlerFor[getServiceHealthInput, ServiceHealth] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, in getServiceHealthInput) (*mcp.CallToolResult, ServiceHealth, error) {
		raw, err := traefik.FetchRawData(ctx, target)
		if err != nil {
			return nil, ServiceHealth{}, err
		}

		info, ok := raw.Services[in.Name]
		if !ok {
			return nil, ServiceHealth{}, fmt.Errorf("service %q not found", in.Name)
		}

		health := ServiceHealth{Name: in.Name, Servers: info.ServerStatus}
		for _, status := range info.ServerStatus {
			if status == "UP" {
				health.Up++
			} else {
				health.Down++
			}
		}
		health.Healthy = health.Down == 0 && health.Up > 0
		return nil, health, nil
	}
}

func serviceSummary(name string, info *traefik.ServiceInfo) ServiceSummary {
	s := ServiceSummary{
		Name:         name,
		ServerStatus: info.ServerStatus,
	}
	if info.ServiceInfo != nil {
		s.Status = info.Status
		s.UsedBy = info.UsedBy
		s.Errors = info.Err
		s.Types = configKeys(info.Service)
		if info.Service != nil && info.LoadBalancer != nil {
			for _, server := range info.LoadBalancer.Servers {
				s.Servers = append(s.Servers, server.URL)
			}
		}
	}
	return s
}
