---
schema_version: 2
kind: static-section
name: HostResolverConfig
id: static.hostresolver
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/types/host_resolver.go#L4
summary: HostResolverConfig contain configuration for CNAME Flattening.
fields:
  - name: cnameFlattening
    go_name: CnameFlattening
    type: boolean
    go_type: bool
    default: false
  - name: resolvConfig
    go_name: ResolvConfig
    type: string
    go_type: string
    default: /etc/resolv.conf
  - name: resolvDepth
    go_name: ResolvDepth
    type: integer
    go_type: int
    default: 5
representations:
  yaml_path: hostResolver
  toml_path: hostResolver
---

# HostResolverConfig

HostResolverConfig contain configuration for CNAME Flattening.
