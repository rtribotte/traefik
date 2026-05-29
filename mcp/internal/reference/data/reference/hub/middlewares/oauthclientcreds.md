---
schema_version: 2
kind: middleware-hub
name: OAuthClientCreds
id: hub.middlewares.oauthclientcreds
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/oauthclientcreds/config.go#L23
summary: Configuration holds the configuration for the OAuth client credentials middleware.
fields:
  - name: clientConfig
    go_name: ClientConfig
    type: object
    go_type: httpclient.Config
  - name: clientID
    go_name: ClientID
    type: string
    go_type: string
  - name: clientSecret
    go_name: ClientSecret
    type: string
    go_type: string
  - name: url
    go_name: URL
    type: string
    go_type: string
  - name: store
    go_name: Store
    type: object
    go_type: '*StoreConfig'
  - name: audience
    go_name: Audience
    type: string
    go_type: string
  - name: scopes
    go_name: Scopes
    type: array
    items: string
    go_type: '[]string'
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders defines headers that should be added to the request and populated with values extracted from the access token.
  - name: claims
    go_name: Claims
    type: string
    go_type: string
    description: 'Claims defines an expression to perform validation on the access token. For example: Equals(`grp`, `admin`) && Equals(`scope`, `deploy`)'
  - name: usernameClaim
    go_name: UsernameClaim
    type: string
    go_type: string
    description: UsernameClaim defines the claim used to set the clientUsername in the accessLog.
representations:
  yaml_path: http.middlewares.<name>.plugin.oauthclientcreds
  toml_path: http.middlewares.<name>.plugin.oauthclientcreds
  label_prefix: traefik.http.middlewares.<name>.plugin.oauthclientcreds
---

# OAuthClientCreds

Configuration holds the configuration for the OAuth client credentials middleware.
