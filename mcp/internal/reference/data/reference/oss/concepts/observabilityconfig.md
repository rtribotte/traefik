---
schema_version: 2
kind: concept
name: ObservabilityConfig
id: concept.observabilityconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L195
summary: ObservabilityConfig holds the observability configuration for an entry point.
fields:
  - name: accessLogs
    go_name: AccessLogs
    type: boolean
    go_type: '*bool'
    default: true
    description: AccessLogs enables access logs for this router.
  - name: metrics
    go_name: Metrics
    type: boolean
    go_type: '*bool'
    default: true
    description: Metrics enables metrics for this router.
  - name: tracing
    go_name: Tracing
    type: boolean
    go_type: '*bool'
    default: true
    description: Tracing enables tracing for this router.
  - name: traceVerbosity
    go_name: TraceVerbosity
    type: string
    go_type: otypes.TracingVerbosity
    default: minimal
    description: TraceVerbosity defines the verbosity level of the tracing for this router.
---

# ObservabilityConfig

ObservabilityConfig holds the observability configuration for an entry point.
