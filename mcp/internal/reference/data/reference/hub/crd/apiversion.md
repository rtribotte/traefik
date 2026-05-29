---
schema_version: 2
kind: crd
name: APIVersion
id: crd.apiversion
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiversions.yaml
summary: APIVersion defines a version of an API.
fields:
  - name: cors
    type: object
  - name: description
    type: object
  - name: openApiSpec
    type: object
  - name: release
    type: object
  - name: title
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIVersion
    spec_path: .spec
---

# APIVersion

APIVersion defines a version of an API.
