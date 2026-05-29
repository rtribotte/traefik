---
schema_version: 2
kind: middleware-hub
name: ContentGuard
id: hub.middlewares.contentguard
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/contentguard/config.go#L17
summary: Config holds the configuration for content-guard middleware.
fields:
  - name: clientRequestFormat
    go_name: ClientRequestFormat
    type: object
    go_type: aiformat.ClientRequestFormat
    description: 'ClientRequestFormat defines the format used by the upstream client. Valid values: "ccr", "custom", "responsesAPI". Default: "custom".'
  - name: engine
    go_name: Engine
    type: object
    go_type: EngineConfig
  - name: request
    go_name: Request
    type: object
    go_type: RulesConfig
  - name: response
    go_name: Response
    type: object
    go_type: RulesConfig
representations:
  yaml_path: http.middlewares.<name>.plugin.contentguard
  toml_path: http.middlewares.<name>.plugin.contentguard
  label_prefix: traefik.http.middlewares.<name>.plugin.contentguard
---

# ContentGuard

Config holds the configuration for content-guard middleware.
