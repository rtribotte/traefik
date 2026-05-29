---
schema_version: 2
kind: rest-endpoint
name: deleteApplication
id: hub.rest.deleteapplication
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Delete application
representations:
  openapi:
    path: /applications/{app-id}
    method: DELETE
    tag: Applications
---

# deleteApplication

Deletes an application and all its associated API keys and subscriptions.
Only non-managed applications can be deleted. Users can only delete applications
they own.
