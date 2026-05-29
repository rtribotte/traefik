---
schema_version: 2
kind: concept
name: ResponseForwarding
id: concept.responseforwarding
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/http_config.go#L453
summary: ResponseForwarding holds the response forwarding configuration.
fields:
  - name: flushInterval
    go_name: FlushInterval
    type: duration
    go_type: ptypes.Duration
    default: 100ms
    description: 'FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body. A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms'
---

# ResponseForwarding

ResponseForwarding holds the response forwarding configuration.
