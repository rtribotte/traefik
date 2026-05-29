---
schema_version: 2
kind: concept
name: HighestRandomWeight
id: concept.highestrandomweight
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L266
summary: HighestRandomWeight is a weighted sticky load-balancer of services.
fields:
  - name: services
    go_name: Services
    type: array
    items: object
    go_type: '[]HRWService'
    type_ref: oss:HRWService
    description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
    fields:
      - name: name
        go_name: Name
        type: string
        go_type: string
        description: Name defines the name of the referenced IngressRoute resource.
      - name: weight
        go_name: Weight
        type: integer
        go_type: '*int'
        default: 1
        description: Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
  - name: healthCheck
    go_name: HealthCheck
    type: object
    go_type: '*HealthCheck'
    type_ref: oss:HealthCheck
    description: HealthCheck enables automatic self-healthcheck for this service, i.e. whenever one of its children is reported as down, this service becomes aware of it, and takes it into account (i.e. it ignores the down child) when running the load-balancing algorithm. In addition, if the parent of this service also has HealthCheck enabled, this service reports to its parent any status change.
---

# HighestRandomWeight

HighestRandomWeight is a weighted sticky load-balancer of services.
