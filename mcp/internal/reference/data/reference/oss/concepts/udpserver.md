---
schema_version: 2
kind: concept
name: UDPServer
id: concept.udpserver
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L110
summary: UDPServer defines a UDP server configuration.
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    description: Address defines the authentication server address.
---

# UDPServer

UDPServer defines a UDP server configuration.
