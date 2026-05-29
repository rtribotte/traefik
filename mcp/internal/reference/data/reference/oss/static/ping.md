---
schema_version: 2
kind: static-section
name: Handler
id: static.ping
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/ping/ping.go#L10
summary: Handler expose ping routes.
fields:
  - name: entryPoint
    go_name: EntryPoint
    type: string
    go_type: string
    default: traefik
  - name: manualRouting
    go_name: ManualRouting
    type: boolean
    go_type: bool
  - name: terminatingStatusCode
    go_name: TerminatingStatusCode
    type: integer
    go_type: int
    default: 503
representations:
  yaml_path: ping
  toml_path: ping
---

# Handler

Handler expose ping routes.
