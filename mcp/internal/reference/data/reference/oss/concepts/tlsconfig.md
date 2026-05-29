---
schema_version: 2
kind: concept
name: TLSConfig
id: concept.tlsconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L145
summary: TLSConfig is the default TLS configuration for all the routers associated to the concerned entry point.
fields:
  - name: options
    go_name: Options
    type: string
    go_type: string
    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
  - name: certResolver
    go_name: CertResolver
    type: string
    go_type: string
    description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
  - name: domains
    go_name: Domains
    type: array
    items: object
    go_type: '[]types.Domain'
    type_ref: oss:Domain
    description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
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

# TLSConfig

TLSConfig is the default TLS configuration for all the routers associated to the concerned entry point.
