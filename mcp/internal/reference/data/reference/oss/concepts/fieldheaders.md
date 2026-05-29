---
schema_version: 2
kind: concept
name: FieldHeaders
id: concept.fieldheaders
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L89
summary: FieldHeaders holds configuration for access log headers.
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

# FieldHeaders

FieldHeaders holds configuration for access log headers.
