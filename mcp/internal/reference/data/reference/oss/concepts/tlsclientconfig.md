---
schema_version: 2
kind: concept
name: TLSClientConfig
id: concept.tlsclientconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L205
summary: TLSClientConfig options to configure TLS communication between Traefik and the servers.
fields:
  - name: serverName
    go_name: ServerName
    type: string
    go_type: string
    description: ServerName defines the server name used to contact the server.
  - name: insecureSkipVerify
    go_name: InsecureSkipVerify
    type: boolean
    go_type: bool
    description: InsecureSkipVerify disables TLS certificate verification.
  - name: rootCAs
    go_name: RootCAs
    type: array
    items: string
    go_type: '[]types.FileOrContent'
    description: RootCAs defines a list of CA certificate Secrets or ConfigMaps used to validate server certificates.
  - name: certificates
    go_name: Certificates
    type: array
    items: object
    go_type: traefiktls.Certificates
    description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
  - name: peerCertURI
    go_name: PeerCertURI
    type: string
    go_type: string
    description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
  - name: spiffe
    go_name: Spiffe
    type: object
    go_type: '*Spiffe'
    type_ref: oss:Spiffe
    description: Spiffe defines the SPIFFE configuration.
    fields:
      - name: ids
        go_name: IDs
        type: array
        items: string
        go_type: '[]string'
        description: IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).
      - name: trustDomain
        go_name: TrustDomain
        type: string
        go_type: string
        description: TrustDomain defines the allowed SPIFFE trust domain.
---

# TLSClientConfig

TLSClientConfig options to configure TLS communication between Traefik and the servers.
