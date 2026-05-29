---
schema_version: 2
kind: concept
name: DevPortalAPIIndex
id: hub.api.devportal
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Index of the developer portal REST API endpoints.
---

# DevPortalAPIIndex

Developer portal REST API endpoints, grouped by tag.

## API Keys

- `POST /applications/{app-id}/api-keys` — `hub.rest.createapikey` — Create API key
- `DELETE /applications/{app-id}/api-keys` — `hub.rest.deleteapikey` — Delete API key
- `GET /applications/{app-id}/api-keys` — `hub.rest.listapikeys` — List API keys
- `POST /applications/{app-id}/api-keys/suspend` — `hub.rest.suspendapikey` — Suspend or unsuspend API key

## APIs

- `GET /apis/{api}` — `hub.rest.getapispec` — Get API specification
- `GET /apis/{api}/versions/{version}` — `hub.rest.getapiversionspec` — Get API version specification

## Applications

- `POST /applications` — `hub.rest.createapplication` — Create application
- `DELETE /applications/{app-id}` — `hub.rest.deleteapplication` — Delete application
- `GET /applications` — `hub.rest.listapplications` — List applications
- `PUT /applications/{app-id}` — `hub.rest.updateapplication` — Update application

## Content

- `GET /content/{content-id}` — `hub.rest.getcontent` — Get content

## Portal

- `GET /` — `hub.rest.getportal` — Get portal information

## Self-Service Subscriptions

- `DELETE /self-service-subscriptions/{id}` — `hub.rest.deleteselfservicesubscription` — Delete self-service subscription
- `POST /self-service-subscriptions` — `hub.rest.upsertselfservicesubscriptions` — Create or update self-service subscriptions
