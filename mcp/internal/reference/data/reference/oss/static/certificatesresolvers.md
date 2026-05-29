---
schema_version: 2
kind: static-section
name: CertificateResolver
id: static.certificatesresolvers
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
    type_ref: oss:static.Configuration
    fields:
      - name: global
        go_name: Global
        type: object
        go_type: '*Global'
        type_ref: oss:Global
        fields:
          - name: checkNewVersion
            go_name: CheckNewVersion
            type: boolean
            go_type: bool
          - name: sendAnonymousUsage
            go_name: SendAnonymousUsage
            type: boolean
            go_type: bool
          - name: notAppendXForwardedFor
            go_name: NotAppendXForwardedFor
            type: boolean
            go_type: bool
      - name: serversTransport
        go_name: ServersTransport
        type: object
        go_type: '*ServersTransport'
        type_ref: oss:ServersTransport
        description: ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.
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
      - name: tcpServersTransport
        go_name: TCPServersTransport
        type: object
        go_type: '*TCPServersTransport'
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
      - name: entryPoints
        go_name: EntryPoints
        type: object
        items: object
        go_type: EntryPoints
        description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
      - name: providers
        go_name: Providers
        type: object
        go_type: '*Providers'
        type_ref: oss:Providers
        fields:
          - name: providersThrottleDuration
            go_name: ProvidersThrottleDuration
            type: duration
            go_type: ptypes.Duration
          - name: precedence
            go_name: Precedence
            type: array
            items: string
            go_type: '[]string'
            default:
              - kubernetesgateway
              - kubernetescrd
              - kubernetes
              - kubernetesingressnginx
              - swarm
              - docker
              - file
              - redis
              - knative
              - consul
              - consulcatalog
              - nomad
              - etcd
              - ecs
              - http
              - zookeeper
              - rest
          - name: docker
            go_name: Docker
            type: object
            go_type: '*docker.Provider'
          - name: swarm
            go_name: Swarm
            type: object
            go_type: '*docker.SwarmProvider'
          - name: file
            go_name: File
            type: object
            go_type: '*file.Provider'
          - name: kubernetesIngress
            go_name: KubernetesIngress
            type: object
            go_type: '*ingress.Provider'
          - name: kubernetesIngressNGINX
            go_name: KubernetesIngressNGINX
            type: object
            go_type: '*ingressnginx.Provider'
          - name: kubernetesCRD
            go_name: KubernetesCRD
            type: object
            go_type: '*crd.Provider'
          - name: kubernetesGateway
            go_name: KubernetesGateway
            type: object
            go_type: '*gateway.Provider'
          - name: knative
            go_name: Knative
            type: object
            go_type: '*knative.Provider'
          - name: rest
            go_name: Rest
            type: object
            go_type: '*rest.Provider'
          - name: consulCatalog
            go_name: ConsulCatalog
            type: object
            go_type: '*consulcatalog.ProviderBuilder'
          - name: nomad
            go_name: Nomad
            type: object
            go_type: '*nomad.ProviderBuilder'
          - name: ecs
            go_name: Ecs
            type: object
            go_type: '*ecs.Provider'
          - name: consul
            go_name: Consul
            type: object
            go_type: '*consul.ProviderBuilder'
          - name: etcd
            go_name: Etcd
            type: object
            go_type: '*etcd.Provider'
          - name: zooKeeper
            go_name: ZooKeeper
            type: object
            go_type: '*zk.Provider'
          - name: redis
            go_name: Redis
            type: object
            go_type: '*redis.Provider'
            description: Redis hold the configs of Redis as bucket in rate limiter.
          - name: http
            go_name: HTTP
            type: object
            go_type: '*http.Provider'
          - name: plugin
            go_name: Plugin
            type: object
            items: object
            go_type: map[string]PluginConf
            description: 'Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/overview/#community-middlewares'
      - name: api
        go_name: API
        type: object
        go_type: '*API'
        type_ref: oss:API
        fields:
          - name: basePath
            go_name: BasePath
            type: string
            go_type: string
            default: /
          - name: insecure
            go_name: Insecure
            type: boolean
            go_type: bool
          - name: dashboard
            go_name: Dashboard
            type: boolean
            go_type: bool
            default: true
          - name: debug
            go_name: Debug
            type: boolean
            go_type: bool
          - name: disableDashboardAd
            go_name: DisableDashboardAd
            type: boolean
            go_type: bool
          - name: dashboardName
            go_name: DashboardName
            type: string
            go_type: string
            default: ""
      - name: metrics
        go_name: Metrics
        type: object
        go_type: '*otypes.Metrics'
        type_ref: oss:Metrics
        description: Metrics enables metrics for this router.
        fields:
          - name: addInternals
            go_name: AddInternals
            type: boolean
            go_type: bool
          - name: prometheus
            go_name: Prometheus
            type: object
            go_type: '*Prometheus'
            type_ref: oss:Prometheus
            fields:
              - name: buckets
                go_name: Buckets
                type: array
                items: number
                go_type: '[]float64'
                default:
                  - 0.1
                  - 0.3
                  - 1.2
                  - 5
              - name: addEntryPointsLabels
                go_name: AddEntryPointsLabels
                type: boolean
                go_type: bool
                default: true
              - name: addRoutersLabels
                go_name: AddRoutersLabels
                type: boolean
                go_type: bool
              - name: addServicesLabels
                go_name: AddServicesLabels
                type: boolean
                go_type: bool
                default: true
              - name: entryPoint
                go_name: EntryPoint
                type: string
                go_type: string
                default: traefik
              - name: manualRouting
                go_name: ManualRouting
                type: boolean
                go_type: bool
              - name: headerLabels
                go_name: HeaderLabels
                type: object
                items: string
                go_type: map[string]string
          - name: datadog
            go_name: Datadog
            type: object
            go_type: '*Datadog'
            type_ref: oss:Datadog
            fields:
              - name: address
                go_name: Address
                type: string
                go_type: string
                default: localhost:8125
                description: Address defines the authentication server address.
              - name: pushInterval
                go_name: PushInterval
                type: duration
                go_type: types.Duration
                default: 10s
              - name: addEntryPointsLabels
                go_name: AddEntryPointsLabels
                type: boolean
                go_type: bool
                default: true
              - name: addRoutersLabels
                go_name: AddRoutersLabels
                type: boolean
                go_type: bool
              - name: addServicesLabels
                go_name: AddServicesLabels
                type: boolean
                go_type: bool
                default: true
              - name: prefix
                go_name: Prefix
                type: string
                go_type: string
                default: traefik
                description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
          - name: statsD
            go_name: StatsD
            type: object
            go_type: '*Statsd'
            type_ref: oss:Statsd
            fields:
              - name: address
                go_name: Address
                type: string
                go_type: string
                default: localhost:8125
                description: Address defines the authentication server address.
              - name: pushInterval
                go_name: PushInterval
                type: duration
                go_type: types.Duration
                default: 10s
              - name: addEntryPointsLabels
                go_name: AddEntryPointsLabels
                type: boolean
                go_type: bool
                default: true
              - name: addRoutersLabels
                go_name: AddRoutersLabels
                type: boolean
                go_type: bool
              - name: addServicesLabels
                go_name: AddServicesLabels
                type: boolean
                go_type: bool
                default: true
              - name: prefix
                go_name: Prefix
                type: string
                go_type: string
                default: traefik
                description: Prefix is the string to add before the current path in the requested URL. It should include a leading slash (/).
          - name: influxDB2
            go_name: InfluxDB2
            type: object
            go_type: '*InfluxDB2'
            type_ref: oss:InfluxDB2
            fields:
              - name: address
                go_name: Address
                type: string
                go_type: string
                default: http://localhost:8086
                description: Address defines the authentication server address.
              - name: token
                go_name: Token
                type: string
                go_type: tTypes.FileOrContent
              - name: pushInterval
                go_name: PushInterval
                type: duration
                go_type: types.Duration
                default: 10s
              - name: org
                go_name: Org
                type: string
                go_type: string
              - name: bucket
                go_name: Bucket
                type: string
                go_type: string
              - name: addEntryPointsLabels
                go_name: AddEntryPointsLabels
                type: boolean
                go_type: bool
                default: true
              - name: addRoutersLabels
                go_name: AddRoutersLabels
                type: boolean
                go_type: bool
              - name: addServicesLabels
                go_name: AddServicesLabels
                type: boolean
                go_type: bool
                default: true
              - name: additionalLabels
                go_name: AdditionalLabels
                type: object
                items: string
                go_type: map[string]string
          - name: otlp
            go_name: OTLP
            type: object
            go_type: '*OTLP'
            type_ref: oss:OTLP
            fields:
              - name: grpc
                go_name: GRPC
                type: object
                go_type: '*OTelGRPC'
                type_ref: oss:OTelGRPC
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: localhost:4317
                  - name: insecure
                    go_name: Insecure
                    type: boolean
                    go_type: bool
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
              - name: http
                go_name: HTTP
                type: object
                go_type: '*OTelHTTP'
                type_ref: oss:OTelHTTP
                default:
                  endpoint: https://localhost:4318
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: https://localhost:4318
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
              - name: addEntryPointsLabels
                go_name: AddEntryPointsLabels
                type: boolean
                go_type: bool
                default: true
              - name: addRoutersLabels
                go_name: AddRoutersLabels
                type: boolean
                go_type: bool
              - name: addServicesLabels
                go_name: AddServicesLabels
                type: boolean
                go_type: bool
                default: true
              - name: explicitBoundaries
                go_name: ExplicitBoundaries
                type: array
                items: number
                go_type: '[]float64'
                default:
                  - 0.005
                  - 0.01
                  - 0.025
                  - 0.05
                  - 0.075
                  - 0.1
                  - 0.25
                  - 0.5
                  - 0.75
                  - 1
                  - 2.5
                  - 5
                  - 7.5
                  - 10
              - name: pushInterval
                go_name: PushInterval
                type: duration
                go_type: types.Duration
                default: 10s
              - name: serviceName
                go_name: ServiceName
                type: string
                go_type: string
                default: traefik
              - name: resourceAttributes
                go_name: ResourceAttributes
                type: object
                items: string
                go_type: map[string]string
      - name: ping
        go_name: Ping
        type: object
        go_type: '*ping.Handler'
        type_ref: oss:Handler
        fields:
          - name: entryPoint
            go_name: EntryPoint
            type: string
            go_type: string
            default: traefik
          - name: manualRouting
            go_name: ManualRouting
            type: boolean
            go_type: bool
          - name: terminatingStatusCode
            go_name: TerminatingStatusCode
            type: integer
            go_type: int
            default: 503
      - name: log
        go_name: Log
        type: object
        go_type: '*otypes.TraefikLog'
        type_ref: oss:TraefikLog
        fields:
          - name: level
            go_name: Level
            type: string
            go_type: string
            default: ERROR
          - name: format
            go_name: Format
            type: string
            go_type: string
            default: common
          - name: noColor
            go_name: NoColor
            type: boolean
            go_type: bool
          - name: filePath
            go_name: FilePath
            type: string
            go_type: string
          - name: maxSize
            go_name: MaxSize
            type: integer
            go_type: int
          - name: maxAge
            go_name: MaxAge
            type: integer
            go_type: int
            description: MaxAge defines the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.
          - name: maxBackups
            go_name: MaxBackups
            type: integer
            go_type: int
          - name: compress
            go_name: Compress
            type: boolean
            go_type: bool
          - name: otlp
            go_name: OTLP
            type: object
            go_type: '*OTelLog'
            type_ref: oss:OTelLog
            fields:
              - name: serviceName
                go_name: ServiceName
                type: string
                go_type: string
                default: traefik
              - name: resourceAttributes
                go_name: ResourceAttributes
                type: object
                items: string
                go_type: map[string]string
              - name: grpc
                go_name: GRPC
                type: object
                go_type: '*OTelGRPC'
                type_ref: oss:OTelGRPC
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: localhost:4317
                  - name: insecure
                    go_name: Insecure
                    type: boolean
                    go_type: bool
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
              - name: http
                go_name: HTTP
                type: object
                go_type: '*OTelHTTP'
                type_ref: oss:OTelHTTP
                default:
                  endpoint: https://localhost:4318
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: https://localhost:4318
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
      - name: accessLog
        go_name: AccessLog
        type: object
        go_type: '*otypes.AccessLog'
        type_ref: oss:AccessLog
        fields:
          - name: filePath
            go_name: FilePath
            type: string
            go_type: string
            default: ""
          - name: format
            go_name: Format
            type: string
            go_type: string
            default: common
          - name: filters
            go_name: Filters
            type: object
            go_type: '*AccessLogFilters'
            type_ref: oss:AccessLogFilters
            default: {}
            fields:
              - name: statusCodes
                go_name: StatusCodes
                type: array
                items: string
                go_type: '[]string'
              - name: retryAttempts
                go_name: RetryAttempts
                type: boolean
                go_type: bool
              - name: minDuration
                go_name: MinDuration
                type: duration
                go_type: types.Duration
          - name: fields
            go_name: Fields
            type: object
            go_type: '*AccessLogFields'
            type_ref: oss:AccessLogFields
            default:
              defaultMode: keep
              headers:
                defaultMode: drop
            fields:
              - name: defaultMode
                go_name: DefaultMode
                type: string
                go_type: string
                default: keep
              - name: names
                go_name: Names
                type: object
                items: string
                go_type: map[string]string
              - name: headers
                go_name: Headers
                type: object
                go_type: '*FieldHeaders'
                type_ref: oss:FieldHeaders
                default:
                  defaultMode: drop
                description: Headers defines custom headers to be sent to the health check endpoint.
                fields:
                  - name: defaultMode
                    go_name: DefaultMode
                    type: string
                    go_type: string
                  - name: names
                    go_name: Names
                    type: object
                    items: string
                    go_type: map[string]string
          - name: bufferingSize
            go_name: BufferingSize
            type: integer
            go_type: int64
          - name: addInternals
            go_name: AddInternals
            type: boolean
            go_type: bool
          - name: dualOutput
            go_name: DualOutput
            type: boolean
            go_type: bool
          - name: otlp
            go_name: OTLP
            type: object
            go_type: '*OTelLog'
            type_ref: oss:OTelLog
            fields:
              - name: serviceName
                go_name: ServiceName
                type: string
                go_type: string
                default: traefik
              - name: resourceAttributes
                go_name: ResourceAttributes
                type: object
                items: string
                go_type: map[string]string
              - name: grpc
                go_name: GRPC
                type: object
                go_type: '*OTelGRPC'
                type_ref: oss:OTelGRPC
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: localhost:4317
                  - name: insecure
                    go_name: Insecure
                    type: boolean
                    go_type: bool
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
              - name: http
                go_name: HTTP
                type: object
                go_type: '*OTelHTTP'
                type_ref: oss:OTelHTTP
                default:
                  endpoint: https://localhost:4318
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: https://localhost:4318
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
      - name: tracing
        go_name: Tracing
        type: object
        go_type: '*Tracing'
        type_ref: oss:Tracing
        description: Tracing enables tracing for this router.
        fields:
          - name: serviceName
            go_name: ServiceName
            type: string
            go_type: string
            default: traefik
          - name: resourceAttributes
            go_name: ResourceAttributes
            type: object
            items: string
            go_type: map[string]string
          - name: capturedRequestHeaders
            go_name: CapturedRequestHeaders
            type: array
            items: string
            go_type: '[]string'
          - name: capturedResponseHeaders
            go_name: CapturedResponseHeaders
            type: array
            items: string
            go_type: '[]string'
          - name: safeQueryParams
            go_name: SafeQueryParams
            type: array
            items: string
            go_type: '[]string'
          - name: sampleRate
            go_name: SampleRate
            type: number
            go_type: float64
            default: 1
          - name: addInternals
            go_name: AddInternals
            type: boolean
            go_type: bool
          - name: otlp
            go_name: OTLP
            type: object
            go_type: '*otypes.OTelTracing'
            type_ref: oss:OTelTracing
            default:
              http:
                endpoint: https://localhost:4318
            fields:
              - name: grpc
                go_name: GRPC
                type: object
                go_type: '*OTelGRPC'
                type_ref: oss:OTelGRPC
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: localhost:4317
                  - name: insecure
                    go_name: Insecure
                    type: boolean
                    go_type: bool
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
              - name: http
                go_name: HTTP
                type: object
                go_type: '*OTelHTTP'
                type_ref: oss:OTelHTTP
                default:
                  endpoint: https://localhost:4318
                fields:
                  - name: endpoint
                    go_name: Endpoint
                    type: string
                    go_type: string
                    default: https://localhost:4318
                  - name: tls
                    go_name: TLS
                    type: object
                    go_type: '*types.ClientTLS'
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
                  - name: headers
                    go_name: Headers
                    type: object
                    items: string
                    go_type: map[string]string
                    description: Headers defines custom headers to be sent to the health check endpoint.
          - name: globalAttributes
            go_name: GlobalAttributes
            type: object
            items: string
            go_type: map[string]string
      - name: hostResolver
        go_name: HostResolver
        type: object
        go_type: '*types.HostResolverConfig'
        type_ref: oss:HostResolverConfig
        fields:
          - name: cnameFlattening
            go_name: CnameFlattening
            type: boolean
            go_type: bool
            default: false
          - name: resolvConfig
            go_name: ResolvConfig
            type: string
            go_type: string
            default: /etc/resolv.conf
          - name: resolvDepth
            go_name: ResolvDepth
            type: integer
            go_type: int
            default: 5
      - name: certificatesResolvers
        go_name: CertificatesResolvers
        type: object
        items: object
        go_type: map[string]CertificateResolver
        type_ref: oss:CertificateResolver
      - name: experimental
        go_name: Experimental
        type: object
        go_type: '*Experimental'
        type_ref: oss:Experimental
        fields:
          - name: plugins
            go_name: Plugins
            type: object
            items: object
            go_type: map[string]plugins.Descriptor
          - name: localPlugins
            go_name: LocalPlugins
            type: object
            items: object
            go_type: map[string]plugins.LocalDescriptor
          - name: abortOnPluginFailure
            go_name: AbortOnPluginFailure
            type: boolean
            go_type: bool
          - name: fastProxy
            go_name: FastProxy
            type: object
            go_type: '*FastProxyConfig'
            type_ref: oss:FastProxyConfig
            fields:
              - name: debug
                go_name: Debug
                type: boolean
                go_type: bool
          - name: otlplogs
            go_name: OTLPLogs
            type: boolean
            go_type: bool
          - name: knative
            go_name: Knative
            type: boolean
            go_type: bool
          - name: kubernetesIngressNGINX
            go_name: KubernetesIngressNGINX
            type: boolean
            go_type: bool
          - name: kubernetesGateway
            go_name: KubernetesGateway
            type: boolean
            go_type: bool
      - name: core
        go_name: Core
        type: object
        go_type: '*Core'
        type_ref: oss:Core
        fields:
          - name: defaultRuleSyntax
            go_name: DefaultRuleSyntax
            type: string
            go_type: string
            default: v3
      - name: spiffe
        go_name: Spiffe
        type: object
        go_type: '*SpiffeClientConfig'
        type_ref: oss:SpiffeClientConfig
        description: Spiffe defines the SPIFFE configuration.
        fields:
          - name: workloadAPIAddr
            go_name: WorkloadAPIAddr
            type: string
            go_type: string
      - name: ocsp
        go_name: OCSP
        type: object
        go_type: '*tls.OCSPConfig'
        type_ref: oss:OCSPConfig
        fields:
          - name: responderOverrides
            go_name: ResponderOverrides
            type: object
            items: string
            go_type: map[string]string
  - name: tailscale
    go_name: Tailscale
    type: object
    go_type: '*struct{}'
representations:
  yaml_path: certificatesResolvers
  toml_path: certificatesResolvers
---

# CertificateResolver

CertificateResolver contains the configuration for the different types of certificates resolver.
