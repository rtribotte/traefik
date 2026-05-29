---
schema_version: 2
kind: provider
name: Provider
id: provider.file
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/file/file.go#L35
summary: Provider holds configurations of the provider.
fields:
  - name: directory
    go_name: Directory
    type: string
    go_type: string
  - name: watch
    go_name: Watch
    type: boolean
    go_type: bool
    default: true
  - name: filename
    go_name: Filename
    type: string
    go_type: string
    default: ""
  - name: debugLogGeneratedTemplate
    go_name: DebugLogGeneratedTemplate
    type: boolean
    go_type: bool
representations:
  yaml_path: providers.file
  toml_path: providers.file
---

# Provider

Provider holds configurations of the provider.
