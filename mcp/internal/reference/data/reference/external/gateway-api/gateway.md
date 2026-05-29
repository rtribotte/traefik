---
schema_version: 2
kind: crd
name: Gateway
id: gateway-api.gateway
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/gateway.json
summary: Gateway represents an instance of a service-traffic handling infrastructure
fields:
  - name: addresses
    type: array
    items: object
    description: Addresses requested for this Gateway. This is optional and behavior can
  - name: allowedListeners
    type: object
    description: AllowedListeners defines which ListenerSets can be attached to this Gateway.
  - name: defaultScope
    type: string
    description: DefaultScope, when set, configures the Gateway as a default Gateway,
  - name: gatewayClassName
    type: string
    description: GatewayClassName used for this Gateway. This is the name of a
  - name: infrastructure
    type: object
    description: Infrastructure defines infrastructure level attributes about this Gateway instance.
  - name: listeners
    type: array
    items: object
    description: Listeners associated with this Gateway. Listeners define
  - name: tls
    type: object
    description: TLS specifies frontend and backend tls configuration for entire gateway.
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1
    kind: Gateway
    spec_path: .spec
---

# Gateway

Gateway represents an instance of a service-traffic handling infrastructure
by binding Listeners to a set of IP addresses.
