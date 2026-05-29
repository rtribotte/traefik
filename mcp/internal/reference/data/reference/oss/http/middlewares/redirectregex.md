---
schema_version: 2
kind: middleware-http
name: RedirectRegex
id: http.middlewares.redirectregex
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L698
summary: 'RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/redirectregex/#regex'
fields:
  - name: regex
    go_name: Regex
    type: string
    go_type: string
    description: Regex defines the regex used to match and capture elements from the request URL.
  - name: replacement
    go_name: Replacement
    type: string
    go_type: string
    description: Replacement defines how to modify the URL to have the new target URL.
  - name: permanent
    go_name: Permanent
    type: boolean
    go_type: bool
    description: Permanent defines whether the redirection is permanent (308).
  - name: StatusCode
    go_name: StatusCode
    type: integer
    go_type: '*int'
    description: StatusCode is for supporting the NGINX annotations related to redirect.
representations:
  yaml_path: http.middlewares.<name>.redirectRegex
  toml_path: http.middlewares.<name>.redirectRegex
  label_prefix: traefik.http.middlewares.<name>.redirectregex
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.redirectRegex
---

# RedirectRegex

RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/redirectregex/#regex
