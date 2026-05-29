---
schema_version: 2
kind: concept
name: TCPServersTransport
id: concept.tcpserverstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L181
summary: TCPServersTransport options to configure communication between Traefik and the servers.
fields:
  - name: dialKeepAlive
    go_name: DialKeepAlive
    type: duration
    go_type: ptypes.Duration
    default: 15s
    description: DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
  - name: dialTimeout
    go_name: DialTimeout
    type: duration
    go_type: ptypes.Duration
    default: 30s
    description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
  - name: proxyProtocol
    go_name: ProxyProtocol
    type: object
    go_type: '*ProxyProtocol'
    type_ref: oss:ProxyProtocol
    description: ProxyProtocol holds the PROXY Protocol configuration.
    fields:
      - name: version
        go_name: Version
        type: integer
        go_type: int
        default: 2
        description: Version defines the PROXY Protocol version to use.
  - name: terminationDelay
    go_name: TerminationDelay
    type: duration
    go_type: ptypes.Duration
    default: 100ms
    description: TerminationDelay, corresponds to the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).
  - name: tls
    go_name: TLS
    type: object
    go_type: '*TLSClientConfig'
    type_ref: oss:TLSClientConfig
    description: TLS defines the configuration used to secure the connection to the authentication server.
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

# TCPServersTransport

TCPServersTransport options to configure communication between Traefik and the servers.
