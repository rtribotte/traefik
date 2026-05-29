---
schema_version: 2
kind: static-section
name: EntryPoint
id: static.entrypoints
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L16
summary: EntryPoint holds the entry point configuration.
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    description: Address defines the authentication server address.
  - name: allowACMEByPass
    go_name: AllowACMEByPass
    type: boolean
    go_type: bool
  - name: reusePort
    go_name: ReusePort
    type: boolean
    go_type: bool
  - name: asDefault
    go_name: AsDefault
    type: boolean
    go_type: bool
  - name: transport
    go_name: Transport
    type: object
    go_type: '*EntryPointsTransport'
    type_ref: oss:static.EntryPointsTransport
    default:
      lifeCycle:
        graceTimeOut: 10s
      respondingTimeouts:
        idleTimeout: 3m0s
        readTimeout: 1m0s
    fields:
      - name: lifeCycle
        go_name: LifeCycle
        type: object
        go_type: '*LifeCycle'
        type_ref: oss:LifeCycle
        default:
          graceTimeOut: 10s
        fields:
          - name: requestAcceptGraceTimeout
            go_name: RequestAcceptGraceTimeout
            type: duration
            go_type: ptypes.Duration
          - name: graceTimeOut
            go_name: GraceTimeOut
            type: duration
            go_type: ptypes.Duration
            default: 10s
      - name: respondingTimeouts
        go_name: RespondingTimeouts
        type: object
        go_type: '*RespondingTimeouts'
        type_ref: oss:RespondingTimeouts
        default:
          idleTimeout: 3m0s
          readTimeout: 1m0s
        fields:
          - name: readTimeout
            go_name: ReadTimeout
            type: duration
            go_type: ptypes.Duration
            default: 1m0s
            description: ReadTimeout defines the timeout for socket read operations. Default value is 3 seconds.
          - name: writeTimeout
            go_name: WriteTimeout
            type: duration
            go_type: ptypes.Duration
            description: WriteTimeout defines the timeout for socket write operations. Default value is 3 seconds.
          - name: idleTimeout
            go_name: IdleTimeout
            type: duration
            go_type: ptypes.Duration
            default: 3m0s
      - name: keepAliveMaxTime
        go_name: KeepAliveMaxTime
        type: duration
        go_type: ptypes.Duration
      - name: keepAliveMaxRequests
        go_name: KeepAliveMaxRequests
        type: integer
        go_type: int
  - name: proxyProtocol
    go_name: ProxyProtocol
    type: object
    go_type: '*ProxyProtocol'
    type_ref: oss:static.ProxyProtocol
    description: ProxyProtocol holds the PROXY Protocol configuration.
    fields:
      - name: insecure
        go_name: Insecure
        type: boolean
        go_type: bool
      - name: trustedIPs
        go_name: TrustedIPs
        type: array
        items: string
        go_type: '[]string'
  - name: forwardedHeaders
    go_name: ForwardedHeaders
    type: object
    go_type: '*ForwardedHeaders'
    type_ref: oss:static.ForwardedHeaders
    default: {}
    fields:
      - name: insecure
        go_name: Insecure
        type: boolean
        go_type: bool
      - name: trustedIPs
        go_name: TrustedIPs
        type: array
        items: string
        go_type: '[]string'
      - name: connection
        go_name: Connection
        type: array
        items: string
        go_type: '[]string'
      - name: notAppendXForwardedFor
        go_name: NotAppendXForwardedFor
        type: boolean
        go_type: bool
  - name: http
    go_name: HTTP
    type: object
    go_type: HTTPConfig
    type_ref: oss:static.HTTPConfig
    default:
      maxHeaderBytes: 1048576
      sanitizePath: true
    fields:
      - name: redirections
        go_name: Redirections
        type: object
        go_type: '*Redirections'
        type_ref: oss:Redirections
        fields:
          - name: entryPoint
            go_name: EntryPoint
            type: object
            go_type: '*RedirectEntryPoint'
            type_ref: oss:RedirectEntryPoint
            fields:
              - name: to
                go_name: To
                type: string
                go_type: string
              - name: scheme
                go_name: Scheme
                type: string
                go_type: string
                default: https
                description: Scheme defines the scheme to use for the request to the upstream Kubernetes Service. It defaults to https when Kubernetes Service port is 443, http otherwise.
              - name: permanent
                go_name: Permanent
                type: boolean
                go_type: bool
                default: true
                description: Permanent defines whether the redirection is permanent (308).
              - name: priority
                go_name: Priority
                type: integer
                go_type: int
                default: 9223372036854775807
                description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
      - name: middlewares
        go_name: Middlewares
        type: array
        items: string
        go_type: '[]string'
        description: Middlewares is the list of MiddlewareRef which composes the chain.
      - name: tls
        go_name: TLS
        type: object
        go_type: '*TLSConfig'
        type_ref: oss:TLSConfig
        description: TLS defines the configuration used to secure the connection to the authentication server.
        fields:
          - name: options
            go_name: Options
            type: string
            go_type: string
            description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
          - name: certResolver
            go_name: CertResolver
            type: string
            go_type: string
            description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
          - name: domains
            go_name: Domains
            type: array
            items: object
            go_type: '[]types.Domain'
            type_ref: oss:Domain
            description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
            fields:
              - name: main
                go_name: Main
                type: string
                go_type: string
                description: Main defines the main domain name.
              - name: sans
                go_name: SANs
                type: array
                items: string
                go_type: '[]string'
                description: SANs defines the subject alternative domain names.
      - name: encodedCharacters
        go_name: EncodedCharacters
        type: object
        go_type: '*EncodedCharacters'
        type_ref: oss:EncodedCharacters
        fields:
          - name: allowEncodedSlash
            go_name: AllowEncodedSlash
            type: boolean
            go_type: bool
            description: AllowEncodedSlash defines whether requests with encoded slash characters in the path are allowed.
          - name: allowEncodedBackSlash
            go_name: AllowEncodedBackSlash
            type: boolean
            go_type: bool
            description: AllowEncodedBackSlash defines whether requests with encoded back slash characters in the path are allowed.
          - name: allowEncodedNullCharacter
            go_name: AllowEncodedNullCharacter
            type: boolean
            go_type: bool
            description: AllowEncodedNullCharacter defines whether requests with encoded null characters in the path are allowed.
          - name: allowEncodedSemicolon
            go_name: AllowEncodedSemicolon
            type: boolean
            go_type: bool
            description: AllowEncodedSemicolon defines whether requests with encoded semicolon characters in the path are allowed.
          - name: allowEncodedPercent
            go_name: AllowEncodedPercent
            type: boolean
            go_type: bool
            description: AllowEncodedPercent defines whether requests with encoded percent characters in the path are allowed.
          - name: allowEncodedQuestionMark
            go_name: AllowEncodedQuestionMark
            type: boolean
            go_type: bool
            description: AllowEncodedQuestionMark defines whether requests with encoded question mark characters in the path are allowed.
          - name: allowEncodedHash
            go_name: AllowEncodedHash
            type: boolean
            go_type: bool
            description: AllowEncodedHash defines whether requests with encoded hash characters in the path are allowed.
      - name: encodeQuerySemicolons
        go_name: EncodeQuerySemicolons
        type: boolean
        go_type: bool
      - name: sanitizePath
        go_name: SanitizePath
        type: boolean
        go_type: '*bool'
        default: true
      - name: maxHeaderBytes
        go_name: MaxHeaderBytes
        type: integer
        go_type: int
        default: 1048576
  - name: http2
    go_name: HTTP2
    type: object
    go_type: '*HTTP2Config'
    type_ref: oss:static.HTTP2Config
    default:
      maxConcurrentStreams: 250
      maxDecoderHeaderTableSize: 4096
      maxEncoderHeaderTableSize: 4096
    fields:
      - name: maxConcurrentStreams
        go_name: MaxConcurrentStreams
        type: integer
        go_type: int32
        default: 250
      - name: maxDecoderHeaderTableSize
        go_name: MaxDecoderHeaderTableSize
        type: integer
        go_type: int32
        default: 4096
      - name: maxEncoderHeaderTableSize
        go_name: MaxEncoderHeaderTableSize
        type: integer
        go_type: int32
        default: 4096
  - name: http3
    go_name: HTTP3
    type: object
    go_type: '*HTTP3Config'
    type_ref: oss:static.HTTP3Config
    fields:
      - name: advertisedPort
        go_name: AdvertisedPort
        type: integer
        go_type: int
  - name: udp
    go_name: UDP
    type: object
    go_type: '*UDPConfig'
    type_ref: oss:static.UDPConfig
    default:
      timeout: 3s
    fields:
      - name: timeout
        go_name: Timeout
        type: duration
        go_type: ptypes.Duration
        default: 3s
        description: Timeout defines how much time the middleware is allowed to retry the request. The value of timeout should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
  - name: observability
    go_name: Observability
    type: object
    go_type: '*ObservabilityConfig'
    type_ref: oss:static.ObservabilityConfig
    description: 'Observability defines the observability configuration for a router. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/observability/'
    fields:
      - name: accessLogs
        go_name: AccessLogs
        type: boolean
        go_type: '*bool'
        default: true
        description: AccessLogs enables access logs for this router.
      - name: metrics
        go_name: Metrics
        type: boolean
        go_type: '*bool'
        default: true
        description: Metrics enables metrics for this router.
      - name: tracing
        go_name: Tracing
        type: boolean
        go_type: '*bool'
        default: true
        description: Tracing enables tracing for this router.
      - name: traceVerbosity
        go_name: TraceVerbosity
        type: string
        go_type: otypes.TracingVerbosity
        default: minimal
        description: TraceVerbosity defines the verbosity level of the tracing for this router.
representations:
  yaml_path: entryPoints
  toml_path: entryPoints
---

# EntryPoint

EntryPoint holds the entry point configuration.
