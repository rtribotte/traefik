---
schema_version: 2
kind: middleware-hub
name: HubAPIKey
id: hub.middlewares.hubapikey
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/hubapikey/config.go#L13
summary: Configuration configures an API Key handler.
fields:
  - name: keySource
    go_name: KeySource
    type: object
    go_type: token.Source
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
representations:
  yaml_path: http.middlewares.<name>.plugin.hubapikey
  toml_path: http.middlewares.<name>.plugin.hubapikey
  label_prefix: traefik.http.middlewares.<name>.plugin.hubapikey
---

# HubAPIKey

Configuration configures an API Key handler.
