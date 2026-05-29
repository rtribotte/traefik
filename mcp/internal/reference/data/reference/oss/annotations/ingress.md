---
schema_version: 2
kind: ingress-annotations
name: TraefikIngressAnnotations
id: annotations.ingress
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/ingress/annotations.go
summary: Annotations supported on Kubernetes Ingress objects by the Traefik Ingress provider.
tags:
  - On Ingress
  - On Service
fields:
  - name: traefik.ingress.kubernetes.io/router.entrypoints
    type: array
    items: string
    go_type: '[]string'
    examples:
      - '`ep1,ep2`'
    description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
  - name: traefik.ingress.kubernetes.io/router.middlewares
    type: array
    items: string
    go_type: '[]string'
    examples:
      - '`auth@file,default-prefix@kubernetescrd`'
    description: Middlewares is the list of MiddlewareRef which composes the chain.
  - name: traefik.ingress.kubernetes.io/router.observability.accesslogs
    type: boolean
    go_type: '*bool'
    examples:
      - '`true`'
    description: AccessLogs enables access logs for this router.
  - name: traefik.ingress.kubernetes.io/router.observability.metadata.ingress.ingressname
    type: string
    go_type: string
  - name: traefik.ingress.kubernetes.io/router.observability.metadata.ingress.namespace
    type: string
    go_type: string
    description: Namespace defines the namespace of the referenced IngressRoute resource.
  - name: traefik.ingress.kubernetes.io/router.observability.metadata.ingress.servicename
    type: string
    go_type: string
  - name: traefik.ingress.kubernetes.io/router.observability.metadata.ingress.serviceport
    type: string
    go_type: string
  - name: traefik.ingress.kubernetes.io/router.observability.metrics
    type: boolean
    go_type: '*bool'
    examples:
      - '`true`'
    description: Metrics enables metrics for this router.
  - name: traefik.ingress.kubernetes.io/router.observability.traceverbosity
    type: object
    go_type: otypes.TracingVerbosity
    description: TraceVerbosity defines the verbosity level of the tracing for this router.
  - name: traefik.ingress.kubernetes.io/router.observability.tracing
    type: boolean
    go_type: '*bool'
    examples:
      - '`true`'
    description: Tracing enables tracing for this router.
  - name: traefik.ingress.kubernetes.io/router.pathmatcher
    type: string
    go_type: string
    examples:
      - '`Path`'
    description: 'Overrides the default router rule type used for a path. Only path-related matcher name should be specified: `Path`, `PathPrefix` or `PathRegexp`. Default: `PathPrefix`'
  - name: traefik.ingress.kubernetes.io/router.priority
    type: integer
    go_type: int
    examples:
      - '`"42"`'
    description: 'Priority defines the router''s priority. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#priority'
  - name: traefik.ingress.kubernetes.io/router.rulesyntax
    type: string
    go_type: string
    examples:
      - '`"v2"`'
    description: See [rule syntax](https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/routing/rules-and-priority/#rulesyntax) for more information. **Deprecated:** RuleSyntax option is deprecated and will be removed in the next major version. Please do not use this field and rewrite the router rules to use the v3 syntax.
  - name: traefik.ingress.kubernetes.io/router.tls.certresolver
    type: string
    go_type: string
    examples:
      - '`myresolver`'
    description: 'CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/tls/certificate-resolvers/acme/'
  - name: traefik.ingress.kubernetes.io/router.tls.domains
    type: array
    items: object
    go_type: '[]types.Domain'
    type_ref: oss:Domain
    description: 'Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-certificates/#domains'
  - name: traefik.ingress.kubernetes.io/router.tls.options
    type: string
    go_type: string
    examples:
      - '`foobar@file`'
    description: 'Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/tls/tls-options/'
  - name: traefik.ingress.kubernetes.io/service.middlewares
    type: array
    items: string
    go_type: '[]string'
    description: Middlewares is the list of MiddlewareRef which composes the chain.
  - name: traefik.ingress.kubernetes.io/service.nativelb
    type: boolean
    go_type: '*bool'
    examples:
      - '`"true"`'
    description: NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
  - name: traefik.ingress.kubernetes.io/service.nodeportlb
    type: boolean
    go_type: bool
    examples:
      - '`"true"`'
    description: NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort. It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes. By default, NodePortLB is false.
  - name: traefik.ingress.kubernetes.io/service.passhostheader
    type: boolean
    go_type: '*bool'
    examples:
      - '`"true"`'
    description: PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service. By default, passHostHeader is true.
  - name: traefik.ingress.kubernetes.io/service.serversscheme
    type: string
    go_type: string
    examples:
      - '`h2c`'
    description: Overrides the default scheme.
  - name: traefik.ingress.kubernetes.io/service.serverstransport
    type: string
    go_type: string
    examples:
      - '`foobar@file`'
    description: ServersTransport defines the name of ServersTransport resource to use. It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.domain
    type: string
    go_type: string
    examples:
      - '`"foo.com"`'
    description: 'Domain defines the host to which the cookie will be sent. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#domaindomain-value'
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.httponly
    type: boolean
    go_type: bool
    examples:
      - '`"true"`'
    description: HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.maxage
    type: integer
    go_type: int
    examples:
      - '`42`'
    description: MaxAge defines the number of seconds until the cookie expires. When set to a negative number, the cookie expires immediately. When set to zero, the cookie never expires.
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.name
    type: string
    go_type: string
    examples:
      - '`foobar`'
    description: Name defines the Cookie name.
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.path
    type: string
    go_type: '*string'
    examples:
      - '`/foobar`'
    description: 'Path defines the path that must exist in the requested URL for the browser to send the Cookie header. When not provided the cookie will be sent on every request to the domain. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#pathpath-value'
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.samesite
    type: string
    go_type: string
    examples:
      - '`"none"`'
    description: 'SameSite defines the same site policy. More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite'
  - name: traefik.ingress.kubernetes.io/service.sticky.cookie.secure
    type: boolean
    go_type: bool
    examples:
      - '`"true"`'
    description: Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).
---

# TraefikIngressAnnotations
