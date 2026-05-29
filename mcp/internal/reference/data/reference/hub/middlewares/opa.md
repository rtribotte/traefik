---
schema_version: 2
kind: middleware-hub
name: OPA
id: hub.middlewares.opa
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/opa/config.go#L15
summary: Configuration holds the OPA middleware configuration.
fields:
  - name: policy
    go_name: Policy
    type: object
    go_type: types.FileOrContent
    description: Path or content of the policy.
  - name: bundlePath
    go_name: BundlePath
    type: string
    go_type: string
  - name: allow
    go_name: Allow
    type: string
    go_type: string
    description: Allow holds an expression that defines if the request is authorized.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: 'ForwardHeaders defines headers that should be added to the request and populated with the result of a given expression, for example: {"HeaderName":"expressionToEvaluate"}'
representations:
  yaml_path: http.middlewares.<name>.plugin.opa
  toml_path: http.middlewares.<name>.plugin.opa
  label_prefix: traefik.http.middlewares.<name>.plugin.opa
---

# OPA

Configuration holds the OPA middleware configuration.
