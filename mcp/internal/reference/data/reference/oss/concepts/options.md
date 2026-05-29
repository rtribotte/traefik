---
schema_version: 2
kind: concept
name: Options
id: concept.options
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/tls/tls.go#L44
summary: Options configures TLS for an entry point.
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
---

# Options

Options configures TLS for an entry point.
