---
schema_version: 2
kind: middleware-hub
name: Coraza
id: hub.middlewares.coraza
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/coraza/config.go#L13
summary: Configuration holds the Coraza middleware configuration.
fields:
  - name: directives
    go_name: Directives
    type: array
    items: string
    go_type: '[]string'
    description: Directives parses the directives from the given string and adds them to the WAF.
  - name: crsEnabled
    go_name: CRSEnabled
    type: boolean
    go_type: bool
    description: CRSEnabled coreruleset configs added to coraza.
  - name: txId
    go_name: TxID
    type: string
    go_type: string
    description: TxID is only for the ingress-nginx provider. It should not be documented.
representations:
  yaml_path: http.middlewares.<name>.plugin.coraza
  toml_path: http.middlewares.<name>.plugin.coraza
  label_prefix: traefik.http.middlewares.<name>.plugin.coraza
---

# Coraza

Configuration holds the Coraza middleware configuration.
