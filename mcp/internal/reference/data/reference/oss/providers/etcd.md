---
schema_version: 2
kind: provider
name: Provider
id: provider.etcd
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kv/etcd/etcd.go#L20
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
  - name: tls
    go_name: TLS
    type: object
    go_type: '*types.ClientTLS'
    type_ref: oss:ClientTLS
    description: TLS defines the configuration used to secure the connection to the authentication server.
    fields:
      - name: ca
        go_name: CA
        type: string
        go_type: string
      - name: cert
        go_name: Cert
        type: string
        go_type: string
      - name: key
        go_name: Key
        type: string
        go_type: string
      - name: insecureSkipVerify
        go_name: InsecureSkipVerify
        type: boolean
        go_type: bool
        description: InsecureSkipVerify defines whether the server certificates should be validated.
      - name: caOptional
        go_name: CAOptional
        type: boolean
        go_type: '*bool'
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
  yaml_path: providers.etcd
  toml_path: providers.etcd
---

# Provider

Provider holds configurations of the provider.
