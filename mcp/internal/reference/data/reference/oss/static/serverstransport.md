---
schema_version: 2
kind: static-section
name: ServersTransport
id: static.serverstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L146
summary: ServersTransport options to configure communication between Traefik and the servers.
fields:
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
  - name: maxIdleConnsPerHost
    go_name: MaxIdleConnsPerHost
    type: integer
    go_type: int
    description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
  - name: forwardingTimeouts
    go_name: ForwardingTimeouts
    type: object
    go_type: '*ForwardingTimeouts'
    type_ref: oss:static.ForwardingTimeouts
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
  - name: spiffe
    go_name: Spiffe
    type: object
    go_type: '*Spiffe'
    type_ref: oss:static.Spiffe
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
  yaml_path: serversTransport
  toml_path: serversTransport
---

# ServersTransport

ServersTransport options to configure communication between Traefik and the servers.
