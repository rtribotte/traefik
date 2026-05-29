---
schema_version: 2
kind: rest-endpoint
name: listAPIKeys
id: hub.rest.listapikeys
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: List API keys
representations:
  openapi:
    path: /applications/{app-id}/api-keys
    method: GET
    tag: API Keys
---

# listAPIKeys

Retrieves all API keys for an application. Users can only list API keys for applications they own.
