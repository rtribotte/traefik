---
schema_version: 2
kind: provider-hub
name: Hub
id: hub.providers.hub
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/provider/hub/hub.go#L65
summary: Provider holds configurations of the provider.
fields:
  - name: entryPoints
    go_name: EntryPoints
    type: array
    items: string
    go_type: '[]string'
  - name: InstanceID
    go_name: InstanceID
    type: string
    go_type: string
  - name: APIMgtMetrics
    go_name: APIMgtMetrics
    type: boolean
    go_type: bool
representations:
  yaml_path: hub.providers.hub
  toml_path: hub.providers.hub
---

# Hub

Provider holds configurations of the provider.
