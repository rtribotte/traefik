---
schema_version: 2
kind: middleware-http
name: ReplacePathRegex
id: http.middlewares.replacepathregex
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L744
summary: 'ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/replacepathregex/'
fields:
  - name: regex
    go_name: Regex
    type: string
    go_type: string
    description: Regex defines the regular expression used to match and capture the path from the request URL.
  - name: replacement
    go_name: Replacement
    type: string
    go_type: string
    description: Replacement defines the replacement path format, which can include captured variables.
representations:
  yaml_path: http.middlewares.<name>.replacePathRegex
  toml_path: http.middlewares.<name>.replacePathRegex
  label_prefix: traefik.http.middlewares.<name>.replacepathregex
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.replacePathRegex
---

# ReplacePathRegex

ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/replacepathregex/
