---
schema_version: 2
kind: static-section
name: MCPGateway
id: hub.static.mcpgateway
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L78
summary: MCPGateway holds the MCP gateway configuration.
fields:
  - name: maxRequestBodySize
    go_name: MaxRequestBodySize
    type: integer
    go_type: int
representations:
  yaml_path: hub.mcpgateway
  toml_path: hub.mcpgateway
---

# MCPGateway

MCPGateway holds the MCP gateway configuration.
