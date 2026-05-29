---
schema_version: 2
kind: static-section
name: Providers
id: hub.static.providers
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/static/static_config.go#L109
summary: Providers contains providers configuration.
fields:
  - name: consulCatalogEnterprise
    go_name: ConsulCatalogEnterprise
    type: object
    go_type: '*consulcatalogenterprise.Configuration'
  - name: microcks
    go_name: Microcks
    type: object
    go_type: '*microcks.Provider'
  - name: multicluster
    go_name: Multicluster
    type: object
    go_type: '*multicluster.Configuration'
  - name: nutanixPrismCentral
    go_name: NutanixPrismCentral
    type: object
    go_type: '*nutanixprismcentral.Provider'
representations:
  yaml_path: hub.providers
  toml_path: hub.providers
---

# Providers

Providers contains providers configuration.
