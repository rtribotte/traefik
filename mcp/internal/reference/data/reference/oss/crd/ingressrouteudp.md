---
schema_version: 2
kind: crd
name: IngressRouteUDP
id: crd.ingressrouteudp
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_ingressrouteudps.yaml
summary: IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
fields:
  - name: entryPoints
    type: object
    description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
  - name: ingressClassName
    type: object
    description: IngressClassName defines the name of the IngressClass cluster resource.
  - name: routes
    type: object
    description: Routes defines the list of routes.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteUDP
    spec_path: .spec
---

# IngressRouteUDP

IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
