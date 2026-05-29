---
schema_version: 2
kind: middleware-hub
name: BasicAuth
id: hub.middlewares.basicauth
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/basicauth/config.go#L7
summary: Config configures a basic auth ACP handler.
fields:
  - name: users
    go_name: Users
    type: object
    go_type: Users
  - name: realm
    go_name: Realm
    type: string
    go_type: string
  - name: stripAuthorizationHeader
    go_name: StripAuthorizationHeader
    type: boolean
    go_type: bool
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
representations:
  yaml_path: http.middlewares.<name>.plugin.basicauth
  toml_path: http.middlewares.<name>.plugin.basicauth
  label_prefix: traefik.http.middlewares.<name>.plugin.basicauth
---

# BasicAuth

Config configures a basic auth ACP handler.
