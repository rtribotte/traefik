---
schema_version: 2
kind: middleware-http
name: ForwardAuth
id: http.middlewares.forwardauth
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L289
summary: 'ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/'
fields:
  - name: address
    go_name: Address
    type: string
    go_type: string
    description: Address defines the authentication server address.
  - name: tls
    go_name: TLS
    type: object
    go_type: '*ClientTLS'
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
  - name: trustForwardHeader
    go_name: TrustForwardHeader
    type: boolean
    go_type: '*bool'
    description: 'TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.'
  - name: authResponseHeaders
    go_name: AuthResponseHeaders
    type: array
    items: string
    go_type: '[]string'
    description: AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.
  - name: authResponseHeadersRegex
    go_name: AuthResponseHeadersRegex
    type: string
    go_type: string
    description: 'AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/#authresponseheadersregex'
  - name: authRequestHeaders
    go_name: AuthRequestHeaders
    type: array
    items: string
    go_type: '[]string'
    description: AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server. If not set or empty then all request headers are passed.
  - name: maxResponseBodySize
    go_name: MaxResponseBodySize
    type: integer
    go_type: '*int64'
    description: MaxResponseBodySize defines the maximum body size in bytes allowed in the response from the authentication server.
  - name: addAuthCookiesToResponse
    go_name: AddAuthCookiesToResponse
    type: array
    items: string
    go_type: '[]string'
    description: AddAuthCookiesToResponse defines the list of cookies to copy from the authentication server response to the response.
  - name: headerField
    go_name: HeaderField
    type: string
    go_type: string
    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/#headerfield'
  - name: forwardBody
    go_name: ForwardBody
    type: boolean
    go_type: bool
    description: ForwardBody defines whether to send the request body to the authentication server.
  - name: maxBodySize
    go_name: MaxBodySize
    type: integer
    go_type: '*int64'
    default: -1
    description: MaxBodySize defines the maximum body size in bytes allowed to be forwarded to the authentication server.
  - name: preserveLocationHeader
    go_name: PreserveLocationHeader
    type: boolean
    go_type: bool
    description: PreserveLocationHeader defines whether to forward the Location header to the client as is or prefix it with the domain name of the authentication server.
  - name: preserveRequestMethod
    go_name: PreserveRequestMethod
    type: boolean
    go_type: bool
    description: PreserveRequestMethod defines whether to preserve the original request method while forwarding the request to the authentication server.
  - name: authSigninURL
    go_name: AuthSigninURL
    type: string
    go_type: string
    description: AuthSigninURL specifies the URL to redirect to when the authentication server returns 401 Unauthorized.
representations:
  yaml_path: http.middlewares.<name>.forwardAuth
  toml_path: http.middlewares.<name>.forwardAuth
  label_prefix: traefik.http.middlewares.<name>.forwardauth
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.forwardAuth
---

# ForwardAuth

ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwardauth/
