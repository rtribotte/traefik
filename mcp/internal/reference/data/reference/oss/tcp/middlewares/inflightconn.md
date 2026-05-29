---
schema_version: 2
kind: middleware-tcp
name: TCPInFlightConn
id: tcp.middlewares.inflightconn
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/tcp_middlewares.go#L19
summary: 'TCPInFlightConn holds the TCP InFlightConn middleware configuration. This middleware prevents services from being overwhelmed with high load, by limiting the number of allowed simultaneous connections for one IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/tcp/inflightconn/'
fields:
  - name: amount
    go_name: Amount
    type: integer
    go_type: int64
    description: Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.
representations:
  yaml_path: tcp.middlewares.<name>.inFlightConn
  toml_path: tcp.middlewares.<name>.inFlightConn
  label_prefix: traefik.tcp.middlewares.<name>.inflightconn
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: MiddlewareTCP
    spec_path: .spec.inFlightConn
---

# TCPInFlightConn

TCPInFlightConn holds the TCP InFlightConn middleware configuration. This middleware prevents services from being overwhelmed with high load, by limiting the number of allowed simultaneous connections for one IP. More info: https://doc.traefik.io/traefik/v3.7/middlewares/tcp/inflightconn/
