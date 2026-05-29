---
schema_version: 2
kind: middleware-hub
name: OIDC
id: hub.middlewares.oidc
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/oidc/config.go#L25
summary: Configuration holds the configuration for the OIDC middleware.
fields:
  - name: clientConfig
    go_name: ClientConfig
    type: object
    go_type: httpclient.Config
  - name: issuer
    go_name: Issuer
    type: string
    go_type: string
  - name: clientId
    go_name: ClientID
    type: string
    go_type: string
  - name: clientSecret
    go_name: ClientSecret
    type: string
    go_type: string
  - name: secret
    go_name: Secret
    type: object
    go_type: '*SecretReference'
  - name: discoveryParams
    go_name: DiscoveryParams
    type: object
    items: string
    go_type: map[string]string
    description: Map of arbitrary query parameters to be added to the openid-configuration well-known URI during the discovery mechanism.
  - name: redirectUrl
    go_name: RedirectURL
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
  - name: postLoginRedirectUrl
    go_name: PostLoginRedirectURL
    type: string
    go_type: string
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: postLogoutRedirectUrl
    go_name: PostLogoutRedirectURL
    type: string
    go_type: '*string'
  - name: backchannelLogoutUrl
    go_name: BackchannelLogoutURL
    type: string
    go_type: string
  - name: backchannelLogoutSessionsRequired
    go_name: BackchannelLogoutSessionsRequired
    type: boolean
    go_type: bool
  - name: disableAuthRedirectionPaths
    go_name: DisableAuthRedirectionPaths
    type: array
    items: string
    go_type: '[]string'
    description: DisableAuthRedirectionPaths disables the automatic redirection to the identity provider when no valid session is found for requests matching this set of path prefixes. It makes the middleware respond with a 401 instead of trying to initialize a new authorization code flow.
  - name: disableLogin
    go_name: DisableLogin
    type: boolean
    go_type: bool
  - name: disableIssuerCheck
    go_name: DisableIssuerCheck
    type: boolean
    go_type: bool
    description: DisableIssuerCheck determines whether the issuer value check between the middleware configuration, the discovery endpoint, and the token's issuer is bypassed.
  - name: trustedIssuer
    go_name: TrustedIssuer
    type: string
    go_type: string
    description: TrustedIssuer allows discovery to work when the issuer_url reported by upstream is mismatched with the discovery URL.
  - name: scopes
    go_name: Scopes
    type: array
    items: string
    go_type: '[]string'
  - name: authParams
    go_name: AuthParams
    type: object
    items: string
    go_type: map[string]string
  - name: stateCookie
    go_name: StateCookie
    type: object
    go_type: '*AuthStateCookie'
  - name: session
    go_name: Session
    type: object
    go_type: '*AuthSession'
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: csrf
    go_name: CSRF
    type: object
    go_type: '*CSRFConfig'
    description: CSRF defines whether the CSRF protection is enabled.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders defines headers that should be added to the request and populated with values extracted from the ID token.
  - name: claims
    go_name: Claims
    type: string
    go_type: string
    description: 'Claims defines an expression to perform validation on the ID token. For example: Equals(`grp`, `admin`) && Equals(`scope`, `deploy`)'
  - name: usernameClaim
    go_name: UsernameClaim
    type: string
    go_type: string
    description: UsernameClaim defines the claim used to set the clientUsername in the accessLog.
  - name: pkce
    go_name: PKCE
    type: boolean
    go_type: bool
    description: PKCE determines whether Proof Key for Code Exchange is enabled for the authorization code flow. When enabled, uses S256 challenge method.
representations:
  yaml_path: http.middlewares.<name>.plugin.oidc
  toml_path: http.middlewares.<name>.plugin.oidc
  label_prefix: traefik.http.middlewares.<name>.plugin.oidc
---

# OIDC

Configuration holds the configuration for the OIDC middleware.
