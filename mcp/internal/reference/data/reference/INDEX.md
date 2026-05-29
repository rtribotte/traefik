# Traefik configuration reference index

> Configuration surface for Traefik v3.7.0 and Traefik Hub, generated from the Go source code. OSS pages live under `reference/oss/`, Hub pages under `reference/hub/`, vendored Kubernetes catalogues under `reference/external/`. Each entry maps to a `reference/<source>/<path>.md` file with the full structure inlined and a `schemas/<source>/<path>.schema.json` ready to feed `tool.input_schema`.

Use this index to pick which concepts you need, then load their detailed pages.

## HTTP middlewares

Per-request transformations applied between routers and services.

- `http.middlewares.addprefix` , AddPrefix , AddPrefix holds the add prefix middleware configuration. This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/addprefix/
- `http.middlewares.basicauth` , BasicAuth , BasicAuth holds the basic auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/basicauth/
- `http.middlewares.buffering` , Buffering , Buffering holds the buffering middleware configuration. This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v3.7/middlew...
- `http.middlewares.chain` , Chain , Chain holds the chain middleware configuration. This middleware enables to define reusable combinations of other pieces of middleware.
- `http.middlewares.circuitbreaker` , CircuitBreaker , CircuitBreaker holds the circuit breaker middleware configuration. This middleware protects the system from stacking requests to unhealthy services, resulting in cascading failures. More info: https:/...
- `http.middlewares.compress` , Compress , Compress holds the compress middleware configuration. This middleware compresses responses before sending them to the client, using gzip, brotli, or zstd compression.
- `http.middlewares.contenttype` , ContentType , ContentType holds the content-type middleware configuration. This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
- `http.middlewares.digestauth` , DigestAuth , DigestAuth holds the digest auth middleware configuration. This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/digestauth/
- `http.middlewares.encodedcharacters` , EncodedCharacters , EncodedCharacters configures which encoded characters are allowed in the request path.
- `http.middlewares.errors` , ErrorPage , ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes.
- `http.middlewares.forwardauth` , ForwardAuth , ForwardAuth holds the forward auth middleware configuration. This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/forwarda...
- `http.middlewares.grpcweb` , GrpcWeb , GrpcWeb holds the gRPC web middleware configuration. This middleware converts a gRPC web request to an HTTP/2 gRPC request.
- `http.middlewares.headers` , Headers , Headers holds the headers middleware configuration. This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/headers/#customrequesthe...
- `http.middlewares.inflightreq` , InFlightReq , InFlightReq holds the in-flight request middleware configuration. This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v3.7/...
- `http.middlewares.ipallowlist` , IPAllowList , IPAllowList holds the IP allowlist middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipallowlist...
- `http.middlewares.ipwhitelist` , IPWhiteList (deprecated: please use IPAllowList instead) , IPWhiteList holds the IP whitelist middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipwhitelist...
- `http.middlewares.passtlsclientcert` , PassTLSClientCert , PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/...
- `http.middlewares.ratelimit` , RateLimit , RateLimit holds the rate limit configuration. This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is.
- `http.middlewares.redirectregex` , RedirectRegex , RedirectRegex holds the redirect regex middleware configuration. This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/middlewares/ht...
- `http.middlewares.redirectscheme` , RedirectScheme , RedirectScheme holds the redirect scheme middleware configuration. This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/red...
- `http.middlewares.replacepath` , ReplacePath , ReplacePath holds the replace path middleware configuration. This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traef...
- `http.middlewares.replacepathregex` , ReplacePathRegex , ReplacePathRegex holds the replace path regex middleware configuration. This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v3.7/...
- `http.middlewares.retry` , Retry , Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middlewa...
- `http.middlewares.stripprefix` , StripPrefix , StripPrefix holds the strip prefix middleware configuration. This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/strippref...
- `http.middlewares.stripprefixregex` , StripPrefixRegex , StripPrefixRegex holds the strip prefix regex middleware configuration. This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http...

