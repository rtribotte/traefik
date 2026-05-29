---
schema_version: 2
kind: crd
name: IngressRoute
id: crd.ingressroute
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_ingressroutes.yaml
summary: IngressRoute is the CRD implementation of a Traefik HTTP Router.
fields:
  - name: entryPoints
    type: object
    description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
  - name: ingressClassName
    type: object
    description: IngressClassName defines the name of the IngressClass cluster resource.
  - name: parentRefs
    type: object
    description: 'ParentRefs defines references to parent IngressRoute resources for multi-layer routing. When set, this IngressRoute''s routers will be children of the referenced parent IngressRoute''s routers. More info: https://doc.traefik.io/traefik/v3.7/routing/routers/#parentrefs'
  - name: routes
    type: object
    description: Routes defines the list of routes.
  - name: tls
    type: object
    description: TLS defines the configuration used to secure the connection to the authentication server.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRoute
    spec_path: .spec
---

# IngressRoute

IngressRoute is the CRD implementation of a Traefik HTTP Router.
