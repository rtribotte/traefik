---
schema_version: 2
kind: concept
name: TLSClientCertificateInfo
id: concept.tlsclientcertificateinfo
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L811
summary: TLSClientCertificateInfo holds the client TLS certificate info configuration.
fields:
  - name: notAfter
    go_name: NotAfter
    type: boolean
    go_type: bool
    description: NotAfter defines whether to add the Not After information from the Validity part.
  - name: notBefore
    go_name: NotBefore
    type: boolean
    go_type: bool
    description: NotBefore defines whether to add the Not Before information from the Validity part.
  - name: sans
    go_name: Sans
    type: boolean
    go_type: bool
    description: Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.
  - name: serialNumber
    go_name: SerialNumber
    type: boolean
    go_type: bool
    description: SerialNumber defines whether to add the client serialNumber information.
  - name: subject
    go_name: Subject
    type: object
    go_type: '*TLSClientCertificateSubjectDNInfo'
    type_ref: oss:TLSClientCertificateSubjectDNInfo
    description: Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.
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
  - name: issuer
    go_name: Issuer
    type: object
    go_type: '*TLSClientCertificateIssuerDNInfo'
    type_ref: oss:TLSClientCertificateIssuerDNInfo
    description: Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.
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

# TLSClientCertificateInfo

TLSClientCertificateInfo holds the client TLS certificate info configuration.
