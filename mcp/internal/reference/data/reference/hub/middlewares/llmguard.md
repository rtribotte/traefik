---
schema_version: 2
kind: middleware-hub
name: LLMGuard
id: hub.middlewares.llmguard
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/llmguard/config.go#L15
summary: Config is the unified configuration for an LLM guard middleware.
fields:
  - name: clientConfig
    go_name: ClientConfig
    type: object
    go_type: httpclient.Config
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: clientRequestFormat
    go_name: ClientRequestFormat
    type: object
    go_type: aiformat.ClientRequestFormat
    description: 'ClientRequestFormat defines the format used by the upstream client. Valid values: "ccr", "custom", "responsesAPI". Default: "custom".'
  - name: format
    go_name: Format
    type: object
    go_type: GuardFormat
  - name: request
    go_name: Request
    type: object
    go_type: '*RequestConfig'
  - name: response
    go_name: Response
    type: object
    go_type: '*ResponseConfig'
  - name: model
    go_name: Model
    type: string
    go_type: string
  - name: params
    go_name: Params
    type: object
    go_type: '*Params'
representations:
  yaml_path: http.middlewares.<name>.plugin.llmguard
  toml_path: http.middlewares.<name>.plugin.llmguard
  label_prefix: traefik.http.middlewares.<name>.plugin.llmguard
---

# LLMGuard

Config is the unified configuration for an LLM guard middleware.
