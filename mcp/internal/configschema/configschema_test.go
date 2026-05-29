package configschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateStaticValid(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	cfg := []byte(`
entryPoints:
  web:
    address: ":80"
providers:
  docker: {}
log:
  level: DEBUG
`)
	problems, err := v.Validate(Static, cfg)
	require.NoError(t, err)
	assert.Empty(t, problems)
}

func TestValidateStaticInvalid(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// log.level is a string enum; a map is the wrong type.
	cfg := []byte(`
log:
  level:
    nested: true
`)
	problems, err := v.Validate(Static, cfg)
	require.NoError(t, err)
	require.NotEmpty(t, problems)
	assert.Contains(t, problems[len(problems)-1].Location, "/log/level")
}

func TestValidateDynamicValid(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	cfg := []byte(`
http:
  routers:
    my-router:
      rule: "Host(` + "`example.com`" + `)"
      service: my-service
  services:
    my-service:
      loadBalancer:
        servers:
          - url: "http://127.0.0.1:8080"
`)
	problems, err := v.Validate(Dynamic, cfg)
	require.NoError(t, err)
	assert.Empty(t, problems)
}

func TestValidateDynamicInvalid(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// servers must be an array of objects, not a string.
	cfg := []byte(`
http:
  services:
    my-service:
      loadBalancer:
        servers: "http://127.0.0.1:8080"
`)
	problems, err := v.Validate(Dynamic, cfg)
	require.NoError(t, err)
	assert.NotEmpty(t, problems)
}

func TestValidateAcceptsJSON(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	cfg := []byte(`{"entryPoints":{"web":{"address":":80"}}}`)
	problems, err := v.Validate(Static, cfg)
	require.NoError(t, err)
	assert.Empty(t, problems)
}

func TestValidateUnparseable(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	_, err = v.Validate(Static, []byte("this: : not: valid: yaml:"))
	assert.Error(t, err)
}
