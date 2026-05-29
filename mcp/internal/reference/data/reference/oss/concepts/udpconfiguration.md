---
schema_version: 2
kind: concept
name: UDPConfiguration
id: concept.udpconfiguration
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L10
summary: UDPConfiguration contains all the UDP configuration parameters.
fields:
  - name: routers
    go_name: Routers
    type: object
    items: object
    go_type: map[string]*UDPRouter
    type_ref: oss:UDPRouter
    fields:
      - name: entryPoints
        go_name: EntryPoints
        type: array
        items: string
        go_type: '[]string'
        description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
      - name: service
        go_name: Service
        type: string
        go_type: string
        description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
  - name: services
    go_name: Services
    type: object
    items: object
    go_type: map[string]*UDPService
    type_ref: oss:UDPService
    description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
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
---

# UDPConfiguration

UDPConfiguration contains all the UDP configuration parameters.
