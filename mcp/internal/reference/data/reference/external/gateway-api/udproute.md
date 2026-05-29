---
schema_version: 2
kind: crd
name: UDPRoute
id: gateway-api.udproute
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/udproute.json
summary: UDPRoute provides a way to route UDP traffic. When combined with a Gateway
fields:
  - name: parentRefs
    type: array
    items: object
    description: ParentRefs references the resources (usually Gateways) that a Route wants
  - name: rules
    type: array
    items: object
    description: Rules are a list of UDP matchers and actions.
  - name: useDefaultGateways
    type: string
    description: UseDefaultGateways indicates the default Gateway scope to use for this
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1alpha2
    kind: UDPRoute
    spec_path: .spec
---

# UDPRoute

UDPRoute provides a way to route UDP traffic. When combined with a Gateway
listener, it can be used to forward traffic on the port specified by the
listener to a set of backends specified by the UDPRoute.
