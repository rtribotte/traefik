---
schema_version: 2
kind: crd
name: ReferenceGrant
id: gateway-api.referencegrant
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/referencegrant.json
summary: ReferenceGrant identifies kinds of resources in other namespaces that are
fields:
  - name: from
    type: array
    items: object
    description: From describes the trusted namespaces and kinds that can reference the
  - name: to
    type: array
    items: object
    description: To describes the resources that may be referenced by the resources
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1beta1
    kind: ReferenceGrant
    spec_path: .spec
---

# ReferenceGrant

ReferenceGrant identifies kinds of resources in other namespaces that are
trusted to reference the specified kinds of resources in the same namespace
as the policy.

Each ReferenceGrant can be used to represent a unique trust relationship.
Additional Reference Grants can be used to add to the set of trusted
sources of inbound references for the namespace they are defined within.

All cross-namespace references in Gateway API (with the exception of cross-namespace
Gateway-route attachment) require a ReferenceGrant.

ReferenceGrant is a form of runtime verification allowing users to assert
which cross-namespace object references are permitted. Implementations that
support ReferenceGrant MUST NOT permit cross-namespace references which have
no grant, and MUST respond to the removal of a grant by revoking the access
that the grant allowed.
