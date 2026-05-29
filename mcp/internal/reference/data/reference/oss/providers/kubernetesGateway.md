---
schema_version: 2
kind: provider
name: Provider
id: provider.kubernetesgateway
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/gateway/kubernetes.go#L65
summary: Provider holds configurations of the provider.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: token
    go_name: Token
    type: string
    go_type: types.FileOrContent
  - name: certAuthFilePath
    go_name: CertAuthFilePath
    type: string
    go_type: string
  - name: namespaces
    go_name: Namespaces
    type: array
    items: string
    go_type: '[]string'
  - name: labelSelector
    go_name: LabelSelector
    type: string
    go_type: string
  - name: throttleDuration
    go_name: ThrottleDuration
    type: duration
    go_type: ptypes.Duration
  - name: experimentalChannel
    go_name: ExperimentalChannel
    type: boolean
    go_type: bool
  - name: statusAddress
    go_name: StatusAddress
    type: object
    go_type: '*StatusAddress'
  - name: nativeLBByDefault
    go_name: NativeLBByDefault
    type: boolean
    go_type: bool
representations:
  yaml_path: providers.kubernetesGateway
  toml_path: providers.kubernetesGateway
---

# Provider

Provider holds configurations of the provider.
