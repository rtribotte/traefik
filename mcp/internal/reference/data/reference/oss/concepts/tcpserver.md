---
schema_version: 2
kind: concept
name: TCPServer
id: concept.tcpserver
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L156
summary: TCPServer holds a TCP Server configuration.
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    description: Address defines the authentication server address.
  - name: tls
    go_name: TLS
    type: boolean
    go_type: bool
    description: TLS defines the configuration used to secure the connection to the authentication server.
---

# TCPServer

TCPServer holds a TCP Server configuration.
