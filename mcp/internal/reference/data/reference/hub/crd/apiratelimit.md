---
schema_version: 2
kind: crd
name: APIRateLimit
id: crd.apiratelimit
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiratelimits.yaml
summary: APIRateLimit defines how group of consumers are rate limited on a set of APIs.
fields:
  - name: apiSelector
    type: object
  - name: apis
    type: object
  - name: everyone
    type: object
  - name: groups
    type: object
  - name: limit
    type: object
  - name: period
    type: object
  - name: strategy
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIRateLimit
    spec_path: .spec
---

# APIRateLimit

APIRateLimit defines how group of consumers are rate limited on a set of APIs.
