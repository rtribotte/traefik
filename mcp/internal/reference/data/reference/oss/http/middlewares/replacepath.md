---
schema_version: 2
kind: middleware-http
name: ReplacePath
id: http.middlewares.replacepath
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L734
summary: 'ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/replacepath/'
fields:
  - name: path
    go_name: Path
    type: string
    go_type: string
    description: Path defines the path to use as replacement in the request URL.
representations:
  yaml_path: http.middlewares.<name>.replacePath
  toml_path: http.middlewares.<name>.replacePath
  label_prefix: traefik.http.middlewares.<name>.replacepath
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.replacePath
---

# ReplacePath

ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/replacepath/
