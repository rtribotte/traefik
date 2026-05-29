---
schema_version: 2
kind: middleware-http
name: EncodedCharacters
id: http.middlewares.encodedcharacters
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L236
summary: EncodedCharacters configures which encoded characters are allowed in the request path.
fields:
  - name: allowEncodedSlash
    go_name: AllowEncodedSlash
    type: boolean
    go_type: bool
    description: AllowEncodedSlash defines whether requests with encoded slash characters in the path are allowed.
  - name: allowEncodedBackSlash
    go_name: AllowEncodedBackSlash
    type: boolean
    go_type: bool
    description: AllowEncodedBackSlash defines whether requests with encoded back slash characters in the path are allowed.
  - name: allowEncodedNullCharacter
    go_name: AllowEncodedNullCharacter
    type: boolean
    go_type: bool
    description: AllowEncodedNullCharacter defines whether requests with encoded null characters in the path are allowed.
  - name: allowEncodedSemicolon
    go_name: AllowEncodedSemicolon
    type: boolean
    go_type: bool
    description: AllowEncodedSemicolon defines whether requests with encoded semicolon characters in the path are allowed.
  - name: allowEncodedPercent
    go_name: AllowEncodedPercent
    type: boolean
    go_type: bool
    description: AllowEncodedPercent defines whether requests with encoded percent characters in the path are allowed.
  - name: allowEncodedQuestionMark
    go_name: AllowEncodedQuestionMark
    type: boolean
    go_type: bool
    description: AllowEncodedQuestionMark defines whether requests with encoded question mark characters in the path are allowed.
  - name: allowEncodedHash
    go_name: AllowEncodedHash
    type: boolean
    go_type: bool
    description: AllowEncodedHash defines whether requests with encoded hash characters in the path are allowed.
representations:
  yaml_path: http.middlewares.<name>.encodedCharacters
  toml_path: http.middlewares.<name>.encodedCharacters
  label_prefix: traefik.http.middlewares.<name>.encodedcharacters
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.encodedCharacters
---

# EncodedCharacters

EncodedCharacters configures which encoded characters are allowed in the request path.
