---
schema_version: 2
kind: provider
name: Provider
id: provider.kubernetescrd
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/crd/kubernetes.go#L53
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
  - name: allowCrossNamespace
    go_name: AllowCrossNamespace
    type: boolean
    go_type: bool
  - name: allowExternalNameServices
    go_name: AllowExternalNameServices
    type: boolean
    go_type: bool
  - name: labelSelector
    go_name: LabelSelector
    type: string
    go_type: string
  - name: ingressClass
    go_name: IngressClass
    type: string
    go_type: string
  - name: throttleDuration
    go_name: ThrottleDuration
    type: duration
    go_type: ptypes.Duration
  - name: allowEmptyServices
    go_name: AllowEmptyServices
    type: boolean
    go_type: bool
  - name: nativeLBByDefault
    go_name: NativeLBByDefault
    type: boolean
    go_type: bool
  - name: disableClusterScopeResources
    go_name: DisableClusterScopeResources
    type: boolean
    go_type: bool
representations:
  yaml_path: providers.kubernetesCRD
  toml_path: providers.kubernetesCRD
---

# Provider

Provider holds configurations of the provider.
