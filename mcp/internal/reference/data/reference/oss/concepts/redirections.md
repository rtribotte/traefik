---
schema_version: 2
kind: concept
name: Redirections
id: concept.redirections
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L125
summary: Redirections is a set of redirection for an entry point.
fields:
  - name: entryPoint
    go_name: EntryPoint
    type: object
    go_type: '*RedirectEntryPoint'
    type_ref: oss:RedirectEntryPoint
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

# Redirections

Redirections is a set of redirection for an entry point.
