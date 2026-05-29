---
schema_version: 2
kind: middleware-http
name: CircuitBreaker
id: http.middlewares.circuitbreaker
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L167
summary: 'CircuitBreaker holds the circuit breaker middleware configuration. This middleware protects the system from stacking requests to unhealthy services, resulting in cascading failures. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/circuitbreaker/'
fields:
  - name: expression
    go_name: Expression
    type: string
    go_type: string
    description: Expression defines the expression that, once matched, opens the circuit breaker and applies the fallback mechanism instead of calling the services.
  - name: checkPeriod
    go_name: CheckPeriod
    type: duration
    go_type: ptypes.Duration
    default: 100ms
    description: CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).
  - name: fallbackDuration
    go_name: FallbackDuration
    type: duration
    go_type: ptypes.Duration
    default: 10s
    description: FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).
  - name: recoveryDuration
    go_name: RecoveryDuration
    type: duration
    go_type: ptypes.Duration
    default: 10s
    description: RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).
  - name: responseCode
    go_name: ResponseCode
    type: integer
    go_type: int
    default: 503
    description: ResponseCode is the status code that the circuit breaker will return while it is in the open state.
representations:
  yaml_path: http.middlewares.<name>.circuitBreaker
  toml_path: http.middlewares.<name>.circuitBreaker
  label_prefix: traefik.http.middlewares.<name>.circuitbreaker
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.circuitBreaker
---

# CircuitBreaker

CircuitBreaker holds the circuit breaker middleware configuration. This middleware protects the system from stacking requests to unhealthy services, resulting in cascading failures. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/circuitbreaker/
