---
schema_version: 2
kind: concept
name: Providers
id: concept.providers
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L261
summary: Providers contains providers configuration.
fields:
  - name: providersThrottleDuration
    go_name: ProvidersThrottleDuration
    type: duration
    go_type: ptypes.Duration
  - name: precedence
    go_name: Precedence
    type: array
    items: string
    go_type: '[]string'
    default:
      - kubernetesgateway
      - kubernetescrd
      - kubernetes
      - kubernetesingressnginx
      - swarm
      - docker
      - file
      - redis
      - knative
      - consul
      - consulcatalog
      - nomad
      - etcd
      - ecs
      - http
      - zookeeper
      - rest
  - name: docker
    go_name: Docker
    type: object
    go_type: '*docker.Provider'
  - name: swarm
    go_name: Swarm
    type: object
    go_type: '*docker.SwarmProvider'
  - name: file
    go_name: File
    type: object
    go_type: '*file.Provider'
  - name: kubernetesIngress
    go_name: KubernetesIngress
    type: object
    go_type: '*ingress.Provider'
  - name: kubernetesIngressNGINX
    go_name: KubernetesIngressNGINX
    type: object
    go_type: '*ingressnginx.Provider'
  - name: kubernetesCRD
    go_name: KubernetesCRD
    type: object
    go_type: '*crd.Provider'
  - name: kubernetesGateway
    go_name: KubernetesGateway
    type: object
    go_type: '*gateway.Provider'
  - name: knative
    go_name: Knative
    type: object
    go_type: '*knative.Provider'
  - name: rest
    go_name: Rest
    type: object
    go_type: '*rest.Provider'
  - name: consulCatalog
    go_name: ConsulCatalog
    type: object
    go_type: '*consulcatalog.ProviderBuilder'
  - name: nomad
    go_name: Nomad
    type: object
    go_type: '*nomad.ProviderBuilder'
  - name: ecs
    go_name: Ecs
    type: object
    go_type: '*ecs.Provider'
  - name: consul
    go_name: Consul
    type: object
    go_type: '*consul.ProviderBuilder'
  - name: etcd
    go_name: Etcd
    type: object
    go_type: '*etcd.Provider'
  - name: zooKeeper
    go_name: ZooKeeper
    type: object
    go_type: '*zk.Provider'
  - name: redis
    go_name: Redis
    type: object
    go_type: '*redis.Provider'
    description: Redis hold the configs of Redis as bucket in rate limiter.
  - name: http
    go_name: HTTP
    type: object
    go_type: '*http.Provider'
  - name: plugin
    go_name: Plugin
    type: object
    items: object
    go_type: map[string]PluginConf
    description: 'Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/overview/#community-middlewares'
---

# Providers

Providers contains providers configuration.
