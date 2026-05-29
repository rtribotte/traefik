---
schema_version: 2
kind: concept
name: RespondingTimeouts
id: concept.respondingtimeouts
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L200
summary: RespondingTimeouts contains timeout configurations for incoming requests to the Traefik instance.
fields:
  - name: readTimeout
    go_name: ReadTimeout
    type: duration
    go_type: ptypes.Duration
    default: 1m0s
    description: ReadTimeout defines the timeout for socket read operations. Default value is 3 seconds.
  - name: writeTimeout
    go_name: WriteTimeout
    type: duration
    go_type: ptypes.Duration
    description: WriteTimeout defines the timeout for socket write operations. Default value is 3 seconds.
  - name: idleTimeout
    go_name: IdleTimeout
    type: duration
    go_type: ptypes.Duration
    default: 3m0s
---

# RespondingTimeouts

RespondingTimeouts contains timeout configurations for incoming requests to the Traefik instance.
