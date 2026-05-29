---
schema_version: 2
kind: rest-endpoint
name: getAPIVersionSpec
id: hub.rest.getapiversionspec
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Get API version specification
representations:
  openapi:
    path: /apis/{api}/versions/{version}
    method: GET
    tag: APIs
---

# getAPIVersionSpec

Retrieves the OpenAPI 3.0 specification for a specific version of an API.
