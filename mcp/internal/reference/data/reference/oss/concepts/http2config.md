---
schema_version: 2
kind: concept
name: HTTP2Config
id: concept.http2config
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L106
summary: HTTP2Config is the HTTP2 configuration of an entry point.
fields:
  - name: maxConcurrentStreams
    go_name: MaxConcurrentStreams
    type: integer
    go_type: int32
    default: 250
  - name: maxDecoderHeaderTableSize
    go_name: MaxDecoderHeaderTableSize
    type: integer
    go_type: int32
    default: 4096
  - name: maxEncoderHeaderTableSize
    go_name: MaxEncoderHeaderTableSize
    type: integer
    go_type: int32
    default: 4096
---

# HTTP2Config

HTTP2Config is the HTTP2 configuration of an entry point.
