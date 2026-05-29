---
schema_version: 2
kind: provider
name: Provider
id: provider.ecs
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/ecs/ecs.go#L36
summary: Provider holds configurations of the provider.
fields:
  - name: constraints
    go_name: Constraints
    type: string
    go_type: string
  - name: exposedByDefault
    go_name: ExposedByDefault
    type: boolean
    go_type: bool
    default: true
  - name: refreshSeconds
    go_name: RefreshSeconds
    type: integer
    go_type: int
    default: 15
  - name: defaultRule
    go_name: DefaultRule
    type: string
    go_type: string
  - name: clusters
    go_name: Clusters
    type: array
    items: string
    go_type: '[]string'
    default:
      - default
    description: Provider lookup parameters.
  - name: autoDiscoverClusters
    go_name: AutoDiscoverClusters
    type: boolean
    go_type: bool
    default: false
  - name: healthyTasksOnly
    go_name: HealthyTasksOnly
    type: boolean
    go_type: bool
    default: false
  - name: ecsAnywhere
    go_name: ECSAnywhere
    type: boolean
    go_type: bool
  - name: region
    go_name: Region
    type: string
    go_type: string
  - name: accessKeyID
    go_name: AccessKeyID
    type: string
    go_type: string
  - name: secretAccessKey
    go_name: SecretAccessKey
    type: string
    go_type: string
representations:
  yaml_path: providers.ecs
  toml_path: providers.ecs
---

# Provider

Provider holds configurations of the provider.
