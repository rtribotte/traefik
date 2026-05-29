---
schema_version: 2
kind: crd
name: TLSStore
id: crd.tlsstore
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_tlsstores.yaml
summary: TLSStore is the CRD implementation of a Traefik TLS Store.
fields:
  - name: certificates
    type: object
    description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
  - name: defaultCertificate
    type: object
    description: DefaultCertificate defines the default certificate configuration.
  - name: defaultGeneratedCert
    type: object
    description: DefaultGeneratedCert defines the default generated certificate configuration.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: TLSStore
    spec_path: .spec
---

# TLSStore

TLSStore is the CRD implementation of a Traefik TLS Store.
For the time being, only the TLSStore named default is supported.
This means that you cannot have two stores that are named default in different Kubernetes namespaces.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#certificates-stores
