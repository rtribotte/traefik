---
schema_version: 2
kind: concept
name: HTTPConfig
id: concept.httpconfig
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/entrypoints.go#L68
summary: HTTPConfig is the HTTP configuration of an entry point.
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
---

# HTTPConfig

HTTPConfig is the HTTP configuration of an entry point.
