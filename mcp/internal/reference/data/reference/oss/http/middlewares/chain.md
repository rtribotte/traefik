---
schema_version: 2
kind: middleware-http
name: Chain
id: http.middlewares.chain
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L157
summary: Chain holds the chain middleware configuration. This middleware enables to define reusable combinations of other pieces of middleware.
fields:
  - name: middlewares
    go_name: Middlewares
    type: array
    items: string
    go_type: '[]string'
    description: Middlewares is the list of middleware names which composes the chain.
representations:
  yaml_path: http.middlewares.<name>.chain
  toml_path: http.middlewares.<name>.chain
  label_prefix: traefik.http.middlewares.<name>.chain
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.chain
---

# Chain

Chain holds the chain middleware configuration. This middleware enables to define reusable combinations of other pieces of middleware.
