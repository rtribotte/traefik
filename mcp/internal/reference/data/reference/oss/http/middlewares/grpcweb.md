---
schema_version: 2
kind: middleware-http
name: GrpcWeb
id: http.middlewares.grpcweb
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L70
summary: GrpcWeb holds the gRPC web middleware configuration. This middleware converts a gRPC web request to an HTTP/2 gRPC request.
fields:
  - name: allowOrigins
    go_name: AllowOrigins
    type: array
    items: string
    go_type: '[]string'
    description: AllowOrigins is a list of allowable origins. Can also be a wildcard origin "*".
representations:
  yaml_path: http.middlewares.<name>.grpcWeb
  toml_path: http.middlewares.<name>.grpcWeb
  label_prefix: traefik.http.middlewares.<name>.grpcweb
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.grpcWeb
---

# GrpcWeb

GrpcWeb holds the gRPC web middleware configuration. This middleware converts a gRPC web request to an HTTP/2 gRPC request.
