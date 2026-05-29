---
schema_version: 2
kind: middleware-http
name: RedirectScheme
id: http.middlewares.redirectscheme
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L715
summary: 'RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/redirectscheme/'
fields:
  - name: scheme
    go_name: Scheme
    type: string
    go_type: string
    description: Scheme defines the scheme of the new URL.
  - name: port
    go_name: Port
    type: string
    go_type: string
    description: Port defines the port of the new URL.
  - name: permanent
    go_name: Permanent
    type: boolean
    go_type: bool
    description: Permanent defines whether the redirection is permanent. For HTTP GET requests a 301 is returned, otherwise a 308 is returned.
  - name: ForcePermanentRedirect
    go_name: ForcePermanentRedirect
    type: boolean
    go_type: bool
    description: ForcePermanentRedirect is an internal field (not exposed in configuration). When set to true, this forces the use of permanent redirects 308, regardless of the request method. Used by the provider ingress-nginx.
representations:
  yaml_path: http.middlewares.<name>.redirectScheme
  toml_path: http.middlewares.<name>.redirectScheme
  label_prefix: traefik.http.middlewares.<name>.redirectscheme
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.redirectScheme
---

# RedirectScheme

RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/redirectscheme/