## TCP middlewares

Per-connection transformations for TCP routers.

- `tcp.middlewares.inflightconn` , TCPInFlightConn , TCPInFlightConn holds the TCP InFlightConn middleware configuration. This middleware prevents services from being overwhelmed with high load, by limiting the number of allowed simultaneous connections...
- `tcp.middlewares.ipallowlist` , TCPIPAllowList , TCPIPAllowList holds the TCP IPAllowList middleware configuration. This middleware limits allowed requests based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/tcp/ipallo...
- `tcp.middlewares.ipwhitelist` , TCPIPWhiteList (deprecated: please use IPAllowList instead) , TCPIPWhiteList holds the TCP IPWhiteList middleware configuration.

## Routers, services, TLS

Top-level dynamic config types.

- `http.routers` , Router , Router holds the router configuration.
- `http.services` , Service , Service holds a service configuration (can only be of one type at the same time).
- `tcp.routers` , TCPRouter , TCPRouter holds the router configuration.
- `tcp.services` , TCPService , TCPService holds a tcp service configuration (can only be of one type at the same time).
- `tls` , TLS , TLSConfiguration contains all the configuration parameters of a TLS connection.
- `udp.routers` , UDPRouter , UDPRouter defines the configuration for an UDP router.
- `udp.services` , UDPService , UDPService defines the configuration for a UDP service. All fields are mutually exclusive.

## Static configuration

Top-level traefik.yaml sections (process-wide).

- `static` , StaticConfiguration , Configuration is the static configuration.
- `static.accesslog` , AccessLog , AccessLog holds the configuration settings for the access logger (middlewares/accesslog).
- `static.api` , API , API holds the API configuration.
- `static.certificatesresolvers` , CertificateResolver , CertificateResolver contains the configuration for the different types of certificates resolver.
- `static.core` , Core (deprecated: Please do not use this field) , Core configures Traefik core behavior.
- `static.entrypoints` , EntryPoint , EntryPoint holds the entry point configuration.
- `static.experimental` , Experimental , Experimental experimental Traefik features.
- `static.global` , Global , Global holds the global configuration.
- `static.hostresolver` , HostResolverConfig , HostResolverConfig contain configuration for CNAME Flattening.
- `static.log` , TraefikLog , TraefikLog holds the configuration settings for the traefik logger.
- `static.metrics` , Metrics , Metrics provides options to expose and send Traefik metrics to different third party monitoring systems.
- `static.ocsp` , OCSPConfig , OCSPConfig contains the OCSP configuration.
- `static.ping` , Handler , Handler expose ping routes.
- `static.providers` , Providers , Provider configuration. One sub-key per enabled provider.
- `static.serverstransport` , ServersTransport , ServersTransport options to configure communication between Traefik and the servers.
- `static.spiffe` , SpiffeClientConfig , SpiffeClientConfig defines the SPIFFE client configuration.
- `static.tcpserverstransport` , TCPServersTransport , TCPServersTransport options to configure communication between Traefik and the servers.
- `static.tracing` , Tracing , Tracing holds the tracing configuration.

## Providers

Discovery sources Traefik watches for dynamic configuration.

