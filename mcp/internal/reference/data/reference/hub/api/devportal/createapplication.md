---
schema_version: 2
kind: rest-endpoint
name: createApplication
id: hub.rest.createapplication
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Create application
representations:
  openapi:
    path: /applications
    method: POST
    tag: Applications
---

# createApplication

Creates a new application for the authenticated user.
