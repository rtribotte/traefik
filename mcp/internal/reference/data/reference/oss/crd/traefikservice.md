---
schema_version: 2
kind: crd
name: TraefikService
id: crd.traefikservice
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_traefikservices.yaml
summary: TraefikService is the CRD implementation of a Traefik Service.
fields:
  - name: failover
    type: object
    description: Failover defines the Failover service configuration.
  - name: highestRandomWeight
    type: object
    description: HighestRandomWeight defines the highest random weight service configuration.
  - name: mirroring
    type: object
    description: Mirroring defines the Mirroring service configuration.
  - name: weighted
    type: object
    description: Weighted defines the Weighted Round Robin configuration.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: TraefikService
    spec_path: .spec
---

# TraefikService

TraefikService is the CRD implementation of a Traefik Service.
TraefikService object allows to:
- Apply weight to Services on load-balancing
- Mirror traffic on services
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/kubernetes/crd/http/traefikservice/
