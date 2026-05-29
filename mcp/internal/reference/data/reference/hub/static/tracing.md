---
schema_version: 2
kind: static-section
name: Tracing
id: hub.static.tracing
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L88
summary: Tracing holds the tracing configuration.
fields:
  - name: additionalTraceHeaders
    go_name: AdditionalTraceHeaders
    type: object
    go_type: AdditionalTraceHeaders
representations:
  yaml_path: hub.tracing
  toml_path: hub.tracing
---

# Tracing

Tracing holds the tracing configuration.
