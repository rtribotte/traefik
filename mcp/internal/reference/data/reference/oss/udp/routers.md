---
schema_version: 2
kind: router-udp
name: UDPRouter
id: udp.routers
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/udp_config.go#L57
summary: UDPRouter defines the configuration for an UDP router.
fields:
  - name: entryPoints
    go_name: EntryPoints
    type: array
    items: string
    go_type: '[]string'
    description: 'EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/install-configuration/entrypoints/ Default: all.'
  - name: service
    go_name: Service
    type: string
    go_type: string
    description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
representations:
  yaml_path: udp.routers.<name>
  toml_path: udp.routers.<name>
  label_prefix: traefik.udp.routers.<name>
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: IngressRouteUDP
    spec_path: .spec
---

# UDPRouter

UDPRouter defines the configuration for an UDP router.
