---
schema_version: 2
kind: middleware-http
name: ContentType
id: http.middlewares.contenttype
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L80
summary: ContentType holds the content-type middleware configuration. This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
fields:
  - name: autoDetect
    go_name: AutoDetect
    type: boolean
    go_type: '*bool'
    description: AutoDetect specifies whether to let the `Content-Type` header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response.
representations:
  yaml_path: http.middlewares.<name>.contentType
  toml_path: http.middlewares.<name>.contentType
  label_prefix: traefik.http.middlewares.<name>.contenttype
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.contentType
---

# ContentType

ContentType holds the content-type middleware configuration. This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
