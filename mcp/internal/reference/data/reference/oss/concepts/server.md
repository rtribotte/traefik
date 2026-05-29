---
schema_version: 2
kind: concept
name: Server
id: concept.server
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L470
summary: Server holds the server configuration.
fields:
  - name: url
    go_name: URL
    type: string
    go_type: string
  - name: weight
    go_name: Weight
    type: integer
    go_type: '*int'
    description: Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
  - name: preservePath
    go_name: PreservePath
    type: boolean
    go_type: bool
---

# Server

Server holds the server configuration.
