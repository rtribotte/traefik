---
schema_version: 2
kind: crd
name: TLSRoute
id: gateway-api.tlsroute
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/tlsroute.json
summary: The TLSRoute resource is similar to TCPRoute, but can be configured
fields:
  - name: hostnames
    type: array
    items: string
    description: Hostnames defines a set of SNI names that should match against the
  - name: parentRefs
    type: array
    items: object
    description: ParentRefs references the resources (usually Gateways) that a Route wants
  - name: rules
    type: array
    items: object
    description: Rules are a list of TLS matchers and actions.
  - name: useDefaultGateways
    type: string
    description: UseDefaultGateways indicates the default Gateway scope to use for this
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1alpha2
    kind: TLSRoute
    spec_path: .spec
---

# TLSRoute

The TLSRoute resource is similar to TCPRoute, but can be configured
to match against TLS-specific metadata. This allows more flexibility
in matching streams for a given TLS listener.
