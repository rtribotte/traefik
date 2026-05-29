---
schema_version: 2
kind: concept
name: EntryPointsTransport
id: concept.entrypointstransport
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L169
summary: EntryPointsTransport configures communication between clients and Traefik.
fields:
  - name: lifeCycle
    go_name: LifeCycle
    type: object
    go_type: '*LifeCycle'
    type_ref: oss:LifeCycle
    default:
      graceTimeOut: 10s
    fields:
      - name: requestAcceptGraceTimeout
        go_name: RequestAcceptGraceTimeout
        type: duration
        go_type: ptypes.Duration
      - name: graceTimeOut
        go_name: GraceTimeOut
        type: duration
        go_type: ptypes.Duration
        default: 10s
  - name: respondingTimeouts
    go_name: RespondingTimeouts
    type: object
    go_type: '*RespondingTimeouts'
    type_ref: oss:RespondingTimeouts
    default:
      idleTimeout: 3m0s
      readTimeout: 1m0s
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
  - name: keepAliveMaxTime
    go_name: KeepAliveMaxTime
    type: duration
    go_type: ptypes.Duration
  - name: keepAliveMaxRequests
    go_name: KeepAliveMaxRequests
    type: integer
    go_type: int
---

# EntryPointsTransport

EntryPointsTransport configures communication between clients and Traefik.
