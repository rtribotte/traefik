---
schema_version: 2
kind: middleware-http
name: Headers
id: http.middlewares.headers
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L349
summary: 'Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/headers/#customrequestheaders'
fields:
  - name: customRequestHeaders
    go_name: CustomRequestHeaders
    type: object
    items: string
    go_type: map[string]string
    description: CustomRequestHeaders defines the header names and values to apply to the request.
  - name: customResponseHeaders
    go_name: CustomResponseHeaders
    type: object
    items: string
    go_type: map[string]string
    description: CustomResponseHeaders defines the header names and values to apply to the response.
  - name: accessControlAllowCredentials
    go_name: AccessControlAllowCredentials
    type: boolean
    go_type: bool
    description: AccessControlAllowCredentials defines whether the request can include user credentials.
  - name: accessControlAllowHeaders
    go_name: AccessControlAllowHeaders
    type: array
    items: string
    go_type: '[]string'
    description: AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.
  - name: accessControlAllowMethods
    go_name: AccessControlAllowMethods
    type: array
    items: string
    go_type: '[]string'
    description: AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.
  - name: accessControlAllowOriginList
    go_name: AccessControlAllowOriginList
    type: array
    items: string
    go_type: '[]string'
    description: AccessControlAllowOriginList is a list of allowable origins. Can also be a wildcard origin "*".
  - name: accessControlAllowOriginListRegex
    go_name: AccessControlAllowOriginListRegex
    type: array
    items: string
    go_type: '[]string'
    description: AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).
  - name: accessControlExposeHeaders
    go_name: AccessControlExposeHeaders
    type: array
    items: string
    go_type: '[]string'
    description: AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.
  - name: accessControlMaxAge
    go_name: AccessControlMaxAge
    type: integer
    go_type: int64
    description: AccessControlMaxAge defines the time that a preflight request may be cached.
  - name: addVaryHeader
    go_name: AddVaryHeader
    type: boolean
    go_type: bool
    description: AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.
  - name: allowedHosts
    go_name: AllowedHosts
    type: array
    items: string
    go_type: '[]string'
    description: AllowedHosts defines the fully qualified list of allowed domain names.
  - name: hostsProxyHeaders
    go_name: HostsProxyHeaders
    type: array
    items: string
    go_type: '[]string'
    description: HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.
  - name: sslProxyHeaders
    go_name: SSLProxyHeaders
    type: object
    items: string
    go_type: map[string]string
    description: 'SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request. It can be useful when using other proxies (example: "X-Forwarded-Proto": "https").'
  - name: stsSeconds
    go_name: STSSeconds
    type: integer
    go_type: '*int64'
    description: STSSeconds defines the max-age of the Strict-Transport-Security header. If set to 0, the header is not set.
  - name: stsIncludeSubdomains
    go_name: STSIncludeSubdomains
    type: boolean
    go_type: bool
    description: STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.
  - name: stsPreload
    go_name: STSPreload
    type: boolean
    go_type: bool
    description: STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.
  - name: forceSTSHeader
    go_name: ForceSTSHeader
    type: boolean
    go_type: bool
    description: ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.
  - name: frameDeny
    go_name: FrameDeny
    type: boolean
    go_type: bool
    description: FrameDeny defines whether to add the X-Frame-Options header with the DENY value.
  - name: customFrameOptionsValue
    go_name: CustomFrameOptionsValue
    type: string
    go_type: string
    description: CustomFrameOptionsValue defines the X-Frame-Options header value. This overrides the FrameDeny option.
  - name: contentTypeNosniff
    go_name: ContentTypeNosniff
    type: boolean
    go_type: bool
    description: ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.
  - name: browserXssFilter
    go_name: BrowserXSSFilter
    type: boolean
    go_type: bool
    description: BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1; mode=block.
  - name: customBrowserXSSValue
    go_name: CustomBrowserXSSValue
    type: string
    go_type: string
    description: CustomBrowserXSSValue defines the X-XSS-Protection header value. This overrides the BrowserXssFilter option.
  - name: contentSecurityPolicy
    go_name: ContentSecurityPolicy
    type: string
    go_type: string
    description: ContentSecurityPolicy defines the Content-Security-Policy header value.
  - name: contentSecurityPolicyReportOnly
    go_name: ContentSecurityPolicyReportOnly
    type: string
    go_type: string
    description: ContentSecurityPolicyReportOnly defines the Content-Security-Policy-Report-Only header value.
  - name: publicKey
    go_name: PublicKey
    type: string
    go_type: string
    description: PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.
  - name: referrerPolicy
    go_name: ReferrerPolicy
    type: string
    go_type: string
    description: ReferrerPolicy defines the Referrer-Policy header value. This allows sites to control whether browsers forward the Referer header to other sites.
  - name: permissionsPolicy
    go_name: PermissionsPolicy
    type: string
    go_type: string
    description: PermissionsPolicy defines the Permissions-Policy header value. This allows sites to control browser features.
  - name: isDevelopment
    go_name: IsDevelopment
    type: boolean
    go_type: bool
    description: IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing. Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.
  - name: featurePolicy
    go_name: FeaturePolicy
    type: string
    go_type: '*string'
  - name: sslRedirect
    go_name: SSLRedirect
    type: boolean
    go_type: '*bool'
  - name: sslTemporaryRedirect
    go_name: SSLTemporaryRedirect
    type: boolean
    go_type: '*bool'
  - name: sslHost
    go_name: SSLHost
    type: string
    go_type: '*string'
  - name: sslForceHost
    go_name: SSLForceHost
    type: boolean
    go_type: '*bool'
representations:
  yaml_path: http.middlewares.<name>.headers
  toml_path: http.middlewares.<name>.headers
  label_prefix: traefik.http.middlewares.<name>.headers
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.headers
---

# Headers

Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/headers/#customrequestheaders
