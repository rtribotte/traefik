---
schema_version: 2
kind: service-udp
name: UDPService
id: udp.services
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L18
summary: UDPService defines the configuration for a UDP service. All fields are mutually exclusive.
fields:
  - name: loadBalancer
    go_name: LoadBalancer
    type: object
    go_type: '*UDPServersLoadBalancer'
    type_ref: oss:UDPServersLoadBalancer
    fields:
      - name: servers
        go_name: Servers
        type: array
        items: object
        go_type: '[]UDPServer'
        type_ref: oss:UDPServer
        fields:
          - name: address
            go_name: Address
            type: string
            go_type: string
            description: Address defines the authentication server address.
  - name: weighted
    go_name: Weighted
    type: object
    go_type: '*UDPWeightedRoundRobin'
    type_ref: oss:UDPWeightedRoundRobin
    description: Weighted defines the Weighted Round Robin configuration.
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
variants:
  - loadBalancer
  - weighted
representations:
  yaml_path: udp.services.<name>
  toml_path: udp.services.<name>
  label_prefix: traefik.udp.services.<name>
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteUDP
    spec_path: .spec
---

# UDPService

UDPService defines the configuration for a UDP service. All fields are mutually exclusive.
