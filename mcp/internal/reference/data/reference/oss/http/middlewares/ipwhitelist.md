---
schema_version: 2
kind: middleware-http
name: IPWhiteList
id: http.middlewares.ipwhitelist
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L536
summary: 'IPWhiteList holds the IP whitelist middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipwhitelist/'
deprecated: true
replaced_by: please use IPAllowList instead
fields:
  - name: sourceRange
    go_name: SourceRange
    type: array
    items: string
    go_type: '[]string'
    description: SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation). Required.
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
representations:
  yaml_path: http.middlewares.<name>.ipWhiteList
  toml_path: http.middlewares.<name>.ipWhiteList
  label_prefix: traefik.http.middlewares.<name>.ipwhitelist
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.ipWhiteList
---

# IPWhiteList

IPWhiteList holds the IP whitelist middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipwhitelist/
