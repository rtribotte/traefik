---
schema_version: 2
kind: concept
name: LifeCycle
id: concept.lifecycle
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L226
summary: LifeCycle contains configurations relevant to the lifecycle (such as the shutdown phase) of Traefik.
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
---

# LifeCycle

LifeCycle contains configurations relevant to the lifecycle (such as the shutdown phase) of Traefik.
