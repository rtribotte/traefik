---
schema_version: 2
kind: concept
name: TLSClientCertificateIssuerDNInfo
id: concept.tlsclientcertificateissuerdninfo
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L830
summary: TLSClientCertificateIssuerDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
fields:
  - name: country
    go_name: Country
    type: boolean
    go_type: bool
    description: Country defines whether to add the country information into the issuer.
  - name: province
    go_name: Province
    type: boolean
    go_type: bool
    description: Province defines whether to add the province information into the issuer.
  - name: locality
    go_name: Locality
    type: boolean
    go_type: bool
    description: Locality defines whether to add the locality information into the issuer.
  - name: organization
    go_name: Organization
    type: boolean
    go_type: bool
    description: Organization defines whether to add the organization information into the issuer.
  - name: commonName
    go_name: CommonName
    type: boolean
    go_type: bool
    description: CommonName defines whether to add the organizationalUnit information into the issuer.
  - name: serialNumber
    go_name: SerialNumber
    type: boolean
    go_type: bool
    description: SerialNumber defines whether to add the serialNumber information into the issuer.
  - name: domainComponent
    go_name: DomainComponent
    type: boolean
    go_type: bool
    description: DomainComponent defines whether to add the domainComponent information into the issuer.
---

# TLSClientCertificateIssuerDNInfo

TLSClientCertificateIssuerDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
