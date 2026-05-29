---
schema_version: 2
kind: middleware-hub
name: SemanticCache
id: hub.middlewares.semanticcache
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/semanticcache/middleware.go#L51
summary: Config holds the config for semantic cache middleware.
fields:
  - name: vectorDB
    go_name: VectorDB
    type: object
    go_type: vectordb.Config
  - name: vectorizer
    go_name: Vectorizer
    type: object
    go_type: vectorizer.Config
  - name: readOnly
    go_name: ReadOnly
    type: boolean
    go_type: bool
  - name: contentTemplate
    go_name: ContentTemplate
    type: string
    go_type: string
  - name: ttl
    go_name: TTL
    type: integer
    go_type: int
  - name: allowBypass
    go_name: AllowBypass
    type: boolean
    go_type: bool
representations:
  yaml_path: http.middlewares.<name>.plugin.semanticcache
  toml_path: http.middlewares.<name>.plugin.semanticcache
  label_prefix: traefik.http.middlewares.<name>.plugin.semanticcache
---

# SemanticCache

Config holds the config for semantic cache middleware.
