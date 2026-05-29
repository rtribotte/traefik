---
schema_version: 2
kind: crd
name: APIBundle
id: crd.apibundle
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apibundles.yaml
summary: APIBundle defines a set of APIs.
fields:
  - name: apiSelector
    type: object
  - name: apis
    type: object
  - name: title
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIBundle
    spec_path: .spec
---

# APIBundle

APIBundle defines a set of APIs.
