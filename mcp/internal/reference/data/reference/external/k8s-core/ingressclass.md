---
schema_version: 2
kind: kubernetes-core
name: IngressClass
id: k8s-core.ingressclass
source: k8s-core
traefik_version: v3.7.0
extracted_from:
  - schemas/external/k8s-core/ingressclass.json
summary: IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.
fields:
  - name: controller
    type: string
    description: controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
  - name: parameters
    type: object
    description: IngressClassParametersReference identifies an API object. This can be used to specify a cluster or namespace-scoped resource.
representations:
  yaml_path: spec
  crd:
    apiVersion: networking.k8s.io/v1
    kind: IngressClass
    spec_path: .spec
---

# IngressClass

IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.
