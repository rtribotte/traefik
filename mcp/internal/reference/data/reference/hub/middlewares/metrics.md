---
schema_version: 2
kind: middleware-hub
name: Metrics
id: hub.middlewares.metrics
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/metrics/metrics.go#L32
summary: Configuration configures an API Management metrics handler.
fields:
  - name: apiName
    go_name: APIName
    type: string
    go_type: string
  - name: apiNamespace
    go_name: APINamespace
    type: string
    go_type: string
  - name: apiVersionName
    go_name: APIVersionName
    type: string
    go_type: string
representations:
  yaml_path: http.middlewares.<name>.plugin.metrics
  toml_path: http.middlewares.<name>.plugin.metrics
  label_prefix: traefik.http.middlewares.<name>.plugin.metrics
---

# Metrics

Configuration configures an API Management metrics handler.
