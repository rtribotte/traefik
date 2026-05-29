---
schema_version: 2
kind: provider-hub
name: KubernetesCRD
id: hub.providers.kubernetescrd
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/provider/kubernetescrd/provider.go#L25
summary: Provider wraps the kubernetescrd provider and adds Uplink CRD support. It watches Uplink resources and populates the http.uplinks part of the dynamic configuration.
representations:
  yaml_path: hub.providers.kubernetesCRD
  toml_path: hub.providers.kubernetesCRD
---

# KubernetesCRD

Provider wraps the kubernetescrd provider and adds Uplink CRD support. It watches Uplink resources and populates the http.uplinks part of the dynamic configuration.
