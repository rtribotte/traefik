---
schema_version: 2
kind: concept
name: SourceCriterion
id: concept.sourcecriterion
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L598
summary: SourceCriterion defines what criterion is used to group requests as originating from a common source. If none are set, the default is to use the request's remote address field. All fields are mutually exclusive.
fields:
  - name: ipStrategy
    go_name: IPStrategy
    type: object
    go_type: '*IPStrategy'
    type_ref: oss:IPStrategy
    fields:
      - name: depth
        go_name: Depth
        type: integer
        go_type: int
        description: Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).
      - name: excludedIPs
        go_name: ExcludedIPs
        type: array
        items: string
        go_type: '[]string'
        description: ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.
      - name: ipv6Subnet
        go_name: IPv6Subnet
        type: integer
        go_type: '*int'
        description: IPv6Subnet configures Traefik to consider all IPv6 addresses from the defined subnet as originating from the same IP. Applies to RemoteAddrStrategy and DepthStrategy.
  - name: requestHeaderName
    go_name: RequestHeaderName
    type: string
    go_type: string
    description: RequestHeaderName defines the name of the header used to group incoming requests.
  - name: requestHost
    go_name: RequestHost
    type: boolean
    go_type: bool
    description: RequestHost defines whether to consider the request Host as the source.
---

# SourceCriterion

SourceCriterion defines what criterion is used to group requests as originating from a common source. If none are set, the default is to use the request's remote address field. All fields are mutually exclusive.