- `provider.consul` , ProviderBuilder , ProviderBuilder is responsible for constructing namespaced instances of the Consul provider.
- `provider.consulcatalog` , ProviderBuilder , ProviderBuilder is responsible for constructing namespaced instances of the Consul Catalog provider.
- `provider.docker` , Provider , Provider holds configurations of the provider.
- `provider.ecs` , Provider , Provider holds configurations of the provider.
- `provider.etcd` , Provider , Provider holds configurations of the provider.
- `provider.file` , Provider , Provider holds configurations of the provider.
- `provider.http` , Provider , Provider is a provider.Provider implementation that queries an HTTP(s) endpoint for a configuration.
- `provider.knative` , Provider , Provider holds configurations of the provider.
- `provider.kubernetescrd` , Provider , Provider holds configurations of the provider.
- `provider.kubernetesgateway` , Provider , Provider holds configurations of the provider.
- `provider.kubernetesingress` , Provider , Provider holds configurations of the provider.
- `provider.kubernetesingressnginx` , Provider , Provider holds configurations of the provider.
- `provider.nomad` , ProviderBuilder , ProviderBuilder is responsible for constructing namespaced instances of the Nomad provider.
- `provider.redis` , Provider , Provider holds configurations of the provider.
- `provider.rest` , Provider , Provider is a provider.Provider implementation that provides a Rest API.
- `provider.swarm` , SwarmProvider , SwarmProvider holds configurations of the provider.
- `provider.zookeeper` , Provider , Provider holds configurations of the provider.

## CRDs

Kubernetes Custom Resource Definitions (Traefik native and Hub).

- `crd.accesscontrolpolicy` , AccessControlPolicy , AccessControlPolicy defines an access control policy.
- `crd.aiservice` , AIService , AIService is a Kubernetes-like Service to interact with a text-based LLM provider. It defines the parameters and credentials required to interact with various LLM providers.
- `crd.api` , API , API defines an HTTP interface that is exposed to external clients. It specifies the supported versions
- `crd.apiauth` , APIAuth , APIAuth defines the authentication configuration for APIs.
- `crd.apibundle` , APIBundle , APIBundle defines a set of APIs.
- `crd.apicatalogitem` , APICatalogItem , APICatalogItem defines APIs that will be part of the API catalog on the portal.
- `crd.apiplan` , APIPlan , APIPlan defines API Plan policy.
- `crd.apiportal` , APIPortal , APIPortal defines a developer portal for accessing the documentation of APIs.
- `crd.apiportalauth` , APIPortalAuth , APIPortalAuth defines the authentication configuration for an APIPortal.
- `crd.apiratelimit` , APIRateLimit , APIRateLimit defines how group of consumers are rate limited on a set of APIs.
- `crd.apiversion` , APIVersion , APIVersion defines a version of an API.
- `crd.contentitem` , ContentItem , ContentItem defines additional documentation for given resource.
- `crd.ingressroute` , IngressRoute , IngressRoute is the CRD implementation of a Traefik HTTP Router.
- `crd.ingressroutetcp` , IngressRouteTCP , IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
- `crd.ingressrouteudp` , IngressRouteUDP , IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
- `crd.managedapplication` , ManagedApplication , ManagedApplication represents a managed application.
- `crd.managedsubscription` , ManagedSubscription , ManagedSubscription defines a Subscription managed by the API manager as the result of a pre-negotiation with its
- `crd.middleware` , Middleware , Middleware is the CRD implementation of a Traefik Middleware.
- `crd.middlewaretcp` , MiddlewareTCP , MiddlewareTCP is the CRD implementation of a Traefik TCP middleware.
- `crd.serverstransport` , ServersTransport , ServersTransport is the CRD implementation of a ServersTransport.
- `crd.serverstransporttcp` , ServersTransportTCP , ServersTransportTCP is the CRD implementation of a TCPServersTransport.
- `crd.tlsoption` , TLSOption , TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection.
- `crd.tlsstore` , TLSStore , TLSStore is the CRD implementation of a Traefik TLS Store.
- `crd.traefikservice` , TraefikService , TraefikService is the CRD implementation of a Traefik Service.
- `crd.uplink` , Uplink , Uplink is an inter-cluster service advertisement: a child cluster declares an Uplink to advertise

## Annotations

Annotations Traefik reads from Kubernetes Ingress objects.

