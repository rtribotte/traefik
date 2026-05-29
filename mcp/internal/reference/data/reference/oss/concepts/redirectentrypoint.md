---
schema_version: 2
kind: concept
name: RedirectEntryPoint
id: concept.redirectentrypoint
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L130
summary: RedirectEntryPoint is the definition of an entry point redirection.
fields:
  - name: to
    go_name: To
    type: string
    go_type: string
  - name: scheme
    go_name: Scheme
    type: string
    go_type: string
    default: https
    description: Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.
  - name: permanent
    go_name: Permanent
    type: boolean
    go_type: bool
    default: true
    description: Permanent defines whether the redirection is permanent (308).
  - name: priority
    go_name: Priority
    type: integer
    go_type: int
    default: 9223372036854775807
    description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
---

# RedirectEntryPoint

RedirectEntryPoint is the definition of an entry point redirection.
