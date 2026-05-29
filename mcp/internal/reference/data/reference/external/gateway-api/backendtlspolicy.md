---
schema_version: 2
kind: crd
name: BackendTLSPolicy
id: gateway-api.backendtlspolicy
source: gateway-api
traefik_version: v3.7.0
extracted_from:
  - schemas/external/gateway-api/backendtlspolicy.json
summary: BackendTLSPolicy provides a way to configure how a Gateway
fields:
  - name: options
    type: object
    description: Options are a list of key/value pairs to enable extended TLS
  - name: targetRefs
    type: array
    items: object
    description: TargetRefs identifies an API object to apply the policy to.
  - name: validation
    type: object
    description: Validation contains backend TLS validation configuration.
representations:
  yaml_path: spec
  crd:
    apiVersion: gateway.networking.k8s.io/v1
    kind: BackendTLSPolicy
    spec_path: .spec
---

# BackendTLSPolicy

BackendTLSPolicy provides a way to configure how a Gateway
connects to a Backend via TLS.
