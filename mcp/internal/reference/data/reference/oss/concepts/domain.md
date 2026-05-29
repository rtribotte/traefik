---
schema_version: 2
kind: concept
name: Domain
id: concept.domain
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/types/domains.go#L10
summary: Domain holds a domain name with SANs.
fields:
  - name: main
    go_name: Main
    type: string
    go_type: string
    description: Main defines the main domain name.
  - name: sans
    go_name: SANs
    type: array
    items: string
    go_type: '[]string'
    description: SANs defines the subject alternative domain names.
---

# Domain

Domain holds a domain name with SANs.
