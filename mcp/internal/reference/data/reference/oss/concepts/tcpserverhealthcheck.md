---
schema_version: 2
kind: concept
name: TCPServerHealthCheck
id: concept.tcpserverhealthcheck
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L217
summary: TCPServerHealthCheck holds the HealthCheck configuration.
fields:
  - name: port
    go_name: Port
    type: integer
    go_type: int
    description: Port defines the port of a Kubernetes Service. This can be a reference to a named port.
  - name: send
    go_name: Send
    type: string
    go_type: string
  - name: expect
    go_name: Expect
    type: string
    go_type: string
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
    description: Timeout defines how much time the middleware is allowed to retry the request. The value of timeout should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
---

# TCPServerHealthCheck

TCPServerHealthCheck holds the HealthCheck configuration.
