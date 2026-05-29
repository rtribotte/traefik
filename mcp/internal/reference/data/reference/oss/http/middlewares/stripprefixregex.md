---
schema_version: 2
kind: middleware-http
name: StripPrefixRegex
id: http.middlewares.stripprefixregex
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L803
summary: 'StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/stripprefixregex/'
fields:
  - name: regex
    go_name: Regex
    type: array
    items: string
    go_type: '[]string'
    description: Regex defines the regular expression to match the path prefix from the request URL.
representations:
  yaml_path: http.middlewares.<name>.stripPrefixRegex
  toml_path: http.middlewares.<name>.stripPrefixRegex
  label_prefix: traefik.http.middlewares.<name>.stripprefixregex
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.stripPrefixRegex
---

# StripPrefixRegex

StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/stripprefixregex/
