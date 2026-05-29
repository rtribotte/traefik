---
schema_version: 2
kind: concept
name: CertAndStores
id: concept.certandstores
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tls.go#L86
summary: CertAndStores allows mapping a TLS certificate to a list of entry points.
fields:
  - name: certFile
    go_name: CertFile
    type: string
    go_type: types.FileOrContent
  - name: keyFile
    go_name: KeyFile
    type: string
    go_type: types.FileOrContent
  - name: stores
    go_name: Stores
    type: array
    items: string
    go_type: '[]string'
---

# CertAndStores

CertAndStores allows mapping a TLS certificate to a list of entry points.
