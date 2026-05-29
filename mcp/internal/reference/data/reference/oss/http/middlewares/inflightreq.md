---
schema_version: 2
kind: middleware-http
name: InFlightReq
id: http.middlewares.inflightreq
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L561
summary: 'InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/inflightreq/'
fields:
  - name: amount
    go_name: Amount
    type: integer
    go_type: int64
    description: Amount defines the maximum amount of allowed simultaneous in-flight request. The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).
  - name: sourceCriterion
    go_name: SourceCriterion
    type: object
    go_type: '*SourceCriterion'
    type_ref: oss:SourceCriterion
    description: 'SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/inflightreq/#sourcecriterion'
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
representations:
  yaml_path: http.middlewares.<name>.inFlightReq
  toml_path: http.middlewares.<name>.inFlightReq
  label_prefix: traefik.http.middlewares.<name>.inflightreq
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.inFlightReq
---

# InFlightReq

InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/inflightreq/
