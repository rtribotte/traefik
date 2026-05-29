package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// addPrompts registers the guided diagnostic workflow. Prompts are explicitly
// selected by the user, so unlike searched tools they always load — making them
// the reliable way to enforce a multi-step procedure (notably: re-fetch live
// state instead of answering from earlier results).
func addPrompts(s *mcp.Server) {
	s.AddPrompt(&mcp.Prompt{
		Name:        "diagnose",
		Description: "Diagnose a Traefik routing problem — a missing/not-routing router, 5xx errors, or unexplained latency — from live data, and ground the fix in the reference.",
		Arguments: []*mcp.PromptArgument{
			{Name: "problem", Description: "What's wrong, in the user's words (e.g. 'api.localhost 404s', 'billing returns 502', 'checkout is slow'). Optional."},
			{Name: "target", Description: "The router, service or host involved if known (e.g. api@docker, billing.localhost). Optional."},
		},
	}, diagnose)
}

func promptResult(description, text string) *mcp.GetPromptResult {
	return &mcp.GetPromptResult{
		Description: description,
		Messages: []*mcp.PromptMessage{
			{Role: "user", Content: &mcp.TextContent{Text: text}},
		},
	}
}

func diagnose(_ context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return promptResult("Guided Traefik diagnosis", buildDiagnose(req.Params.Arguments["problem"], req.Params.Arguments["target"])), nil
}

// buildDiagnose renders the general triage playbook. Shared by the user-invoked
// prompt and the model-invoked tool so both stay in lockstep. It covers the three
// symptom families (missing route, 5xx, latency) behind one entry point and ends
// by grounding the fix in the embedded reference.
func buildDiagnose(problem, target string) string {
	subject := "the problem the user is describing"
	if problem != "" {
		subject = fmt.Sprintf("this problem: %q", problem)
	}
	if target != "" {
		subject += fmt.Sprintf(" (involving %q)", target)
	}

	return fmt.Sprintf(`Diagnose %s in Traefik.

Work only from live data. Traefik's configuration is dynamic and may have changed,
so re-fetch at every step; never answer from earlier results. Begin every path with
list_routers (note its configHash) and get_overview for the active providers.

First classify the symptom, then follow the matching path:

A — the route is missing, 404s, or never takes effect:
  1. If the router is absent, it was likely never registered: a misspelled
     label/key (e.g. "rulee" for "rule"), the wrong provider, or
     exposedByDefault=false without traefik.enable.
  2. If present but its status is "warning" or "disabled", call get_router and read
     its errors. Common causes: it references a middleware or service that does not
     exist, or uses Traefik v2 rule syntax on v3.
  3. Cross-check referenced names with list_middlewares and list_services.
  4. Use tail_access_logs to see whether any request ever reached the router.

B — the route returns 5xx errors:
  1. Confirm a router actually matches the request (get_router). If none matches,
     this is a routing/404 problem, not a 5xx — say so.
  2. Call get_service and get_service_health: are the backend servers UP? Is the
     configured server list/port what you expect?
  3. Call tail_access_logs with minStatus=500 (and the service name) for the actual
     5xx entries, their count and paths.
  4. Interpret: 502 with backends unreachable -> the backend is down or the service
     points at the wrong port (a service-config problem, not Traefik); 503 with no
     healthy servers -> all backends failing health checks; 500 passed through ->
     the application itself errored, Traefik is fine.

C — the route works but is slow:
  1. Confirm routing and health are fine (router present, servers UP, 200s).
  2. Call tail_access_logs filtered to the host/service, ideally with minDurationMs,
     to read the request durations.
  3. If durations are high but the backend is fast, look in the traces:
     search_traces with {duration>...} then get_trace. A large gap between the
     entrypoint span and the ReverseProxy span means the time is in the middleware
     chain, not the backend or the network.

Once you know the cause, ground the fix before proposing it: search_traefik_docs
for the relevant concept, get_traefik_concept for its exact v3 contract, and if you
produce corrected configuration, run validate_traefik_config on it before telling
the user to apply it.

Report the single most likely cause and the concrete fix. Rank the candidates if
more than one is plausible.`, subject)
}
