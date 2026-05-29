---
schema_version: 2
kind: middleware-hub
name: HMAC
id: hub.middlewares.hmac
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/hmac/config.go#L9
summary: Configuration holds the HMAC Authentication Middleware configuration.
fields:
  - name: keys
    go_name: Keys
    type: array
    items: object
    go_type: '[]Key'
    description: 'TODO: redactor should handle slices of struct.'
  - name: validateDigest
    go_name: ValidateDigest
    type: boolean
    go_type: '*bool'
  - name: enforcedHeaders
    go_name: EnforcedHeaders
    type: array
    items: string
    go_type: '[]string'
representations:
  yaml_path: http.middlewares.<name>.plugin.hmac
  toml_path: http.middlewares.<name>.plugin.hmac
  label_prefix: traefik.http.middlewares.<name>.plugin.hmac
---

# HMAC

Configuration holds the HMAC Authentication Middleware configuration.
