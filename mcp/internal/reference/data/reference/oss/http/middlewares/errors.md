---
schema_version: 2
kind: middleware-http
name: ErrorPage
id: http.middlewares.errors
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L257
summary: ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes.
fields:
  - name: status
    go_name: Status
    type: array
    items: string
    go_type: '[]string'
    description: Status defines which status or range of statuses should result in an error page. It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).
  - name: statusRewrites
    go_name: StatusRewrites
    type: object
    items: integer
    go_type: map[string]int
    description: 'StatusRewrites defines a mapping of status codes that should be returned instead of the original error status codes. For example: "418": 404 or "410-418": 404'
  - name: service
    go_name: Service
    type: string
    go_type: string
    description: Service defines the name of the service that will serve the error page.
  - name: query
    go_name: Query
    type: string
    go_type: string
    description: Query defines the URL for the error page (hosted by service). The {status} variable can be used in order to insert the status code in the URL. The {originalStatus} variable can be used in order to insert the upstream status code in the URL. The {url} variable can be used in order to insert the escaped request URL.
  - name: errorRequestHeaders
    go_name: ErrorRequestHeaders
    type: array
    items: string
    go_type: '[]string'
    description: ErrorRequestHeaders defines the list of request headers forwarded to the error page service. When nil (not set), all original request headers are forwarded. Set to an empty list to forward no headers, or list specific headers to forward only those.
  - name: nginxHeaders
    go_name: NginxHeaders
    type: object
    go_type: '*http.Header'
    description: NginxHeaders defines the headers to forward to the Error page service. NginxHeaders option is unexposed to other providers than the IngressNGINX one.
representations:
  yaml_path: http.middlewares.<name>.errors
  toml_path: http.middlewares.<name>.errors
  label_prefix: traefik.http.middlewares.<name>.errors
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.errors
---

# ErrorPage

ErrorPage holds the custom error middleware configuration. This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes.
