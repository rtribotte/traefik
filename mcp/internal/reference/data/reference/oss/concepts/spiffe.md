---
schema_version: 2
kind: concept
name: Spiffe
id: concept.spiffe
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L546
summary: Spiffe holds the SPIFFE configuration.
fields:
  - name: ids
    go_name: IDs
    type: array
    items: string
    go_type: '[]string'
    description: IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).
  - name: trustDomain
    go_name: TrustDomain
    type: string
    go_type: string
    description: TrustDomain defines the allowed SPIFFE trust domain.
---

# Spiffe

Spiffe holds the SPIFFE configuration.
