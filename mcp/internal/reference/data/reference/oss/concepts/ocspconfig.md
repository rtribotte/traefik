---
schema_version: 2
kind: concept
name: OCSPConfig
id: concept.ocspconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tlsmanager.go#L51
summary: OCSPConfig contains the OCSP configuration.
fields:
  - name: responderOverrides
    go_name: ResponderOverrides
    type: object
    items: string
    go_type: map[string]string
---

# OCSPConfig

OCSPConfig contains the OCSP configuration.
