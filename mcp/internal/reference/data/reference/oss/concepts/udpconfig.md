---
schema_version: 2
kind: concept
name: UDPConfig
id: concept.udpconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L185
summary: UDPConfig is the UDP configuration of an entry point.
fields:
  - name: timeout
    go_name: Timeout
    type: duration
    go_type: ptypes.Duration
    default: 3s
    description: Timeout defines how much time the middleware is allowed to retry the request. The value of timeout should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
---

# UDPConfig

UDPConfig is the UDP configuration of an entry point.
