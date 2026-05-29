---
schema_version: 2
kind: concept
name: Store
id: concept.store
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tls.go#L68
summary: Store holds the options for a given Store.
fields:
  - name: defaultCertificate
    go_name: DefaultCertificate
    type: object
    go_type: '*Certificate'
    type_ref: oss:Certificate
    description: DefaultCertificate defines the default certificate configuration.
    fields:
      - name: certFile
        go_name: CertFile
        type: string
        go_type: types.FileOrContent
      - name: keyFile
        go_name: KeyFile
        type: string
        go_type: types.FileOrContent
  - name: defaultGeneratedCert
    go_name: DefaultGeneratedCert
    type: object
    go_type: '*GeneratedCert'
    type_ref: oss:GeneratedCert
    description: DefaultGeneratedCert defines the default generated certificate configuration.
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

# Store

Store holds the options for a given Store.
