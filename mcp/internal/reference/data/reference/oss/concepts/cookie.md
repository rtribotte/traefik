---
schema_version: 2
kind: concept
name: Cookie
id: concept.cookie
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L333
summary: Cookie holds the sticky configuration based on cookie.
fields:
  - name: name
    go_name: Name
    type: string
    go_type: string
    description: Name defines the Cookie name.
  - name: secure
    go_name: Secure
    type: boolean
    go_type: bool
    description: Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).
  - name: httpOnly
    go_name: HTTPOnly
    type: boolean
    go_type: bool
    description: HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.
  - name: sameSite
    go_name: SameSite
    type: string
    go_type: string
    description: 'SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite'
  - name: maxAge
    go_name: MaxAge
    type: integer
    go_type: int
    description: MaxAge defines the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.
  - name: path
    go_name: Path
    type: string
    go_type: '*string'
    default: /
    description: 'Path defines the path that must exist in the requested URL for the browser to send the Cookie header. When not provided the cookie will be sent on every request to the domain. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#pathpath-value'
  - name: domain
    go_name: Domain
    type: string
    go_type: string
    description: 'Domain defines the host to which the cookie will be sent. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#domaindomain-value'
---

# Cookie

Cookie holds the sticky configuration based on cookie.
