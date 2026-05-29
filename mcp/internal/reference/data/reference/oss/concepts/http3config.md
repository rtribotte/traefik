---
schema_version: 2
kind: concept
name: HTTP3Config
id: concept.http3config
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L120
summary: HTTP3Config is the HTTP3 configuration of an entry point.
fields:
  - name: advertisedPort
    go_name: AdvertisedPort
    type: integer
    go_type: int
---

# HTTP3Config

HTTP3Config is the HTTP3 configuration of an entry point.
