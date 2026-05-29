---
schema_version: 2
kind: crd
name: TCPRoute
id: gateway-api.tcproute
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/tcproute.json
summary: TCPRoute provides a way to route TCP requests. When combined with a Gateway
fields:
  - name: parentRefs
    type: array
    items: object
    description: ParentRefs references the resources (usually Gateways) that a Route wants
  - name: rules
    type: array
    items: object
    description: Rules are a list of TCP matchers and actions.
  - name: useDefaultGateways
    type: string
    description: UseDefaultGateways indicates the default Gateway scope to use for this
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1alpha2
    kind: TCPRoute
    spec_path: .spec
---

# TCPRoute

TCPRoute provides a way to route TCP requests. When combined with a Gateway
listener, it can be used to forward connections on the port specified by the
listener to a set of backends specified by the TCPRoute.
