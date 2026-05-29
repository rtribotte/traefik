---
schema_version: 2
kind: concept
name: UplinkHealthCheck
id: hub.concept.uplinkhealthcheck
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/dynamic/ext/ext.go#L55
summary: 'UplinkHealthCheck mirrors Traefik''s ServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go'
fields:
  - name: scheme
    go_name: Scheme
    type: string
    go_type: string
  - name: mode
    go_name: Mode
    type: string
    go_type: string
  - name: path
    go_name: Path
    type: string
    go_type: string
  - name: method
    go_name: Method
    type: string
    go_type: string
  - name: status
    go_name: Status
    type: integer
    go_type: int
  - name: port
    go_name: Port
    type: integer
    go_type: int
  - name: interval
    go_name: Interval
    type: duration
    go_type: ptypes.Duration
  - name: unhealthyInterval
    go_name: UnhealthyInterval
    type: duration
    go_type: '*ptypes.Duration'
  - name: timeout
    go_name: Timeout
    type: duration
    go_type: ptypes.Duration
  - name: hostname
    go_name: Hostname
    type: string
    go_type: string
  - name: followRedirects
    go_name: FollowRedirects
    type: boolean
    go_type: '*bool'
  - name: headers
    go_name: Headers
    type: object
    items: string
    go_type: map[string]string
---

# UplinkHealthCheck

UplinkHealthCheck mirrors Traefik's ServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
