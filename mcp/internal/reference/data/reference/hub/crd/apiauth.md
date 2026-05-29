---
schema_version: 2
kind: crd
name: APIAuth
id: crd.apiauth
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiauths.yaml
summary: APIAuth defines the authentication configuration for APIs.
fields:
  - name: apiKey
    type: object
  - name: isDefault
    type: object
  - name: jwt
    type: object
  - name: ldap
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIAuth
    spec_path: .spec
---

# APIAuth

APIAuth defines the authentication configuration for APIs.
