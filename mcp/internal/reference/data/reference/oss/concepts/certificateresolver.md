---
schema_version: 2
kind: concept
name: CertificateResolver
id: concept.certificateresolver
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L133
summary: CertificateResolver contains the configuration for the different types of certificates resolver.
fields:
  - name: acme
    go_name: ACME
    type: object
    go_type: '*acmeprovider.Configuration'
    type_ref: oss:Configuration
    fields:
      - name: http
        go_name: HTTP
        type: object
        go_type: '*HTTPConfiguration'
        type_ref: oss:HTTPConfiguration
        fields:
          - name: routers
            go_name: Routers
            type: object
            items: object
            go_type: map[string]*Router
            type_ref: oss:Router
            fields:
              - name: entryPoints
                go_name: EntryPoints
                type: array
                items: string
                go_type: '[]string'
                description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
              - name: middlewares
                go_name: Middlewares
                type: array
                items: string
                go_type: '[]string'
                description: Middlewares is the list of MiddlewareRef which composes the chain.
              - name: service
                go_name: Service
                type: string
                go_type: string
                description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
              - name: rule
                go_name: Rule
                type: string
                go_type: string
              - name: parentRefs
                go_name: ParentRefs
                type: array
                items: string
                go_type: '[]string'
                description: 'ParentRefs defines references to parent IngressRoute resources for multi-layer routing. When set, this IngressRoute''s routers will be children of the referenced parent IngressRoute''s routers. More info: https://doc.traefik.io/traefik/v3.7/routing/routers/#parentrefs'
              - name: ruleSyntax
                go_name: RuleSyntax
                type: string
                go_type: string
              - name: priority
                go_name: Priority
                type: integer
                go_type: int
                description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
              - name: tls
                go_name: TLS
                type: object
                go_type: '*RouterTLSConfig'
                type_ref: oss:RouterTLSConfig
                description: TLS defines the configuration used to secure the connection to the authentication server.
                fields:
                  - name: options
                    go_name: Options
                    type: string
                    go_type: string
                    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
                  - name: certResolver
                    go_name: CertResolver
                    type: string
                    go_type: string
                    description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
                  - name: domains
                    go_name: Domains
                    type: array
                    items: object
                    go_type: '[]types.Domain'
                    type_ref: oss:Domain
                    description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
                    fields:
                      - name: main
                        go_name: Main
                        type: string
                        go_type: string
                        description: Main defines the main domain name.
                      - name: sans
                        go_name: SANs
                        type: array
                        items: string
                        go_type: '[]string'
                        description: SANs defines the subject alternative domain names.
              - name: observability
                go_name: Observability
                type: object
                go_type: '*RouterObservabilityConfig'
                type_ref: oss:RouterObservabilityConfig
                description: 'Observability defines the observability configuration for a router. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/observability/'
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
          - name: services
            go_name: Services
            type: object
            items: object
            go_type: map[string]*Service
            type_ref: oss:Service
            description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
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
          - name: middlewares
            go_name: Middlewares
            type: object
            items: object
            go_type: map[string]*Middleware
            type_ref: oss:Middleware
            description: Middlewares is the list of MiddlewareRef which composes the chain.
            fields:
              - name: addPrefix
                go_name: AddPrefix
                type: object
                go_type: '*AddPrefix'
                type_ref: oss:AddPrefix
                fields:
                  - name: prefix
                    go_name: Prefix
                    type: string
                    go_type: string
                    description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
              - name: stripPrefix
                go_name: StripPrefix
                type: object
                go_type: '*StripPrefix'
                type_ref: oss:StripPrefix
                fields:
                  - name: prefixes
                    go_name: Prefixes
                    type: array
                    items: string
                    go_type: '[]string'
                    description: Prefixes defines the prefixes to strip from the request URL.
                  - name: forceSlash
                    go_name: ForceSlash
                    type: boolean
                    go_type: '*bool'
              - name: stripPrefixRegex
                go_name: StripPrefixRegex
                type: object
                go_type: '*StripPrefixRegex'
                type_ref: oss:StripPrefixRegex
                fields:
                  - name: regex
                    go_name: Regex
                    type: array
                    items: string
                    go_type: '[]string'
                    description: Regex defines the regular expression to match the path prefix from the request URL.
              - name: replacePath
                go_name: ReplacePath
                type: object
                go_type: '*ReplacePath'
                type_ref: oss:ReplacePath
                fields:
                  - name: path
                    go_name: Path
                    type: string
                    go_type: string
                    description: Path defines the path to use as replacement in the request URL.
              - name: replacePathRegex
                go_name: ReplacePathRegex
                type: object
                go_type: '*ReplacePathRegex'
                type_ref: oss:ReplacePathRegex
                fields:
                  - name: regex
                    go_name: Regex
                    type: string
                    go_type: string
                    description: Regex defines the regular expression used to match and capture the path from the request URL.
                  - name: replacement
                    go_name: Replacement
                    type: string
                    go_type: string
                    description: Replacement defines the replacement path format, which can include captured variables.
              - name: chain
                go_name: Chain
                type: object
                go_type: '*Chain'
                type_ref: oss:Chain
                fields:
                  - name: middlewares
                    go_name: Middlewares
                    type: array
                    items: string
                    go_type: '[]string'
                    description: Middlewares is the list of middleware names which composes the chain.
              - name: ipWhiteList
                go_name: IPWhiteList
                type: object
                go_type: '*IPWhiteList'
                type_ref: oss:IPWhiteList
                description: 'IPWhiteList defines the IPWhiteList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipwhitelist/'
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
              - name: ipAllowList
                go_name: IPAllowList
                type: object
                go_type: '*IPAllowList'
                type_ref: oss:IPAllowList
                description: 'IPAllowList defines the IPAllowList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipallowlist/'
                fields:
                  - name: sourceRange
                    go_name: SourceRange
                    type: array
                    items: string
                    go_type: '[]string'
                    description: SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).
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
                  - name: rejectStatusCode
                    go_name: RejectStatusCode
                    type: integer
                    go_type: int
                    description: RejectStatusCode defines the HTTP status code used for refused requests. If not set, the default is 403 (Forbidden).
              - name: headers
                go_name: Headers
                type: object
                go_type: '*Headers'
                type_ref: oss:Headers
                description: Headers defines custom headers to be sent to the health check endpoint.
                fields:
                  - name: customRequestHeaders
                    go_name: CustomRequestHeaders
                    type: object
                    items: string
                    go_type: map[string]string
                    description: CustomRequestHeaders defines the header names and values to apply to the request.
                  - name: customResponseHeaders
                    go_name: CustomResponseHeaders
                    type: object
                    items: string
                    go_type: map[string]string
                    description: CustomResponseHeaders defines the header names and values to apply to the response.
                  - name: accessControlAllowCredentials
                    go_name: AccessControlAllowCredentials
                    type: boolean
                    go_type: bool
                    description: AccessControlAllowCredentials defines whether the request can include user credentials.
                  - name: accessControlAllowHeaders
                    go_name: AccessControlAllowHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.
                  - name: accessControlAllowMethods
                    go_name: AccessControlAllowMethods
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.
                  - name: accessControlAllowOriginList
                    go_name: AccessControlAllowOriginList
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin "*".
                  - name: accessControlAllowOriginListRegex
                    go_name: AccessControlAllowOriginListRegex
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).
                  - name: accessControlExposeHeaders
                    go_name: AccessControlExposeHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.
                  - name: accessControlMaxAge
                    go_name: AccessControlMaxAge
                    type: integer
                    go_type: int64
                    description: AccessControlMaxAge defines the time that a preflight request may be cached.
                  - name: addVaryHeader
                    go_name: AddVaryHeader
                    type: boolean
                    go_type: bool
                    description: AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.
                  - name: allowedHosts
                    go_name: AllowedHosts
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AllowedHosts defines the fully qualified list of allowed domain names.
                  - name: hostsProxyHeaders
                    go_name: HostsProxyHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.
                  - name: sslProxyHeaders
                    go_name: SSLProxyHeaders
                    type: object
                    items: string
                    go_type: map[string]string
                    description: 'SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: "X-Forwarded-Proto": "https").'
                  - name: stsSeconds
                    go_name: STSSeconds
                    type: integer
                    go_type: '*int64'
                    description: STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.
                  - name: stsIncludeSubdomains
                    go_name: STSIncludeSubdomains
                    type: boolean
                    go_type: bool
                    description: STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.
                  - name: stsPreload
                    go_name: STSPreload
                    type: boolean
                    go_type: bool
                    description: STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.
                  - name: forceSTSHeader
                    go_name: ForceSTSHeader
                    type: boolean
                    go_type: bool
                    description: ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.
                  - name: frameDeny
                    go_name: FrameDeny
                    type: boolean
                    go_type: bool
                    description: FrameDeny defines whether to add the X-Frame-Options header with the DENY value.
                  - name: customFrameOptionsValue
                    go_name: CustomFrameOptionsValue
                    type: string
                    go_type: string
                    description: CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.
                  - name: contentTypeNosniff
                    go_name: ContentTypeNosniff
                    type: boolean
                    go_type: bool
                    description: ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.
                  - name: browserXssFilter
                    go_name: BrowserXSSFilter
                    type: boolean
                    go_type: bool
                    description: BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.
                  - name: customBrowserXSSValue
                    go_name: CustomBrowserXSSValue
                    type: string
                    go_type: string
                    description: CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.
                  - name: contentSecurityPolicy
                    go_name: ContentSecurityPolicy
                    type: string
                    go_type: string
                    description: ContentSecurityPolicy defines the Content-Security-Policy header value.
                  - name: contentSecurityPolicyReportOnly
                    go_name: ContentSecurityPolicyReportOnly
                    type: string
                    go_type: string
                    description: ContentSecurityPolicyReportOnly defines the Content-Security-Policy-Report-Only header value.
                  - name: publicKey
                    go_name: PublicKey
                    type: string
                    go_type: string
                    description: PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.
                  - name: referrerPolicy
                    go_name: ReferrerPolicy
                    type: string
                    go_type: string
                    description: ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.
                  - name: permissionsPolicy
                    go_name: PermissionsPolicy
                    type: string
                    go_type: string
                    description: PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.
                  - name: isDevelopment
                    go_name: IsDevelopment
                    type: boolean
                    go_type: bool
                    description: IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.
                  - name: featurePolicy
                    go_name: FeaturePolicy
                    type: string
                    go_type: '*string'
                  - name: sslRedirect
                    go_name: SSLRedirect
                    type: boolean
                    go_type: '*bool'
                  - name: sslTemporaryRedirect
                    go_name: SSLTemporaryRedirect
                    type: boolean
                    go_type: '*bool'
                  - name: sslHost
                    go_name: SSLHost
                    type: string
                    go_type: '*string'
                  - name: sslForceHost
                    go_name: SSLForceHost
                    type: boolean
                    go_type: '*bool'
              - name: encodedCharacters
                go_name: EncodedCharacters
                type: object
                go_type: '*EncodedCharacters'
                type_ref: oss:EncodedCharacters
                fields:
                  - name: allowEncodedSlash
                    go_name: AllowEncodedSlash
                    type: boolean
                    go_type: bool
                    description: AllowEncodedSlash defines whether requests with encoded slash characters in the path are allowed.
                  - name: allowEncodedBackSlash
                    go_name: AllowEncodedBackSlash
                    type: boolean
                    go_type: bool
                    description: AllowEncodedBackSlash defines whether requests with encoded back slash characters in the path are allowed.
                  - name: allowEncodedNullCharacter
                    go_name: AllowEncodedNullCharacter
                    type: boolean
                    go_type: bool
                    description: AllowEncodedNullCharacter defines whether requests with encoded null characters in the path are allowed.
                  - name: allowEncodedSemicolon
                    go_name: AllowEncodedSemicolon
                    type: boolean
                    go_type: bool
                    description: AllowEncodedSemicolon defines whether requests with encoded semicolon characters in the path are allowed.
                  - name: allowEncodedPercent
                    go_name: AllowEncodedPercent
                    type: boolean
                    go_type: bool
                    description: AllowEncodedPercent defines whether requests with encoded percent characters in the path are allowed.
                  - name: allowEncodedQuestionMark
                    go_name: AllowEncodedQuestionMark
                    type: boolean
                    go_type: bool
                    description: AllowEncodedQuestionMark defines whether requests with encoded question mark characters in the path are allowed.
                  - name: allowEncodedHash
                    go_name: AllowEncodedHash
                    type: boolean
                    go_type: bool
                    description: AllowEncodedHash defines whether requests with encoded hash characters in the path are allowed.
              - name: errors
                go_name: Errors
                type: object
                go_type: '*ErrorPage'
                type_ref: oss:ErrorPage
                description: Errors defines which errors should trigger the use of the fallback service.
                fields:
                  - name: status
                    go_name: Status
                    type: array
                    items: string
                    go_type: '[]string'
                    description: Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).
                  - name: statusRewrites
                    go_name: StatusRewrites
                    type: object
                    items: integer
                    go_type: map[string]int
                    description: 'StatusRewrites defines a mapping of status codes that should be returned instead of the original error status codes. For example: "418": 404 or "410-418": 404'
                  - name: service
                    go_name: Service
                    type: string
                    go_type: string
                    description: Service defines the name of the service that will serve the error page.
                  - name: query
                    go_name: Query
                    type: string
                    go_type: string
                    description: Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL. The {originalStatus} variable can be used in order to insert the upstream status code in the URL. The {url} variable can be used in order to insert the escaped request URL.
                  - name: errorRequestHeaders
                    go_name: ErrorRequestHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: ErrorRequestHeaders defines the list of request headers forwarded to the error page service. When nil (not set), all original request headers are forwarded. Set to an empty list to forward no headers, or list specific headers to forward only those.
              - name: rateLimit
                go_name: RateLimit
                type: object
                go_type: '*RateLimit'
                type_ref: oss:RateLimit
                fields:
                  - name: average
                    go_name: Average
                    type: integer
                    go_type: int64
                    description: Average is the maximum rate, by default in requests/s, allowed for the given source. It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.
                  - name: period
                    go_name: Period
                    type: duration
                    go_type: ptypes.Duration
                    default: 1s
                    description: 'Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period. It defaults to a second.'
                  - name: burst
                    go_name: Burst
                    type: integer
                    go_type: int64
                    default: 1
                    description: Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time. It defaults to 1.
                  - name: sourceCriterion
                    go_name: SourceCriterion
                    type: object
                    go_type: '*SourceCriterion'
                    type_ref: oss:SourceCriterion
                    description: SourceCriterion defines what criterion is used to group requests as originating from a common source. If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).
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
                  - name: redis
                    go_name: Redis
                    type: object
                    go_type: '*Redis'
                    type_ref: oss:Redis
                    description: Redis stores the configuration for using Redis as a bucket in the rate-limiting algorithm. If not specified, Traefik will default to an in-memory bucket for the algorithm.
                    fields:
                      - name: endpoints
                        go_name: Endpoints
                        type: array
                        items: string
                        go_type: '[]string'
                        default:
                          - localhost:6379
                        description: Endpoints contains either a single address or a seed list of host:port addresses. Default value is ["localhost:6379"].
                      - name: tls
                        go_name: TLS
                        type: object
                        go_type: '*types.ClientTLS'
                        type_ref: oss:ClientTLS
                        description: TLS defines TLS-specific configurations, including the CA, certificate, and key, which can be provided as a file path or file content.
                        fields:
                          - name: ca
                            go_name: CA
                            type: string
                            go_type: string
                          - name: cert
                            go_name: Cert
                            type: string
                            go_type: string
                          - name: key
                            go_name: Key
                            type: string
                            go_type: string
                          - name: insecureSkipVerify
                            go_name: InsecureSkipVerify
                            type: boolean
                            go_type: bool
                            description: InsecureSkipVerify defines whether the server certificates should be validated.
                          - name: caOptional
                            go_name: CAOptional
                            type: boolean
                            go_type: '*bool'
                      - name: username
                        go_name: Username
                        type: string
                        go_type: string
                        description: Username defines the username to connect to the Redis server.
                      - name: password
                        go_name: Password
                        type: string
                        go_type: string
                        description: Password defines the password to connect to the Redis server.
                      - name: db
                        go_name: DB
                        type: integer
                        go_type: int
                        description: DB defines the Redis database that will be selected after connecting to the server.
                      - name: poolSize
                        go_name: PoolSize
                        type: integer
                        go_type: int
                        description: PoolSize defines the initial number of socket connections. If the pool runs out of available connections, additional ones will be created beyond PoolSize. This can be limited using MaxActiveConns. Default value is 0, meaning 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
                      - name: minIdleConns
                        go_name: MinIdleConns
                        type: integer
                        go_type: int
                        description: MinIdleConns defines the minimum number of idle connections. Default value is 0, and idle connections are not closed by default.
                      - name: maxActiveConns
                        go_name: MaxActiveConns
                        type: integer
                        go_type: int
                        description: MaxActiveConns defines the maximum number of connections allocated by the pool at a given time. Default value is 0, meaning there is no limit.
                      - name: readTimeout
                        go_name: ReadTimeout
                        type: duration
                        go_type: '*ptypes.Duration'
                        default: 3s
                        description: ReadTimeout defines the timeout for socket read operations. Default value is 3 seconds.
                      - name: writeTimeout
                        go_name: WriteTimeout
                        type: duration
                        go_type: '*ptypes.Duration'
                        default: 3s
                        description: WriteTimeout defines the timeout for socket write operations. Default value is 3 seconds.
                      - name: dialTimeout
                        go_name: DialTimeout
                        type: duration
                        go_type: '*ptypes.Duration'
                        default: 5s
                        description: DialTimeout sets the timeout for establishing new connections. Default value is 5 seconds.
              - name: redirectRegex
                go_name: RedirectRegex
                type: object
                go_type: '*RedirectRegex'
                type_ref: oss:RedirectRegex
                fields:
                  - name: regex
                    go_name: Regex
                    type: string
                    go_type: string
                    description: Regex defines the regex used to match and capture elements from the request URL.
                  - name: replacement
                    go_name: Replacement
                    type: string
                    go_type: string
                    description: Replacement defines how to modify the URL to have the new target URL.
                  - name: permanent
                    go_name: Permanent
                    type: boolean
                    go_type: bool
                    description: Permanent defines whether the redirection is permanent (308).
              - name: redirectScheme
                go_name: RedirectScheme
                type: object
                go_type: '*RedirectScheme'
                type_ref: oss:RedirectScheme
                fields:
                  - name: scheme
                    go_name: Scheme
                    type: string
                    go_type: string
                    description: Scheme defines the scheme of the new URL.
                  - name: port
                    go_name: Port
                    type: string
                    go_type: string
                    description: Port defines the port of the new URL.
                  - name: permanent
                    go_name: Permanent
                    type: boolean
                    go_type: bool
                    description: Permanent defines whether the redirection is permanent. For HTTP GET requests a 301 is returned, otherwise a 308 is returned.
              - name: basicAuth
                go_name: BasicAuth
                type: object
                go_type: '*BasicAuth'
                type_ref: oss:BasicAuth
                fields:
                  - name: users
                    go_name: Users
                    type: array
                    items: string
                    go_type: Users
                    description: 'Users is an array of authorized users. Each user must be declared using the name:hashed-password format. Tip: Use htpasswd to generate the passwords.'
                  - name: usersFile
                    go_name: UsersFile
                    type: string
                    go_type: string
                    description: UsersFile is the path to an external file that contains the authorized users.
                  - name: realm
                    go_name: Realm
                    type: string
                    go_type: string
                    description: 'Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.'
                  - name: removeHeader
                    go_name: RemoveHeader
                    type: boolean
                    go_type: bool
                    description: 'RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.'
                  - name: headerField
                    go_name: HeaderField
                    type: string
                    go_type: string
                    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/#headerfield'
              - name: digestAuth
                go_name: DigestAuth
                type: object
                go_type: '*DigestAuth'
                type_ref: oss:DigestAuth
                fields:
                  - name: users
                    go_name: Users
                    type: array
                    items: string
                    go_type: Users
                    description: Users defines the authorized users. Each user should be declared using the name:realm:encoded-password format.
                  - name: usersFile
                    go_name: UsersFile
                    type: string
                    go_type: string
                    description: UsersFile is the path to an external file that contains the authorized users for the middleware.
                  - name: removeHeader
                    go_name: RemoveHeader
                    type: boolean
                    go_type: bool
                    description: RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.
                  - name: realm
                    go_name: Realm
                    type: string
                    go_type: string
                    description: 'Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.'
                  - name: headerField
                    go_name: HeaderField
                    type: string
                    go_type: string
                    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/#headerfield'
              - name: forwardAuth
                go_name: ForwardAuth
                type: object
                go_type: '*ForwardAuth'
                type_ref: oss:ForwardAuth
                fields:
                  - name: address
                    go_name: Address
                    type: string
                    go_type: string
                    description: Address defines the authentication server address.
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*ClientTLS'
                    type_ref: oss:ClientTLS
                    description: TLS defines the configuration used to secure the connection to the authentication server.
                    fields:
                      - name: ca
                        go_name: CA
                        type: string
                        go_type: string
                      - name: cert
                        go_name: Cert
                        type: string
                        go_type: string
                      - name: key
                        go_name: Key
                        type: string
                        go_type: string
                      - name: insecureSkipVerify
                        go_name: InsecureSkipVerify
                        type: boolean
                        go_type: bool
                        description: InsecureSkipVerify defines whether the server certificates should be validated.
                      - name: caOptional
                        go_name: CAOptional
                        type: boolean
                        go_type: '*bool'
                  - name: trustForwardHeader
                    go_name: TrustForwardHeader
                    type: boolean
                    go_type: '*bool'
                    description: 'TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.'
                  - name: authResponseHeaders
                    go_name: AuthResponseHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.
                  - name: authResponseHeadersRegex
                    go_name: AuthResponseHeadersRegex
                    type: string
                    go_type: string
                    description: 'AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/#authresponseheadersregex'
                  - name: authRequestHeaders
                    go_name: AuthRequestHeaders
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.
                  - name: maxResponseBodySize
                    go_name: MaxResponseBodySize
                    type: integer
                    go_type: '*int64'
                    description: MaxResponseBodySize defines the maximum body size in bytes allowed in the response from the authentication server.
                  - name: addAuthCookiesToResponse
                    go_name: AddAuthCookiesToResponse
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AddAuthCookiesToResponse defines the list of cookies to copy from the authentication server response to the response.
                  - name: headerField
                    go_name: HeaderField
                    type: string
                    go_type: string
                    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/#headerfield'
                  - name: forwardBody
                    go_name: ForwardBody
                    type: boolean
                    go_type: bool
                    description: ForwardBody defines whether to send the request body to the authentication server.
                  - name: maxBodySize
                    go_name: MaxBodySize
                    type: integer
                    go_type: '*int64'
                    default: -1
                    description: MaxBodySize defines the maximum body size in bytes allowed to be forwarded to the authentication server.
                  - name: preserveLocationHeader
                    go_name: PreserveLocationHeader
                    type: boolean
                    go_type: bool
                    description: PreserveLocationHeader defines whether to forward the Location header to the client as is or prefix it with the domain name of the authentication server.
                  - name: preserveRequestMethod
                    go_name: PreserveRequestMethod
                    type: boolean
                    go_type: bool
                    description: PreserveRequestMethod defines whether to preserve the original request method while forwarding the request to the authentication server.
                  - name: authSigninURL
                    go_name: AuthSigninURL
                    type: string
                    go_type: string
                    description: AuthSigninURL specifies the URL to redirect to when the authentication server returns 401 Unauthorized.
              - name: inFlightReq
                go_name: InFlightReq
                type: object
                go_type: '*InFlightReq'
                type_ref: oss:InFlightReq
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
              - name: buffering
                go_name: Buffering
                type: object
                go_type: '*Buffering'
                type_ref: oss:Buffering
                fields:
                  - name: maxRequestBodyBytes
                    go_name: MaxRequestBodyBytes
                    type: integer
                    go_type: int64
                    description: 'MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes). If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).'
                  - name: memRequestBodyBytes
                    go_name: MemRequestBodyBytes
                    type: integer
                    go_type: int64
                    description: 'MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory. Default: 1048576 (1Mi).'
                  - name: maxResponseBodyBytes
                    go_name: MaxResponseBodyBytes
                    type: integer
                    go_type: int64
                    description: 'MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes). If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).'
                  - name: memResponseBodyBytes
                    go_name: MemResponseBodyBytes
                    type: integer
                    go_type: int64
                    description: 'MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory. Default: 1048576 (1Mi).'
                  - name: retryExpression
                    go_name: RetryExpression
                    type: string
                    go_type: string
                    description: 'RetryExpression defines the retry conditions. It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/buffering/#retryexpression'
              - name: circuitBreaker
                go_name: CircuitBreaker
                type: object
                go_type: '*CircuitBreaker'
                type_ref: oss:CircuitBreaker
                fields:
                  - name: expression
                    go_name: Expression
                    type: string
                    go_type: string
                    description: Expression defines the expression that, once matched, opens the circuit breaker and applies the fallback mechanism instead of calling the services.
                  - name: checkPeriod
                    go_name: CheckPeriod
                    type: duration
                    go_type: ptypes.Duration
                    default: 100ms
                    description: CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).
                  - name: fallbackDuration
                    go_name: FallbackDuration
                    type: duration
                    go_type: ptypes.Duration
                    default: 10s
                    description: FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).
                  - name: recoveryDuration
                    go_name: RecoveryDuration
                    type: duration
                    go_type: ptypes.Duration
                    default: 10s
                    description: RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).
                  - name: responseCode
                    go_name: ResponseCode
                    type: integer
                    go_type: int
                    default: 503
                    description: ResponseCode is the status code that the circuit breaker will return while it is in the open state.
              - name: compress
                go_name: Compress
                type: object
                go_type: '*Compress'
                type_ref: oss:Compress
                fields:
                  - name: excludedContentTypes
                    go_name: ExcludedContentTypes
                    type: array
                    items: string
                    go_type: '[]string'
                    description: ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing. `application/grpc` is always excluded.
                  - name: includedContentTypes
                    go_name: IncludedContentTypes
                    type: array
                    items: string
                    go_type: '[]string'
                    description: IncludedContentTypes defines the list of content types to compare the Content-Type header of the responses before compressing.
                  - name: minResponseBodyBytes
                    go_name: MinResponseBodyBytes
                    type: integer
                    go_type: int
                    description: 'MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed. Default: 1024.'
                  - name: encodings
                    go_name: Encodings
                    type: array
                    items: string
                    go_type: '[]string'
                    default:
                      - gzip
                      - br
                      - zstd
                    description: Encodings defines the list of supported compression algorithms.
                  - name: defaultEncoding
                    go_name: DefaultEncoding
                    type: string
                    go_type: string
                    description: DefaultEncoding specifies the default encoding if the `Accept-Encoding` header is not in the request or contains a wildcard (`*`).
              - name: passTLSClientCert
                go_name: PassTLSClientCert
                type: object
                go_type: '*PassTLSClientCert'
                type_ref: oss:PassTLSClientCert
                fields:
                  - name: pem
                    go_name: PEM
                    type: boolean
                    go_type: bool
                    description: PEM sets the X-Forwarded-Tls-Client-Cert header with the certificate.
                  - name: info
                    go_name: Info
                    type: object
                    go_type: '*TLSClientCertificateInfo'
                    type_ref: oss:TLSClientCertificateInfo
                    description: Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.
                    fields:
                      - name: notAfter
                        go_name: NotAfter
                        type: boolean
                        go_type: bool
                        description: NotAfter defines whether to add the Not After information from the Validity part.
                      - name: notBefore
                        go_name: NotBefore
                        type: boolean
                        go_type: bool
                        description: NotBefore defines whether to add the Not Before information from the Validity part.
                      - name: sans
                        go_name: Sans
                        type: boolean
                        go_type: bool
                        description: Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.
                      - name: serialNumber
                        go_name: SerialNumber
                        type: boolean
                        go_type: bool
                        description: SerialNumber defines whether to add the client serialNumber information.
                      - name: subject
                        go_name: Subject
                        type: object
                        go_type: '*TLSClientCertificateSubjectDNInfo'
                        type_ref: oss:TLSClientCertificateSubjectDNInfo
                        description: Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.
                        fields:
                          - name: country
                            go_name: Country
                            type: boolean
                            go_type: bool
                            description: Country defines whether to add the country information into the subject.
                          - name: province
                            go_name: Province
                            type: boolean
                            go_type: bool
                            description: Province defines whether to add the province information into the subject.
                          - name: locality
                            go_name: Locality
                            type: boolean
                            go_type: bool
                            description: Locality defines whether to add the locality information into the subject.
                          - name: organization
                            go_name: Organization
                            type: boolean
                            go_type: bool
                            description: Organization defines whether to add the organization information into the subject.
                          - name: organizationalUnit
                            go_name: OrganizationalUnit
                            type: boolean
                            go_type: bool
                            description: OrganizationalUnit defines whether to add the organizationalUnit information into the subject.
                          - name: commonName
                            go_name: CommonName
                            type: boolean
                            go_type: bool
                            description: CommonName defines whether to add the organizationalUnit information into the subject.
                          - name: serialNumber
                            go_name: SerialNumber
                            type: boolean
                            go_type: bool
                            description: SerialNumber defines whether to add the serialNumber information into the subject.
                          - name: domainComponent
                            go_name: DomainComponent
                            type: boolean
                            go_type: bool
                            description: DomainComponent defines whether to add the domainComponent information into the subject.
                      - name: issuer
                        go_name: Issuer
                        type: object
                        go_type: '*TLSClientCertificateIssuerDNInfo'
                        type_ref: oss:TLSClientCertificateIssuerDNInfo
                        description: Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.
                        fields:
                          - name: country
                            go_name: Country
                            type: boolean
                            go_type: bool
                            description: Country defines whether to add the country information into the issuer.
                          - name: province
                            go_name: Province
                            type: boolean
                            go_type: bool
                            description: Province defines whether to add the province information into the issuer.
                          - name: locality
                            go_name: Locality
                            type: boolean
                            go_type: bool
                            description: Locality defines whether to add the locality information into the issuer.
                          - name: organization
                            go_name: Organization
                            type: boolean
                            go_type: bool
                            description: Organization defines whether to add the organization information into the issuer.
                          - name: commonName
                            go_name: CommonName
                            type: boolean
                            go_type: bool
                            description: CommonName defines whether to add the organizationalUnit information into the issuer.
                          - name: serialNumber
                            go_name: SerialNumber
                            type: boolean
                            go_type: bool
                            description: SerialNumber defines whether to add the serialNumber information into the issuer.
                          - name: domainComponent
                            go_name: DomainComponent
                            type: boolean
                            go_type: bool
                            description: DomainComponent defines whether to add the domainComponent information into the issuer.
              - name: retry
                go_name: Retry
                type: object
                go_type: '*Retry'
                type_ref: oss:Retry
                fields:
                  - name: attempts
                    go_name: Attempts
                    type: integer
                    go_type: int
                    description: Attempts defines how many times the request should be retried.
                  - name: timeout
                    go_name: Timeout
                    type: duration
                    go_type: ptypes.Duration
                    description: Timeout defines how much time the middleware is allowed to retry the request.
                  - name: initialInterval
                    go_name: InitialInterval
                    type: duration
                    go_type: ptypes.Duration
                    description: InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
                  - name: maxRequestBodyBytes
                    go_name: MaxRequestBodyBytes
                    type: integer
                    go_type: '*int64'
                    default: 2097152
                    description: MaxRequestBodyBytes defines the maximum size for the request body.
                  - name: status
                    go_name: Status
                    type: array
                    items: string
                    go_type: '[]string'
                    description: Status defines the range of HTTP status codes to retry on.
                  - name: disableRetryOnNetworkError
                    go_name: DisableRetryOnNetworkError
                    type: boolean
                    go_type: bool
                    description: DisableRetryOnNetworkError defines whether to disable the retry if an error occurs when transmitting the request to the server.
                  - name: retryNonIdempotentMethod
                    go_name: RetryNonIdempotentMethod
                    type: boolean
                    go_type: bool
                    description: RetryNonIdempotentMethod activates the retry for non-idempotent methods (POST, LOCK, PATCH)
              - name: contentType
                go_name: ContentType
                type: object
                go_type: '*ContentType'
                type_ref: oss:ContentType
                fields:
                  - name: autoDetect
                    go_name: AutoDetect
                    type: boolean
                    go_type: '*bool'
                    description: AutoDetect specifies whether to let the `Content-Type` header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response.
              - name: grpcWeb
                go_name: GrpcWeb
                type: object
                go_type: '*GrpcWeb'
                type_ref: oss:GrpcWeb
                fields:
                  - name: allowOrigins
                    go_name: AllowOrigins
                    type: array
                    items: string
                    go_type: '[]string'
                    description: AllowOrigins is a list of allowable origins. Can also be a wildcard origin "*".
              - name: plugin
                go_name: Plugin
                type: object
                items: object
                go_type: map[string]PluginConf
                description: 'Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/overview/#community-middlewares'
          - name: models
            go_name: Models
            type: object
            items: object
            go_type: map[string]*Model
            type_ref: oss:Model
            fields:
              - name: middlewares
                go_name: Middlewares
                type: array
                items: string
                go_type: '[]string'
                description: Middlewares is the list of MiddlewareRef which composes the chain.
              - name: tls
                go_name: TLS
                type: object
                go_type: '*RouterTLSConfig'
                type_ref: oss:RouterTLSConfig
                description: TLS defines the configuration used to secure the connection to the authentication server.
                fields:
                  - name: options
                    go_name: Options
                    type: string
                    go_type: string
                    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
                  - name: certResolver
                    go_name: CertResolver
                    type: string
                    go_type: string
                    description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
                  - name: domains
                    go_name: Domains
                    type: array
                    items: object
                    go_type: '[]types.Domain'
                    type_ref: oss:Domain
                    description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
                    fields:
                      - name: main
                        go_name: Main
                        type: string
                        go_type: string
                        description: Main defines the main domain name.
                      - name: sans
                        go_name: SANs
                        type: array
                        items: string
                        go_type: '[]string'
                        description: SANs defines the subject alternative domain names.
              - name: observability
                go_name: Observability
                type: object
                go_type: RouterObservabilityConfig
                type_ref: oss:RouterObservabilityConfig
                description: 'Observability defines the observability configuration for a router. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/observability/'
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
          - name: serversTransports
            go_name: ServersTransports
            type: object
            items: object
            go_type: map[string]*ServersTransport
            type_ref: oss:ServersTransport
            fields:
              - name: serverName
                go_name: ServerName
                type: string
                go_type: string
                description: ServerName defines the server name used to contact the server.
              - name: insecureSkipVerify
                go_name: InsecureSkipVerify
                type: boolean
                go_type: bool
                description: InsecureSkipVerify defines whether the server certificates should be validated.
              - name: rootCAs
                go_name: RootCAs
                type: array
                items: string
                go_type: '[]types.FileOrContent'
                description: RootCAs defines a list of CA certificate Secrets or ConfigMaps used to validate server certificates.
              - name: certificates
                go_name: Certificates
                type: array
                items: object
                go_type: traefiktls.Certificates
                description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
              - name: cipherSuites
                go_name: CipherSuites
                type: array
                items: string
                go_type: '[]string'
                description: CipherSuites defines the cipher suites to use when contacting backend servers.
              - name: minVersion
                go_name: MinVersion
                type: string
                go_type: string
                description: MinVersion defines the minimum TLS version to use when contacting backend servers.
              - name: maxVersion
                go_name: MaxVersion
                type: string
                go_type: string
                description: MaxVersion defines the maximum TLS version to use when contacting backend servers.
              - name: maxIdleConnsPerHost
                go_name: MaxIdleConnsPerHost
                type: integer
                go_type: int
                description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
              - name: forwardingTimeouts
                go_name: ForwardingTimeouts
                type: object
                go_type: '*ForwardingTimeouts'
                type_ref: oss:ForwardingTimeouts
                description: ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
                fields:
                  - name: dialTimeout
                    go_name: DialTimeout
                    type: duration
                    go_type: ptypes.Duration
                    default: 30s
                    description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
                  - name: responseHeaderTimeout
                    go_name: ResponseHeaderTimeout
                    type: duration
                    go_type: ptypes.Duration
                    description: ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
                  - name: idleConnTimeout
                    go_name: IdleConnTimeout
                    type: duration
                    go_type: ptypes.Duration
                    default: 1m30s
                    description: IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.
                  - name: readIdleTimeout
                    go_name: ReadIdleTimeout
                    type: duration
                    go_type: ptypes.Duration
                    description: ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.
                  - name: pingTimeout
                    go_name: PingTimeout
                    type: duration
                    go_type: ptypes.Duration
                    default: 15s
                    description: PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.
              - name: disableHTTP2
                go_name: DisableHTTP2
                type: boolean
                go_type: bool
                description: DisableHTTP2 disables HTTP/2 for connections with backend servers.
              - name: peerCertURI
                go_name: PeerCertURI
                type: string
                go_type: string
                description: PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
              - name: spiffe
                go_name: Spiffe
                type: object
                go_type: '*Spiffe'
                type_ref: oss:Spiffe
                description: Spiffe defines the SPIFFE configuration.
                fields:
                  - name: ids
                    go_name: IDs
                    type: array
                    items: string
                    go_type: '[]string'
                    description: IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).
                  - name: trustDomain
                    go_name: TrustDomain
                    type: string
                    go_type: string
                    description: TrustDomain defines the allowed SPIFFE trust domain.
      - name: tcp
        go_name: TCP
        type: object
        go_type: '*TCPConfiguration'
        type_ref: oss:TCPConfiguration
        fields:
          - name: routers
            go_name: Routers
            type: object
            items: object
            go_type: map[string]*TCPRouter
            type_ref: oss:TCPRouter
            fields:
              - name: entryPoints
                go_name: EntryPoints
                type: array
                items: string
                go_type: '[]string'
                description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
              - name: middlewares
                go_name: Middlewares
                type: array
                items: string
                go_type: '[]string'
                description: Middlewares is the list of MiddlewareRef which composes the chain.
              - name: service
                go_name: Service
                type: string
                go_type: string
                description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
              - name: rule
                go_name: Rule
                type: string
                go_type: string
              - name: ruleSyntax
                go_name: RuleSyntax
                type: string
                go_type: string
              - name: priority
                go_name: Priority
                type: integer
                go_type: int
                description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
              - name: tls
                go_name: TLS
                type: object
                go_type: '*RouterTCPTLSConfig'
                type_ref: oss:RouterTCPTLSConfig
                description: TLS defines the configuration used to secure the connection to the authentication server.
                fields:
                  - name: passthrough
                    go_name: Passthrough
                    type: boolean
                    go_type: bool
                    description: Passthrough defines whether a TLS router will terminate the TLS connection.
                  - name: options
                    go_name: Options
                    type: string
                    go_type: string
                    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
                  - name: certResolver
                    go_name: CertResolver
                    type: string
                    go_type: string
                    description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
                  - name: domains
                    go_name: Domains
                    type: array
                    items: object
                    go_type: '[]types.Domain'
                    type_ref: oss:Domain
                    description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
                    fields:
                      - name: main
                        go_name: Main
                        type: string
                        go_type: string
                        description: Main defines the main domain name.
                      - name: sans
                        go_name: SANs
                        type: array
                        items: string
                        go_type: '[]string'
                        description: SANs defines the subject alternative domain names.
          - name: services
            go_name: Services
            type: object
            items: object
            go_type: map[string]*TCPService
            type_ref: oss:TCPService
            description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
            fields:
              - name: loadBalancer
                go_name: LoadBalancer
                type: object
                go_type: '*TCPServersLoadBalancer'
                type_ref: oss:TCPServersLoadBalancer
                fields:
                  - name: servers
                    go_name: Servers
                    type: array
                    items: object
                    go_type: '[]TCPServer'
                    type_ref: oss:TCPServer
                    fields:
                      - name: address
                        go_name: Address
                        type: string
                        go_type: string
                        description: Address defines the authentication server address.
                      - name: tls
                        go_name: TLS
                        type: boolean
                        go_type: bool
                        description: TLS defines the configuration used to secure the connection to the authentication server.
                  - name: serversTransport
                    go_name: ServersTransport
                    type: string
                    go_type: string
                    description: ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.
                  - name: proxyProtocol
                    go_name: ProxyProtocol
                    type: object
                    go_type: '*ProxyProtocol'
                    type_ref: oss:ProxyProtocol
                    description: ProxyProtocol holds the PROXY Protocol configuration.
                    fields:
                      - name: version
                        go_name: Version
                        type: integer
                        go_type: int
                        default: 2
                        description: Version defines the PROXY Protocol version to use.
                  - name: terminationDelay
                    go_name: TerminationDelay
                    type: integer
                    go_type: '*int'
                    description: TerminationDelay, corresponds to the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).
                  - name: healthCheck
                    go_name: HealthCheck
                    type: object
                    go_type: '*TCPServerHealthCheck'
                    type_ref: oss:TCPServerHealthCheck
                    description: Healthcheck defines health checks for ExternalName services.
                    fields:
                      - name: port
                        go_name: Port
                        type: integer
                        go_type: int
                        description: Port defines the port of a Kubernetes Service. This can be a reference to a named port.
                      - name: send
                        go_name: Send
                        type: string
                        go_type: string
                      - name: expect
                        go_name: Expect
                        type: string
                        go_type: string
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
                        description: Timeout defines how much time the middleware is allowed to retry the request. The value of timeout should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
              - name: weighted
                go_name: Weighted
                type: object
                go_type: '*TCPWeightedRoundRobin'
                type_ref: oss:TCPWeightedRoundRobin
                description: Weighted defines the Weighted Round Robin configuration.
                fields:
                  - name: services
                    go_name: Services
                    type: array
                    items: object
                    go_type: '[]TCPWRRService'
                    type_ref: oss:TCPWRRService
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
                    description: Healthcheck defines health checks for ExternalName services.
          - name: middlewares
            go_name: Middlewares
            type: object
            items: object
            go_type: map[string]*TCPMiddleware
            type_ref: oss:TCPMiddleware
            description: Middlewares is the list of MiddlewareRef which composes the chain.
            fields:
              - name: inFlightConn
                go_name: InFlightConn
                type: object
                go_type: '*TCPInFlightConn'
                type_ref: oss:TCPInFlightConn
                description: InFlightConn defines the InFlightConn middleware configuration.
                fields:
                  - name: amount
                    go_name: Amount
                    type: integer
                    go_type: int64
                    description: Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.
              - name: ipWhiteList
                go_name: IPWhiteList
                type: object
                go_type: '*TCPIPWhiteList'
                type_ref: oss:TCPIPWhiteList
                description: 'IPWhiteList defines the IPWhiteList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipwhitelist/'
                fields:
                  - name: sourceRange
                    go_name: SourceRange
                    type: array
                    items: string
                    go_type: '[]string'
                    description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
              - name: ipAllowList
                go_name: IPAllowList
                type: object
                go_type: '*TCPIPAllowList'
                type_ref: oss:TCPIPAllowList
                description: 'IPAllowList defines the IPAllowList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipallowlist/'
                fields:
                  - name: sourceRange
                    go_name: SourceRange
                    type: array
                    items: string
                    go_type: '[]string'
                    description: SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
          - name: serversTransports
            go_name: ServersTransports
            type: object
            items: object
            go_type: map[string]*TCPServersTransport
            type_ref: oss:TCPServersTransport
            fields:
              - name: dialKeepAlive
                go_name: DialKeepAlive
                type: duration
                go_type: ptypes.Duration
                default: 15s
                description: DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
              - name: dialTimeout
                go_name: DialTimeout
                type: duration
                go_type: ptypes.Duration
                default: 30s
                description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
              - name: proxyProtocol
                go_name: ProxyProtocol
                type: object
                go_type: '*ProxyProtocol'
                type_ref: oss:ProxyProtocol
                description: ProxyProtocol holds the PROXY Protocol configuration.
                fields:
                  - name: version
                    go_name: Version
                    type: integer
                    go_type: int
                    default: 2
                    description: Version defines the PROXY Protocol version to use.
              - name: terminationDelay
                go_name: TerminationDelay
                type: duration
                go_type: ptypes.Duration
                default: 100ms
                description: TerminationDelay, corresponds to the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).
              - name: tls
                go_name: TLS
                type: object
                go_type: '*TLSClientConfig'
                type_ref: oss:TLSClientConfig
                description: TLS defines the configuration used to secure the connection to the authentication server.
                fields:
                  - name: serverName
                    go_name: ServerName
                    type: string
                    go_type: string
                    description: ServerName defines the server name used to contact the server.
                  - name: insecureSkipVerify
                    go_name: InsecureSkipVerify
                    type: boolean
                    go_type: bool
                    description: InsecureSkipVerify disables TLS certificate verification.
                  - name: rootCAs
                    go_name: RootCAs
                    type: array
                    items: string
                    go_type: '[]types.FileOrContent'
                    description: RootCAs defines a list of CA certificate Secrets or ConfigMaps used to validate server certificates.
                  - name: certificates
                    go_name: Certificates
                    type: array
                    items: object
                    go_type: traefiktls.Certificates
                    description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
                  - name: peerCertURI
                    go_name: PeerCertURI
                    type: string
                    go_type: string
                    description: MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
                  - name: spiffe
                    go_name: Spiffe
                    type: object
                    go_type: '*Spiffe'
                    type_ref: oss:Spiffe
                    description: Spiffe defines the SPIFFE configuration.
                    fields:
                      - name: ids
                        go_name: IDs
                        type: array
                        items: string
                        go_type: '[]string'
                        description: IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).
                      - name: trustDomain
                        go_name: TrustDomain
                        type: string
                        go_type: string
                        description: TrustDomain defines the allowed SPIFFE trust domain.
      - name: udp
        go_name: UDP
        type: object
        go_type: '*UDPConfiguration'
        type_ref: oss:UDPConfiguration
        fields:
          - name: routers
            go_name: Routers
            type: object
            items: object
            go_type: map[string]*UDPRouter
            type_ref: oss:UDPRouter
            fields:
              - name: entryPoints
                go_name: EntryPoints
                type: array
                items: string
                go_type: '[]string'
                description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
              - name: service
                go_name: Service
                type: string
                go_type: string
                description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
          - name: services
            go_name: Services
            type: object
            items: object
            go_type: map[string]*UDPService
            type_ref: oss:UDPService
            description: Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
            fields:
              - name: loadBalancer
                go_name: LoadBalancer
                type: object
                go_type: '*UDPServersLoadBalancer'
                type_ref: oss:UDPServersLoadBalancer
                fields:
                  - name: servers
                    go_name: Servers
                    type: array
                    items: object
                    go_type: '[]UDPServer'
                    type_ref: oss:UDPServer
                    fields:
                      - name: address
                        go_name: Address
                        type: string
                        go_type: string
                        description: Address defines the authentication server address.
              - name: weighted
                go_name: Weighted
                type: object
                go_type: '*UDPWeightedRoundRobin'
                type_ref: oss:UDPWeightedRoundRobin
                description: Weighted defines the Weighted Round Robin configuration.
                fields:
                  - name: services
                    go_name: Services
                    type: array
                    items: object
                    go_type: '[]UDPWRRService'
                    type_ref: oss:UDPWRRService
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
      - name: tls
        go_name: TLS
        type: object
        go_type: '*TLSConfiguration'
        type_ref: oss:TLSConfiguration
        description: TLS defines the configuration used to secure the connection to the authentication server.
        fields:
          - name: certificates
            go_name: Certificates
            type: array
            items: object
            go_type: '[]*tls.CertAndStores'
            type_ref: oss:CertAndStores
            description: Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
            fields:
              - name: certFile
                go_name: CertFile
                type: string
                go_type: types.FileOrContent
              - name: keyFile
                go_name: KeyFile
                type: string
                go_type: types.FileOrContent
              - name: stores
                go_name: Stores
                type: array
                items: string
                go_type: '[]string'
          - name: options
            go_name: Options
            type: object
            items: object
            go_type: map[string]tls.Options
            type_ref: oss:Options
            description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
            fields:
              - name: minVersion
                go_name: MinVersion
                type: string
                go_type: string
                description: MinVersion defines the minimum TLS version to use when contacting backend servers.
              - name: maxVersion
                go_name: MaxVersion
                type: string
                go_type: string
                description: MaxVersion defines the maximum TLS version to use when contacting backend servers.
              - name: cipherSuites
                go_name: CipherSuites
                type: array
                items: string
                go_type: '[]string'
                default:
                  - TLS_AES_128_GCM_SHA256
                  - TLS_AES_256_GCM_SHA384
                  - TLS_CHACHA20_POLY1305_SHA256
                  - TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
                  - TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
                  - TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
                  - TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
                  - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
                  - TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
                  - TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
                  - TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
                  - TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
                  - TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
                description: CipherSuites defines the cipher suites to use when contacting backend servers.
              - name: curvePreferences
                go_name: CurvePreferences
                type: array
                items: string
                go_type: '[]string'
                description: 'CurvePreferences defines the preferred elliptic curves. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#curve-preferences'
              - name: clientAuth
                go_name: ClientAuth
                type: object
                go_type: ClientAuth
                type_ref: oss:ClientAuth
                description: ClientAuth defines the server's policy for TLS Client Authentication.
                fields:
                  - name: caFiles
                    go_name: CAFiles
                    type: array
                    items: string
                    go_type: '[]types.FileOrContent'
                  - name: clientAuthType
                    go_name: ClientAuthType
                    type: string
                    go_type: string
                    description: 'ClientAuthType defines the client authentication type to apply. The available values are: "NoClientCert", "RequestClientCert", "VerifyClientCertIfGiven" and "RequireAndVerifyClientCert".'
              - name: sniStrict
                go_name: SniStrict
                type: boolean
                go_type: bool
                description: SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.
              - name: alpnProtocols
                go_name: ALPNProtocols
                type: array
                items: string
                go_type: '[]string'
                default:
                  - h2
                  - http/1.1
                  - acme-tls/1
                description: 'ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#certificates-stores#alpn-protocols'
              - name: disableSessionTickets
                go_name: DisableSessionTickets
                type: boolean
                go_type: bool
                description: DisableSessionTickets disables TLS session resumption via session tickets.
              - name: preferServerCipherSuites
                go_name: PreferServerCipherSuites
                type: boolean
                go_type: '*bool'
                description: PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's. It is enabled automatically when minVersion or maxVersion is set.
          - name: stores
            go_name: Stores
            type: object
            items: object
            go_type: map[string]tls.Store
            type_ref: oss:Store
            fields:
              - name: defaultCertificate
                go_name: DefaultCertificate
                type: object
                go_type: '*Certificate'
                type_ref: oss:Certificate
                description: DefaultCertificate defines the default certificate configuration.
                fields:
                  - name: certFile
                    go_name: CertFile
                    type: string
                    go_type: types.FileOrContent
                  - name: keyFile
                    go_name: KeyFile
                    type: string
                    go_type: types.FileOrContent
              - name: defaultGeneratedCert
                go_name: DefaultGeneratedCert
                type: object
                go_type: '*GeneratedCert'
                type_ref: oss:GeneratedCert
                description: DefaultGeneratedCert defines the default generated certificate configuration.
                fields:
                  - name: resolver
                    go_name: Resolver
                    type: string
                    go_type: string
                    description: Resolver is the name of the resolver that will be used to issue the DefaultCertificate.
                  - name: domain
                    go_name: Domain
                    type: object
                    go_type: '*types.Domain'
                    type_ref: oss:Domain
                    description: Domain is the domain definition for the DefaultCertificate.
                    fields:
                      - name: main
                        go_name: Main
                        type: string
                        go_type: string
                        description: Main defines the main domain name.
                      - name: sans
                        go_name: SANs
                        type: array
                        items: string
                        go_type: '[]string'
                        description: SANs defines the subject alternative domain names.
  - name: tailscale
    go_name: Tailscale
    type: object
    go_type: '*struct{}'
---

# CertificateResolver

CertificateResolver contains the configuration for the different types of certificates resolver.
