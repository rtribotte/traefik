---
schema_version: 2
kind: middleware-http
name: AddPrefix
id: http.middlewares.addprefix
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L93
summary: 'AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/addprefix/'
fields:
  - name: prefix
    go_name: Prefix
    type: string
    go_type: string
    description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
representations:
  yaml_path: http.middlewares.<name>.addPrefix
  toml_path: http.middlewares.<name>.addPrefix
  label_prefix: traefik.http.middlewares.<name>.addprefix
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.addPrefix
---

# AddPrefix

AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/addprefix/
