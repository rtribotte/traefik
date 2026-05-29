---
schema_version: 2
kind: concept
name: Prometheus
id: concept.prometheus
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/metrics.go#L24
summary: Prometheus can contain specific configuration used by the Prometheus Metrics exporter.
fields:
  - name: buckets
    go_name: Buckets
    type: array
    items: number
    go_type: '[]float64'
    default:
      - 0.1
      - 0.3
      - 1.2
      - 5
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
  - name: entryPoint
    go_name: EntryPoint
    type: string
    go_type: string
    default: traefik
  - name: manualRouting
    go_name: ManualRouting
    type: boolean
    go_type: bool
  - name: headerLabels
    go_name: HeaderLabels
    type: object
    items: string
    go_type: map[string]string
---

# Prometheus

Prometheus can contain specific configuration used by the Prometheus Metrics exporter.
