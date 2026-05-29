---
schema_version: 2
kind: provider-hub
name: Traefik
id: hub.providers.traefik
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/provider/traefik/internal.go#L35
summary: Provider wraps Traefik's internal provider to handle distributed ACME challenges. It extends Traefik's internal provider functionality by handling both standard and distributed ACME HTTP challenges. This provider should be used instead of Traefik's internal provider when distributed ACME is required.
representations:
  yaml_path: hub.providers.traefik
  toml_path: hub.providers.traefik
---

# Traefik

Provider wraps Traefik's internal provider to handle distributed ACME challenges. It extends Traefik's internal provider functionality by handling both standard and distributed ACME HTTP challenges. This provider should be used instead of Traefik's internal provider when distributed ACME is required.
