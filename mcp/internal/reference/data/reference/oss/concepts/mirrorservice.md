---
schema_version: 2
kind: concept
name: MirrorService
id: concept.mirrorservice
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L244
summary: MirrorService holds the MirrorService configuration.
fields:
  - name: name
    go_name: Name
    type: string
    go_type: string
    description: Name defines the name of the referenced IngressRoute resource.
  - name: percent
    go_name: Percent
    type: integer
    go_type: int
    description: 'Percent defines the part of the traffic to mirror. Supported values: 0 to 100.'
---

# MirrorService

MirrorService holds the MirrorService configuration.
