---
schema_version: 2
kind: middleware-http
name: Retry
id: http.middlewares.retry
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L757
summary: 'Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/retry/'
fields:
  - name: attempts
    go_name: Attempts
    type: integer
    go_type: int
    description: Attempts defines how many times the request should be retried.
  - name: timeout
    go_name: Timeout
    type: duration
    go_type: ptypes.Duration
    description: Timeout defines how much time the middleware is allowed to retry the request.
  - name: initialInterval
    go_name: InitialInterval
    type: duration
    go_type: ptypes.Duration
    description: InitialInterval defines the first wait time in the exponential backoff series. The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
  - name: maxRequestBodyBytes
    go_name: MaxRequestBodyBytes
    type: integer
    go_type: '*int64'
    default: 2097152
    description: MaxRequestBodyBytes defines the maximum size for the request body.
  - name: status
    go_name: Status
    type: array
    items: string
    go_type: '[]string'
    description: Status defines the range of HTTP status codes to retry on.
  - name: disableRetryOnNetworkError
    go_name: DisableRetryOnNetworkError
    type: boolean
    go_type: bool
    description: DisableRetryOnNetworkError defines whether to disable the retry if an error occurs when transmitting the request to the server.
  - name: retryNonIdempotentMethod
    go_name: RetryNonIdempotentMethod
    type: boolean
    go_type: bool
    description: RetryNonIdempotentMethod activates the retry for non-idempotent methods (POST, LOCK, PATCH)
representations:
  yaml_path: http.middlewares.<name>.retry
  toml_path: http.middlewares.<name>.retry
  label_prefix: traefik.http.middlewares.<name>.retry
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.retry
---

# Retry

Retry holds the retry middleware configuration. This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/retry/
