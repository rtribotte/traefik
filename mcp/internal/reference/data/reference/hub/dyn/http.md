---
schema_version: 2
kind: concept
name: DynExt
id: hub.dyn.http
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/config/dynamic/ext/hub/pkg/config/dynamic/ext/ext.go#L6
summary: Extensions Hub adds on top of Traefik dynamic.HTTPConfiguration and dynamic.Router.
representations:
  yaml_path: http
  toml_path: http
---

# DynExt

Hub overlays two extra fields on the Traefik dynamic configuration:

- `http.uplinks` — map of inter cluster service advertisements. See `hub.concept.uplink`.
- `http.routers.<name>.uplinks` — list of uplink names this router advertises to a parent cluster.
