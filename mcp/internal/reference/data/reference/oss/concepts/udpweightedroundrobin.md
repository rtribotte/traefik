---
schema_version: 2
kind: concept
name: UDPWeightedRoundRobin
id: concept.udpweightedroundrobin
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L36
summary: UDPWeightedRoundRobin is a weighted round robin UDP load-balancer of services.
fields:
  - name: services
    go_name: Services
    type: array
    items: object
    go_type: '[]UDPWRRService'
    type_ref: oss:UDPWRRService
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
---

# UDPWeightedRoundRobin

UDPWeightedRoundRobin is a weighted round robin UDP load-balancer of services.
