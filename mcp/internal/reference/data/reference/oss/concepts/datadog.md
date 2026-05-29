---
schema_version: 2
kind: concept
name: Datadog
id: concept.datadog
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/metrics.go#L43
summary: Datadog contains address and metrics pushing interval configuration.
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    default: localhost:8125
    description: Address defines the authentication server address.
  - name: pushInterval
    go_name: PushInterval
    type: duration
    go_type: types.Duration
    default: 10s
  - name: addEntryPointsLabels
    go_name: AddEntryPointsLabels
    type: boolean
    go_type: bool
    default: true
  - name: addRoutersLabels
    go_name: AddRoutersLabels
    type: boolean
    go_type: bool
  - name: addServicesLabels
    go_name: AddServicesLabels
    type: boolean
    go_type: bool
    default: true
  - name: prefix
    go_name: Prefix
    type: string
    go_type: string
    default: traefik
    description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
---

# Datadog

Datadog contains address and metrics pushing interval configuration.
