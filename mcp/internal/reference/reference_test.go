package reference

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newCatalogue(t *testing.T) *Catalogue {
	t.Helper()
	c, err := New()
	require.NoError(t, err)
	return c
}

func TestNewLoadsCatalogue(t *testing.T) {
	c := newCatalogue(t)

	assert.NotEmpty(t, c.Version())
	assert.NotEmpty(t, c.nodes)
	assert.NotEmpty(t, c.schemas)

	// A well-known concept is present and resolvable.
	n, ok := c.nodes["http.middlewares.forwardauth"]
	require.True(t, ok)
	assert.Equal(t, "oss", n.Source)
	assert.NotEmpty(t, n.path)
}

func TestSearchRanksConceptByID(t *testing.T) {
	c := newCatalogue(t)

	res, err := c.Search(context.Background(), "forwardauth middleware", "", 5)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	assert.Contains(t, res[0].ID, "forwardauth")
}

func TestSearchFiltersBySource(t *testing.T) {
	c := newCatalogue(t)

	res, err := c.Search(context.Background(), "middleware", "hub", 10)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	for _, r := range res {
		assert.Equal(t, "hub", r.Source)
	}
}

func TestConceptReturnsMarkdown(t *testing.T) {
	c := newCatalogue(t)

	md, err := c.Concept("http.middlewares.forwardauth")
	require.NoError(t, err)
	assert.Contains(t, md, "id: http.middlewares.forwardauth")
	assert.Contains(t, md, "address")
}

func TestConceptUnknownID(t *testing.T) {
	c := newCatalogue(t)

	_, err := c.Concept("does.not.exist")
	require.Error(t, err)
}

func TestSchemaReturnsJSONSchema(t *testing.T) {
	c := newCatalogue(t)

	s, err := c.Schema("http.middlewares.forwardauth")
	require.NoError(t, err)
	assert.Contains(t, s, "\"$schema\"")
	assert.Contains(t, s, "ForwardAuth")
}

func TestValidateDynamicValid(t *testing.T) {
	c := newCatalogue(t)

	const cfg = `
http:
  routers:
    r:
      rule: "Host(` + "`x.localhost`" + `)"
      service: s
  services:
    s:
      loadBalancer:
        servers:
          - url: "http://127.0.0.1:80"
`
	res, err := c.Validate([]byte(cfg), "")
	require.NoError(t, err)
	assert.True(t, res.Valid)
	require.Len(t, res.Docs, 1)
	assert.Contains(t, res.Docs[0].Matches, "oss/dynamic.schema.json")
}

func TestValidateDynamicInvalid(t *testing.T) {
	c := newCatalogue(t)

	const cfg = `
http:
  services:
    s:
      loadBalancer:
        servers: "not-an-array"
`
	res, err := c.Validate([]byte(cfg), "")
	require.NoError(t, err)
	assert.False(t, res.Valid)
	require.Len(t, res.Docs, 1)
	assert.NotEmpty(t, res.Docs[0].Problems)
	assert.NotEmpty(t, res.Docs[0].ClosestSchema)
}

func TestValidateAcceptsJSON(t *testing.T) {
	c := newCatalogue(t)

	const cfg = `{"http":{"services":{"s":{"loadBalancer":{"servers":[{"url":"http://127.0.0.1:80"}]}}}}}`
	res, err := c.Validate([]byte(cfg), "")
	require.NoError(t, err)
	assert.True(t, res.Valid)
}

func TestValidateConceptFragment(t *testing.T) {
	c := newCatalogue(t)

	// A forwardauth fragment validated against its own schema.
	const cfg = `address: "http://auth:8080/verify"`
	res, err := c.Validate([]byte(cfg), "http.middlewares.forwardauth")
	require.NoError(t, err)
	assert.True(t, res.Valid)
}

func TestValidateUnknownConcept(t *testing.T) {
	c := newCatalogue(t)

	_, err := c.Validate([]byte(`{}`), "no.such.concept")
	require.Error(t, err)
}

func TestValidateUnparseable(t *testing.T) {
	c := newCatalogue(t)

	_, err := c.Validate([]byte("this: : not: valid: yaml:"), "")
	require.Error(t, err)
}

func TestDocExceptionReturnsNote(t *testing.T) {
	c := newCatalogue(t)

	// ipwhitelist is a deprecated alias listed under DOC_INDEX exceptions.
	require.True(t, c.docException["http.middlewares.ipwhitelist"])
	res, err := c.Doc(context.Background(), "http.middlewares.ipwhitelist")
	require.NoError(t, err)
	assert.Empty(t, res.Markdown)
	assert.NotEmpty(t, res.Note)
}

func TestSchemaPathForMapping(t *testing.T) {
	got := schemaPathFor("data/reference/oss/http/middlewares/forwardauth.md")
	assert.Equal(t, "data/schemas/oss/http/middlewares/forwardauth.schema.json", got)
	assert.True(t, strings.HasPrefix(got, schemasRoot))
}
