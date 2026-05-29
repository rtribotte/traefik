package rag

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmbeddedParsesCorpus(t *testing.T) {
	r, err := NewEmbedded()
	require.NoError(t, err)
	assert.NotEmpty(t, r.entries)

	var oss, hub bool
	for _, e := range r.entries {
		switch e.source {
		case "oss":
			oss = true
		case "hub":
			hub = true
		}
		assert.NotEmpty(t, e.id)
		assert.NotEmpty(t, e.url)
	}
	assert.True(t, oss, "expected oss entries")
	assert.True(t, hub, "expected hub entries")
}

func TestSearchRanksConceptByID(t *testing.T) {
	r, err := NewEmbedded()
	require.NoError(t, err)

	res, err := r.Search(context.Background(), "forwardauth middleware", "", 5)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	assert.True(t, strings.Contains(res[0].ID, "forwardauth"), "top result %q should be forwardauth", res[0].ID)
	assert.NotEmpty(t, res[0].URL)
}

func TestSearchFiltersBySource(t *testing.T) {
	r, err := NewEmbedded()
	require.NoError(t, err)

	res, err := r.Search(context.Background(), "middleware", "oss", 50)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	for _, e := range res {
		assert.Equal(t, "oss", e.Source)
	}
}

func TestSearchRespectsLimit(t *testing.T) {
	r, err := NewEmbedded()
	require.NoError(t, err)

	res, err := r.Search(context.Background(), "middleware", "", 3)
	require.NoError(t, err)
	assert.LessOrEqual(t, len(res), 3)
}

func TestSearchNoMatch(t *testing.T) {
	r, err := NewEmbedded()
	require.NoError(t, err)

	res, err := r.Search(context.Background(), "zzzznotaconcept", "", 5)
	require.NoError(t, err)
	assert.Empty(t, res)
}
