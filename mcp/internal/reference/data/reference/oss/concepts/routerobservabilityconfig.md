---
schema_version: 2
kind: concept
name: RouterObservabilityConfig
id: concept.routerobservabilityconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L160
summary: RouterObservabilityConfig holds the observability configuration for a router.
fields:
  - name: accessLogs
    go_name: AccessLogs
    type: boolean
    go_type: '*bool'
    description: AccessLogs enables access logs for this router.
  - name: metrics
    go_name: Metrics
    type: boolean
    go_type: '*bool'
    description: Metrics enables metrics for this router.
  - name: tracing
    go_name: Tracing
    type: boolean
    go_type: '*bool'
    description: Tracing enables tracing for this router.
  - name: traceVerbosity
    go_name: TraceVerbosity
    type: string
    go_type: otypes.TracingVerbosity
    default: minimal
    description: TraceVerbosity defines the verbosity level of the tracing for this router.
---

# RouterObservabilityConfig

RouterObservabilityConfig holds the observability configuration for a router.
