---
schema_version: 2
kind: middleware-http
name: Buffering
id: http.middlewares.buffering
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L128
summary: 'Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/buffering/#maxrequestbodybytes'
fields:
  - name: maxRequestBodyBytes
    go_name: MaxRequestBodyBytes
    type: integer
    go_type: int64
    description: 'MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).'
  - name: memRequestBodyBytes
    go_name: MemRequestBodyBytes
    type: integer
    go_type: int64
    description: 'MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).'
  - name: maxResponseBodyBytes
    go_name: MaxResponseBodyBytes
    type: integer
    go_type: int64
    description: 'MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).'
  - name: memResponseBodyBytes
    go_name: MemResponseBodyBytes
    type: integer
    go_type: int64
    description: 'MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).'
  - name: retryExpression
    go_name: RetryExpression
    type: string
    go_type: string
    description: 'RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/buffering/#retryexpression'
  - name: disableRequestBuffer
    go_name: DisableRequestBuffer
    type: boolean
    go_type: bool
    description: Only configurable via code, not via configuration files.
  - name: disableResponseBuffer
    go_name: DisableResponseBuffer
    type: boolean
    go_type: bool
representations:
  yaml_path: http.middlewares.<name>.buffering
  toml_path: http.middlewares.<name>.buffering
  label_prefix: traefik.http.middlewares.<name>.buffering
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.buffering
---

# Buffering

Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/buffering/#maxrequestbodybytes
