package reference

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/santhosh-tekuri/jsonschema/v6"
	yamlv3 "gopkg.in/yaml.v3"
	"sigs.k8s.io/yaml"
)

// Problem is a single schema violation located within a configuration.
type Problem struct {
	// Location is a JSON Pointer into the document (empty for the root).
	Location string `json:"location"`
	// Message describes the violation.
	Message string `json:"message"`
}

// DocValidation is the validation outcome for one document of the input.
type DocValidation struct {
	// Doc is the zero-based index of the document within a multi-document input.
	Doc int `json:"doc"`
	// Valid reports whether at least one registry schema accepts the document.
	Valid bool `json:"valid"`
	// Matches lists the schemas that accept the document (paths in the registry).
	Matches []string `json:"matches,omitempty"`
	// ClosestSchema is the best-fitting schema when none matches (fewest errors).
	ClosestSchema string `json:"closestSchema,omitempty"`
	// Problems are the violations against ClosestSchema.
	Problems []Problem `json:"problems,omitempty"`
}

// Validation is the result of validating a (possibly multi-document) input.
type Validation struct {
	Valid bool            `json:"valid"`
	Docs  []DocValidation `json:"docs"`
}

// rotationScopes is the set of registry scopes tried when no concept hint is
// given: the document-level schemas (whole file configs, CRDs, annotated
// manifests, core resources). Per-concept fragments are reached via a hint.
var rotationScopes = map[string]bool{
	"traefik-file":        true,
	"kubernetes-crd":      true,
	"traefik-hub-crd":     true,
	"kubernetes-manifest": true,
	"kubernetes-core":     true,
}

// Validate validates a YAML or JSON configuration against the reference schema
// registry by catalogue rotation. When concept is set, the input is validated
// against that single concept's fragment schema; otherwise it is matched
// against every document-level schema and the closest fit is reported when none
// accepts it. Multi-document YAML is validated document by document. The error
// is non-nil only when the input cannot be parsed at all.
func (c *Catalogue) Validate(config []byte, concept string) (Validation, error) {
	candidates, err := c.candidates(concept)
	if err != nil {
		return Validation{}, err
	}

	docs, err := splitYAMLDocs(config)
	if err != nil {
		return Validation{}, fmt.Errorf("parse configuration: %w", err)
	}

	result := Validation{Valid: true}
	for i, raw := range docs {
		inst, err := toJSONValue(raw)
		if err != nil {
			return Validation{}, fmt.Errorf("parse document %d: %w", i, err)
		}

		dv := c.validateDoc(i, inst, candidates)
		if !dv.Valid {
			result.Valid = false
		}
		result.Docs = append(result.Docs, dv)
	}
	return result, nil
}

func (c *Catalogue) candidates(concept string) ([]SchemaRef, error) {
	if concept != "" {
		ref, ok := c.schemaByConcept[concept]
		if !ok {
			return nil, fmt.Errorf("unknown concept %q: use search_traefik_docs to find a valid id", concept)
		}
		return []SchemaRef{ref}, nil
	}

	var out []SchemaRef
	for _, s := range c.schemas {
		if rotationScopes[s.Scope] {
			out = append(out, s)
		}
	}
	return out, nil
}

func (c *Catalogue) validateDoc(index int, inst any, candidates []SchemaRef) DocValidation {
	dv := DocValidation{Doc: index}

	var (
		bestPath     string
		bestProblems []Problem
		haveBest     bool
	)
	for _, ref := range candidates {
		schema, err := c.compiledFor(ref.Path)
		if err != nil {
			continue
		}
		problems := validateAgainst(schema, inst)
		if len(problems) == 0 {
			dv.Matches = append(dv.Matches, ref.Path)
			continue
		}
		if !haveBest || len(problems) < len(bestProblems) {
			bestPath, bestProblems, haveBest = ref.Path, problems, true
		}
	}

	if len(dv.Matches) > 0 {
		dv.Valid = true
		sort.Strings(dv.Matches)
		return dv
	}

	dv.ClosestSchema = bestPath
	dv.Problems = bestProblems
	return dv
}

func (c *Catalogue) compiledFor(relPath string) (*jsonschema.Schema, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if s, ok := c.compiled[relPath]; ok {
		return s, nil
	}

	full := schemasRoot + "/" + relPath
	body, err := dataFS.ReadFile(full)
	if err != nil {
		return nil, err
	}
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	comp := jsonschema.NewCompiler()
	if err := comp.AddResource(full, doc); err != nil {
		return nil, err
	}
	s, err := comp.Compile(full)
	if err != nil {
		return nil, err
	}
	c.compiled[relPath] = s
	return s, nil
}

func validateAgainst(schema *jsonschema.Schema, inst any) []Problem {
	err := schema.Validate(inst)
	if err == nil {
		return nil
	}
	var verr *jsonschema.ValidationError
	if !errors.As(err, &verr) {
		return []Problem{{Message: err.Error()}}
	}
	return flatten(verr)
}

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

// splitYAMLDocs splits a YAML stream into its documents, re-serialising each so
// it can be normalised independently. A JSON input is a single YAML document.
func splitYAMLDocs(b []byte) ([][]byte, error) {
	dec := yamlv3.NewDecoder(bytes.NewReader(b))
	var docs [][]byte
	for {
		var node yamlv3.Node
		err := dec.Decode(&node)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		if node.Kind == 0 {
			continue
		}
		out, err := yamlv3.Marshal(&node)
		if err != nil {
			return nil, err
		}
		docs = append(docs, out)
	}
	return docs, nil
}

// toJSONValue normalises one YAML document into the JSON-native value the
// validator requires.
func toJSONValue(raw []byte) (any, error) {
	jsonBytes, err := yaml.YAMLToJSON(raw)
	if err != nil {
		return nil, err
	}
	return jsonschema.UnmarshalJSON(bytes.NewReader(jsonBytes))
}
