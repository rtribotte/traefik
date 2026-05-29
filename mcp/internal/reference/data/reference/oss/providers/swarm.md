---
schema_version: 2
kind: provider
name: SwarmProvider
id: provider.swarm
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/docker/pswarm.go#L29
summary: SwarmProvider holds configurations of the provider.
fields:
  - name: exposedByDefault
    go_name: ExposedByDefault
    type: boolean
    go_type: bool
  - name: constraints
    go_name: Constraints
    type: string
    go_type: string
  - name: allowEmptyServices
    go_name: AllowEmptyServices
    type: boolean
    go_type: bool
  - name: network
    go_name: Network
    type: string
    go_type: string
  - name: useBindPortIP
    go_name: UseBindPortIP
    type: boolean
    go_type: bool
  - name: watch
    go_name: Watch
    type: boolean
    go_type: bool
  - name: defaultRule
    go_name: DefaultRule
    type: string
    go_type: string
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
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
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
  - name: httpClientTimeout
    go_name: HTTPClientTimeout
    type: duration
    go_type: ptypes.Duration
  - name: refreshSeconds
    go_name: RefreshSeconds
    type: duration
    go_type: ptypes.Duration
representations:
  yaml_path: providers.swarm
  toml_path: providers.swarm
---

# SwarmProvider

SwarmProvider holds configurations of the provider.
