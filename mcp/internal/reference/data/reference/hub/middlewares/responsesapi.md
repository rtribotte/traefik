---
schema_version: 2
kind: middleware-hub
name: ResponsesAPI
id: hub.middlewares.responsesapi
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/responsesapi/config.go#L23
summary: Config holds the ResponsesAPI Middleware configuration.
fields:
  - name: token
    go_name: Token
    type: string
    go_type: string
  - name: model
    go_name: Model
    type: string
    go_type: string
  - name: allowModelOverride
    go_name: AllowModelOverride
    type: boolean
    go_type: '*bool'
  - name: allowParamsOverride
    go_name: AllowParamsOverride
    type: boolean
    go_type: '*bool'
  - name: params
    go_name: Params
    type: object
    go_type: '*Params'
  - name: instructions
    go_name: Instructions
    type: string
    go_type: string
representations:
  yaml_path: http.middlewares.<name>.plugin.responsesapi
  toml_path: http.middlewares.<name>.plugin.responsesapi
  label_prefix: traefik.http.middlewares.<name>.plugin.responsesapi
---

# ResponsesAPI

Config holds the ResponsesAPI Middleware configuration.
