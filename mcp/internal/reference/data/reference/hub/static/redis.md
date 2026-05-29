---
schema_version: 2
kind: static-section
name: Config
id: hub.static.redis
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/redis/config.go#L11
summary: Config is the Redis configuration.
fields:
  - name: cluster
    go_name: Cluster
    type: object
    go_type: '*ClusterConfig'
  - name: sentinel
    go_name: Sentinel
    type: object
    go_type: '*SentinelConfig'
  - name: endpoints
    go_name: Endpoints
    type: array
    items: string
    go_type: '[]string'
  - name: username
    go_name: Username
    type: string
    go_type: string
  - name: password
    go_name: Password
    type: string
    go_type: string
  - name: database
    go_name: Database
    type: integer
    go_type: int
  - name: unstableResp3
    go_name: UnstableResp3
    type: boolean
    go_type: bool
  - name: protocol
    go_name: Protocol
    type: integer
    go_type: int
  - name: tls
    go_name: TLS
    type: object
    go_type: '*ttypes.ClientTLS'
  - name: timeout
    go_name: Timeout
    type: duration
    go_type: '*types.Duration'
representations:
  yaml_path: hub.redis
  toml_path: hub.redis
---

# Config

Config is the Redis configuration.
