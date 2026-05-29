---
schema_version: 2
kind: concept
name: TCPWeightedRoundRobin
id: concept.tcpweightedroundrobin
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L51
summary: TCPWeightedRoundRobin is a weighted round robin tcp load-balancer of services.
fields:
  - name: services
    go_name: Services
    type: array
    items: object
    go_type: '[]TCPWRRService'
    type_ref: oss:TCPWRRService
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
    description: Healthcheck defines health checks for ExternalName services.
---

# TCPWeightedRoundRobin

TCPWeightedRoundRobin is a weighted round robin tcp load-balancer of services.
