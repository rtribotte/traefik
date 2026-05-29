---
schema_version: 2
kind: concept
name: TCPMiddleware
id: concept.tcpmiddleware
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_middlewares.go#L6
summary: TCPMiddleware holds the TCPMiddleware configuration.
fields:
  - name: inFlightConn
    go_name: InFlightConn
    type: object
    go_type: '*TCPInFlightConn'
    type_ref: oss:TCPInFlightConn
    description: InFlightConn defines the InFlightConn middleware configuration.
    fields:
      - name: amount
        go_name: Amount
        type: integer
        go_type: int64
        description: Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.
  - name: ipWhiteList
    go_name: IPWhiteList
    type: object
    go_type: '*TCPIPWhiteList'
    type_ref: oss:TCPIPWhiteList
    description: 'IPWhiteList defines the IPWhiteList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipwhitelist/'
    fields:
      - name: sourceRange
        go_name: SourceRange
        type: array
        items: string
        go_type: '[]string'
        description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
  - name: ipAllowList
    go_name: IPAllowList
    type: object
    go_type: '*TCPIPAllowList'
    type_ref: oss:TCPIPAllowList
    description: 'IPAllowList defines the IPAllowList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipallowlist/'
    fields:
      - name: sourceRange
        go_name: SourceRange
        type: array
        items: string
        go_type: '[]string'
        description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
variants:
  - inFlightConn
  - ipWhiteList
  - ipAllowList
---

# TCPMiddleware

TCPMiddleware holds the TCPMiddleware configuration.
