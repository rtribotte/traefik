---
schema_version: 2
kind: static-section
name: SpiffeClientConfig
id: static.spiffe
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L128
summary: SpiffeClientConfig defines the SPIFFE client configuration.
fields:
  - name: workloadAPIAddr
    go_name: WorkloadAPIAddr
    type: string
    go_type: string
representations:
  yaml_path: spiffe
  toml_path: spiffe
---

# SpiffeClientConfig

SpiffeClientConfig defines the SPIFFE client configuration.
