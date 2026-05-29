---
schema_version: 2
kind: crd
name: ContentItem
id: crd.contentitem
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_contentitems.yaml
summary: ContentItem defines additional documentation for given resource.
fields:
  - name: content
    type: object
  - name: link
    type: object
  - name: order
    type: object
  - name: parentRef
    type: object
  - name: title
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: ContentItem
    spec_path: .spec
---

# ContentItem

ContentItem defines additional documentation for given resource.
