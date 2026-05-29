---
schema_version: 2
kind: concept
name: WRRService
id: concept.wrrservice
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L279
summary: WRRService is a reference to a service load-balanced with weighted round-robin.
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

# WRRService

WRRService is a reference to a service load-balanced with weighted round-robin.
