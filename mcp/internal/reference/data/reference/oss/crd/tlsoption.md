---
schema_version: 2
kind: crd
name: TLSOption
id: crd.tlsoption
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_tlsoptions.yaml
summary: TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection.
fields:
  - name: alpnProtocols
    type: object
    description: 'ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#alpn-protocols'
  - name: cipherSuites
    type: object
    description: CipherSuites defines the cipher suites to use when contacting backend servers.
  - name: clientAuth
    type: object
    description: ClientAuth defines the server's policy for TLS Client Authentication.
  - name: curvePreferences
    type: object
    description: 'CurvePreferences defines the preferred elliptic curves. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#curve-preferences'
  - name: disableSessionTickets
    type: object
    description: DisableSessionTickets disables TLS session resumption via session tickets.
  - name: maxVersion
    type: object
    description: MaxVersion defines the maximum TLS version to use when contacting backend servers.
  - name: minVersion
    type: object
    description: MinVersion defines the minimum TLS version to use when contacting backend servers.
  - name: preferServerCipherSuites
    type: object
    description: PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's. It is enabled automatically when minVersion or maxVersion is set.
  - name: sniStrict
    type: object
    description: SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: TLSOption
    spec_path: .spec
---

# TLSOption

TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#tls-options
