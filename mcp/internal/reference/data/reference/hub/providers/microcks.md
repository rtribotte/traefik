---
schema_version: 2
kind: provider-hub
name: Microcks
id: hub.providers.microcks
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/provider/microcks/provider.go#L52
summary: Provider is a provider.Provider implementation that queries a Microcks instance for service configurations.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: auth
    go_name: Auth
    type: object
    go_type: '*Auth'
  - name: pollInterval
    go_name: PollInterval
    type: duration
    go_type: ptypes.Duration
  - name: pollTimeout
    go_name: PollTimeout
    type: duration
    go_type: ptypes.Duration
  - name: tls
    go_name: TLS
    type: object
    go_type: '*types.ClientTLS'
representations:
  yaml_path: hub.providers.microcks
  toml_path: hub.providers.microcks
---

# Microcks

Provider is a provider.Provider implementation that queries a Microcks instance for service configurations.
