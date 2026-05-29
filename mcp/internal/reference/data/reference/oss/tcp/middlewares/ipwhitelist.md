---
schema_version: 2
kind: middleware-tcp
name: TCPIPWhiteList
id: tcp.middlewares.ipwhitelist
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_middlewares.go#L31
summary: TCPIPWhiteList holds the TCP IPWhiteList middleware configuration.
deprecated: true
replaced_by: please use IPAllowList instead
fields:
  - name: sourceRange
    go_name: SourceRange
    type: array
    items: string
    go_type: '[]string'
    description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
representations:
  yaml_path: tcp.middlewares.<name>.ipWhiteList
  toml_path: tcp.middlewares.<name>.ipWhiteList
  label_prefix: traefik.tcp.middlewares.<name>.ipwhitelist
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: MiddlewareTCP
    spec_path: .spec.ipWhiteList
---

# TCPIPWhiteList

TCPIPWhiteList holds the TCP IPWhiteList middleware configuration.
