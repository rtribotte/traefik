---
schema_version: 2
kind: concept
name: InfluxDB2
id: concept.influxdb2
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/metrics.go#L90
summary: InfluxDB2 contains address, token and metrics pushing interval configuration.
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    default: http://localhost:8086
    description: Address defines the authentication server address.
  - name: token
    go_name: Token
    type: string
    go_type: tTypes.FileOrContent
  - name: pushInterval
    go_name: PushInterval
    type: duration
    go_type: types.Duration
    default: 10s
  - name: org
    go_name: Org
    type: string
    go_type: string
  - name: bucket
    go_name: Bucket
    type: string
    go_type: string
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
  - name: additionalLabels
    go_name: AdditionalLabels
    type: object
    items: string
    go_type: map[string]string
---

# InfluxDB2

InfluxDB2 contains address, token and metrics pushing interval configuration.
