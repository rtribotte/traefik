---
schema_version: 2
kind: service-http
name: Service
id: http.services
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L63
summary: Service holds a service configuration (can only be of one type at the same time).
fields:
  - name: middlewares
    go_name: Middlewares
    type: array
    items: string
    go_type: '[]string'
    description: Middlewares is the list of MiddlewareRef which composes the chain.
  - name: loadBalancer
    go_name: LoadBalancer
    type: object
    go_type: '*ServersLoadBalancer'
    type_ref: oss:ServersLoadBalancer
    fields:
      - name: sticky
        go_name: Sticky
        type: object
        go_type: '*Sticky'
        type_ref: oss:Sticky
        description: 'Sticky defines the sticky sessions configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/load-balancing/service/#sticky-sessions'
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
      - name: servers
        go_name: Servers
        type: array
        items: object
        go_type: '[]Server'
        type_ref: oss:Server
        fields:
          - name: url
            go_name: URL
            type: string
            go_type: string
          - name: weight
            go_name: Weight
            type: integer
            go_type: '*int'
            description: Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
          - name: preservePath
            go_name: PreservePath
            type: boolean
            go_type: bool
      - name: strategy
        go_name: Strategy
        type: string
        go_type: BalancerStrategy
        default: wrr
        description: 'Strategy defines the load balancing strategy between the servers. Supported values are: wrr (Weighed round-robin), p2c (Power of two choices), hrw (Highest Random Weight), and leasttime (Least-Time). RoundRobin value is deprecated and supported for backward compatibility. TODO: when the deprecated RoundRobin value will be removed, set the default kubebuilder value to wrr.'
      - name: healthCheck
        go_name: HealthCheck
        type: object
        go_type: '*ServerHealthCheck'
        type_ref: oss:ServerHealthCheck
        description: HealthCheck enables regular active checks of the responsiveness of the children servers of this load-balancer. To propagate status changes (e.g. all servers of this service are down) upwards, HealthCheck must also be enabled on the parent(s) of this service.
        fields:
          - name: scheme
            go_name: Scheme
            type: string
            go_type: string
            description: Scheme replaces the server URL scheme for the health check endpoint.
          - name: mode
            go_name: Mode
            type: string
            go_type: string
            default: http
            description: 'Mode defines the health check mode. If defined to grpc, will use the gRPC health check protocol to probe the server. Default: http'
          - name: path
            go_name: Path
            type: string
            go_type: string
            description: Path defines the server URL path for the health check endpoint.
          - name: method
            go_name: Method
            type: string
            go_type: string
            description: Method defines the healthcheck method.
          - name: status
            go_name: Status
            type: integer
            go_type: int
            description: Status defines the expected HTTP status code of the response to the health check request.
          - name: port
            go_name: Port
            type: integer
            go_type: int
            description: Port defines the server URL port for the health check endpoint.
          - name: interval
            go_name: Interval
            type: duration
            go_type: ptypes.Duration
            default: 30s
            description: 'Interval defines the frequency of the health check calls for healthy targets. Default: 30s'
          - name: unhealthyInterval
            go_name: UnhealthyInterval
            type: duration
            go_type: '*ptypes.Duration'
            description: 'UnhealthyInterval defines the frequency of the health check calls for unhealthy targets. When UnhealthyInterval is not defined, it defaults to the Interval value. Default: 30s'
          - name: timeout
            go_name: Timeout
            type: duration
            go_type: ptypes.Duration
            default: 5s
            description: 'Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy. Default: 5s'
          - name: hostname
            go_name: Hostname
            type: string
            go_type: string
            description: Hostname defines the value of hostname in the Host header of the health check request.
          - name: followRedirects
            go_name: FollowRedirects
            type: boolean
            go_type: '*bool'
            default: true
            description: 'FollowRedirects defines whether redirects should be followed during the health check calls. Default: true'
          - name: headers
            go_name: Headers
            type: object
            items: string
            go_type: map[string]string
            description: Headers defines custom headers to be sent to the health check endpoint.
      - name: passiveHealthCheck
        go_name: PassiveHealthCheck
        type: object
        go_type: '*PassiveServerHealthCheck'
        type_ref: oss:PassiveServerHealthCheck
        description: PassiveHealthCheck enables passive health checks for children servers of this load-balancer.
        fields:
          - name: failureWindow
            go_name: FailureWindow
            type: duration
            go_type: ptypes.Duration
            default: 10s
            description: FailureWindow defines the time window during which the failed attempts must occur for the server to be marked as unhealthy. It also defines for how long the server will be considered unhealthy.
          - name: maxFailedAttempts
            go_name: MaxFailedAttempts
            type: integer
            go_type: int
            default: 1
            description: MaxFailedAttempts is the number of consecutive failed attempts allowed within the failure window before marking the server as unhealthy.
      - name: passHostHeader
        go_name: PassHostHeader
        type: boolean
        go_type: '*bool'
        default: true
        description: PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.
      - name: responseForwarding
        go_name: ResponseForwarding
        type: object
        go_type: '*ResponseForwarding'
        type_ref: oss:ResponseForwarding
        default:
          flushInterval: 100ms
        description: ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.
        fields:
          - name: flushInterval
            go_name: FlushInterval
            type: duration
            go_type: ptypes.Duration
            default: 100ms
            description: 'FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms'
      - name: serversTransport
        go_name: ServersTransport
        type: string
        go_type: string
        description: ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.
  - name: highestRandomWeight
    go_name: HighestRandomWeight
    type: object
    go_type: '*HighestRandomWeight'
    type_ref: oss:HighestRandomWeight
    description: HighestRandomWeight defines the highest random weight service configuration.
    fields:
      - name: services
        go_name: Services
        type: array
        items: object
        go_type: '[]HRWService'
        type_ref: oss:HRWService
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
      - name: healthCheck
        go_name: HealthCheck
        type: object
        go_type: '*HealthCheck'
        type_ref: oss:HealthCheck
        description: HealthCheck enables automatic self-healthcheck for this service, i.e. whenever one of its children is reported as down, this service becomes aware of it, and takes it into account (i.e. it ignores the down child) when running the load-balancing algorithm. In addition, if the parent of this service also has HealthCheck enabled, this service reports to its parent any status change.
  - name: weighted
    go_name: Weighted
    type: object
    go_type: '*WeightedRoundRobin'
    type_ref: oss:WeightedRoundRobin
    description: Weighted defines the Weighted Round Robin configuration.
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
  - name: mirroring
    go_name: Mirroring
    type: object
    go_type: '*Mirroring'
    type_ref: oss:Mirroring
    description: Mirroring defines the Mirroring service configuration.
    fields:
      - name: service
        go_name: Service
        type: string
        go_type: string
        description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
      - name: mirrorBody
        go_name: MirrorBody
        type: boolean
        go_type: '*bool'
        default: true
        description: MirrorBody defines whether the body of the request should be mirrored. Default value is true.
      - name: maxBodySize
        go_name: MaxBodySize
        type: integer
        go_type: '*int64'
        default: -1
        description: MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.
      - name: mirrors
        go_name: Mirrors
        type: array
        items: object
        go_type: '[]MirrorService'
        type_ref: oss:MirrorService
        description: Mirrors defines the list of mirrors where Traefik will duplicate the traffic.
        fields:
          - name: name
            go_name: Name
            type: string
            go_type: string
            description: Name defines the name of the referenced IngressRoute resource.
          - name: percent
            go_name: Percent
            type: integer
            go_type: int
            description: 'Percent defines the part of the traffic to mirror. Supported values: 0 to 100.'
      - name: healthCheck
        go_name: HealthCheck
        type: object
        go_type: '*HealthCheck'
        type_ref: oss:HealthCheck
        description: Healthcheck defines health checks for ExternalName services.
  - name: failover
    go_name: Failover
    type: object
    go_type: '*Failover'
    type_ref: oss:Failover
    description: Failover defines the Failover service configuration.
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
variants:
  - loadBalancer
  - mirroring
  - weighted
  - failover
  - highestRandomWeight
representations:
  yaml_path: http.services.<name>
  toml_path: http.services.<name>
  label_prefix: traefik.http.services.<name>
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRoute
    spec_path: .spec
---

# Service

Service holds a service configuration (can only be of one type at the same time).
