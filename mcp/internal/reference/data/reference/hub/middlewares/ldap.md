---
schema_version: 2
kind: middleware-hub
name: LDAP
id: hub.middlewares.ldap
source: hub
traefik_version: v3.20.2
extracted_from:
  - hub/pkg/middleware/ldap/config.go#L20
summary: Configuration holds the LDAP Middleware configuration.
fields:
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
  - name: baseDN
    go_name: BaseDN
    type: string
    go_type: string
    description: BaseDN is the base domain name that should be used for bind and search queries.
  - name: attribute
    go_name: Attribute
    type: string
    go_type: string
    description: 'Attribute is the LDAP object attribute used to form a bind DN when sending bind queries: <Attribute>=<Username>,<BaseDN> where the Username is extracted from the Authorization header in the request.'
  - name: searchFilter
    go_name: SearchFilter
    type: string
    go_type: string
    description: 'SearchFilter can be set to enable search and bind mode. When set, this value will be used to filter the results of a search query. Example of a search query: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s)). "%s" can be used as a placeholder that will be replaced by the Username.'
  - name: forwardUsername
    go_name: ForwardUsername
    type: boolean
    go_type: bool
    description: ForwardUsername determines whether a "Username" header should be added to the request, containing the value of the username used to authenticate to the LDAP server.
  - name: forwardUsernameHeader
    go_name: ForwardUsernameHeader
    type: string
    go_type: string
    description: ForwardUsernameHeader sets the name of the header to use to forward the username.
  - name: forwardAuthorization
    go_name: ForwardAuthorization
    type: boolean
    go_type: bool
    description: ForwardAuthorization determines whether the "Authorization" header should be forwarded or stripped from the request.
  - name: wwwAuthenticateHeader
    go_name: WWWAuthenticateHeader
    type: boolean
    go_type: bool
    description: WWWAuthenticateHeader determines whether a "WWW-Authenticate" header should be added to the request if it fails with a 401 Unauthorized status code in order to instruct the User-Agent he should try to authenticate. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate for more information.
  - name: wwwAuthenticateHeaderRealm
    go_name: WWWAuthenticateHeaderRealm
    type: string
    go_type: string
    description: WWWAuthenticateHeaderRealm sets a realm in the "WWW-Authenticate" header.
  - name: groups
    go_name: Groups
    type: object
    go_type: '*Groups'
    description: Groups configuration for group extraction and forwarding.
  - name: forwardHeaders
    go_name: ForwardHeaders
    type: object
    items: string
    go_type: map[string]string
    description: ForwardHeaders maps HTTP header names to LDAP attribute names for forwarding attributes as headers.
representations:
  yaml_path: http.middlewares.<name>.plugin.ldap
  toml_path: http.middlewares.<name>.plugin.ldap
  label_prefix: traefik.http.middlewares.<name>.plugin.ldap
---

# LDAP

Configuration holds the LDAP Middleware configuration.
