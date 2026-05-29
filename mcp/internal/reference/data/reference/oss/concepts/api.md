---
schema_version: 2
kind: concept
name: API
id: concept.api
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/static/static_config.go#L181
summary: API holds the API configuration.
fields:
  - name: basePath
    go_name: BasePath
    type: string
    go_type: string
    default: /
  - name: insecure
    go_name: Insecure
    type: boolean
    go_type: bool
  - name: dashboard
    go_name: Dashboard
    type: boolean
    go_type: bool
    default: true
  - name: debug
    go_name: Debug
    type: boolean
    go_type: bool
  - name: disableDashboardAd
    go_name: DisableDashboardAd
    type: boolean
    go_type: bool
  - name: dashboardName
    go_name: DashboardName
    type: string
    go_type: string
    default: ""
---

# API

API holds the API configuration.
