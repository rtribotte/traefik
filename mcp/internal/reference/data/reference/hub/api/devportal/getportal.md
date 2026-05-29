---
schema_version: 2
kind: rest-endpoint
name: getPortal
id: hub.rest.getportal
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Get portal information
representations:
  openapi:
    path: /
    method: GET
    tag: Portal
---

# getPortal

Retrieves the developer portal configuration, including available APIs, bundles,
plans, and user-specific subscriptions. This endpoint returns only APIs and bundles
that the authenticated user has access to based on their group memberships.
