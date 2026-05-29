package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCertificates(t *testing.T) {
	target := newFixtureTarget(t, map[string]string{"/api/certificates": "certificates.json"})
	handler := listCertificates(target)

	_, out, err := handler(context.Background(), nil, listCertificatesInput{})
	require.NoError(t, err)
	require.Len(t, out.Certificates, 2)

	byCN := map[string]int{}
	for i, c := range out.Certificates {
		byCN[c.CommonName] = i
	}

	require.Contains(t, byCN, "admin.localhost")
	admin := out.Certificates[byCN["admin.localhost"]]
	assert.Equal(t, []string{"admin.localhost", "shop.localhost"}, admin.SANs)
	assert.Equal(t, "enabled", admin.Status)
	assert.Equal(t, "Acme Co", admin.Issuer) // issuerOrg preferred over issuerCN.
	assert.Equal(t, "ECDSA", admin.KeyType)
	assert.Positive(t, admin.DaysUntilExpiry) // notAfter is in 2099.

	require.Contains(t, byCN, "legacy.localhost")
	legacy := out.Certificates[byCN["legacy.localhost"]]
	assert.Equal(t, "expired", legacy.Status)
	assert.Equal(t, "Old CA", legacy.Issuer) // falls back to issuerCN.
	assert.Negative(t, legacy.DaysUntilExpiry) // notAfter is in 2020.
}

func TestListCertificates_Empty(t *testing.T) {
	target := &fakeTarget{responses: map[string]json.RawMessage{"/api/certificates": json.RawMessage(`[]`)}}
	_, out, err := listCertificates(target)(context.Background(), nil, listCertificatesInput{})
	require.NoError(t, err)
	assert.Empty(t, out.Certificates)
}
