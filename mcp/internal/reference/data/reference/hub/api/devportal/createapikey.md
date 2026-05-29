---
schema_version: 2
kind: rest-endpoint
name: createAPIKey
id: hub.rest.createapikey
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Create API key
representations:
  openapi:
    path: /applications/{app-id}/api-keys
    method: POST
    tag: API Keys
---

# createAPIKey

Creates a new API key for an application. The API key token is returned only once
at creation time and cannot be retrieved later. Only non-managed applications can
have API keys created.
