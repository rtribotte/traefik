---
schema_version: 2
kind: tls-option
name: TLS
id: tls
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/config.go#L33
summary: TLSConfiguration contains all the configuration parameters of a TLS connection.
fields:
  - name: certificates
    go_name: Certificates
    type: array
    items: object
    go_type: '[]*tls.CertAndStores'
    type_ref: oss:CertAndStores
    description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
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
  - name: options
    go_name: Options
    type: object
    items: object
    go_type: map[string]tls.Options
    type_ref: oss:Options
    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
    fields:
      - name: minVersion
        go_name: MinVersion
        type: string
        go_type: string
        description: MinVersion defines the minimum TLS version to use when contacting backend servers.
      - name: maxVersion
        go_name: MaxVersion
        type: string
        go_type: string
        description: MaxVersion defines the maximum TLS version to use when contacting backend servers.
      - name: cipherSuites
        go_name: CipherSuites
        type: array
        items: string
        go_type: '[]string'
        default:
          - TLS_AES_128_GCM_SHA256
          - TLS_AES_256_GCM_SHA384
          - TLS_CHACHA20_POLY1305_SHA256
          - TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
          - TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
          - TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
          - TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
          - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
          - TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
          - TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
          - TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
          - TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
          - TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
        description: CipherSuites defines the cipher suites to use when contacting backend servers.
      - name: curvePreferences
        go_name: CurvePreferences
        type: array
        items: string
        go_type: '[]string'
        description: 'CurvePreferences defines the preferred elliptic curves. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#curve-preferences'
      - name: clientAuth
        go_name: ClientAuth
        type: object
        go_type: ClientAuth
        type_ref: oss:ClientAuth
        description: ClientAuth defines the server's policy for TLS Client Authentication.
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
      - name: sniStrict
        go_name: SniStrict
        type: boolean
        go_type: bool
        description: SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.
      - name: alpnProtocols
        go_name: ALPNProtocols
        type: array
        items: string
        go_type: '[]string'
        default:
          - h2
          - http/1.1
          - acme-tls/1
        description: 'ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#alpn-protocols'
      - name: disableSessionTickets
        go_name: DisableSessionTickets
        type: boolean
        go_type: bool
        description: DisableSessionTickets disables TLS session resumption via session tickets.
      - name: preferServerCipherSuites
        go_name: PreferServerCipherSuites
        type: boolean
        go_type: '*bool'
        description: PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's. It is enabled automatically when minVersion or maxVersion is set.
  - name: stores
    go_name: Stores
    type: object
    items: object
    go_type: map[string]tls.Store
    type_ref: oss:Store
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
representations:
  yaml_path: tls
  toml_path: tls
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: TLSOption
    spec_path: .spec
---

# TLS

TLSConfiguration contains all the configuration parameters of a TLS connection.
