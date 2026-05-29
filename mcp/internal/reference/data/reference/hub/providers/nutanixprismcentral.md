---
schema_version: 2
kind: provider-hub
name: NutanixPrismCentral
id: hub.providers.nutanixprismcentral
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/provider/nutanixprismcentral/provider.go#L36
summary: Provider is the Nutanix Prism Central provider implementation.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: apiKey
    go_name: APIKey
    type: string
    go_type: string
  - name: username
    go_name: Username
    type: string
    go_type: string
  - name: password
    go_name: Password
    type: string
    go_type: string
  - name: tls
    go_name: TLS
    type: object
    go_type: '*types.ClientTLS'
  - name: pollInterval
    go_name: PollInterval
    type: duration
    go_type: ptypes.Duration
  - name: pollTimeout
    go_name: PollTimeout
    type: duration
    go_type: ptypes.Duration
  - name: filename
    go_name: Filename
    type: string
    go_type: string
  - name: serviceNameCategoryKey
    go_name: ServiceNameCategoryKey
    type: string
    go_type: string
    default: TraefikServiceName
  - name: allowedVPCs
    go_name: AllowedVPCs
    type: array
    items: object
    go_type: '[]VPCReference'
representations:
  yaml_path: hub.providers.nutanixPrismCentral
  toml_path: hub.providers.nutanixPrismCentral
---

# NutanixPrismCentral

Provider is the Nutanix Prism Central provider implementation.
