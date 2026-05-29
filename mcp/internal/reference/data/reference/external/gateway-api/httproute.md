---
schema_version: 2
kind: crd
name: HTTPRoute
id: gateway-api.httproute
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/httproute.json
summary: HTTPRoute provides a way to route HTTP requests. This includes the capability
fields:
  - name: hostnames
    type: array
    items: string
    description: Hostnames defines a set of hostnames that should match against the HTTP Host
  - name: parentRefs
    type: array
    items: object
    description: ParentRefs references the resources (usually Gateways) that a Route wants
  - name: rules
    type: array
    items: object
    description: Rules are a list of HTTP matchers, filters and actions.
  - name: useDefaultGateways
    type: string
    description: UseDefaultGateways indicates the default Gateway scope to use for this
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1
    kind: HTTPRoute
    spec_path: .spec
---

# HTTPRoute

HTTPRoute provides a way to route HTTP requests. This includes the capability
to match requests by hostname, path, header, or query param. Filters can be
used to specify additional processing steps. Backends specify where matching
requests should be routed.
