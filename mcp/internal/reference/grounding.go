package reference

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

// SearchResult is a concept matched against a query, with enough metadata to
// decide whether to fetch its full page or schema next.
type SearchResult struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Summary string  `json:"summary"`
	Kind    string  `json:"kind"`
	Source  string  `json:"source"`
	Score   float64 `json:"score"`
}

// Search ranks reference concepts by keyword overlap with query. When source is
// "oss", "hub" or "external" only that product is searched; empty searches all.
// This is the discovery step: it returns concept ids to feed Concept and Schema.
func (c *Catalogue) Search(_ context.Context, query, source string, limit int) ([]SearchResult, error) {
	if limit <= 0 {
		limit = 5
	}
	queryTerms := tokenize(query)

	var results []SearchResult
	for _, n := range c.ordered {
		if source != "" && n.Source != source {
			continue
		}
		s := scoreNode(queryTerms, n)
		if s <= 0 {
			continue
		}
		results = append(results, SearchResult{
			ID:      n.ID,
			Title:   n.Title,
			Summary: n.Summary,
			Kind:    n.Kind,
			Source:  n.Source,
			Score:   s,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].Score != results[j].Score {
			return results[i].Score > results[j].Score
		}
		return results[i].ID < results[j].ID
	})
	if len(results) > limit {
		results = results[:limit]
	}
	return results, nil
}

// Concept returns the full reference Markdown page for a concept id: the field
// contract (names, types, defaults, descriptions) the model should ground on.
func (c *Catalogue) Concept(id string) (string, error) {
	n, ok := c.nodes[id]
	if !ok {
		return "", fmt.Errorf("unknown concept id %q: use search_traefik_docs to find a valid id", id)
	}
	body, err := dataFS.ReadFile(n.path)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Schema returns the JSON Schema for a concept id, suitable for tool input
// schemas or structured generation. It errors when the concept is navigation-
// only and has no parallel schema.
func (c *Catalogue) Schema(id string) (string, error) {
	n, ok := c.nodes[id]
	if !ok {
		return "", fmt.Errorf("unknown concept id %q: use search_traefik_docs to find a valid id", id)
	}
	body, err := dataFS.ReadFile(schemaPathFor(n.path))
	if err != nil {
		return "", fmt.Errorf("no schema for %q (navigation-only concept)", id)
	}
	return string(body), nil
}

// DocResult is the narrative documentation page for a concept.
type DocResult struct {
	ID  string `json:"id"`
	URL string `json:"url"`
	// Markdown is the fetched page body, empty when only a URL is available.
	Markdown string `json:"markdown,omitempty"`
	// Note explains why Markdown is empty (Hub-only, exception, fetch error).
	Note string `json:"note,omitempty"`
}

var rawDocClient = &http.Client{Timeout: 15 * time.Second}

// Doc resolves a concept id to its narrative documentation page via
// DOC_INDEX.json. For OSS concepts it fetches the raw Markdown from the pinned
// Traefik tag; Hub pages are not publicly fetchable, so only a URL is returned.
func (c *Catalogue) Doc(ctx context.Context, id string) (DocResult, error) {
	if _, ok := c.nodes[id]; !ok {
		return DocResult{}, fmt.Errorf("unknown concept id %q: use search_traefik_docs to find a valid id", id)
	}
	if c.docException[id] {
		return DocResult{ID: id, Note: "no dedicated narrative doc page for this concept; use get_traefik_concept for its reference page"}, nil
	}
	entry, ok := c.docIndex[id]
	if !ok {
		return DocResult{ID: id, Note: "no DOC_INDEX entry for this concept; use get_traefik_concept for its reference page"}, nil
	}

	if entry.Source == "hub" {
		return DocResult{
			ID:   id,
			URL:  "https://doc.traefik.io/traefik-hub/" + strings.TrimSuffix(entry.DocPath, ".md"),
			Note: "Hub documentation is not publicly fetchable; use get_traefik_concept for the reference contract",
		}, nil
	}

	version := c.version
	if version == "" {
		version = "master"
	}
	url := fmt.Sprintf("https://raw.githubusercontent.com/traefik/traefik/%s/docs/content/%s", version, entry.DocPath)

	body, err := fetchRaw(ctx, url)
	if err != nil {
		return DocResult{ID: id, URL: url, Note: fmt.Sprintf("fetch failed: %v", err)}, nil
	}
	return DocResult{ID: id, URL: url, Markdown: body}, nil
}

func fetchRaw(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := rawDocClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var wordRE = regexp.MustCompile(`[a-z0-9]+`)

func tokenize(s string) map[string]int {
	terms := map[string]int{}
	for _, w := range wordRE.FindAllString(strings.ToLower(s), -1) {
		terms[w]++
	}
	return terms
}

// scoreNode weights matches in the concept id higher than matches in the title
// or summary, so a query naming a concept surfaces that concept first.
func scoreNode(queryTerms map[string]int, n *Node) float64 {
	idTerms := tokenize(n.ID)
	var total float64
	for term := range queryTerms {
		if _, ok := idTerms[term]; ok {
			total += 3
		}
		if cnt, ok := n.terms[term]; ok {
			total += float64(cnt)
		}
	}
	return total
}
