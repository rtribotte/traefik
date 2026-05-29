---
schema_version: 2
kind: crd
name: Middleware
id: crd.middleware
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_middlewares.yaml
summary: Middleware is the CRD implementation of a Traefik Middleware.
fields:
  - name: addPrefix
    type: object
  - name: basicAuth
    type: object
  - name: buffering
    type: object
  - name: chain
    type: object
  - name: circuitBreaker
    type: object
  - name: compress
    type: object
  - name: contentType
    type: object
  - name: digestAuth
    type: object
  - name: encodedCharacters
    type: object
  - name: errors
    type: object
    description: Errors defines which errors should trigger the use of the fallback service.
  - name: forwardAuth
    type: object
  - name: grpcWeb
    type: object
  - name: headers
    type: object
    description: Headers defines custom headers to be sent to the health check endpoint.
  - name: inFlightReq
    type: object
  - name: ipAllowList
    type: object
    description: 'IPAllowList defines the IPAllowList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipallowlist/'
  - name: ipWhiteList
    type: object
    description: 'IPWhiteList defines the IPWhiteList middleware configuration. This middleware accepts/refuses connections based on the client IP. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/middlewares/ipwhitelist/'
  - name: passTLSClientCert
    type: object
  - name: plugin
    type: object
    description: 'Plugin defines the middleware plugin configuration. More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/overview/#community-middlewares'
  - name: rateLimit
    type: object
  - name: redirectRegex
    type: object
  - name: redirectScheme
    type: object
  - name: replacePath
    type: object
  - name: replacePathRegex
    type: object
  - name: retry
    type: object
  - name: stripPrefix
    type: object
  - name: stripPrefixRegex
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec
---

# Middleware

Middleware is the CRD implementation of a Traefik Middleware.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/http/middlewares/overview/
