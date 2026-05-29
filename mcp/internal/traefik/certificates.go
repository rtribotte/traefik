package traefik

import (
	"context"
	"math"
	"time"
)

// Certificate is a lean projection of one entry from GET /api/certificates.
// It keeps the operator-relevant fields — identity, validity window and the
// status Traefik itself computes — and drops the fingerprints and DER details
// that only matter for byte-level inspection.
type Certificate struct {
	CommonName         string    `json:"commonName"`
	SANs               []string  `json:"sans,omitempty"`
	NotBefore          time.Time `json:"notBefore"`
	NotAfter           time.Time `json:"notAfter"`
	DaysUntilExpiry    int       `json:"daysUntilExpiry"`
	Status             string    `json:"status"` // enabled, warning (<30d left) or expired, as computed by Traefik.
	Issuer             string    `json:"issuer,omitempty"`
	SerialNumber       string    `json:"serialNumber,omitempty"`
	KeyType            string    `json:"keyType,omitempty"`
	KeySize            int       `json:"keySize,omitempty"`
	SignatureAlgorithm string    `json:"signatureAlgorithm,omitempty"`
}

// apiCertificate mirrors the fields of the API's certificateRepresentation we
// consume. The type is unexported in pkg/api, so it is restated here rather
// than imported.
type apiCertificate struct {
	SANs               []string  `json:"sans"`
	NotAfter           time.Time `json:"notAfter"`
	NotBefore          time.Time `json:"notBefore"`
	SerialNumber       string    `json:"serialNumber"`
	CommonName         string    `json:"commonName"`
	IssuerOrg          string    `json:"issuerOrg"`
	IssuerCN           string    `json:"issuerCN"`
	KeyType            string    `json:"keyType"`
	KeySize            int       `json:"keySize"`
	SignatureAlgorithm string    `json:"signatureAlgorithm"`
	Status             string    `json:"status"`
}

// FetchCertificates retrieves the certificates Traefik is serving. The endpoint
// returns an empty list when TLS is not configured, never an error.
func FetchCertificates(ctx context.Context, target Target) ([]Certificate, error) {
	var raw []apiCertificate
	if err := target.Get(ctx, "/api/certificates", &raw); err != nil {
		return nil, err
	}

	now := time.Now()
	certs := make([]Certificate, 0, len(raw))
	for _, c := range raw {
		certs = append(certs, projectCertificate(c, now))
	}
	return certs, nil
}

func projectCertificate(c apiCertificate, now time.Time) Certificate {
	issuer := c.IssuerOrg
	if issuer == "" {
		issuer = c.IssuerCN
	}

	return Certificate{
		CommonName:         c.CommonName,
		SANs:               c.SANs,
		NotBefore:          c.NotBefore,
		NotAfter:           c.NotAfter,
		DaysUntilExpiry:    int(math.Floor(c.NotAfter.Sub(now).Hours() / 24)),
		Status:             c.Status,
		Issuer:             issuer,
		SerialNumber:       c.SerialNumber,
		KeyType:            c.KeyType,
		KeySize:            c.KeySize,
		SignatureAlgorithm: c.SignatureAlgorithm,
	}
}
