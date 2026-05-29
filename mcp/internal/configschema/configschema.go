// Package configschema validates Traefik static and dynamic configuration
// against the JSON Schemas published by github.com/traefik/reference, the
// machine-readable reference generated from Traefik's source. The schemas are
// embedded so validation works offline and matches the Traefik version the
// schemas were generated from.
package configschema

import (
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"sigs.k8s.io/yaml"
)

//go:embed schemas/static.schema.json schemas/dynamic.schema.json
var schemaFS embed.FS

// Kind selects which schema a configuration is validated against.
type Kind string

const (
	// Static is the install configuration (traefik.yaml/.toml).
	Static Kind = "static"
	// Dynamic is the routing configuration (routers, services, middlewares).
	Dynamic Kind = "dynamic"
)

// Problem is a single schema violation located within the configuration.
type Problem struct {
	// Location is a JSON Pointer into the configuration (empty for the root).
	Location string `json:"location"`
	// Message describes the violation.
	Message string `json:"message"`
}

// Validator holds the compiled static and dynamic schemas.
type Validator struct {
	static  *jsonschema.Schema
	dynamic *jsonschema.Schema
}

// New compiles the embedded schemas. It fails only if the embedded schemas are
// malformed, which would be a build-time regression.
func New() (*Validator, error) {
	static, err := compile("schemas/static.schema.json")
	if err != nil {
		return nil, fmt.Errorf("static schema: %w", err)
	}
	dynamic, err := compile("schemas/dynamic.schema.json")
	if err != nil {
		return nil, fmt.Errorf("dynamic schema: %w", err)
	}
	return &Validator{static: static, dynamic: dynamic}, nil
}

func compile(name string) (*jsonschema.Schema, error) {
	body, err := schemaFS.ReadFile(name)
	if err != nil {
		return nil, err
	}
	doc, err := jsonschema.UnmarshalJSON(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource(name, doc); err != nil {
		return nil, err
	}
	return c.Compile(name)
}

// Validate parses a YAML or JSON configuration and validates it against the
// schema for kind. It returns the schema violations (empty when valid). The
// error is non-nil only when the input cannot be parsed at all.
func (v *Validator) Validate(kind Kind, config []byte) ([]Problem, error) {
	schema := v.static
	if kind == Dynamic {
		schema = v.dynamic
	}

	// YAMLToJSON accepts JSON as-is (JSON is valid YAML) and produces only
	// JSON-native types, which the validator requires.
	jsonBytes, err := yaml.YAMLToJSON(config)
	if err != nil {
		return nil, fmt.Errorf("parse configuration: %w", err)
	}

	inst, err := jsonschema.UnmarshalJSON(strings.NewReader(string(jsonBytes)))
	if err != nil {
		return nil, fmt.Errorf("parse configuration: %w", err)
	}

	err = schema.Validate(inst)
	if err == nil {
		return nil, nil
	}

	var verr *jsonschema.ValidationError
	if !errors.As(err, &verr) {
		return nil, err
	}

	return flatten(verr), nil
}

// flatten turns the validator's basic output into a sorted-by-appearance list
// of leaf violations.
func flatten(verr *jsonschema.ValidationError) []Problem {
	out := verr.BasicOutput()

	var problems []Problem
	if out.Error != nil {
		problems = append(problems, Problem{Location: out.InstanceLocation, Message: out.Error.String()})
	}
	for _, unit := range out.Errors {
		if unit.Error == nil {
			continue
		}
		problems = append(problems, Problem{Location: unit.InstanceLocation, Message: unit.Error.String()})
	}
	return problems
}
