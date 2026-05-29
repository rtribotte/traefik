---
schema_version: 2
kind: provider
name: Provider
id: provider.zookeeper
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kv/zk/zk.go#L17
summary: Provider holds configurations of the provider.
fields:
  - name: rootKey
    go_name: RootKey
    type: string
    go_type: string
    default: traefik
  - name: endpoints
    go_name: Endpoints
    type: array
    items: string
    go_type: '[]string'
    description: Endpoints contains either a single address or a seed list of host:port addresses. Default value is ["localhost:6379"].
  - name: username
    go_name: Username
    type: string
    go_type: string
    description: Username defines the username to connect to the Redis server.
  - name: password
    go_name: Password
    type: string
    go_type: string
    description: Password defines the password to connect to the Redis server.
representations:
  yaml_path: providers.zooKeeper
  toml_path: providers.zooKeeper
---

# Provider

Provider holds configurations of the provider.
