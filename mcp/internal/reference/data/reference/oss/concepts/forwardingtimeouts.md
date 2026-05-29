---
schema_version: 2
kind: concept
name: ForwardingTimeouts
id: concept.forwardingtimeouts
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L556
summary: ForwardingTimeouts contains timeout configurations for forwarding requests to the backend servers.
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
---

# ForwardingTimeouts

ForwardingTimeouts contains timeout configurations for forwarding requests to the backend servers.
