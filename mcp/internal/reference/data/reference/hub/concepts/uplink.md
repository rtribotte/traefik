---
schema_version: 2
kind: concept
name: Uplink
id: hub.concept.uplink
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/dynamic/ext/ext.go#L33
summary: Uplink represents an inter-cluster service advertisement. A child cluster declares an Uplink to advertise to a parent cluster that it can handle a particular workload. This advertisement gets automatically materialized into a service on the parent layer. The parent can then route traffic to the advertised workload.
fields:
  - name: entryPoints
    go_name: EntryPoints
    type: array
    items: string
    go_type: '[]string'
    description: EntryPoints lists the uplink entrypoint names associated with this uplink. If not specified, defaults to all uplink entrypoints marked as default, or all uplink entrypoints if none are explicitly marked.
  - name: weight
    go_name: Weight
    type: integer
    go_type: '*int'
    description: Weight is used for load balancing across multiple clusters referencing the same uplink. Higher weights receive proportionally more traffic.
  - name: healthCheck
    go_name: HealthCheck
    type: object
    go_type: '*UplinkHealthCheck'
    description: HealthCheck configures the active health check on the load balancer service generated for this uplink.
  - name: passiveHealthCheck
    go_name: PassiveHealthCheck
    type: object
    go_type: '*UplinkPassiveHealthCheck'
    description: PassiveHealthCheck configures the passive health check on the load balancer service generated for this uplink.
---

# Uplink

Uplink represents an inter-cluster service advertisement. A child cluster declares an Uplink to advertise to a parent cluster that it can handle a particular workload. This advertisement gets automatically materialized into a service on the parent layer. The parent can then route traffic to the advertised workload.
