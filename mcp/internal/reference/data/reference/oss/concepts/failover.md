---
schema_version: 2
kind: concept
name: Failover
id: concept.failover
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L221
summary: Failover holds the Failover configuration.
fields:
  - name: service
    go_name: Service
    type: string
    go_type: string
    description: Service defines the main service to use.
  - name: fallback
    go_name: Fallback
    type: string
    go_type: string
    description: Fallback defines the fallback service to use when the main service returns an error.
  - name: healthCheck
    go_name: HealthCheck
    type: object
    go_type: '*HealthCheck'
    type_ref: oss:HealthCheck
    description: Healthcheck defines health checks for ExternalName services.
  - name: errors
    go_name: Errors
    type: object
    go_type: '*FailoverError'
    type_ref: oss:FailoverError
    description: Errors defines which errors should trigger the use of the fallback service.
    fields:
      - name: maxRequestBodyBytes
        go_name: MaxRequestBodyBytes
        type: integer
        go_type: '*int64'
        default: -1
        description: MaxRequestBodyBytes defines the maximum size allowed for the body of the request. Default value is -1, which means unlimited size.
      - name: status
        go_name: Status
        type: array
        items: string
        go_type: '[]string'
        description: Status defines the list of status code ranges for which the fallback service should be used.
---

# Failover

Failover holds the Failover configuration.
