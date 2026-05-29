---
schema_version: 2
kind: crd
name: IngressRouteTCP
id: crd.ingressroutetcp
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_ingressroutetcps.yaml
summary: IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
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
  - name: tls
    type: object
    description: TLS defines the configuration used to secure the connection to the authentication server.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteTCP
    spec_path: .spec
---

# IngressRouteTCP

IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
