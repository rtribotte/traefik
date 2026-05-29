---
schema_version: 2
kind: ingress-nginx-annotations
name: IngressNGINXAnnotations
id: annotations.ingress-nginx
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/ingress-nginx/annotations.go
summary: nginx-style annotations on Kubernetes Ingress objects supported by the Traefik ingress-nginx provider.
tags:
  - Authentication
  - Buffering
  - CORS
  - IP Whitelist
  - Load Balancing & Backend
  - Observability
  - Rate Limiting
  - Retry
  - Routing
  - SSL/TLS
  - Session Affinity
  - Timeout
  - Unsupported Annotations
fields:
  - name: nginx.ingress.kubernetes.io/affinity
    go_name: Affinity
    type: string
    go_type: '*string'
    examples:
      - cookie
  - name: nginx.ingress.kubernetes.io/affinity-canary-behavior
    go_name: AffinityCanaryBehavior
    type: string
    go_type: '*string'
    description: Only the sticky behavior is supported; legacy behavior is not supported.
  - name: nginx.ingress.kubernetes.io/allowlist-source-range
    go_name: AllowlistSourceRange
    type: string
    go_type: '*string'
    examples:
      - ""
      - 192.168.1.0/24, 10.0.0.0/8, 192.168.20.1
      - 192.168.1.0/24
      - 192.168.20.1
  - name: nginx.ingress.kubernetes.io/app-root
    go_name: AppRoot
    type: string
    go_type: '*string'
    examples:
      - foo
      - /foo
  - name: nginx.ingress.kubernetes.io/auth-method
    go_name: AuthMethod
    type: string
    go_type: '*string'
    examples:
      - POST
      - GET
    description: This annotation uses the `proxy_method` directive in Nginx. Thus, it can't be defined on an ingress that already have an `auth-snippet` annotation with the `proxy_method` directive.
  - name: nginx.ingress.kubernetes.io/auth-realm
    go_name: AuthRealm
    type: string
    go_type: '*string'
    examples:
      - Authentication Required
  - name: nginx.ingress.kubernetes.io/auth-response-headers
    go_name: AuthResponseHeaders
    type: string
    go_type: '*string'
    examples:
      - X-Auth-Snippet
      - X-Foo, X-Bar
  - name: nginx.ingress.kubernetes.io/auth-secret
    go_name: AuthSecret
    type: string
    go_type: '*string'
    examples:
      - default/basic-auth
  - name: nginx.ingress.kubernetes.io/auth-secret-type
    go_name: AuthSecretType
    type: string
    go_type: '*string'
    examples:
      - auth-file
  - name: nginx.ingress.kubernetes.io/auth-signin
    go_name: AuthSignin
    type: string
    go_type: '*string'
    examples:
      - https://auth.example.com/oauth2/start?rd=foo
    description: 'Redirects to signin URL on 401 response. It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.   Like ingress-nginx, Traefik automatically appends `rd=$scheme://$best_http_host$escaped_request_uri` so the auth service can redirect back after sign-in; pass an empty `rd` to disable it. On routes without a `Host` matcher, the request''s `Host` header feeds the interpolation and can be abused for open redirects. Scoping routers with a `Host` rule is strongly recommended when relying on this behavior.'
  - name: nginx.ingress.kubernetes.io/auth-snippet
    go_name: AuthSnippet
    type: string
    go_type: '*string'
    examples:
      - |
        add_header X-Auth-Snippet "auth-value";
    description: 'Supported directives: `proxy_method`, `more_set_headers`, `proxy_set_header`, `more_set_input_headers`, `set`, `if`, `return code [text]`. It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.'
  - name: nginx.ingress.kubernetes.io/auth-tls-pass-certificate-to-upstream
    go_name: AuthTLSPassCertificateToUpstream
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
  - name: nginx.ingress.kubernetes.io/auth-tls-secret
    go_name: AuthTLSSecret
    type: string
    go_type: '*string'
    examples:
      - default/ca-secret
      - default/nonexistent-secret
    description: When validation fails, the rejection happens during the TLS handshake rather than returning a 400 Bad Request.
  - name: nginx.ingress.kubernetes.io/auth-tls-verify-client
    go_name: AuthTLSVerifyClient
    type: string
    go_type: '*string'
    examples:
      - optional_no_ca
      - optional
    description: When validation fails, the rejection happens during the TLS handshake rather than returning a 400 Bad Request.
  - name: nginx.ingress.kubernetes.io/auth-type
    go_name: AuthType
    type: string
    go_type: '*string'
    examples:
      - basic
  - name: nginx.ingress.kubernetes.io/auth-url
    go_name: AuthURL
    type: string
    go_type: '*string'
    examples:
      - http://whoami.default.svc/
    description: 'Only URL and response headers copy supported. Forward auth behaves differently than NGINX. It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.'
  - name: nginx.ingress.kubernetes.io/backend-protocol
    go_name: BackendProtocol
    type: string
    go_type: '*string'
    examples:
      - HTTPS
    description: FCGI and AUTO_HTTP not supported.
  - name: nginx.ingress.kubernetes.io/canary
    go_name: Canary
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
  - name: nginx.ingress.kubernetes.io/canary-by-cookie
    go_name: CanaryCookie
    type: string
    go_type: '*string'
    examples:
      - foo.bar
      - foo
  - name: nginx.ingress.kubernetes.io/canary-by-header
    go_name: CanaryHeader
    type: string
    go_type: '*string'
    examples:
      - Foo
  - name: nginx.ingress.kubernetes.io/canary-by-header-pattern
    go_name: CanaryHeaderPattern
    type: string
    go_type: '*string'
    examples:
      - bar(.*)
  - name: nginx.ingress.kubernetes.io/canary-by-header-value
    go_name: CanaryHeaderValue
    type: string
    go_type: '*string'
    examples:
      - bar
  - name: nginx.ingress.kubernetes.io/canary-weight
    go_name: CanaryWeight
    type: integer
    go_type: '*int'
    examples:
      - "10"
  - name: nginx.ingress.kubernetes.io/canary-weight-total
    go_name: CanaryWeightTotal
    type: integer
    go_type: '*int'
    examples:
      - "120"
  - name: nginx.ingress.kubernetes.io/client-body-buffer-size
    go_name: ClientBodyBufferSize
    type: string
    go_type: '*string'
    examples:
      - 10M
      - 10K
    description: ClientBodyBufferSize sets the size of the buffer used for reading request body.
  - name: nginx.ingress.kubernetes.io/configuration-snippet
    go_name: ConfigurationSnippet
    type: string
    go_type: '*string'
    examples:
      - |
        add_header X-Configuration-Snippet "configuration-value";
    description: 'Supported directives: `add_header`, `proxy_method`, `more_set_headers`, `proxy_set_header`, `more_set_input_headers`, `set`, `if`, `return code [text]`. It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.'
  - name: nginx.ingress.kubernetes.io/cors-allow-credentials
    go_name: EnableCORSAllowCredentials
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
  - name: nginx.ingress.kubernetes.io/cors-allow-headers
    go_name: CORSAllowHeaders
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - X-Foo
  - name: nginx.ingress.kubernetes.io/cors-allow-methods
    go_name: CORSAllowMethods
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - PUT, GET, POST, OPTIONS
  - name: nginx.ingress.kubernetes.io/cors-allow-origin
    go_name: CORSAllowOrigin
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - '*'
  - name: nginx.ingress.kubernetes.io/cors-expose-headers
    go_name: CORSExposeHeaders
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - X-Forwarded-For, X-Forwarded-Host
  - name: nginx.ingress.kubernetes.io/cors-max-age
    go_name: CORSMaxAge
    type: integer
    go_type: '*int'
    examples:
      - "42"
  - name: nginx.ingress.kubernetes.io/custom-headers
    go_name: CustomHeaders
    type: string
    go_type: '*string'
    examples:
      - default/custom-headers-configmap
      - other-namespace/custom-headers-configmap
      - default/invalid-header-value-configmap
    description: Header whitelisting, similar to `global-allowed-response-headers` NGINX config is not supported.
  - name: nginx.ingress.kubernetes.io/custom-http-errors
    go_name: CustomHTTPErrors
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - 404,415
      - 404,500
    description: Specifies a comma-separated list of HTTP status codes that should be intercepted and served by an error page backend. When any of these status codes occur, the request is forwarded to the global default backend, or to the backend defined by the [default-backend](#opt-nginx-ingress-kubernetes-iodefault-backend) annotation if specified.
  - name: nginx.ingress.kubernetes.io/default-backend
    go_name: DefaultBackend
    type: string
    go_type: '*string'
    examples:
      - whoami-b
    description: Specifies a fallback service within the same namespace as the Ingress resource used to handle requests when the primary backend service has no active endpoints. If the specified service exposes multiple ports, the first port will receive the traffic.
  - name: nginx.ingress.kubernetes.io/enable-access-log
    go_name: EnableAccessLog
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
      - "false"
    description: Access logs must first be enabled in the [install configuration](../../../install-configuration/observability/logs-and-accesslogs/#access-logs) (globally or per entrypoint) for this annotation to take effect. When access logs are enabled, this annotation allows opting out specific Ingresses by setting it to `"false"`. Conversely, when access logs are disabled on an entrypoint, setting this annotation to `"true"` allows opting in specific Ingresses.
  - name: nginx.ingress.kubernetes.io/enable-cors
    go_name: EnableCORS
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
    description: Partial support.
  - name: nginx.ingress.kubernetes.io/enable-global-auth
    go_name: EnableGlobalAuth
    type: boolean
    go_type: '*bool'
    examples:
      - "false"
  - name: nginx.ingress.kubernetes.io/enable-modsecurity
    go_name: EnableModSecurity
    type: boolean
    go_type: '*bool'
  - name: nginx.ingress.kubernetes.io/enable-owasp-core-rules
    go_name: EnableOWASPCoreRules
    type: boolean
    go_type: '*bool'
  - name: nginx.ingress.kubernetes.io/force-ssl-redirect
    go_name: ForceSSLRedirect
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
    description: Cannot opt-out per route if enabled globally.
  - name: nginx.ingress.kubernetes.io/from-to-www-redirect
    go_name: FromToWwwRedirect
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
    description: Doesn't support wildcard hosts.
  - name: nginx.ingress.kubernetes.io/limit-burst-multiplier
    go_name: LimitBurstMultiplier
    type: integer
    go_type: '*int'
    examples:
      - "10"
      - "0"
    description: Default to a multiplier of 5 if the configured value is less than 1. Exceeding the limit returns `429 Too Many Requests` instead of NGINX's default `503 Service Unavailable`.
  - name: nginx.ingress.kubernetes.io/limit-connections
    go_name: LimitConnections
    type: integer
    go_type: '*int'
    examples:
      - "10"
    description: Exceeding the limit returns `429 Too Many Requests` instead of NGINX's default `503 Service Unavailable`. The concurrent connection limit is evaluated per client IP address. Values less than or equal to `0` are safely ignored.
  - name: nginx.ingress.kubernetes.io/limit-rpm
    go_name: LimitRPM
    type: integer
    go_type: '*int'
    examples:
      - "10"
      - "0"
    description: Exceeding the limit returns `429 Too Many Requests` instead of NGINX's default `503 Service Unavailable`.
  - name: nginx.ingress.kubernetes.io/limit-rps
    go_name: LimitRPS
    type: integer
    go_type: '*int'
    examples:
      - "10"
      - "0"
    description: Exceeding the limit returns `429 Too Many Requests` instead of NGINX's default `503 Service Unavailable`.
  - name: nginx.ingress.kubernetes.io/modsecurity-snippet
    go_name: ModSecuritySnippet
    type: string
    go_type: '*string'
  - name: nginx.ingress.kubernetes.io/modsecurity-transaction-id
    go_name: ModSecurityTransactionID
    type: string
    go_type: '*string'
  - name: nginx.ingress.kubernetes.io/permanent-redirect
    go_name: PermanentRedirect
    type: string
    go_type: '*string'
    examples:
      - https://www.google.com
      - https://foo.bar.com
      - https://www.traefik.io
    description: Defaults to a 301 Moved Permanently status code.
  - name: nginx.ingress.kubernetes.io/permanent-redirect-code
    go_name: PermanentRedirectCode
    type: integer
    go_type: '*int'
    examples:
      - "300"
      - "500"
    description: Only valid 3XX HTTP Status Codes are accepted.
  - name: nginx.ingress.kubernetes.io/proxy-body-size
    go_name: ProxyBodySize
    type: string
    go_type: '*string'
    examples:
      - 10M
    description: ProxyBodySize sets the maximum allowed size of the client request body.
  - name: nginx.ingress.kubernetes.io/proxy-buffer-size
    go_name: ProxyBufferSize
    type: string
    go_type: '*string'
    examples:
      - 16k
    description: ProxyBufferSize sets the size of the memory buffer used for reading the response.
  - name: nginx.ingress.kubernetes.io/proxy-buffering
    go_name: ProxyBuffering
    type: string
    go_type: '*string'
    examples:
      - "on"
    description: ProxyBuffering controls whether response buffering is enabled.
  - name: nginx.ingress.kubernetes.io/proxy-buffers-number
    go_name: ProxyBuffersNumber
    type: integer
    go_type: '*int'
    examples:
      - "8"
    description: ProxyBuffersNumber sets the number of memory buffers used for reading the response.
  - name: nginx.ingress.kubernetes.io/proxy-connect-timeout
    go_name: ProxyConnectTimeout
    type: integer
    go_type: '*int'
    examples:
      - "30"
    description: Timeout can be defined globally at the provider level using the [`proxyConnectTimeout` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxyConnectTimeout).
  - name: nginx.ingress.kubernetes.io/proxy-http-version
    go_name: ProxyHTTPVersion
    type: string
    go_type: '*string'
    examples:
      - "1.0"
      - "1.1"
    description: 'Controls HTTP protocol version for backend communication. Supported value: `"1.1"` (disables HTTP/2 to backend). Value `"1.0"` is not supported and will log a warning.'
  - name: nginx.ingress.kubernetes.io/proxy-max-temp-file-size
    go_name: ProxyMaxTempFileSize
    type: string
    go_type: '*string'
    examples:
      - 100m
    description: ProxyMaxTempFileSize sets the maximum size of a temporary file used to buffer responses.
  - name: nginx.ingress.kubernetes.io/proxy-next-upstream
    go_name: ProxyNextUpstream
    type: string
    go_type: '*string'
    examples:
      - error http_400 non_idempotent
      - "off"
    description: Unlike NGINX, Traefik does not guarantee that retries are sent to a different server. There is no difference between `error` and `timeout`, both are treated as TCP level failure. This configuration can be defined globally at the provider level using the [`proxyNextUpstream` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxyNextUpstream).
  - name: nginx.ingress.kubernetes.io/proxy-next-upstream-timeout
    go_name: ProxyNextUpstreamTimeout
    type: integer
    go_type: '*int'
    examples:
      - "30"
    description: The timeout can be defined globally at the provider level using the [`proxyNextUpstreamTimeout` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxyNextUpstreamTimeout).
  - name: nginx.ingress.kubernetes.io/proxy-next-upstream-tries
    go_name: ProxyNextUpstreamTries
    type: integer
    go_type: '*int'
    examples:
      - "0"
      - "5"
    description: Unlimited retry (0) will be capped to the number of available servers to avoid infinite retries. The value can be defined globally at the provider level using the [`proxyNextUpstreamTries` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxyNextUpstreamTries).
  - name: nginx.ingress.kubernetes.io/proxy-read-timeout
    go_name: ProxyReadTimeout
    type: integer
    go_type: '*int'
    examples:
      - "30"
    description: Timeout can be defined globally at the provider level using the [`proxyReadTimeout` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxyReadTimeout).
  - name: nginx.ingress.kubernetes.io/proxy-request-buffering
    go_name: ProxyRequestBuffering
    type: string
    go_type: '*string'
    examples:
      - "on"
    description: ProxyRequestBuffering controls whether request buffering is enabled.
  - name: nginx.ingress.kubernetes.io/proxy-send-timeout
    go_name: ProxySendTimeout
    type: integer
    go_type: '*int'
    examples:
      - "30"
    description: Timeout can be defined globally at the provider level using the [`proxySendTimeout` option](../../../install-configuration/providers/kubernetes/kubernetes-ingress-nginx/#opt-providers-kubernetesIngressNGINX-proxySendTimeout).
  - name: nginx.ingress.kubernetes.io/proxy-ssl-name
    go_name: ProxySSLName
    type: string
    go_type: '*string'
    examples:
      - whoami.localhost
  - name: nginx.ingress.kubernetes.io/proxy-ssl-secret
    go_name: ProxySSLSecret
    type: string
    go_type: '*string'
    examples:
      - default/ingress-with-proxy-ssl
  - name: nginx.ingress.kubernetes.io/proxy-ssl-server-name
    go_name: ProxySSLServerName
    type: string
    go_type: '*string'
    examples:
      - whoami.localhost
  - name: nginx.ingress.kubernetes.io/proxy-ssl-verify
    go_name: ProxySSLVerify
    type: string
    go_type: '*string'
    examples:
      - "on"
  - name: nginx.ingress.kubernetes.io/rewrite-target
    go_name: RewriteTarget
    type: string
    go_type: '*string'
    examples:
      - /$2
      - /rewritten
      - ""
      - /path
      - $2
  - name: nginx.ingress.kubernetes.io/server-alias
    go_name: ServerAlias
    type: array
    items: string
    go_type: '*[]string'
    examples:
      - shared.localhost
      - alias1.localhost,conflict.localhost
      - alias1.localhost,alias2.localhost
    description: Ignored if the alias conflicts with an existing Ingress Host rule. Ingress Host rules always take precedence.
  - name: nginx.ingress.kubernetes.io/server-snippet
    go_name: ServerSnippet
    type: string
    go_type: '*string'
    examples:
      - |
        add_header X-Server-Snippet "server-value";
    description: 'Supported directives: `add_header`, `proxy_method`, `more_set_headers`, `proxy_set_header`, `more_set_input_headers`, `set`, `if`, `return code [text]`. It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.'
  - name: nginx.ingress.kubernetes.io/service-upstream
    go_name: ServiceUpstream
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
  - name: nginx.ingress.kubernetes.io/session-cookie-domain
    go_name: SessionCookieDomain
    type: string
    go_type: '*string'
    examples:
      - foo.localhost
  - name: nginx.ingress.kubernetes.io/session-cookie-expires
    go_name: SessionCookieExpires
    type: integer
    go_type: '*int'
    examples:
      - "42"
  - name: nginx.ingress.kubernetes.io/session-cookie-max-age
    go_name: SessionCookieMaxAge
    type: integer
    go_type: '*int'
    examples:
      - "42"
  - name: nginx.ingress.kubernetes.io/session-cookie-name
    go_name: SessionCookieName
    type: string
    go_type: '*string'
    examples:
      - MYSTICKYNESS
      - foobar
  - name: nginx.ingress.kubernetes.io/session-cookie-path
    go_name: SessionCookiePath
    type: string
    go_type: '*string'
    examples:
      - /foobar
  - name: nginx.ingress.kubernetes.io/session-cookie-samesite
    go_name: SessionCookieSameSite
    type: string
    go_type: '*string'
    examples:
      - None
  - name: nginx.ingress.kubernetes.io/session-cookie-secure
    go_name: SessionCookieSecure
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
  - name: nginx.ingress.kubernetes.io/ssl-passthrough
    go_name: SSLPassthrough
    type: boolean
    go_type: '*bool'
    examples:
      - "true"
    description: Some differences in SNI/default backend handling.
  - name: nginx.ingress.kubernetes.io/ssl-redirect
    go_name: SSLRedirect
    type: boolean
    go_type: '*bool'
    examples:
      - "false"
    description: Cannot opt-out per route if enabled globally.
  - name: nginx.ingress.kubernetes.io/temporal-redirect
    go_name: TemporalRedirect
    type: string
    go_type: '*string'
    examples:
      - https://www.google.com
    description: Takes precedence over the `permanent-redirect` annotation. Defaults to a 302 Found status code.
  - name: nginx.ingress.kubernetes.io/temporal-redirect-code
    go_name: TemporalRedirectCode
    type: integer
    go_type: '*int'
    examples:
      - "308"
      - "429"
    description: Only valid 3XX HTTP Status Codes are accepted.
  - name: nginx.ingress.kubernetes.io/upstream-hash-by
    go_name: UpstreamHashBy
    type: string
    go_type: '*string'
    examples:
      - $request_uri
    description: 'It supports minimal variable interpolation by using the following NGINX variables: `$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`.'
  - name: nginx.ingress.kubernetes.io/upstream-vhost
    go_name: UpstreamVHost
    type: string
    go_type: '*string'
    examples:
      - upstream-host-header-value
    description: Supports NGINX variable interpolation. Request-time variables (`$scheme`, `$host`, `$http_*`, `$hostname`, `$request_uri`, `$request_method`, `$query_string`, `$args`, `$arg_*`, `$remote_addr`, `$uri`, `$document_uri`, `$server_name`, `$server_port`, `$content_type`, `$content_length`, `$cookie_*`, `$is_args`, `$best_http_host`, `$escaped_request_uri`, `$proxy_add_x_forwarded_for`) and the provider-resolved per-location variables (`$namespace`, `$ingress_name`, `$service_name`, `$service_port`, `$location_path`) are supported. The NGINX-internal variable `$proxy_upstream_name` is not available.
  - name: nginx.ingress.kubernetes.io/use-regex
    go_name: UseRegex
    type: boolean
    go_type: '*bool'
    examples:
      - "false"
      - "true"
  - name: nginx.ingress.kubernetes.io/whitelist-source-range
    go_name: WhitelistSourceRange
    type: string
    go_type: '*string'
    examples:
      - ""
      - 192.168.1.0/24, 10.0.0.0/8, 192.168.20.1
      - 192.168.1.0/24
      - 192.168.20.1
  - name: nginx.ingress.kubernetes.io/x-forwarded-prefix
    go_name: XForwardedPrefix
    type: string
    go_type: '*string'
    examples:
      - x-forwarded-prefix-header-value
      - $1
      - /$1/$2
---

# IngressNGINXAnnotations
