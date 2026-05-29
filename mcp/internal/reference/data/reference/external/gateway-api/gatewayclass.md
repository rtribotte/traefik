---
schema_version: 2
kind: crd
name: GatewayClass
id: gateway-api.gatewayclass
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/gatewayclass.json
summary: GatewayClass describes a class of Gateways available to the user for creating
fields:
  - name: controllerName
    type: string
    description: ControllerName is the name of the controller that is managing Gateways of
  - name: description
    type: string
    description: Description helps describe a GatewayClass with more details.
  - name: parametersRef
    type: object
    description: ParametersRef is a reference to a resource that contains the configuration
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1
    kind: GatewayClass
    spec_path: .spec
---

# GatewayClass

GatewayClass describes a class of Gateways available to the user for creating
Gateway resources.

It is recommended that this resource be used as a template for Gateways. This
means that a Gateway is based on the state of the GatewayClass at the time it
was created and changes to the GatewayClass or associated parameters are not
propagated down to existing Gateways. This recommendation is intended to
limit the blast radius of changes to GatewayClass or associated parameters.
If implementations choose to propagate GatewayClass changes to existing
Gateways, that MUST be clearly documented by the implementation.

Whenever one or more Gateways are using a GatewayClass, implementations SHOULD
add the `gateway-exists-finalizer.gateway.networking.k8s.io` finalizer on the
associated GatewayClass. This ensures that a GatewayClass associated with a
Gateway is not deleted while in use.

GatewayClass is a Cluster level resource.
