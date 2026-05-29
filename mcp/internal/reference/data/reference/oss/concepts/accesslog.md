---
schema_version: 2
kind: concept
name: AccessLog
id: concept.accesslog
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L60
summary: AccessLog holds the configuration settings for the access logger (middlewares/accesslog).
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
---

# AccessLog

AccessLog holds the configuration settings for the access logger (middlewares/accesslog).