- `annotations.gateway-service` , GatewayServiceAnnotations , Annotations on Kubernetes Service objects read by the Traefik Kubernetes Gateway API provider. This is an exhaustive list: any other annotation under the traefik.io/service.* prefix is silently ignore...
- `annotations.ingress` , TraefikIngressAnnotations , Annotations supported on Kubernetes Ingress objects by the Traefik Ingress provider.
- `annotations.ingress-nginx` , IngressNGINXAnnotations , nginx-style annotations on Kubernetes Ingress objects supported by the Traefik ingress-nginx provider.

## Shared concepts

Reusable building blocks referenced from multiple kinds.

- `concept.accesslog` , AccessLog , AccessLog holds the configuration settings for the access logger (middlewares/accesslog).
- `concept.accesslogfields` , AccessLogFields , AccessLogFields holds configuration for access log fields.
- `concept.accesslogfilters` , AccessLogFilters , AccessLogFilters holds filters configuration.
- `concept.api` , API , API holds the API configuration.
- `concept.certandstores` , CertAndStores , CertAndStores allows mapping a TLS certificate to a list of entry points.
- `concept.certificate` , Certificate , Certificate holds a SSL cert/key pair Certs and Key could be either a file path, or the file content itself.
- `concept.certificateresolver` , CertificateResolver , CertificateResolver contains the configuration for the different types of certificates resolver.
- `concept.clientauth` , ClientAuth , ClientAuth defines the parameters of the client authentication part of the TLS connection, if any.
- `concept.clienttls` , ClientTLS , ClientTLS holds TLS specific configurations as client CA, Cert and Key can be either path or file contents. TODO: remove this struct when CAOptional option will be removed.
- `concept.configuration` , Configuration , Configuration is the root of the dynamic configuration.
- `concept.cookie` , Cookie , Cookie holds the sticky configuration based on cookie.
- `concept.core` , Core , Core configures Traefik core behavior.
- `concept.datadog` , Datadog , Datadog contains address and metrics pushing interval configuration.
- `concept.domain` , Domain , Domain holds a domain name with SANs.
- `concept.entrypointstransport` , EntryPointsTransport , EntryPointsTransport configures communication between clients and Traefik.
- `concept.experimental` , Experimental , Experimental experimental Traefik features.
- `concept.failover` , Failover , Failover holds the Failover configuration.
- `concept.failovererror` , FailoverError , FailoverError holds errors configuration.
- `concept.fastproxyconfig` , FastProxyConfig , FastProxyConfig holds the FastProxy configuration.
- `concept.fieldheaders` , FieldHeaders , FieldHeaders holds configuration for access log headers.
- `concept.forwardedheaders` , ForwardedHeaders , ForwardedHeaders Trust client forwarding headers.
- `concept.forwardingtimeouts` , ForwardingTimeouts , ForwardingTimeouts contains timeout configurations for forwarding requests to the backend servers.
- `concept.generatedcert` , GeneratedCert , GeneratedCert defines the default generated certificate configuration.
- `concept.global` , Global , Global holds the global configuration.
- `concept.handler` , Handler , Handler expose ping routes.
- `concept.healthcheck` , HealthCheck , HealthCheck controls healthcheck awareness and propagation at the services level.
- `concept.highestrandomweight` , HighestRandomWeight , HighestRandomWeight is a weighted sticky load-balancer of services.
- `concept.hostresolverconfig` , HostResolverConfig , HostResolverConfig contain configuration for CNAME Flattening.
- `concept.hrwservice` , HRWService , HRWService is a reference to a service load-balanced with highest random weight.
- `concept.http2config` , HTTP2Config , HTTP2Config is the HTTP2 configuration of an entry point.
- `concept.http3config` , HTTP3Config , HTTP3Config is the HTTP3 configuration of an entry point.
- `concept.httpconfig` , HTTPConfig , HTTPConfig is the HTTP configuration of an entry point.
- `concept.httpconfiguration` , HTTPConfiguration , HTTPConfiguration contains all the HTTP configuration parameters.
- `concept.influxdb2` , InfluxDB2 , InfluxDB2 contains address, token and metrics pushing interval configuration.
- `concept.ipstrategy` , IPStrategy , IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/ipallowlist/#ipstrategy
- `concept.lifecycle` , LifeCycle , LifeCycle contains configurations relevant to the lifecycle (such as the shutdown phase) of Traefik.
- `concept.metrics` , Metrics , Metrics provides options to expose and send Traefik metrics to different third party monitoring systems.
- `concept.middleware` , Middleware , Middleware holds the Middleware configuration.
- `concept.mirroring` , Mirroring , Mirroring holds the Mirroring configuration.
- `concept.mirrorservice` , MirrorService , MirrorService holds the MirrorService configuration.
- `concept.model` , Model , Model holds model configuration.
- `concept.observabilityconfig` , ObservabilityConfig , ObservabilityConfig holds the observability configuration for an entry point.
- `concept.ocspconfig` , OCSPConfig , OCSPConfig contains the OCSP configuration.
- `concept.options` , Options , Options configures TLS for an entry point.
- `concept.otelgrpc` , OTelGRPC , OTelGRPC provides configuration settings for the gRPC open-telemetry.
- `concept.otelhttp` , OTelHTTP , OTelHTTP provides configuration settings for the HTTP open-telemetry.
- `concept.otellog` , OTelLog , OTelLog provides configuration settings for the open-telemetry logger.
- `concept.oteltracing` , OTelTracing , OTelTracing provides configuration settings for the open-telemetry tracer.
- `concept.otlp` , OTLP , OTLP contains specific configuration used by the OpenTelemetry Metrics exporter.
- `concept.passiveserverhealthcheck` , PassiveServerHealthCheck , Shared type referenced from configuration. See Go source for details.
- `concept.prometheus` , Prometheus , Prometheus can contain specific configuration used by the Prometheus Metrics exporter.
- `concept.providers` , Providers , Providers contains providers configuration.
- `concept.proxyprotocol` , ProxyProtocol , ProxyProtocol holds the PROXY Protocol configuration. More info: https://doc.traefik.io/traefik/v3.7/routing/services/#proxy-protocol
- `concept.redirectentrypoint` , RedirectEntryPoint , RedirectEntryPoint is the definition of an entry point redirection.
- `concept.redirections` , Redirections , Redirections is a set of redirection for an entry point.
- `concept.redis` , Redis , Redis holds the Redis configuration.
- `concept.respondingtimeouts` , RespondingTimeouts , RespondingTimeouts contains timeout configurations for incoming requests to the Traefik instance.
- `concept.responseforwarding` , ResponseForwarding , ResponseForwarding holds the response forwarding configuration.
- `concept.routerobservabilityconfig` , RouterObservabilityConfig , RouterObservabilityConfig holds the observability configuration for a router.
- `concept.routertcptlsconfig` , RouterTCPTLSConfig , RouterTCPTLSConfig holds the TLS configuration for a router.
- `concept.routertlsconfig` , RouterTLSConfig , RouterTLSConfig holds the TLS configuration for a router.
- `concept.server` , Server , Server holds the server configuration.
- `concept.serverhealthcheck` , ServerHealthCheck , ServerHealthCheck holds the HealthCheck configuration.
- `concept.serversloadbalancer` , ServersLoadBalancer , ServersLoadBalancer holds the ServersLoadBalancer configuration.
- `concept.serverstransport` , ServersTransport , ServersTransport options to configure communication between Traefik and the servers.
- `concept.sourcecriterion` , SourceCriterion , SourceCriterion defines what criterion is used to group requests as originating from a common source. If none are set, the default is to use the request's remote address field. All fields are mutually...
- `concept.spiffe` , Spiffe , Spiffe holds the SPIFFE configuration.
- `concept.spiffeclientconfig` , SpiffeClientConfig , SpiffeClientConfig defines the SPIFFE client configuration.
- `concept.statsd` , Statsd , Statsd contains address and metrics pushing interval configuration.
- `concept.sticky` , Sticky , Sticky holds the sticky configuration.
- `concept.store` , Store , Store holds the options for a given Store.
- `concept.tcpconfiguration` , TCPConfiguration , TCPConfiguration contains all the TCP configuration parameters.
- `concept.tcpmiddleware` , TCPMiddleware , TCPMiddleware holds the TCPMiddleware configuration.
- `concept.tcpserver` , TCPServer , TCPServer holds a TCP Server configuration.
- `concept.tcpserverhealthcheck` , TCPServerHealthCheck , TCPServerHealthCheck holds the HealthCheck configuration.
- `concept.tcpserversloadbalancer` , TCPServersLoadBalancer , TCPServersLoadBalancer holds the LoadBalancerService configuration.
- `concept.tcpserverstransport` , TCPServersTransport , TCPServersTransport options to configure communication between Traefik and the servers.
- `concept.tcpweightedroundrobin` , TCPWeightedRoundRobin , TCPWeightedRoundRobin is a weighted round robin tcp load-balancer of services.
- `concept.tcpwrrservice` , TCPWRRService , TCPWRRService is a reference to a tcp service load-balanced with weighted round robin.
- `concept.tlsclientcertificateinfo` , TLSClientCertificateInfo , TLSClientCertificateInfo holds the client TLS certificate info configuration.
- `concept.tlsclientcertificateissuerdninfo` , TLSClientCertificateIssuerDNInfo , TLSClientCertificateIssuerDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
- `concept.tlsclientcertificatesubjectdninfo` , TLSClientCertificateSubjectDNInfo , TLSClientCertificateSubjectDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
- `concept.tlsclientconfig` , TLSClientConfig , TLSClientConfig options to configure TLS communication between Traefik and the servers.
- `concept.tlsconfig` , TLSConfig , TLSConfig is the default TLS configuration for all the routers associated to the concerned entry point.
- `concept.tlsconfiguration` , TLSConfiguration , TLSConfiguration contains all the configuration parameters of a TLS connection.
- `concept.tracing` , Tracing , Tracing holds the tracing configuration.
- `concept.traefiklog` , TraefikLog , TraefikLog holds the configuration settings for the traefik logger.
- `concept.udpconfig` , UDPConfig , UDPConfig is the UDP configuration of an entry point.
- `concept.udpconfiguration` , UDPConfiguration , UDPConfiguration contains all the UDP configuration parameters.
- `concept.udpserver` , UDPServer , UDPServer defines a UDP server configuration.
- `concept.udpserversloadbalancer` , UDPServersLoadBalancer , UDPServersLoadBalancer defines the configuration for a load-balancer of UDP servers.
- `concept.udpweightedroundrobin` , UDPWeightedRoundRobin , UDPWeightedRoundRobin is a weighted round robin UDP load-balancer of services.
- `concept.udpwrrservice` , UDPWRRService , UDPWRRService is a reference to a UDP service load-balanced with weighted round robin.
- `concept.weightedroundrobin` , WeightedRoundRobin , WeightedRoundRobin is a weighted round robin load-balancer of services.
- `concept.wrrservice` , WRRService , WRRService is a reference to a service load-balanced with weighted round-robin.

