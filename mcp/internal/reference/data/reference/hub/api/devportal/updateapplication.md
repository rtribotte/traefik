---
schema_version: 2
kind: rest-endpoint
name: updateApplication
id: hub.rest.updateapplication
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Update application
representations:
  openapi:
    path: /applications/{app-id}
    method: PUT
    tag: Applications
---

# updateApplication

Updates application notes. Only non-managed applications can be updated.
Users can only update applications they own.
