---
schema_version: 2
kind: middleware-hub
name: OAuthIntrospection
id: hub.middlewares.oauthintro
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/oauthintro/config.go#L19
summary: Configuration configures an OAuth 2.0 Token Introspection middleware.
fields:
  - name: clientConfig
    go_name: ClientConfig
    type: object
    go_type: ClientConfig
  - name: tokenSource
    go_name: TokenSource
    type: object
    go_type: token.Source
  - name: claims
    go_name: Claims
    type: string
    go_type: string
  - name: usernameClaim
    go_name: UsernameClaim
    type: string
    go_type: string
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
  - name: wwwAuthenticate
    go_name: WWWAuthenticate
    type: string
    go_type: string
representations:
  yaml_path: http.middlewares.<name>.plugin.oauthintro
  toml_path: http.middlewares.<name>.plugin.oauthintro
  label_prefix: traefik.http.middlewares.<name>.plugin.oauthintro
---

# OAuthIntrospection

Configuration configures an OAuth 2.0 Token Introspection middleware.
