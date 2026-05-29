// Package rag retrieves Traefik configuration knowledge from the
// machine-readable reference published at github.com/traefik/reference. The
// reference's per-concept index (its llms.txt corpus) is embedded so lookups
// work offline; each entry carries a concept id, a summary and the canonical
// documentation URL.
package rag

import (
	"context"
	"embed"
	"regexp"
	"sort"
	"strings"
)

//go:embed corpus/oss.llms.txt corpus/hub.llms.txt
var corpusFS embed.FS

// Result is a single reference entry matched against a query.
type Result struct {
	// ID is the concept id, e.g. "http.middlewares.forwardauth".
	ID string `json:"id"`
	// Title is the human-readable concept name.
	Title string `json:"title"`
	// Summary describes the concept.
	Summary string `json:"summary"`
	// URL points at the canonical reference/documentation page.
	URL string `json:"url"`
	// Section is the corpus heading the entry sits under.
	Section string `json:"section"`
	// Source is the product the entry belongs to: "oss" or "hub".
	Source string `json:"source"`
	// Score is the relevance score for the query (higher is better).
	Score float64 `json:"score"`
}

// Retriever searches the Traefik reference corpus. The embedded implementation
// is offline; an HTTP-backed implementation can swap in later behind this
// interface without touching the tools.
type Retriever interface {
	Search(ctx context.Context, query, source string, limit int) ([]Result, error)
}

type entry struct {
	id, title, summary, url, section, source string
	terms                                    map[string]int
}

// EmbeddedRetriever searches the embedded reference corpus by keyword overlap.
type EmbeddedRetriever struct {
	entries []entry
}

// NewEmbedded parses the embedded corpus into a searchable retriever.
func NewEmbedded() (*EmbeddedRetriever, error) {
	var entries []entry
	for _, src := range []struct{ file, source string }{
		{"corpus/oss.llms.txt", "oss"},
		{"corpus/hub.llms.txt", "hub"},
	} {
		body, err := corpusFS.ReadFile(src.file)
		if err != nil {
			return nil, err
		}
		entries = append(entries, parse(string(body), src.source)...)
	}
	return &EmbeddedRetriever{entries: entries}, nil
}

var entryLine = regexp.MustCompile(`^- \[([^\]]+)\]\(([^)]+)\):\s*(.*)$`)

func parse(body, source string) []entry {
	var out []entry
	var section string
	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(line, "## ") {
			section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			continue
		}
		m := entryLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		id, url, rest := m[1], m[2], strings.TrimSpace(m[3])

		title := rest
		if i := strings.Index(rest, ". "); i >= 0 {
			title = rest[:i]
		}

		e := entry{
			id:      id,
			title:   title,
			summary: rest,
			url:     url,
			section: section,
			source:  source,
		}
		e.terms = tokenize(id + " " + rest + " " + section)
		out = append(out, e)
	}
	return out
}

var wordRE = regexp.MustCompile(`[a-z0-9]+`)

func tokenize(s string) map[string]int {
	terms := map[string]int{}
	for _, w := range wordRE.FindAllString(strings.ToLower(s), -1) {
		terms[w]++
	}
	return terms
}

// Search returns up to limit entries ranked by keyword overlap with query. When
// source is "oss" or "hub" only that product is searched; empty searches both.
func (r *EmbeddedRetriever) Search(_ context.Context, query, source string, limit int) ([]Result, error) {
	if limit <= 0 {
		limit = 5
	}
	queryTerms := tokenize(query)

	var results []Result
	for _, e := range r.entries {
		if source != "" && e.source != source {
			continue
		}
		score := score(queryTerms, e)
		if score <= 0 {
			continue
		}
		results = append(results, Result{
			ID:      e.id,
			Title:   e.title,
			Summary: e.summary,
			URL:     e.url,
			Section: e.section,
			Source:  e.source,
			Score:   score,
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

// score weights matches in the concept id higher than matches in the summary,
// so a query naming a concept surfaces that concept first.
func score(queryTerms map[string]int, e entry) float64 {
	idTerms := tokenize(e.id)
	var total float64
	for term := range queryTerms {
		if _, ok := idTerms[term]; ok {
			total += 3
		}
		if n, ok := e.terms[term]; ok {
			total += float64(n)
		}
	}
	return total
}
