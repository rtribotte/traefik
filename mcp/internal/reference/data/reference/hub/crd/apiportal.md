---
schema_version: 2
kind: crd
name: APIPortal
id: crd.apiportal
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiportals.yaml
summary: APIPortal defines a developer portal for accessing the documentation of APIs.
fields:
  - name: auth
    type: object
  - name: description
    type: object
  - name: title
    type: object
  - name: trustedUrls
    type: object
  - name: ui
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIPortal
    spec_path: .spec
---

# APIPortal

APIPortal defines a developer portal for accessing the documentation of APIs.
