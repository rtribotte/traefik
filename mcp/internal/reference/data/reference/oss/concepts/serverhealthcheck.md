---
schema_version: 2
kind: concept
name: ServerHealthCheck
id: concept.serverhealthcheck
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L483
summary: ServerHealthCheck holds the HealthCheck configuration.
fields:
  - name: scheme
    go_name: Scheme
    type: string
    go_type: string
    description: Scheme replaces the server URL scheme for the health check endpoint.
  - name: mode
    go_name: Mode
    type: string
    go_type: string
    default: http
    description: 'Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http'
  - name: path
    go_name: Path
    type: string
    go_type: string
    description: Path defines the server URL path for the health check endpoint.
  - name: method
    go_name: Method
    type: string
    go_type: string
    description: Method defines the healthcheck method.
  - name: status
    go_name: Status
    type: integer
    go_type: int
    description: Status defines the expected HTTP status code of the response to the health check request.
  - name: port
    go_name: Port
    type: integer
    go_type: int
    description: Port defines the server URL port for the health check endpoint.
  - name: interval
    go_name: Interval
    type: duration
    go_type: ptypes.Duration
    default: 30s
    description: 'Interval defines the frequency of the health check calls for healthy targets. Default: 30s'
  - name: unhealthyInterval
    go_name: UnhealthyInterval
    type: duration
    go_type: '*ptypes.Duration'
    description: 'UnhealthyInterval defines the frequency of the health check calls for unhealthy targets. When UnhealthyInterval is not defined, it defaults to the Interval value. Default: 30s'
  - name: timeout
    go_name: Timeout
    type: duration
    go_type: ptypes.Duration
    default: 5s
    description: 'Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s'
  - name: hostname
    go_name: Hostname
    type: string
    go_type: string
    description: Hostname defines the value of hostname in the Host header of the health check request.
  - name: followRedirects
    go_name: FollowRedirects
    type: boolean
    go_type: '*bool'
    default: true
    description: 'FollowRedirects defines whether redirects should be followed during the health check calls. Default: true'
  - name: headers
    go_name: Headers
    type: object
    items: string
    go_type: map[string]string
    description: Headers defines custom headers to be sent to the health check endpoint.
---

# ServerHealthCheck

ServerHealthCheck holds the HealthCheck configuration.