## Hub middlewares

Traefik Hub middlewares (API gateway, AI gateway, security).

- `hub.middlewares.apikey` , APIKey , Configuration holds the API Key middleware configuration.
- `hub.middlewares.basicauth` , BasicAuth , Config configures a basic auth ACP handler.
- `hub.middlewares.cache` , Cache , Configuration holds the Cache Middleware configuration.
- `hub.middlewares.chatcompletion` , ChatCompletion , Config holds the ChatCompletion Middleware configuration.
- `hub.middlewares.contentguard` , ContentGuard , Config holds the configuration for content-guard middleware.
- `hub.middlewares.coraza` , Coraza , Configuration holds the Coraza middleware configuration.
- `hub.middlewares.distributedratelimit` , DistributedRateLimit , Configuration holds the DistributedRateLimit middleware configuration.
- `hub.middlewares.forcecase` , ForceCase , Configuration holds the Force Case middleware configuration.
- `hub.middlewares.hmac` , HMAC , Configuration holds the HMAC Authentication Middleware configuration.
- `hub.middlewares.hubapikey` , HubAPIKey , Configuration configures an API Key handler.
- `hub.middlewares.hubldap` , HubLDAP , Configuration holds the configuration for the Hub LDAP middleware.
- `hub.middlewares.jwt` , JWT , Configuration configures a JWT ACP handler.
- `hub.middlewares.ldap` , LDAP , Configuration holds the LDAP Middleware configuration.
- `hub.middlewares.llmguard` , LLMGuard , Config is the unified configuration for an LLM guard middleware.
- `hub.middlewares.mcp` , MCP , Config holds MCP middleware configuration.
- `hub.middlewares.metrics` , Metrics , Configuration configures an API Management metrics handler.
- `hub.middlewares.oauthclientcreds` , OAuthClientCreds , Configuration holds the configuration for the OAuth client credentials middleware.
- `hub.middlewares.oauthintro` , OAuthIntrospection , Configuration configures an OAuth 2.0 Token Introspection middleware.
- `hub.middlewares.oidc` , OIDC , Configuration holds the configuration for the OIDC middleware.
- `hub.middlewares.opa` , OPA , Configuration holds the OPA middleware configuration.
- `hub.middlewares.plan` , Plan , Configuration holds the configuration of the plan middleware.
- `hub.middlewares.queryparam` , QueryParam , Configuration holds the QueryParam middleware configuration.
- `hub.middlewares.responsesapi` , ResponsesAPI , Config holds the ResponsesAPI Middleware configuration.
- `hub.middlewares.semanticcache` , SemanticCache , Config holds the config for semantic cache middleware.
- `hub.middlewares.tokenratelimit` , TokenRateLimit , Config holds the configuration for the token rate limit middleware.

