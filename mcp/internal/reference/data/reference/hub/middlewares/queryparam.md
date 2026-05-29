---
schema_version: 2
kind: middleware-hub
name: QueryParam
id: hub.middlewares.queryparam
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/queryparam/config.go#L22
summary: Configuration holds the QueryParam middleware configuration.
fields:
  - name: set
    go_name: Set
    type: array
    items: object
    go_type: '[]SetOperation'
  - name: append
    go_name: Append
    type: array
    items: object
    go_type: '[]AppendOperation'
  - name: remove
    go_name: Remove
    type: array
    items: object
    go_type: '[]RemoveOperation'
  - name: rename
    go_name: Rename
    type: array
    items: object
    go_type: '[]RenameOperation'
representations:
  yaml_path: http.middlewares.<name>.plugin.queryparam
  toml_path: http.middlewares.<name>.plugin.queryparam
  label_prefix: traefik.http.middlewares.<name>.plugin.queryparam
---

# QueryParam

Configuration holds the QueryParam middleware configuration.
