---
schema_version: 2
kind: middleware-http
name: RateLimit
id: http.middlewares.ratelimit
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L610
summary: RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is.
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
representations:
  yaml_path: http.middlewares.<name>.rateLimit
  toml_path: http.middlewares.<name>.rateLimit
  label_prefix: traefik.http.middlewares.<name>.ratelimit
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.rateLimit
---

# RateLimit

RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is.
