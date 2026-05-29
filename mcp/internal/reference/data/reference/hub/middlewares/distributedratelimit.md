---
schema_version: 2
kind: middleware-hub
name: DistributedRateLimit
id: hub.middlewares.distributedratelimit
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/distributedratelimit/config.go#L26
summary: Configuration holds the DistributedRateLimit middleware configuration.
fields:
  - name: store
    go_name: Store
    type: object
    go_type: StoreConfig
  - name: limit
    go_name: Limit
    type: integer
    go_type: int64
  - name: period
    go_name: Period
    type: duration
    go_type: types.Duration
  - name: burst
    go_name: Burst
    type: integer
    go_type: int64
  - name: sourceCriterion
    go_name: SourceCriterion
    type: object
    go_type: '*dynamic.SourceCriterion'
  - name: denyOnError
    go_name: DenyOnError
    type: boolean
    go_type: bool
  - name: responseHeaders
    go_name: ResponseHeaders
    type: boolean
    go_type: bool
representations:
  yaml_path: http.middlewares.<name>.plugin.distributedratelimit
  toml_path: http.middlewares.<name>.plugin.distributedratelimit
  label_prefix: traefik.http.middlewares.<name>.plugin.distributedratelimit
---

# DistributedRateLimit

Configuration holds the DistributedRateLimit middleware configuration.
