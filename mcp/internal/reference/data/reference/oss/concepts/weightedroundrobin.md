---
schema_version: 2
kind: concept
name: WeightedRoundRobin
id: concept.weightedroundrobin
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L252
summary: WeightedRoundRobin is a weighted round robin load-balancer of services.
fields:
  - name: services
    go_name: Services
    type: array
    items: object
    go_type: '[]WRRService'
    type_ref: oss:WRRService
    description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
    fields:
      - name: name
        go_name: Name
        type: string
        go_type: string
        description: Name defines the name of the referenced IngressRoute resource.
      - name: weight
        go_name: Weight
        type: integer
        go_type: '*int'
        default: 1
        description: Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
  - name: sticky
    go_name: Sticky
    type: object
    go_type: '*Sticky'
    type_ref: oss:Sticky
    description: 'Sticky defines whether sticky sessions are enabled. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/kubernetes/crd/http/traefikservice/#stickiness-and-load-balancing'
    fields:
      - name: cookie
        go_name: Cookie
        type: object
        go_type: '*Cookie'
        type_ref: oss:Cookie
        description: Cookie defines the sticky cookie configuration.
        fields:
          - name: name
            go_name: Name
            type: string
            go_type: string
            description: Name defines the Cookie name.
          - name: secure
            go_name: Secure
            type: boolean
            go_type: bool
            description: Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).
          - name: httpOnly
            go_name: HTTPOnly
            type: boolean
            go_type: bool
            description: HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.
          - name: sameSite
            go_name: SameSite
            type: string
            go_type: string
            description: 'SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite'
          - name: maxAge
            go_name: MaxAge
            type: integer
            go_type: int
            description: MaxAge defines the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.
          - name: path
            go_name: Path
            type: string
            go_type: '*string'
            default: /
            description: 'Path defines the path that must exist in the requested URL for the browser to send the Cookie header. When not provided the cookie will be sent on every request to the domain. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#pathpath-value'
          - name: domain
            go_name: Domain
            type: string
            go_type: string
            description: 'Domain defines the host to which the cookie will be sent. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#domaindomain-value'
  - name: healthCheck
    go_name: HealthCheck
    type: object
    go_type: '*HealthCheck'
    type_ref: oss:HealthCheck
    description: HealthCheck enables automatic self-healthcheck for this service, i.e. whenever one of its children is reported as down, this service becomes aware of it, and takes it into account (i.e. it ignores the down child) when running the load-balancing algorithm. In addition, if the parent of this service also has HealthCheck enabled, this service reports to its parent any status change.
---

# WeightedRoundRobin

WeightedRoundRobin is a weighted round robin load-balancer of services.
