---
schema_version: 2
kind: static-section
name: PluginRegistry
id: hub.static.pluginregistry
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L155
summary: PluginRegistry holds the plugin registry configuration.
fields:
  - name: sources
    go_name: Sources
    type: object
    items: object
    go_type: map[string]PluginRegistrySource
representations:
  yaml_path: hub.pluginRegistry
  toml_path: hub.pluginRegistry
---

# PluginRegistry

PluginRegistry holds the plugin registry configuration.
