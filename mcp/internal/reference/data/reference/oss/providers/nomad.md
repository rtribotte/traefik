---
schema_version: 2
kind: provider
name: ProviderBuilder
id: provider.nomad
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/nomad/nomad.go#L60
summary: ProviderBuilder is responsible for constructing namespaced instances of the Nomad provider.
fields:
  - name: defaultRule
    go_name: DefaultRule
    type: string
    go_type: string
  - name: constraints
    go_name: Constraints
    type: string
    go_type: string
  - name: endpoint
    go_name: Endpoint
    type: object
    go_type: '*EndpointConfig'
  - name: prefix
    go_name: Prefix
    type: string
    go_type: string
    description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
  - name: stale
    go_name: Stale
    type: boolean
    go_type: bool
  - name: exposedByDefault
    go_name: ExposedByDefault
    type: boolean
    go_type: bool
    default: true
  - name: refreshInterval
    go_name: RefreshInterval
    type: duration
    go_type: ptypes.Duration
  - name: allowEmptyServices
    go_name: AllowEmptyServices
    type: boolean
    go_type: bool
  - name: watch
    go_name: Watch
    type: boolean
    go_type: bool
  - name: throttleDuration
    go_name: ThrottleDuration
    type: duration
    go_type: ptypes.Duration
  - name: namespaces
    go_name: Namespaces
    type: array
    items: string
    go_type: '[]string'
representations:
  yaml_path: providers.nomad
  toml_path: providers.nomad
---

# ProviderBuilder

ProviderBuilder is responsible for constructing namespaced instances of the Nomad provider.
