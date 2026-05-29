---
schema_version: 2
kind: concept
name: OTLP
id: concept.otlp
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/metrics.go#L111
summary: OTLP contains specific configuration used by the OpenTelemetry Metrics exporter.
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
---

# OTLP

OTLP contains specific configuration used by the OpenTelemetry Metrics exporter.
