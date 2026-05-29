---
schema_version: 2
kind: rest-endpoint
name: deleteAPIKey
id: hub.rest.deleteapikey
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Delete API key
representations:
  openapi:
    path: /applications/{app-id}/api-keys
    method: DELETE
    tag: API Keys
---

# deleteAPIKey

Deletes an API key for an application.
Only non-managed applications can have API keys deleted.
Users can only delete API keys from applications they own.
