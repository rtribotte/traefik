---
schema_version: 2
kind: crd
name: ManagedApplication
id: crd.managedapplication
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_managedapplications.yaml
summary: ManagedApplication represents a managed application.
fields:
  - name: apiKeys
    type: object
  - name: appId
    type: object
  - name: notes
    type: object
  - name: owner
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: ManagedApplication
    spec_path: .spec
---

# ManagedApplication

ManagedApplication represents a managed application.
