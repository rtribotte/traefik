---
schema_version: 2
kind: gateway-service-annotations
name: GatewayServiceAnnotations
id: annotations.gateway-service
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/gateway/annotations.go
summary: 'Annotations on Kubernetes Service objects read by the Traefik Kubernetes Gateway API provider. This is an exhaustive list: any other annotation under the traefik.io/service.* prefix is silently ignored in the Gateway API context.'
fields:
  - name: traefik.io/service.nativelb
    go_name: NativeLB
    type: boolean
    go_type: '*bool'
    description: NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
---

# GatewayServiceAnnotations
