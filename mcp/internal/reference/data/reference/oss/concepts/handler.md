---
schema_version: 2
kind: concept
name: Handler
id: concept.handler
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
---

# Handler

Handler expose ping routes.
