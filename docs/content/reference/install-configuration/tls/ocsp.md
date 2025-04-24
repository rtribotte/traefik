---
title: "Traefik OCSP Documentation"
description: "Learn how to configure Traefik to use OCSP. Read the technical documentation."
---

# OCSP

Check certificate statuses and do OCSP stapling.
{: .subtitle }

The Online Certificate Status Protocol (OCSP) is an Internet protocol used for obtaining the revocation status of an X.509 digital certificate.

When OCSP is enabled, Traefik will check the status of any configured certificate, and stape the OCSP response to the TLS handshake.

The OCSP check is performed when the certificate is loaded, and once every hour until it is successful at the halfway point before the update date.

## Configuration

### General

Enabling OCSP is part of the [static configuration](../getting-started/configuration-overview.md#the-static-configuration).
It can be defined by using a file (YAML or TOML) or CLI arguments:

```yaml tab="File (YAML)"
## Static configuration
ocsp: {}
```

```toml tab="File (TOML)"
## Static configuration
[ocsp]
```

```bash tab="CLI"
## Static configuration
--ocsp=true
```

### Responder Overrides

The `responderOverrides` option defines the OCSP responder URLs to use instead of the one provided by the certificate.
This is useful when you want to use a different OCSP responder.

```yaml tab="File (YAML)"
## Static configuration
ocsp:
	responderOverrides:
		foo.com: bar.com
```

```toml tab="File (TOML)"
## Static configuration

[ocsp]
  [ocsp.responderOverrides]
    foo.com = "bar.com"
```

```bash tab="CLI"
## Static configuration
-ocsp.responderoverrides.foo.com=bar.com
```
