---
schema_version: 2
kind: crd
name: ServersTransportTCP
id: crd.serverstransporttcp
source: oss
traefik_version: v3.7.0
extracted_from:
  - docs/content/reference/dynamic-configuration/traefik.io_serverstransporttcps.yaml
summary: ServersTransportTCP is the CRD implementation of a TCPServersTransport.
fields:
  - name: dialKeepAlive
    type: object
    description: DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
  - name: dialTimeout
    type: object
    description: DialTimeout is the amount of time to wait until a connection to a backend server can be established.
  - name: proxyProtocol
    type: object
    description: ProxyProtocol holds the PROXY Protocol configuration.
  - name: terminationDelay
    type: object
    description: TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.
  - name: tls
    type: object
    description: TLS defines the configuration used to secure the connection to the authentication server.
representations:
  yaml_path: spec
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: ServersTransportTCP
    spec_path: .spec
---

# ServersTransportTCP

ServersTransportTCP is the CRD implementation of a TCPServersTransport.
If no tcpServersTransport is specified, a default one named default@internal will be used.
The default@internal tcpServersTransport can be configured in the static configuration.
More info: https://doc.traefik.io/traefik/v3.7/reference/routing-configuration/tcp/serverstransport/
