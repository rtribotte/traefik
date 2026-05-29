---
schema_version: 2
kind: provider
name: Provider
id: provider.http
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/http/http.go#L34
summary: Provider is a provider.Provider implementation that queries an HTTP(s) endpoint for a configuration.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: pollInterval
    go_name: PollInterval
    type: duration
    go_type: ptypes.Duration
  - name: pollTimeout
    go_name: PollTimeout
    type: duration
    go_type: ptypes.Duration
  - name: headers
    go_name: Headers
    type: object
    items: string
    go_type: map[string]string
    description: Headers defines custom headers to be sent to the health check endpoint.
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
  - name: maxResponseBodySize
    go_name: MaxResponseBodySize
    type: integer
    go_type: int64
    description: MaxResponseBodySize defines the maximum body size in bytes allowed in the response from the authentication server.
representations:
  yaml_path: providers.http
  toml_path: providers.http
---

# Provider

Provider is a provider.Provider implementation that queries an HTTP(s) endpoint for a configuration.