## Hub static configuration

Hub static config sections under the `hub` key of traefik.yaml.

- `hub.api.devportal` , DevPortalAPIIndex , Index of the developer portal REST API endpoints.
- `hub.concept.uplink` , Uplink , Uplink represents an inter-cluster service advertisement. A child cluster declares an Uplink to advertise to a parent cluster that it can handle a particular workload. This advertisement gets automati...
- `hub.concept.uplinkhealthcheck` , UplinkHealthCheck , UplinkHealthCheck mirrors Traefik's ServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
- `hub.concept.uplinkpassivehealthcheck` , UplinkPassiveHealthCheck , UplinkPassiveHealthCheck mirrors Traefik's PassiveServerHealthCheck. Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
- `hub.dyn.http` , DynExt , Extensions Hub adds on top of Traefik dynamic.HTTPConfiguration and dynamic.Router.
- `hub.static` , Hub , Hub static configuration. Lives under the top level "hub" key of the Traefik static configuration file.
- `hub.static.aigateway` , AIGateway , AIGateway holds the ai gateway configuration.
- `hub.static.apimanagement` , APIManagement , APIManagement holds the API management configuration.
- `hub.static.experimental` , Experimental , Experimental holds experimental features.
- `hub.static.mcpgateway` , MCPGateway , MCPGateway holds the MCP gateway configuration.
- `hub.static.pluginregistry` , PluginRegistry , PluginRegistry holds the plugin registry configuration.
- `hub.static.providers` , Providers , Providers contains providers configuration.
- `hub.static.redis` , Config , Config is the Redis configuration.
- `hub.static.tracing` , Tracing , Tracing holds the tracing configuration.

