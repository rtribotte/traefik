---
schema_version: 2
kind: rest-endpoint
name: suspendAPIKey
id: hub.rest.suspendapikey
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Suspend or unsuspend API key
representations:
  openapi:
    path: /applications/{app-id}/api-keys/suspend
    method: POST
    tag: API Keys
---

# suspendAPIKey

Suspends or unsuspends an API key. Suspended keys cannot be used for authentication.
Only non-managed applications can have API keys suspended/unsuspended.
Users can only modify API keys from applications they own.
