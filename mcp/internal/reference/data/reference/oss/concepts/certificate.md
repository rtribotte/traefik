---
schema_version: 2
kind: concept
name: Certificate
id: concept.certificate
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/certificate.go#L72
summary: Certificate holds a SSL cert/key pair Certs and Key could be either a file path, or the file content itself.
fields:
  - name: certFile
    go_name: CertFile
    type: string
    go_type: types.FileOrContent
  - name: keyFile
    go_name: KeyFile
    type: string
    go_type: types.FileOrContent
---

# Certificate

Certificate holds a SSL cert/key pair Certs and Key could be either a file path, or the file content itself.
