---
schema_version: 2
kind: crd
name: Uplink
id: crd.uplink
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_uplinks.yaml
summary: 'Uplink is an inter-cluster service advertisement: a child cluster declares an Uplink to advertise'
fields:
  - name: entryPoints
    type: object
  - name: exposeName
    type: object
  - name: healthCheck
    type: object
  - name: passiveHealthCheck
    type: object
  - name: weight
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: Uplink
    spec_path: .spec
---

# Uplink

Uplink is an inter-cluster service advertisement: a child cluster declares an Uplink to advertise
to a parent cluster that it can handle a particular workload.