## Hub providers

Hub specific provider integrations.

- `hub.providers.consulcatalogenterprise` , ConsulCatalogEnterprise , Provider is the Consul Catalog Enterprise provider implementation.
- `hub.providers.hub` , Hub , Provider holds configurations of the provider.
- `hub.providers.kubernetescrd` , KubernetesCRD , Provider wraps the kubernetescrd provider and adds Uplink CRD support. It watches Uplink resources and populates the http.uplinks part of the dynamic configuration.
- `hub.providers.microcks` , Microcks , Provider is a provider.Provider implementation that queries a Microcks instance for service configurations.
- `hub.providers.multicluster` , MultiCluster , Provider is the multicluster provider.
- `hub.providers.nutanixprismcentral` , NutanixPrismCentral , Provider is the Nutanix Prism Central provider implementation.
- `hub.providers.traefik` , Traefik , Provider wraps Traefik's internal provider to handle distributed ACME challenges. It extends Traefik's internal provider functionality by handling both standard and distributed ACME HTTP challenges. T...

## Hub developer portal REST

Endpoints exposed by the Hub developer portal HTTP API.

- `hub.rest.createapikey` , createAPIKey , Create API key
- `hub.rest.createapplication` , createApplication , Create application
- `hub.rest.deleteapikey` , deleteAPIKey , Delete API key
- `hub.rest.deleteapplication` , deleteApplication , Delete application
- `hub.rest.deleteselfservicesubscription` , deleteSelfServiceSubscription , Delete self-service subscription
- `hub.rest.getapispec` , getAPISpec , Get API specification
- `hub.rest.getapiversionspec` , getAPIVersionSpec , Get API version specification
- `hub.rest.getcontent` , getContent , Get content
- `hub.rest.getportal` , getPortal , Get portal information
- `hub.rest.listapikeys` , listAPIKeys , List API keys
- `hub.rest.listapplications` , listApplications , List applications
- `hub.rest.suspendapikey` , suspendAPIKey , Suspend or unsuspend API key
- `hub.rest.updateapplication` , updateApplication , Update application
- `hub.rest.upsertselfservicesubscriptions` , upsertSelfServiceSubscriptions , Create or update self-service subscriptions

