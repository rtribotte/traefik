---
schema_version: 2
kind: crd
name: ServersTransport
id: crd.serverstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_serverstransports.yaml
summary: ServersTransport is the CRD implementation of a ServersTransport.
fields:
  - name: certificatesSecrets
    type: object
    description: CertificatesSecrets defines a list of secret storing client certificates for mTLS.
  - name: cipherSuites
    type: object
    description: CipherSuites defines the cipher suites to use when contacting backend servers.
  - name: disableHTTP2
    type: object
    description: DisableHTTP2 disables HTTP/2 for connections with backend servers.
  - name: forwardingTimeouts
    type: object
    description: ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
  - name: insecureSkipVerify
    type: object
    description: InsecureSkipVerify defines whether the server certificates should be validated.
  - name: maxIdleConnsPerHost
    type: object
    description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
  - name: maxVersion
    type: object
    description: MaxVersion defines the maximum TLS version to use when contacting backend servers.
  - name: minVersion
    type: object
    description: MinVersion defines the minimum TLS version to use when contacting backend servers.
  - name: peerCertURI
    type: object
    description: PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
  - name: rootCAs
    type: object
    description: RootCAs defines a list of CA certificate Secrets or ConfigMaps used to validate server certificates.
  - name: rootCAsSecrets
    type: object
    description: RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.
  - name: serverName
    type: object
    description: ServerName defines the server name used to contact the server.
  - name: spiffe
    type: object
    description: Spiffe defines the SPIFFE configuration.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: ServersTransport
    spec_path: .spec
---

# ServersTransport

ServersTransport is the CRD implementation of a ServersTransport.
If no serversTransport is specified, the default@internal will be used.
The default@internal serversTransport is created from the static configuration.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/load-balancing/serverstransport/
