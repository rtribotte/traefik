---
schema_version: 2
kind: crd
name: AccessControlPolicy
id: crd.accesscontrolpolicy
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_accesscontrolpolicies.yaml
summary: AccessControlPolicy defines an access control policy.
fields:
  - name: apiKey
    type: object
  - name: basicAuth
    type: object
  - name: jwt
    type: object
  - name: oAuthIntro
    type: object
  - name: oidc
    type: object
  - name: oidcGoogle
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: AccessControlPolicy
    spec_path: .spec
---

# AccessControlPolicy

AccessControlPolicy defines an access control policy.