## Gateway API CRDs

Gateway API resources (gateway.networking.k8s.io) vendored from upstream catalogues.

- `gateway-api.backendtlspolicy` , BackendTLSPolicy , BackendTLSPolicy provides a way to configure how a Gateway
- `gateway-api.gateway` , Gateway , Gateway represents an instance of a service-traffic handling infrastructure
- `gateway-api.gatewayclass` , GatewayClass , GatewayClass describes a class of Gateways available to the user for creating
- `gateway-api.grpcroute` , GRPCRoute , GRPCRoute provides a way to route gRPC requests. This includes the capability
- `gateway-api.httproute` , HTTPRoute , HTTPRoute provides a way to route HTTP requests. This includes the capability
- `gateway-api.referencegrant` , ReferenceGrant , ReferenceGrant identifies kinds of resources in other namespaces that are
- `gateway-api.tcproute` , TCPRoute , TCPRoute provides a way to route TCP requests. When combined with a Gateway
- `gateway-api.tlsroute` , TLSRoute , The TLSRoute resource is similar to TCPRoute, but can be configured
- `gateway-api.udproute` , UDPRoute , UDPRoute provides a way to route UDP traffic. When combined with a Gateway

## Kubernetes core resources

Core Kubernetes objects that ship next to Traefik routing manifests (Service, Ingress, ...).

- `k8s-core.endpoints` , Endpoints , Endpoints is a collection of endpoints that implement the actual service. Example:
- `k8s-core.endpointslice` , EndpointSlice , EndpointSlice represents a set of service endpoints. Most EndpointSlices are created by the EndpointSlice controller to represent the Pods selected by Service objects. For a given service there may be...
- `k8s-core.ingress` , Ingress , Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traff...
- `k8s-core.ingressclass` , IngressClass , IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be co...
- `k8s-core.service` , Service , Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer r...

