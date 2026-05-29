---
schema_version: 2
kind: service-tcp
name: TCPService
id: tcp.services
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L33
summary: TCPService holds a tcp service configuration (can only be of one type at the same time).
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
variants:
  - loadBalancer
  - weighted
representations:
  yaml_path: tcp.services.<name>
  toml_path: tcp.services.<name>
  label_prefix: traefik.tcp.services.<name>
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteTCP
    spec_path: .spec
---

# TCPService

TCPService holds a tcp service configuration (can only be of one type at the same time).
