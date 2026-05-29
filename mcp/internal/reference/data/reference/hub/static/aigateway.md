---
schema_version: 2
kind: static-section
name: AIGateway
id: hub.static.aigateway
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L68
summary: AIGateway holds the ai gateway configuration.
fields:
  - name: maxRequestBodySize
    go_name: MaxRequestBodySize
    type: integer
    go_type: int
representations:
  yaml_path: hub.aigateway
  toml_path: hub.aigateway
---

# AIGateway

AIGateway holds the ai gateway configuration.
