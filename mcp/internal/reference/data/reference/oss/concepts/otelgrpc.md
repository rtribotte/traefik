---
schema_version: 2
kind: concept
name: OTelGRPC
id: concept.otelgrpc
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/observability/types/otel.go#L6
summary: OTelGRPC provides configuration settings for the gRPC open-telemetry.
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
---

# OTelGRPC

OTelGRPC provides configuration settings for the gRPC open-telemetry.
