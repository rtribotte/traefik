---
schema_version: 2
kind: crd
name: ManagedSubscription
id: crd.managedsubscription
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_managedsubscriptions.yaml
summary: ManagedSubscription defines a Subscription managed by the API manager as the result of a pre-negotiation with its
fields:
  - name: apiBundles
    type: object
  - name: apiPlan
    type: object
  - name: apiSelector
    type: object
  - name: apis
    type: object
  - name: applications
    type: object
  - name: claims
    type: object
  - name: managedApplications
    type: object
  - name: operationFilter
    type: object
  - name: weight
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: ManagedSubscription
    spec_path: .spec
---

# ManagedSubscription

ManagedSubscription defines a Subscription managed by the API manager as the result of a pre-negotiation with its
API consumers. This subscription grant consuming access to a set of APIs to a set of Applications.
