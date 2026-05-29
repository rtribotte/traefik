---
schema_version: 2
kind: middleware-http
name: StripPrefix
id: http.middlewares.stripprefix
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L788
summary: 'StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/stripprefix/'
fields:
  - name: prefixes
    go_name: Prefixes
    type: array
    items: string
    go_type: '[]string'
    description: Prefixes defines the prefixes to strip from the request URL.
  - name: forceSlash
    go_name: ForceSlash
    type: boolean
    go_type: '*bool'
representations:
  yaml_path: http.middlewares.<name>.stripPrefix
  toml_path: http.middlewares.<name>.stripPrefix
  label_prefix: traefik.http.middlewares.<name>.stripprefix
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.stripPrefix
---

# StripPrefix

StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/stripprefix/
