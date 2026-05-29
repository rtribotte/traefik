---
schema_version: 2
kind: concept
name: Mirroring
id: concept.mirroring
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L202
summary: Mirroring holds the Mirroring configuration.
fields:
  - name: service
    go_name: Service
    type: string
    go_type: string
    description: 'Service defines the reference to a Kubernetes Service that will serve the error page. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/errorpages/#service'
  - name: mirrorBody
    go_name: MirrorBody
    type: boolean
    go_type: '*bool'
    default: true
    description: MirrorBody defines whether the body of the request should be mirrored. Default value is true.
  - name: maxBodySize
    go_name: MaxBodySize
    type: integer
    go_type: '*int64'
    default: -1
    description: MaxBodySize defines the maximum size allowed for the body of the request. If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.
  - name: mirrors
    go_name: Mirrors
    type: array
    items: object
    go_type: '[]MirrorService'
    type_ref: oss:MirrorService
    description: Mirrors defines the list of mirrors where Traefik will duplicate the traffic.
    fields:
      - name: name
        go_name: Name
        type: string
        go_type: string
        description: Name defines the name of the referenced IngressRoute resource.
      - name: percent
        go_name: Percent
        type: integer
        go_type: int
        description: 'Percent defines the part of the traffic to mirror. Supported values: 0 to 100.'
  - name: healthCheck
    go_name: HealthCheck
    type: object
    go_type: '*HealthCheck'
    type_ref: oss:HealthCheck
    description: Healthcheck defines health checks for ExternalName services.
---

# Mirroring

Mirroring holds the Mirroring configuration.
