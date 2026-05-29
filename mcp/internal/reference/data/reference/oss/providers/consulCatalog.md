---
schema_version: 2
kind: provider
name: ProviderBuilder
id: provider.consulcatalog
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/consulcatalog/consul_catalog.go#L50
summary: ProviderBuilder is responsible for constructing namespaced instances of the Consul Catalog provider.
fields:
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
    default: traefik
    description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
  - name: refreshInterval
    go_name: RefreshInterval
    type: duration
    go_type: ptypes.Duration
  - name: requireConsistent
    go_name: RequireConsistent
    type: boolean
    go_type: bool
  - name: stale
    go_name: Stale
    type: boolean
    go_type: bool
  - name: cache
    go_name: Cache
    type: boolean
    go_type: bool
  - name: exposedByDefault
    go_name: ExposedByDefault
    type: boolean
    go_type: bool
    default: true
  - name: defaultRule
    go_name: DefaultRule
    type: string
    go_type: string
  - name: connectAware
    go_name: ConnectAware
    type: boolean
    go_type: bool
  - name: connectByDefault
    go_name: ConnectByDefault
    type: boolean
    go_type: bool
  - name: serviceName
    go_name: ServiceName
    type: string
    go_type: string
    default: traefik
  - name: watch
    go_name: Watch
    type: boolean
    go_type: bool
  - name: strictChecks
    go_name: StrictChecks
    type: array
    items: string
    go_type: '[]string'
  - name: namespaces
    go_name: Namespaces
    type: array
    items: string
    go_type: '[]string'
representations:
  yaml_path: providers.consulCatalog
  toml_path: providers.consulCatalog
---

# ProviderBuilder

ProviderBuilder is responsible for constructing namespaced instances of the Consul Catalog provider.
