---
schema_version: 2
kind: rest-endpoint
name: listApplications
id: hub.rest.listapplications
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: List applications
representations:
  openapi:
    path: /applications
    method: GET
    tag: Applications
---

# listApplications

Retrieves all applications owned by the authenticated user. Each application includes
its API keys and subscriptions.
