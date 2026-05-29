---
schema_version: 2
kind: middleware-http
name: BasicAuth
id: http.middlewares.basicauth
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L105
summary: 'BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/'
fields:
  - name: users
    go_name: Users
    type: array
    items: string
    go_type: Users
    description: 'Users is an array of authorized users. Each user must be declared using the name:hashed-password format. Tip: Use htpasswd to generate the passwords.'
  - name: usersFile
    go_name: UsersFile
    type: string
    go_type: string
    description: UsersFile is the path to an external file that contains the authorized users.
  - name: realm
    go_name: Realm
    type: string
    go_type: string
    description: 'Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme. Default: traefik.'
  - name: removeHeader
    go_name: RemoveHeader
    type: boolean
    go_type: bool
    description: 'RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service. Default: false.'
  - name: headerField
    go_name: HeaderField
    type: string
    go_type: string
    description: 'HeaderField defines a header field to store the authenticated user. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/#headerfield'
representations:
  yaml_path: http.middlewares.<name>.basicAuth
  toml_path: http.middlewares.<name>.basicAuth
  label_prefix: traefik.http.middlewares.<name>.basicauth
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.basicAuth
---

# BasicAuth

BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/
