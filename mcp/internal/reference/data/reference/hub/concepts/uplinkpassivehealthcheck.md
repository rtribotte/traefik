---
schema_version: 2
kind: concept
name: UplinkPassiveHealthCheck
id: hub.concept.uplinkpassivehealthcheck
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/dynamic/ext/ext.go#L72
summary: 'UplinkPassiveHealthCheck mirrors Traefik''s PassiveServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go'
fields:
  - name: failureWindow
    go_name: FailureWindow
    type: duration
    go_type: ptypes.Duration
  - name: maxFailedAttempts
    go_name: MaxFailedAttempts
    type: integer
    go_type: int
---

# UplinkPassiveHealthCheck

UplinkPassiveHealthCheck mirrors Traefik's PassiveServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
