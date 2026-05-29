---
schema_version: 2
kind: static-section
name: Experimental
id: static.experimental
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/experimental.go#L6
summary: Experimental experimental Traefik features.
fields:
  - name: plugins
    go_name: Plugins
    type: object
    items: object
    go_type: map[string]plugins.Descriptor
  - name: localPlugins
    go_name: LocalPlugins
    type: object
    items: object
    go_type: map[string]plugins.LocalDescriptor
  - name: abortOnPluginFailure
    go_name: AbortOnPluginFailure
    type: boolean
    go_type: bool
  - name: fastProxy
    go_name: FastProxy
    type: object
    go_type: '*FastProxyConfig'
    type_ref: oss:static.FastProxyConfig
    fields:
      - name: debug
        go_name: Debug
        type: boolean
        go_type: bool
  - name: otlplogs
    go_name: OTLPLogs
    type: boolean
    go_type: bool
  - name: knative
    go_name: Knative
    type: boolean
    go_type: bool
  - name: kubernetesIngressNGINX
    go_name: KubernetesIngressNGINX
    type: boolean
    go_type: bool
  - name: kubernetesGateway
    go_name: KubernetesGateway
    type: boolean
    go_type: bool
representations:
  yaml_path: experimental
  toml_path: experimental
---

# Experimental

Experimental experimental Traefik features.
