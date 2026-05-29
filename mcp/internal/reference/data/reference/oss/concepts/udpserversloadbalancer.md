---
schema_version: 2
kind: concept
name: UDPServersLoadBalancer
id: concept.udpserversloadbalancer
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L65
summary: UDPServersLoadBalancer defines the configuration for a load-balancer of UDP servers.
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
---

# UDPServersLoadBalancer

UDPServersLoadBalancer defines the configuration for a load-balancer of UDP servers.
