---
schema_version: 2
kind: concept
name: TLSClientCertificateSubjectDNInfo
id: concept.tlsclientcertificatesubjectdninfo
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L851
summary: TLSClientCertificateSubjectDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
fields:
  - name: country
    go_name: Country
    type: boolean
    go_type: bool
    description: Country defines whether to add the country information into the subject.
  - name: province
    go_name: Province
    type: boolean
    go_type: bool
    description: Province defines whether to add the province information into the subject.
  - name: locality
    go_name: Locality
    type: boolean
    go_type: bool
    description: Locality defines whether to add the locality information into the subject.
  - name: organization
    go_name: Organization
    type: boolean
    go_type: bool
    description: Organization defines whether to add the organization information into the subject.
  - name: organizationalUnit
    go_name: OrganizationalUnit
    type: boolean
    go_type: bool
    description: OrganizationalUnit defines whether to add the organizationalUnit information into the subject.
  - name: commonName
    go_name: CommonName
    type: boolean
    go_type: bool
    description: CommonName defines whether to add the organizationalUnit information into the subject.
  - name: serialNumber
    go_name: SerialNumber
    type: boolean
    go_type: bool
    description: SerialNumber defines whether to add the serialNumber information into the subject.
  - name: domainComponent
    go_name: DomainComponent
    type: boolean
    go_type: bool
    description: DomainComponent defines whether to add the domainComponent information into the subject.
---

# TLSClientCertificateSubjectDNInfo

TLSClientCertificateSubjectDNInfo holds the client TLS certificate distinguished name info configuration. cf https://tools.ietf.org/html/rfc3739
