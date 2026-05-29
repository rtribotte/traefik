---
schema_version: 2
kind: concept
name: PassiveServerHealthCheck
id: concept.passiveserverhealthcheck
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L508
summary: Shared type referenced from configuration. See Go source for details.
fields:
  - name: failureWindow
    go_name: FailureWindow
    type: duration
    go_type: ptypes.Duration
    default: 10s
    description: FailureWindow defines the time window during which the failed attempts must occur for the server to be marked as unhealthy. It also defines for how long the server will be considered unhealthy.
  - name: maxFailedAttempts
    go_name: MaxFailedAttempts
    type: integer
    go_type: int
    default: 1
    description: MaxFailedAttempts is the number of consecutive failed attempts allowed within the failure window before marking the server as unhealthy.
---

# PassiveServerHealthCheck
