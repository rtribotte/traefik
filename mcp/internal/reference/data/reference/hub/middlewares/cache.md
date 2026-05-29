---
schema_version: 2
kind: middleware-hub
name: Cache
id: hub.middlewares.cache
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/cache/config.go#L13
summary: Configuration holds the Cache Middleware configuration.
fields:
  - name: store
    go_name: Store
    type: object
    go_type: '*StoreConfig'
  - name: maxTtl
    go_name: MaxTTL
    type: integer
    go_type: '*int'
  - name: disableCacheStatusHeader
    go_name: DisableCacheStatusHeader
    type: boolean
    go_type: bool
  - name: maxStale
    go_name: MaxStale
    type: integer
    go_type: int
  - name: excludedResponseCodes
    go_name: ExcludedResponseCodes
    type: array
    items: string
    go_type: '[]string'
representations:
  yaml_path: http.middlewares.<name>.plugin.cache
  toml_path: http.middlewares.<name>.plugin.cache
  label_prefix: traefik.http.middlewares.<name>.plugin.cache
---

# Cache

Configuration holds the Cache Middleware configuration.
