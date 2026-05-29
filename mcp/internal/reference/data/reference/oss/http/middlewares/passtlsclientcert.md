---
schema_version: 2
kind: middleware-http
name: PassTLSClientCert
id: http.middlewares.passtlsclientcert
source: oss
traefik_version: v3.7.0
extracted_from:
  - pkg/config/dynamic/middlewares.go#L578
summary: 'PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/passtlsclientcert/'
fields:
  - name: pem
    go_name: PEM
    type: boolean
    go_type: bool
    description: PEM sets the X-Forwarded-Tls-Client-Cert header with the certificate.
  - name: info
    go_name: Info
    type: object
    go_type: '*TLSClientCertificateInfo'
    type_ref: oss:TLSClientCertificateInfo
    description: Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.
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
representations:
  yaml_path: http.middlewares.<name>.passTLSClientCert
  toml_path: http.middlewares.<name>.passTLSClientCert
  label_prefix: traefik.http.middlewares.<name>.passtlsclientcert
  crd:
    apiVersion: traefik.io/v1alpha1
    kind: Middleware
    spec_path: .spec.passTLSClientCert
---

# PassTLSClientCert

PassTLSClientCert holds the pass TLS client cert middleware configuration. This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v3.7/middlewares/http/passtlsclientcert/
