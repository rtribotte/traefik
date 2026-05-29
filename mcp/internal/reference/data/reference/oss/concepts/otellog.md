---
schema_version: 2
kind: concept
name: OTelLog
id: concept.otellog
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/logs.go#L154
summary: OTelLog provides configuration settings for the open-telemetry logger.
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

# OTelLog

OTelLog provides configuration settings for the open-telemetry logger.
