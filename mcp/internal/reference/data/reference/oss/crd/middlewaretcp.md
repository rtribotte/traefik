---
schema_version: 2
kind: crd
name: MiddlewareTCP
id: crd.middlewaretcp
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_middlewaretcps.yaml
summary: MiddlewareTCP is the CRD implementation of a Traefik TCP middleware.
fields:
  - name: inFlightConn
    type: object
    description: InFlightConn defines the InFlightConn middleware configuration.
  - name: ipAllowList
    type: object
    description: 'IPAllowList defines the IPAllowList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipallowlist/'
  - name: ipWhiteList
    type: object
    description: 'IPWhiteList defines the IPWhiteList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipwhitelist/'
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: MiddlewareTCP
    spec_path: .spec
---

# MiddlewareTCP

MiddlewareTCP is the CRD implementation of a Traefik TCP middleware.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/overview/
