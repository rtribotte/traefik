---
schema_version: 2
kind: concept
name: ClientTLS
id: concept.clienttls
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L335
summary: 'ClientTLS holds TLS specific configurations as client CA, Cert and Key can be either path or file contents. TODO: remove this struct when CAOptional option will be removed.'
fields:
  - name: ca
    go_name: CA
    type: string
    go_type: string
  - name: cert
    go_name: Cert
    type: string
    go_type: string
  - name: key
    go_name: Key
    type: string
    go_type: string
  - name: insecureSkipVerify
    go_name: InsecureSkipVerify
    type: boolean
    go_type: bool
    description: InsecureSkipVerify defines whether the server certificates should be validated.
  - name: caOptional
    go_name: CAOptional
    type: boolean
    go_type: '*bool'
---

# ClientTLS

ClientTLS holds TLS specific configurations as client CA, Cert and Key can be either path or file contents. TODO: remove this struct when CAOptional option will be removed.
