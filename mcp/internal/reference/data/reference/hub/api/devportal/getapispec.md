---
schema_version: 2
kind: rest-endpoint
name: getAPISpec
id: hub.rest.getapispec
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Get API specification
representations:
  openapi:
    path: /apis/{api}
    method: GET
    tag: APIs
---

# getAPISpec

Retrieves the OpenAPI 3.0 specification for an API.
