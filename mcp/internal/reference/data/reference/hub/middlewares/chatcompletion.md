---
schema_version: 2
kind: middleware-hub
name: ChatCompletion
id: hub.middlewares.chatcompletion
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/chatcompletion/config.go#L7
summary: Config holds the ChatCompletion Middleware configuration.
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
  - name: params
    go_name: Params
    type: object
    go_type: '*Params'
  - name: allowParamsOverride
    go_name: AllowParamsOverride
    type: boolean
    go_type: '*bool'
representations:
  yaml_path: http.middlewares.<name>.plugin.chatcompletion
  toml_path: http.middlewares.<name>.plugin.chatcompletion
  label_prefix: traefik.http.middlewares.<name>.plugin.chatcompletion
---

# ChatCompletion

Config holds the ChatCompletion Middleware configuration.
