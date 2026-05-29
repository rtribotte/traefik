---
schema_version: 2
kind: middleware-http
name: DigestAuth
id: http.middlewares.digestauth
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L217
summary: 'DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/digestauth/'
fields:
  - name: users
    go_name: Users
    type: array
    items: string
    go_type: Users
    description: Users defines the authorized users. Each user should be declared using the name:realm:encoded-password format.
  - name: usersFile
    go_name: UsersFile
    type: string
    go_type: string
    description: UsersFile is the path to an external file that contains the authorized users for the middleware.
  - name: removeHeader
    go_name: RemoveHeader
    type: boolean
    go_type: bool
    description: RemoveHeader defines whether to remove the authorization header before forwarding the request to the backend.
  - name: realm
    go_name: Realm
    type: string
    go_type: string
    description: 'Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.'
  - name: headerField
    go_name: HeaderField
    type: string
    go_type: string
    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/#headerfield'
representations:
  yaml_path: http.middlewares.<name>.digestAuth
  toml_path: http.middlewares.<name>.digestAuth
  label_prefix: traefik.http.middlewares.<name>.digestauth
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.digestAuth
---

# DigestAuth

DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/digestauth/
