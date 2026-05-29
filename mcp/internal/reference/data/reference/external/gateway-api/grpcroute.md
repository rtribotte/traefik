---
schema_version: 2
kind: crd
name: GRPCRoute
id: gateway-api.grpcroute
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/grpcroute.json
summary: GRPCRoute provides a way to route gRPC requests. This includes the capability
fields:
  - name: hostnames
    type: array
    items: string
    description: Hostnames defines a set of hostnames to match against the GRPC
  - name: parentRefs
    type: array
    items: object
    description: ParentRefs references the resources (usually Gateways) that a Route wants
  - name: rules
    type: array
    items: object
    description: Rules are a list of GRPC matchers, filters and actions.
  - name: useDefaultGateways
    type: string
    description: UseDefaultGateways indicates the default Gateway scope to use for this
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1
    kind: GRPCRoute
    spec_path: .spec
---

# GRPCRoute

GRPCRoute provides a way to route gRPC requests. This includes the capability
to match requests by hostname, gRPC service, gRPC method, or HTTP/2 header.
Filters can be used to specify additional processing steps. Backends specify
where matching requests will be routed.

GRPCRoute falls under extended support within the Gateway API. Within the
following specification, the word "MUST" indicates that an implementation
supporting GRPCRoute must conform to the indicated requirement, but an
implementation not supporting this route type need not follow the requirement
unless explicitly indicated.

Implementations supporting `GRPCRoute` with the `HTTPS` `ProtocolType` MUST
accept HTTP/2 connections without an initial upgrade from HTTP/1.1, i.e. via
ALPN. If the implementation does not support this, then it MUST set the
"Accepted" condition to "False" for the affected listener with a reason of
"UnsupportedProtocol".  Implementations MAY also accept HTTP/2 connections
with an upgrade from HTTP/1.

Implementations supporting `GRPCRoute` with the `HTTP` `ProtocolType` MUST
support HTTP/2 over cleartext TCP (h2c,
https://www.rfc-editor.org/rfc/rfc7540#section-3.1) without an initial
upgrade from HTTP/1.1, i.e. with prior knowledge
(https://www.rfc-editor.org/rfc/rfc7540#section-3.4). If the implementation
does not support this, then it MUST set the "Accepted" condition to "False"
for the affected listener with a reason of "UnsupportedProtocol".
Implementations MAY also accept HTTP/2 connections with an upgrade from
HTTP/1, i.e. without prior knowledge.
