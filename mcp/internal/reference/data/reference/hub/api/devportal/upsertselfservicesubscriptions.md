---
schema_version: 2
kind: rest-endpoint
name: upsertSelfServiceSubscriptions
id: hub.rest.upsertselfservicesubscriptions
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/hub/api/devportal/openapi/openapi.json
summary: Create or update self-service subscriptions
representations:
  openapi:
    path: /self-service-subscriptions
    method: POST
    tag: Self-Service Subscriptions
---

# upsertSelfServiceSubscriptions

Creates or updates self-service subscriptions in batch. This endpoint allows users to
subscribe their applications to APIs with specific plans. The operation filter, rate limit,
and quota must exactly match the catalog item configuration for the selected plan.

If a subscription ID is provided, the subscription will be updated; otherwise, a new
subscription will be created.

This endpoint requires the self-service-subscription feature to be enabled.
