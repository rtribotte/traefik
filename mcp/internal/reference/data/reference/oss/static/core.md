---
schema_version: 2
kind: static-section
name: Core
id: static.core
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L117
summary: Core configures Traefik core behavior.
deprecated: true
replaced_by: Please do not use this field
fields:
  - name: defaultRuleSyntax
    go_name: DefaultRuleSyntax
    type: string
    go_type: string
    default: v3
representations:
  yaml_path: core
  toml_path: core
---

# Core

Core configures Traefik core behavior.
