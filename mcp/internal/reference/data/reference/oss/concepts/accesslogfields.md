---
schema_version: 2
kind: concept
name: AccessLogFields
id: concept.accesslogfields
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L95
summary: AccessLogFields holds configuration for access log fields.
fields:
  - name: defaultMode
    go_name: DefaultMode
    type: string
    go_type: string
    default: keep
  - name: names
    go_name: Names
    type: object
    items: string
    go_type: map[string]string
  - name: headers
    go_name: Headers
    type: object
    go_type: '*FieldHeaders'
    type_ref: oss:FieldHeaders
    default:
      defaultMode: drop
    description: Headers defines custom headers to be sent to the health check endpoint.
    fields:
      - name: defaultMode
        go_name: DefaultMode
        type: string
        go_type: string
      - name: names
        go_name: Names
        type: object
        items: string
        go_type: map[string]string
---

# AccessLogFields

AccessLogFields holds configuration for access log fields.
