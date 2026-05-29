---
schema_version: 2
kind: middleware-hub
name: ForceCase
id: hub.middlewares.forcecase
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/forcecase/config.go#L9
summary: Configuration holds the Force Case middleware configuration.
fields:
  - name: headers
    go_name: Headers
    type: array
    items: string
    go_type: '[]string'
    description: Headers is the list of headers on which to force case.
representations:
  yaml_path: http.middlewares.<name>.plugin.forcecase
  toml_path: http.middlewares.<name>.plugin.forcecase
  label_prefix: traefik.http.middlewares.<name>.plugin.forcecase
---

# ForceCase

Configuration holds the Force Case middleware configuration.
