---
schema_version: 2
kind: concept
name: ProxyProtocol
id: concept.proxyprotocol
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_config.go#L166
summary: 'ProxyProtocol holds the PROXY Protocol configuration. More info: https://doc.traefik.io/traefik/v3.7/routing/services/#proxy-protocol'
fields:
  - name: version
    go_name: Version
    type: integer
    go_type: int
    default: 2
    description: Version defines the PROXY Protocol version to use.
---

# ProxyProtocol

ProxyProtocol holds the PROXY Protocol configuration. More info: https://doc.traefik.io/traefik/v3.7/routing/services/#proxy-protocol
