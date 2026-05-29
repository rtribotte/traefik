---
schema_version: 2
kind: concept
name: FailoverError
id: concept.failovererror
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L231
summary: FailoverError holds errors configuration.
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

# FailoverError

FailoverError holds errors configuration.
