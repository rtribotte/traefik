---
schema_version: 2
kind: middleware-tcp
name: TCPIPAllowList
id: tcp.middlewares.ipallowlist
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_middlewares.go#L41
summary: 'TCPIPAllowList holds the TCP IPAllowList middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/tcp/ipallowlist/'
fields:
  - name: sourceRange
    go_name: SourceRange
    type: array
    items: string
    go_type: '[]string'
    description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
representations:
  yaml_path: tcp.middlewares.<name>.ipAllowList
  toml_path: tcp.middlewares.<name>.ipAllowList
  label_prefix: traefik.tcp.middlewares.<name>.ipallowlist
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: MiddlewareTCP
    spec_path: .spec.ipAllowList
---

# TCPIPAllowList

TCPIPAllowList holds the TCP IPAllowList middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/tcp/ipallowlist/
