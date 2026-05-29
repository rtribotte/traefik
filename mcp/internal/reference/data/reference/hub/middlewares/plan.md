---
schema_version: 2
kind: middleware-hub
name: Plan
id: hub.middlewares.plan
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/plan/config.go#L15
summary: Configuration holds the configuration of the plan middleware.
fields:
  - name: apiId
    go_name: APIID
    type: string
    go_type: string
  - name: availableOperations
    go_name: AvailableOperations
    type: object
    go_type: '*resource.OperationMatchers'
  - name: subscriptions
    go_name: Subscriptions
    type: array
    items: object
    go_type: '[]Subscription'
  - name: applications
    go_name: Applications
    type: object
    items: object
    go_type: map[string]struct{}
  - name: openApiSpec
    go_name: OpenAPISpec
    type: object
    go_type: '*OpenAPISpec'
  - name: validateRequestMethodAndPath
    go_name: ValidateRequestMethodAndPath
    type: boolean
    go_type: bool
  - name: validateRequestBodySchema
    go_name: ValidateRequestBodySchema
    type: boolean
    go_type: bool
representations:
  yaml_path: http.middlewares.<name>.plugin.plan
  toml_path: http.middlewares.<name>.plugin.plan
  label_prefix: traefik.http.middlewares.<name>.plugin.plan
---

# Plan

Configuration holds the configuration of the plan middleware.
