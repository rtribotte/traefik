---
schema_version: 2
kind: crd
name: APICatalogItem
id: crd.apicatalogitem
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apicatalogitems.yaml
summary: APICatalogItem defines APIs that will be part of the API catalog on the portal.
fields:
  - name: apiBundles
    type: object
  - name: apiPlan
    type: object
  - name: apiSelector
    type: object
  - name: apis
    type: object
  - name: everyone
    type: object
  - name: groups
    type: object
  - name: operationFilter
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APICatalogItem
    spec_path: .spec
---

# APICatalogItem

APICatalogItem defines APIs that will be part of the API catalog on the portal.
