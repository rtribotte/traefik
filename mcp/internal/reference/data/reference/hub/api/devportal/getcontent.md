---
schema_version: 2
kind: rest-endpoint
name: getContent
id: hub.rest.getcontent
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Get content
representations:
  openapi:
    path: /content/{content-id}
    method: GET
    tag: Content
---

# getContent

Retrieves the raw markdown content.
