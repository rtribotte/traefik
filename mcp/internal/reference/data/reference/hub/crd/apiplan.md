---
schema_version: 2
kind: crd
name: APIPlan
id: crd.apiplan
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiplans.yaml
summary: APIPlan defines API Plan policy.
fields:
  - name: description
    type: object
  - name: quota
    type: object
  - name: rateLimit
    type: object
  - name: title
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIPlan
    spec_path: .spec
---

# APIPlan

APIPlan defines API Plan policy.
