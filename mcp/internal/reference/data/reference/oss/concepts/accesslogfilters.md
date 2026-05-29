---
schema_version: 2
kind: concept
name: AccessLogFilters
id: concept.accesslogfilters
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L82
summary: AccessLogFilters holds filters configuration.
fields:
  - name: statusCodes
    go_name: StatusCodes
    type: array
    items: string
    go_type: '[]string'
  - name: retryAttempts
    go_name: RetryAttempts
    type: boolean
    go_type: bool
  - name: minDuration
    go_name: MinDuration
    type: duration
    go_type: types.Duration
---

# AccessLogFilters

AccessLogFilters holds filters configuration.
