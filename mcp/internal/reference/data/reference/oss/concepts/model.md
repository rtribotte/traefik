---
schema_version: 2
kind: concept
name: Model
id: concept.model
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L52
summary: Model holds model configuration.
fields:
  - name: middlewares
    go_name: Middlewares
    type: array
    items: string
    go_type: '[]string'
    description: Middlewares is the list of MiddlewareRef which composes the chain.
  - name: tls
    go_name: TLS
    type: object
    go_type: '*RouterTLSConfig'
    type_ref: oss:RouterTLSConfig
    description: TLS defines the configuration used to secure the connection to the authentication server.
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
  - name: observability
    go_name: Observability
    type: object
    go_type: RouterObservabilityConfig
    type_ref: oss:RouterObservabilityConfig
    description: 'Observability defines the observability configuration for a router. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/observability/'
    fields:
      - name: accessLogs
        go_name: AccessLogs
        type: boolean
        go_type: '*bool'
        description: AccessLogs enables access logs for this router.
      - name: metrics
        go_name: Metrics
        type: boolean
        go_type: '*bool'
        description: Metrics enables metrics for this router.
      - name: tracing
        go_name: Tracing
        type: boolean
        go_type: '*bool'
        description: Tracing enables tracing for this router.
      - name: traceVerbosity
        go_name: TraceVerbosity
        type: string
        go_type: otypes.TracingVerbosity
        default: minimal
        description: TraceVerbosity defines the verbosity level of the tracing for this router.
---

# Model

Model holds model configuration.
