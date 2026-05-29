---
schema_version: 2
kind: static-section
name: Hub
id: hub.static
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/config/static/hub/pkg/config/static/static_config.go#L19
summary: Hub static configuration. Lives under the top level "hub" key of the Traefik static configuration file.
fields:
  - name: offline
    go_name: Offline
    type: boolean
    go_type: '*bool'
  - name: token
    go_name: Token
    type: string
    go_type: string
  - name: tokenFilePath
    go_name: TokenFilePath
    type: string
    go_type: string
  - name: namespaces
    go_name: Namespaces
    type: array
    items: string
    go_type: '[]string'
  - name: apiManagement
    go_name: APIManagement
    type: object
    go_type: '*APIManagement'
  - name: redis
    go_name: Redis
    type: object
    go_type: '*redis.Config'
  - name: sendLogs
    go_name: SendLogs
    type: boolean
    go_type: bool
    default: true
  - name: providers
    go_name: Providers
    type: object
    go_type: '*Providers'
  - name: experimental
    go_name: Experimental
    type: object
    go_type: '*Experimental'
  - name: tracing
    go_name: Tracing
    type: object
    go_type: '*Tracing'
  - name: aigateway
    go_name: AIGateway
    type: object
    go_type: '*AIGateway'
  - name: mcpgateway
    go_name: MCPGateway
    type: object
    go_type: '*MCPGateway'
  - name: pluginRegistry
    go_name: PluginRegistry
    type: object
    go_type: '*PluginRegistry'
  - name: filterHubInternalsFromApi
    go_name: FilterHubInternalsFromAPI
    type: boolean
    go_type: bool
    default: true
  - name: uplinkEntryPoints
    go_name: UplinkEntryPoints
    type: object
    go_type: tstatic.EntryPoints
  - name: platformURL
    go_name: PlatformURL
    type: string
    go_type: string
    description: PlatformURL is the URL of the hub-manager. Only available with an offline license.
representations:
  yaml_path: hub
  toml_path: hub
---

# Hub

Hub is the static configuration of Hub.
