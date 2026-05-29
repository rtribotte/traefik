---
schema_version: 2
kind: crd
name: APIPortalAuth
id: crd.apiportalauth
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_apiportalauths.yaml
summary: APIPortalAuth defines the authentication configuration for an APIPortal.
fields:
  - name: ldap
    type: object
  - name: oidc
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: APIPortalAuth
    spec_path: .spec
---

# APIPortalAuth

APIPortalAuth defines the authentication configuration for an APIPortal.
