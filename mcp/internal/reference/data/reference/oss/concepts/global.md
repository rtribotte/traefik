---
schema_version: 2
kind: concept
name: Global
id: concept.global
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L139
summary: Global holds the global configuration.
fields:
  - name: checkNewVersion
    go_name: CheckNewVersion
    type: boolean
    go_type: bool
  - name: sendAnonymousUsage
    go_name: SendAnonymousUsage
    type: boolean
    go_type: bool
  - name: notAppendXForwardedFor
    go_name: NotAppendXForwardedFor
    type: boolean
    go_type: bool
---

# Global

Global holds the global configuration.
