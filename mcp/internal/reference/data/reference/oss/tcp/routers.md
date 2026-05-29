---
schema_version: 2
kind: router-tcp
name: TCPRouter
id: tcp.routers
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L73
summary: TCPRouter holds the router configuration.
fields:
  - name: entryPoints
    go_name: EntryPoints
    type: array
    items: string
    go_type: '[]string'
    description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
  - name: middlewares
    go_name: Middlewares
    type: array
    items: string
    go_type: '[]string'
    description: Middlewares is the list of MiddlewareRef which composes the chain.
  - name: service
    go_name: Service
    type: string
    go_type: string
    description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
  - name: rule
    go_name: Rule
    type: string
    go_type: string
  - name: ruleSyntax
    go_name: RuleSyntax
    type: string
    go_type: string
  - name: priority
    go_name: Priority
    type: integer
    go_type: int
    description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
  - name: tls
    go_name: TLS
    type: object
    go_type: '*RouterTCPTLSConfig'
    type_ref: oss:RouterTCPTLSConfig
    description: TLS defines the configuration used to secure the connection to the authentication server.
    fields:
      - name: passthrough
        go_name: Passthrough
        type: boolean
        go_type: bool
        description: Passthrough defines whether a TLS router will terminate the TLS connection.
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
representations:
  yaml_path: tcp.routers.<name>
  toml_path: tcp.routers.<name>
  label_prefix: traefik.tcp.routers.<name>
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteTCP
    spec_path: .spec
---

# TCPRouter

TCPRouter holds the router configuration.
