---
schema_version: 2
kind: concept
name: Redis
id: concept.redis
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L644
summary: Redis holds the Redis configuration.
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
---

# Redis

Redis holds the Redis configuration.
