---
schema_version: 2
kind: concept
name: TCPConfiguration
id: concept.tcpconfiguration
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L15
summary: TCPConfiguration contains all the TCP configuration parameters.
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
---

# TCPConfiguration

TCPConfiguration contains all the TCP configuration parameters.
