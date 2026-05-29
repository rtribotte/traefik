---
schema_version: 2
kind: provider
name: Provider
id: provider.knative
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/knative/kubernetes.go#L47
summary: Provider holds configurations of the provider.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: token
    go_name: Token
    type: string
    go_type: string
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
  - name: publicEntrypoints
    go_name: PublicEntrypoints
    type: array
    items: string
    go_type: '[]string'
  - name: publicService
    go_name: PublicService
    type: object
    go_type: ServiceRef
  - name: privateEntrypoints
    go_name: PrivateEntrypoints
    type: array
    items: string
    go_type: '[]string'
  - name: privateService
    go_name: PrivateService
    type: object
    go_type: ServiceRef
  - name: throttleDuration
    go_name: ThrottleDuration
    type: duration
    go_type: ptypes.Duration
representations:
  yaml_path: providers.knative
  toml_path: providers.knative
---

# Provider

Provider holds configurations of the provider.
