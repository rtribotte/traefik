---
schema_version: 2
kind: concept
name: ForwardedHeaders
id: concept.forwardedheaders
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L152
summary: ForwardedHeaders Trust client forwarding headers.
fields:
  - name: insecure
    go_name: Insecure
    type: boolean
    go_type: bool
  - name: trustedIPs
    go_name: TrustedIPs
    type: array
    items: string
    go_type: '[]string'
  - name: connection
    go_name: Connection
    type: array
    items: string
    go_type: '[]string'
  - name: notAppendXForwardedFor
    go_name: NotAppendXForwardedFor
    type: boolean
    go_type: bool
---

# ForwardedHeaders

ForwardedHeaders Trust client forwarding headers.
