---
schema_version: 2
kind: middleware-hub
name: APIKey
id: hub.middlewares.apikey
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/apikey/config.go#L18
summary: Configuration holds the API Key middleware configuration.
fields:
  - name: keySource
    go_name: KeySource
    type: object
    go_type: token.Source
  - name: secretNonBase64Encoded
    go_name: SecretNonBase64Encoded
    type: boolean
    go_type: bool
  - name: secretValues
    go_name: SecretValues
    type: array
    items: string
    go_type: '[]string'
representations:
  yaml_path: http.middlewares.<name>.plugin.apikey
  toml_path: http.middlewares.<name>.plugin.apikey
  label_prefix: traefik.http.middlewares.<name>.plugin.apikey
---

# APIKey

Configuration holds the API Key middleware configuration.
