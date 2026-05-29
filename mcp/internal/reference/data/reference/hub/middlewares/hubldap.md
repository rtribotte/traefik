---
schema_version: 2
kind: middleware-hub
name: HubLDAP
id: hub.middlewares.hubldap
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/hubldap/config.go#L19
summary: Configuration holds the configuration for the Hub LDAP middleware.
fields:
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
  - name: session
    go_name: Session
    type: object
    go_type: '*SessionConfig'
    description: Session configuration
  - name: sessionKey
    go_name: SessionKey
    type: string
    go_type: string
  - name: loginUrl
    go_name: LoginURL
    type: string
    go_type: string
    description: URL configuration.
  - name: logoutUrl
    go_name: LogoutURL
    type: string
    go_type: string
  - name: portalLogoUrl
    go_name: PortalLogoURL
    type: string
    go_type: string
    description: Portal configuration.
  - name: portalTitle
    go_name: PortalTitle
    type: string
    go_type: string
representations:
  yaml_path: http.middlewares.<name>.plugin.hubldap
  toml_path: http.middlewares.<name>.plugin.hubldap
  label_prefix: traefik.http.middlewares.<name>.plugin.hubldap
---

# HubLDAP

Configuration holds the configuration for the Hub LDAP middleware.
