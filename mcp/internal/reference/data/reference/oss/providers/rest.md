---
schema_version: 2
kind: provider
name: Provider
id: provider.rest
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/rest/rest.go#L22
summary: Provider is a provider.Provider implementation that provides a Rest API.
fields:
  - name: insecure
    go_name: Insecure
    type: boolean
    go_type: bool
representations:
  yaml_path: providers.rest
  toml_path: providers.rest
---

# Provider

Provider is a provider.Provider implementation that provides a Rest API.
