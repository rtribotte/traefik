---
schema_version: 2
kind: static-section
name: TraefikLog
id: static.log
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L39
summary: TraefikLog holds the configuration settings for the traefik logger.
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
representations:
  yaml_path: log
  toml_path: log
---

# TraefikLog

TraefikLog holds the configuration settings for the traefik logger.
