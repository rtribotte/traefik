---
schema_version: 2
kind: kubernetes-core
name: Ingress
id: k8s-core.ingress
source: k8s-core
traefik_version: v3.7.0
extracted_from:
  - schemas/external/k8s-core/ingress.json
summary: Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
fields:
  - name: defaultBackend
    type: object
    description: IngressBackend describes all endpoints for a given service and port.
  - name: ingressClassName
    type: string
    description: ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the `kubernetes.io/ingress.class` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.
  - name: rules
    type: array
    items: object
    description: rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
  - name: tls
    type: array
    items: object
    description: tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
representations:
  yaml_path: spec
  crd:
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    spec_path: .spec
---

# Ingress

Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
