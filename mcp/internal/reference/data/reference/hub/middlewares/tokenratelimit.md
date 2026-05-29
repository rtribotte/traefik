---
schema_version: 2
kind: middleware-hub
name: TokenRateLimit
id: hub.middlewares.tokenratelimit
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/tokenratelimit/config.go#L38
summary: Config holds the configuration for the token rate limit middleware.
fields:
  - name: store
    go_name: Store
    type: object
    go_type: StoreConfig
  - name: inputTokenLimit
    go_name: InputTokenLimit
    type: object
    go_type: '*TokenLimitConfig'
  - name: outputTokenLimit
    go_name: OutputTokenLimit
    type: object
    go_type: '*TokenLimitConfig'
  - name: totalTokenLimit
    go_name: TotalTokenLimit
    type: object
    go_type: '*TokenLimitConfig'
  - name: estimateStrategy
    go_name: EstimateStrategy
    type: object
    go_type: '*EstimateStrategy'
  - name: clientRequestFormat
    go_name: ClientRequestFormat
    type: object
    go_type: aiformat.ClientRequestFormat
  - name: onDenyResponse
    go_name: OnDenyResponse
    type: object
    go_type: '*aiformat.DenyResponse'
  - name: sourceCriterion
    go_name: SourceCriterion
    type: object
    go_type: '*dynamic.SourceCriterion'
    description: Source criterion for bucket naming.
representations:
  yaml_path: http.middlewares.<name>.plugin.tokenratelimit
  toml_path: http.middlewares.<name>.plugin.tokenratelimit
  label_prefix: traefik.http.middlewares.<name>.plugin.tokenratelimit
---

# TokenRateLimit

Config holds the configuration for the token rate limit middleware.
