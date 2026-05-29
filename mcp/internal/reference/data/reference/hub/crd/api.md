---
schema_version: 2
kind: crd
name: API
id: crd.api
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apis.yaml
summary: API defines an HTTP interface that is exposed to external clients. It specifies the supported versions
fields:
  - name: cors
    type: object
  - name: description
    type: object
  - name: openApiSpec
    type: object
  - name: title
    type: object
  - name: versions
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: API
    spec_path: .spec
---

# API

API defines an HTTP interface that is exposed to external clients. It specifies the supported versions
and provides instructions for accessing its documentation. Once instantiated, an API object is associated
with an Ingress, IngressRoute, or HTTPRoute resource, enabling the exposure of the described API to the outside world.
