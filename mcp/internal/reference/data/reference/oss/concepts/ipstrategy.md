---
schema_version: 2
kind: concept
name: IPStrategy
id: concept.ipstrategy
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L479
summary: 'IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipallowlist/#ipstrategy'
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
---

# IPStrategy

IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipallowlist/#ipstrategy
