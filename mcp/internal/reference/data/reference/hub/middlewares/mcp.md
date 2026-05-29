---
schema_version: 2
kind: middleware-hub
name: MCP
id: hub.middlewares.mcp
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/mcp/config.go#L67
summary: Config holds MCP middleware configuration.
fields:
  - name: resourceMetadata
    go_name: ResourceMetadata
    type: object
    go_type: '*ResourceMetadata'
  - name: policies
    go_name: Policies
    type: array
    items: object
    go_type: '[]Policy'
  - name: defaultAction
    go_name: DefaultAction
    type: object
    go_type: action
  - name: listPolicies
    go_name: ListPolicies
    type: array
    items: object
    go_type: '[]ListPolicy'
  - name: listDefaultAction
    go_name: ListDefaultAction
    type: object
    go_type: listAction
  - name: statusCodeOnDeny
    go_name: StatusCodeOnDeny
    type: integer
    go_type: int
representations:
  yaml_path: http.middlewares.<name>.plugin.mcp
  toml_path: http.middlewares.<name>.plugin.mcp
  label_prefix: traefik.http.middlewares.<name>.plugin.mcp
---

# MCP

Config holds MCP middleware configuration.
