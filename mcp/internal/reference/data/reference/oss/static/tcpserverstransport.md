---
schema_version: 2
kind: static-section
name: TCPServersTransport
id: static.tcpserverstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L161
summary: TCPServersTransport options to configure communication between Traefik and the servers.
fields:
  - name: dialKeepAlive
    go_name: DialKeepAlive
    type: duration
    go_type: ptypes.Duration
    description: DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
  - name: dialTimeout
    go_name: DialTimeout
    type: duration
    go_type: ptypes.Duration
    description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
  - name: terminationDelay
    go_name: TerminationDelay
    type: duration
    go_type: ptypes.Duration
    description: TerminationDelay, corresponds to the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).
  - name: tls
    go_name: TLS
    type: object
    go_type: '*TLSClientConfig'
    type_ref: oss:static.TLSClientConfig
    description: TLS defines the configuration used to secure the connection to the authentication server.
    fields:
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
representations:
  yaml_path: tcpServersTransport
  toml_path: tcpServersTransport
---

# TCPServersTransport

TCPServersTransport options to configure communication between Traefik and the servers.
