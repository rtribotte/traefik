---
schema_version: 2
kind: concept
name: Core
id: concept.core
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L117
summary: Core configures Traefik core behavior.
fields:
  - name: defaultRuleSyntax
    go_name: DefaultRuleSyntax
    type: string
    go_type: string
    default: v3
---

# Core

Core configures Traefik core behavior.
