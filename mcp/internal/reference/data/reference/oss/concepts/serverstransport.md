---
schema_version: 2
kind: concept
name: ServersTransport
id: concept.serverstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L528
summary: ServersTransport options to configure communication between Traefik and the servers.
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
    description: InsecureSkipVerify defines whether the server certificates should be validated.
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
  - name: cipherSuites
    go_name: CipherSuites
    type: array
    items: string
    go_type: '[]string'
    description: CipherSuites defines the cipher suites to use when contacting backend servers.
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
  - name: maxIdleConnsPerHost
    go_name: MaxIdleConnsPerHost
    type: integer
    go_type: int
    description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
  - name: forwardingTimeouts
    go_name: ForwardingTimeouts
    type: object
    go_type: '*ForwardingTimeouts'
    type_ref: oss:ForwardingTimeouts
    description: ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
    fields:
      - name: dialTimeout
        go_name: DialTimeout
        type: duration
        go_type: ptypes.Duration
        default: 30s
        description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
      - name: responseHeaderTimeout
        go_name: ResponseHeaderTimeout
        type: duration
        go_type: ptypes.Duration
        description: ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
      - name: idleConnTimeout
        go_name: IdleConnTimeout
        type: duration
        go_type: ptypes.Duration
        default: 1m30s
        description: IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.
      - name: readIdleTimeout
        go_name: ReadIdleTimeout
        type: duration
        go_type: ptypes.Duration
        description: ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.
      - name: pingTimeout
        go_name: PingTimeout
        type: duration
        go_type: ptypes.Duration
        default: 15s
        description: PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.
  - name: disableHTTP2
    go_name: DisableHTTP2
    type: boolean
    go_type: bool
    description: DisableHTTP2 disables HTTP/2 for connections with backend servers.
  - name: peerCertURI
    go_name: PeerCertURI
    type: string
    go_type: string
    description: PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
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

# ServersTransport

ServersTransport options to configure communication between Traefik and the servers.
