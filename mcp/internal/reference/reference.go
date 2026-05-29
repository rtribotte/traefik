// Package reference exposes the machine-readable Traefik configuration
// reference published at github.com/traefik/reference. The whole reference
// (per-concept Markdown pages, the parallel JSON Schemas, and the discovery,
// documentation and schema indexes) is vendored under data/ and embedded, so
// grounding and schema validation work offline and stay pinned to the Traefik
// version the reference was generated from.
//
// It implements the two consumption patterns the reference is designed for:
// structured grounding (an index of concepts, each resolvable to its Markdown
// contract, its JSON Schema, and its narrative documentation page) and schema
// validation by catalogue rotation (an unknown configuration is matched against
// the schema registry rather than a single hard-coded schema).
package reference

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"sigs.k8s.io/yaml"
)

//go:embed data
var dataFS embed.FS

const (
	referenceRoot = "data/reference"
	schemasRoot   = "data/schemas"
)

// Node is a single concept in the reference: a configuration object with a
// stable id, resolvable to a Markdown contract page and a parallel JSON Schema.
type Node struct {
	// ID is the concept id, e.g. "http.middlewares.forwardauth".
	ID string `json:"id"`
	// Title is the Go type name behind the concept, e.g. "ForwardAuth".
	Title string `json:"title"`
	// Summary is the one-line description of the concept.
	Summary string `json:"summary"`
	// Kind classifies the concept, e.g. "middleware-http", "crd", "provider".
	Kind string `json:"kind"`
	// Source is the product the concept belongs to: "oss", "hub" or "external".
	Source string `json:"source"`

	path  string
	terms map[string]int
}

// SchemaRef is one entry of the schema registry (schemas/INDEX.json): a schema
// plus the metadata used to decide whether a given document should be validated
// against it.
type SchemaRef struct {
	Path          string   `json:"path"`
	Title         string   `json:"title"`
	Source        string   `json:"source"`
	Scope         string   `json:"scope"`
	ConceptID     string   `json:"concept_id"`
	Kinds         []string `json:"kinds"`
	APIVersions   []string `json:"api_versions"`
	FilenameGlobs []string `json:"filename_globs"`
}

type docEntry struct {
	ConceptID string `json:"concept_id"`
	Source    string `json:"source"`
	DocPath   string `json:"doc_path"`
}

// Catalogue is the in-memory view over the embedded reference.
type Catalogue struct {
	version string

	nodes   map[string]*Node
	ordered []*Node

	docIndex     map[string]docEntry
	docException map[string]bool

	schemas         []SchemaRef
	schemaByConcept map[string]SchemaRef

	mu       sync.Mutex
	compiled map[string]*jsonschema.Schema
}

// New parses the embedded reference into a queryable catalogue. It fails only
// if the embedded data is malformed, which would be a build-time regression.
func New() (*Catalogue, error) {
	c := &Catalogue{
		nodes:           map[string]*Node{},
		docIndex:        map[string]docEntry{},
		docException:    map[string]bool{},
		schemaByConcept: map[string]SchemaRef{},
		compiled:        map[string]*jsonschema.Schema{},
	}
	if err := c.loadNodes(); err != nil {
		return nil, fmt.Errorf("load reference nodes: %w", err)
	}
	if err := c.loadDocIndex(); err != nil {
		return nil, fmt.Errorf("load DOC_INDEX.json: %w", err)
	}
	if err := c.loadSchemaIndex(); err != nil {
		return nil, fmt.Errorf("load schemas/INDEX.json: %w", err)
	}
	return c, nil
}

// Version is the Traefik OSS version the embedded reference was generated from.
func (c *Catalogue) Version() string { return c.version }

type frontmatter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Kind    string `json:"kind"`
	Source  string `json:"source"`
}

func (c *Catalogue) loadNodes() error {
	return fs.WalkDir(dataFS, referenceRoot, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".md") {
			return nil
		}
		body, err := dataFS.ReadFile(p)
		if err != nil {
			return err
		}
		fmBytes, ok := splitFrontmatter(body)
		if !ok {
			return nil
		}
		var fm frontmatter
		// Navigation-only pages have no parseable frontmatter or no id; skip them.
		if err := yaml.Unmarshal(fmBytes, &fm); err != nil || fm.ID == "" {
			return nil
		}
		n := &Node{
			ID:      fm.ID,
			Title:   fm.Name,
			Summary: fm.Summary,
			Kind:    fm.Kind,
			Source:  fm.Source,
			path:    p,
			terms:   tokenize(fm.ID + " " + fm.Name + " " + fm.Summary),
		}
		c.nodes[n.ID] = n
		c.ordered = append(c.ordered, n)
		return nil
	})
}

func (c *Catalogue) loadDocIndex() error {
	body, err := dataFS.ReadFile(referenceRoot + "/DOC_INDEX.json")
	if err != nil {
		return err
	}
	var doc struct {
		Entries    []docEntry `json:"entries"`
		Exceptions []string   `json:"exceptions"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return err
	}
	for _, e := range doc.Entries {
		c.docIndex[e.ConceptID] = e
	}
	for _, id := range doc.Exceptions {
		c.docException[id] = true
	}
	return nil
}

func (c *Catalogue) loadSchemaIndex() error {
	body, err := dataFS.ReadFile(schemasRoot + "/INDEX.json")
	if err != nil {
		return err
	}
	var idx struct {
		Version string      `json:"version"`
		Schemas []SchemaRef `json:"schemas"`
	}
	if err := json.Unmarshal(body, &idx); err != nil {
		return err
	}
	c.version = idx.Version
	c.schemas = idx.Schemas
	for _, s := range idx.Schemas {
		if s.ConceptID != "" {
			c.schemaByConcept[s.ConceptID] = s
		}
	}
	return nil
}

// splitFrontmatter returns the YAML frontmatter block of a Markdown page, or
// false when the page does not start with one.
func splitFrontmatter(b []byte) ([]byte, bool) {
	s := string(b)
	if !strings.HasPrefix(s, "---\n") {
		return nil, false
	}
	rest := s[len("---\n"):]
	i := strings.Index(rest, "\n---")
	if i < 0 {
		return nil, false
	}
	return []byte(rest[:i]), true
}

// schemaPathFor maps a reference Markdown path to its parallel schema path:
// data/reference/<rel>.md -> data/schemas/<rel>.schema.json.
func schemaPathFor(nodePath string) string {
	rel := strings.TrimPrefix(nodePath, referenceRoot+"/")
	rel = strings.TrimSuffix(rel, ".md") + ".schema.json"
	return schemasRoot + "/" + rel
}
