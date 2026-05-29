---
schema_version: 2
kind: concept
name: ClientAuth
id: concept.clientauth
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tls.go#L34
summary: ClientAuth defines the parameters of the client authentication part of the TLS connection, if any.
fields:
  - name: caFiles
    go_name: CAFiles
    type: array
    items: string
    go_type: '[]types.FileOrContent'
  - name: clientAuthType
    go_name: ClientAuthType
    type: string
    go_type: string
    description: 'ClientAuthType defines the client authentication type to apply. The available values are: "NoClientCert", "RequestClientCert", "VerifyClientCertIfGiven" and "RequireAndVerifyClientCert".'
---

# ClientAuth

ClientAuth defines the parameters of the client authentication part of the TLS connection, if any.
