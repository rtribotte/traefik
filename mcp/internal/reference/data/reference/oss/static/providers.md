---
schema_version: 2
kind: static-section
name: Providers
id: static.providers
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go
summary: Provider configuration. One sub-key per enabled provider.
related:
  - id: oss:provider.docker
    relation: details
  - id: oss:provider.swarm
    relation: details
  - id: oss:provider.file
    relation: details
  - id: oss:provider.kubernetesingress
    relation: details
  - id: oss:provider.kubernetesingressnginx
    relation: details
  - id: oss:provider.kubernetescrd
    relation: details
  - id: oss:provider.kubernetesgateway
    relation: details
  - id: oss:provider.knative
    relation: details
  - id: oss:provider.rest
    relation: details
  - id: oss:provider.consulcatalog
    relation: details
  - id: oss:provider.nomad
    relation: details
  - id: oss:provider.ecs
    relation: details
  - id: oss:provider.consul
    relation: details
  - id: oss:provider.etcd
    relation: details
  - id: oss:provider.zookeeper
    relation: details
  - id: oss:provider.redis
    relation: details
  - id: oss:provider.http
    relation: details
representations:
  yaml_path: providers
  toml_path: providers
---

# Providers

Each provider has its own dedicated page under reference/providers/.
