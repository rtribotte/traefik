---
schema_version: 2
kind: concept
name: GeneratedCert
id: concept.generatedcert
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tls.go#L76
summary: GeneratedCert defines the default generated certificate configuration.
fields:
  - name: resolver
    go_name: Resolver
    type: string
    go_type: string
    description: Resolver is the name of the resolver that will be used to issue the DefaultCertificate.
  - name: domain
    go_name: Domain
    type: object
    go_type: '*types.Domain'
    type_ref: oss:Domain
    description: Domain is the domain definition for the DefaultCertificate.
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

# GeneratedCert

GeneratedCert defines the default generated certificate configuration.
