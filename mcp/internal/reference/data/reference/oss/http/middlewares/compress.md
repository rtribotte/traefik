---
schema_version: 2
kind: middleware-http
name: Compress
id: http.middlewares.compress
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L192
summary: Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip, brotli, or zstd compression.
fields:
  - name: excludedContentTypes
    go_name: ExcludedContentTypes
    type: array
    items: string
    go_type: '[]string'
    description: ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing. `application/grpc` is always excluded.
  - name: includedContentTypes
    go_name: IncludedContentTypes
    type: array
    items: string
    go_type: '[]string'
    description: IncludedContentTypes defines the list of content types to compare the Content-Type header of the responses before compressing.
  - name: minResponseBodyBytes
    go_name: MinResponseBodyBytes
    type: integer
    go_type: int
    description: 'MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.'
  - name: encodings
    go_name: Encodings
    type: array
    items: string
    go_type: '[]string'
    default:
      - gzip
      - br
      - zstd
    description: Encodings defines the list of supported compression algorithms.
  - name: defaultEncoding
    go_name: DefaultEncoding
    type: string
    go_type: string
    description: DefaultEncoding specifies the default encoding if the `Accept-Encoding` header is not in the request or contains a wildcard (`*`).
representations:
  yaml_path: http.middlewares.<name>.compress
  toml_path: http.middlewares.<name>.compress
  label_prefix: traefik.http.middlewares.<name>.compress
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.compress
---

# Compress

Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip, brotli, or zstd compression.
