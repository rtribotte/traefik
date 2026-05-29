---
schema_version: 2
kind: middleware-hub
name: JWT
id: hub.middlewares.jwt
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/jwt/config.go#L28
summary: Configuration configures a JWT ACP handler.
fields:
  - name: clientConfig
    go_name: ClientConfig
    type: object
    go_type: httpclient.Config
  - name: signingSecret
    go_name: SigningSecret
    type: string
    go_type: string
  - name: signingSecretBase64Encoded
    go_name: SigningSecretBase64Encoded
    type: boolean
    go_type: bool
  - name: publicKey
    go_name: PublicKey
    type: object
    go_type: FileOrContent
  - name: jwksFile
    go_name: JWKSFile
    type: object
    go_type: FileOrContent
  - name: jwksUrl
    go_name: JWKSURL
    type: string
    go_type: string
  - name: trustedIssuers
    go_name: TrustedIssuers
    type: array
    items: object
    go_type: '[]TrustedIssuer'
  - name: wwwAuthenticate
    go_name: WWWAuthenticate
    type: string
    go_type: string
  - name: headerName
    go_name: HeaderName
    type: string
    go_type: string
    description: HeaderName defines the HTTP header name from which to extract the JWT token. Defaults to "Authorization" when not set.
  - name: stripAuthorizationHeader
    go_name: StripAuthorizationHeader
    type: boolean
    go_type: '*bool'
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
  - name: tokenQueryKey
    go_name: TokenQueryKey
    type: string
    go_type: string
  - name: tokenKey
    go_name: TokenKey
    type: string
    go_type: string
    description: TokenKey defines the form key or query key where to look for the token if not found in the Authorization header.
  - name: claims
    go_name: Claims
    type: string
    go_type: string
  - name: usernameClaim
    go_name: UsernameClaim
    type: string
    go_type: string
    description: UsernameClaim defines the claim used to set the clientUsername in the accessLog.
representations:
  yaml_path: http.middlewares.<name>.plugin.jwt
  toml_path: http.middlewares.<name>.plugin.jwt
  label_prefix: traefik.http.middlewares.<name>.plugin.jwt
---

# JWT

Configuration configures a JWT ACP handler.
