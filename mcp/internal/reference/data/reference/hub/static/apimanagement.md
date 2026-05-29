---
schema_version: 2
kind: static-section
name: APIManagement
id: hub.static.apimanagement
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L117
summary: APIManagement holds the API management configuration.
fields:
  - name: admission
    go_name: Admission
    type: object
    go_type: '*Admission'
  - name: openApi
    go_name: OpenAPI
    type: object
    go_type: '*OpenAPI'
  - name: apiKey
    go_name: APIKey
    type: object
    go_type: '*APIKey'
representations:
  yaml_path: hub.apiManagement
  toml_path: hub.apiManagement
---

# APIManagement

APIManagement holds the API management configuration.
