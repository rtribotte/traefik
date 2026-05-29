---
schema_version: 2
kind: provider
name: Provider
id: provider.kubernetesingressnginx
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/provider/kubernetes/ingress-nginx/kubernetes.go#L79
summary: Provider holds configurations of the provider.
fields:
  - name: endpoint
    go_name: Endpoint
    type: string
    go_type: string
  - name: token
    go_name: Token
    type: string
    go_type: types.FileOrContent
  - name: certAuthFilePath
    go_name: CertAuthFilePath
    type: string
    go_type: string
  - name: throttleDuration
    go_name: ThrottleDuration
    type: duration
    go_type: ptypes.Duration
  - name: globalAuthURL
    go_name: GlobalAuthURL
    type: string
    go_type: string
  - name: watchNamespace
    go_name: WatchNamespace
    type: string
    go_type: string
  - name: watchNamespaceSelector
    go_name: WatchNamespaceSelector
    type: string
    go_type: string
  - name: ingressClass
    go_name: IngressClass
    type: string
    go_type: string
  - name: controllerClass
    go_name: ControllerClass
    type: string
    go_type: string
  - name: watchIngressWithoutClass
    go_name: WatchIngressWithoutClass
    type: boolean
    go_type: bool
  - name: ingressClassByName
    go_name: IngressClassByName
    type: boolean
    go_type: bool
  - name: publishService
    go_name: PublishService
    type: string
    go_type: string
    description: 'TODO: support report-node-internal-ip-address and update-status.'
  - name: publishStatusAddress
    go_name: PublishStatusAddress
    type: array
    items: string
    go_type: '[]string'
  - name: defaultBackendService
    go_name: DefaultBackendService
    type: string
    go_type: string
  - name: disableSvcExternalName
    go_name: DisableSvcExternalName
    type: boolean
    go_type: bool
  - name: ipAllowListStrategy
    go_name: IPAllowListStrategy
    type: object
    go_type: '*dynamic.IPStrategy'
    type_ref: oss:IPStrategy
    fields:
      - name: depth
        go_name: Depth
        type: integer
        go_type: int
        description: Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).
      - name: excludedIPs
        go_name: ExcludedIPs
        type: array
        items: string
        go_type: '[]string'
        description: ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.
      - name: ipv6Subnet
        go_name: IPv6Subnet
        type: integer
        go_type: '*int'
        description: IPv6Subnet configures Traefik to consider all IPv6 addresses from the defined subnet as originating from the same IP. Applies to RemoteAddrStrategy and DepthStrategy.
  - name: httpEntryPoint
    go_name: HTTPEntryPoint
    type: string
    go_type: string
  - name: httpsEntryPoint
    go_name: HTTPSEntryPoint
    type: string
    go_type: string
  - name: proxyRequestBuffering
    go_name: ProxyRequestBuffering
    type: boolean
    go_type: bool
    description: Configuration options available within the NGINX Ingress Controller ConfigMap.
  - name: clientBodyBufferSize
    go_name: ClientBodyBufferSize
    type: integer
    go_type: int64
  - name: proxyBodySize
    go_name: ProxyBodySize
    type: integer
    go_type: int64
  - name: proxyBuffering
    go_name: ProxyBuffering
    type: boolean
    go_type: bool
  - name: proxyBufferSize
    go_name: ProxyBufferSize
    type: integer
    go_type: int64
  - name: proxyBuffersNumber
    go_name: ProxyBuffersNumber
    type: integer
    go_type: int
  - name: proxyConnectTimeout
    go_name: ProxyConnectTimeout
    type: integer
    go_type: int
  - name: proxyReadTimeout
    go_name: ProxyReadTimeout
    type: integer
    go_type: int
  - name: proxySendTimeout
    go_name: ProxySendTimeout
    type: integer
    go_type: int
  - name: proxyNextUpstream
    go_name: ProxyNextUpstream
    type: string
    go_type: string
  - name: proxyNextUpstreamTries
    go_name: ProxyNextUpstreamTries
    type: integer
    go_type: int
  - name: proxyNextUpstreamTimeout
    go_name: ProxyNextUpstreamTimeout
    type: integer
    go_type: int
  - name: customHTTPErrors
    go_name: CustomHTTPErrors
    type: array
    items: string
    go_type: '[]string'
  - name: upstreamKeepaliveTimeout
    go_name: UpstreamKeepaliveTimeout
    type: integer
    go_type: int
  - name: allowCrossNamespaceResources
    go_name: AllowCrossNamespaceResources
    type: boolean
    go_type: bool
  - name: globalAllowedResponseHeaders
    go_name: GlobalAllowedResponseHeaders
    type: array
    items: string
    go_type: '[]string'
  - name: allowSnippetAnnotations
    go_name: AllowSnippetAnnotations
    type: boolean
    go_type: bool
  - name: strictValidatePathType
    go_name: StrictValidatePathType
    type: boolean
    go_type: bool
    default: true
representations:
  yaml_path: providers.kubernetesIngressNGINX
  toml_path: providers.kubernetesIngressNGINX
---

# Provider

Provider holds configurations of the provider.
