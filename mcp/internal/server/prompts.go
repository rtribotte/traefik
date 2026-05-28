package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// addPrompts registers the guided diagnostic workflows. Prompts are explicitly
// selected by the user, so unlike searched tools they always load — making them
// the reliable way to enforce a multi-step procedure (notably: re-fetch live
// state instead of answering from earlier results).
func addPrompts(s *mcp.Server) {
	s.AddPrompt(&mcp.Prompt{
		Name:        "diagnose_router_missing",
		Description: "Diagnose why a Traefik router is missing or not routing traffic.",
		Arguments: []*mcp.PromptArgument{
			{Name: "router", Description: "Fully qualified router name if known (e.g. api@docker). Optional."},
		},
	}, diagnoseRouterMissing)

	s.AddPrompt(&mcp.Prompt{
		Name:        "diagnose_5xx",
		Description: "Diagnose 5xx errors and determine whether the fault is Traefik, the service config, or the app.",
		Arguments: []*mcp.PromptArgument{
			{Name: "service", Description: "Service name if known (e.g. billing@docker). Optional."},
			{Name: "host", Description: "Request host if known (e.g. billing.localhost). Optional."},
		},
	}, diagnose5xx)
}

func promptResult(description, text string) *mcp.GetPromptResult {
	return &mcp.GetPromptResult{
		Description: description,
		Messages: []*mcp.PromptMessage{
			{Role: "user", Content: &mcp.TextContent{Text: text}},
		},
	}
}

func diagnoseRouterMissing(_ context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return promptResult("Guided router diagnosis", buildRouterMissing(req.Params.Arguments["router"])), nil
}

// buildRouterMissing renders the router-diagnosis playbook. Shared by the
// user-invoked prompt and the model-invoked tool so both stay in lockstep.
func buildRouterMissing(router string) string {
	target := "the router the user is asking about"
	if router != "" {
		target = fmt.Sprintf("router %q", router)
	}

	return fmt.Sprintf(`Diagnose why %s is missing or not routing in Traefik.

Work only from live data. Do not answer from earlier results — Traefik's config is
dynamic and may have changed. Re-fetch at every step.

1. Call list_routers to get the current routers (note its configHash).
2. If the router is absent entirely, the most likely cause is that it was never
   registered: a misspelled label/key (e.g. "rulee" instead of "rule"), the wrong
   provider, or exposedByDefault=false without traefik.enable. Confirm by checking
   get_overview for the active providers.
3. If the router is present but its status is "warning" or "disabled", call
   get_router on it and read its errors. Common causes: it references a middleware
   that does not exist, references a missing service, or uses v2 rule syntax on v3.
4. Cross-check referenced names with list_middlewares and list_services.
5. Use tail_access_logs to see whether any request ever reached the router.

Then report the single most likely cause and the concrete fix. Rank the
candidates if more than one is plausible.`, target)
}

func diagnose5xx(_ context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return promptResult("Guided 5xx diagnosis", buildDiagnose5xx(req.Params.Arguments["service"], req.Params.Arguments["host"])), nil
}

// buildDiagnose5xx renders the 5xx-diagnosis playbook. Shared by the user-invoked
// prompt and the model-invoked tool so both stay in lockstep.
func buildDiagnose5xx(service, host string) string {
	subject := "the affected service"
	if service != "" {
		subject = fmt.Sprintf("service %q", service)
	} else if host != "" {
		subject = fmt.Sprintf("the service behind host %q", host)
	}

	return fmt.Sprintf(`Diagnose 5xx errors for %s and determine whether the fault is
Traefik, the service configuration, or the backend application.

Work only from live data. Re-fetch at every step; do not answer from earlier results.

1. Call list_routers (and get_router on the matching one) to confirm a router
   actually matches the request%s. If none matches, this is a routing/404 problem,
   not a 5xx — say so.
2. Call get_service and get_service_health for the service. Are the backend
   servers UP? Is the configured server list/port what you expect?
3. Call tail_access_logs with minStatus=500 (and the service name) to see the
   actual 5xx entries, their count and paths.
4. Interpret:
   - 502 with the router/service present but backends unreachable -> the backend
     is down or the service points at the wrong port. This is a service
     configuration problem, not Traefik.
   - 503 with no healthy servers -> all backends failing health checks.
   - 500 passed through -> the application itself errored; Traefik is fine.

Report whether the fault lies in Traefik routing, the service config, or the app,
and give the concrete fix.`, subject, hostClause(host))
}

func hostClause(host string) string {
	if host == "" {
		return ""
	}
	return fmt.Sprintf(" for host %q", host)
}
