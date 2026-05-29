---
schema_version: 2
kind: rest-endpoint
name: deleteSelfServiceSubscription
id: hub.rest.deleteselfservicesubscription
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Delete self-service subscription
representations:
  openapi:
    path: /self-service-subscriptions/{id}
    method: DELETE
    tag: Self-Service Subscriptions
---

# deleteSelfServiceSubscription

Deletes a self-service subscription. Users can only delete subscriptions they own.
