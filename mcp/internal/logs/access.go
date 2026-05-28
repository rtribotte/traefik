// Package logs reads and parses Traefik's file-based access and application logs.
package logs

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"

	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
)

// AccessEntry is a parsed JSON access-log line, projected to the fields an
// operator diagnosing traffic actually needs.
type AccessEntry struct {
	Time       string  `json:"time"`
	Status     int     `json:"status"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	Host       string  `json:"host"`
	Router     string  `json:"router"`
	Service    string  `json:"service"`
	DurationMs float64 `json:"durationMs"`
	ClientHost string  `json:"clientHost"`
}

// AccessFilter narrows which access-log entries are returned.
type AccessFilter struct {
	MinStatus int    // keep entries with status >= MinStatus (0 = no filter).
	Service   string // keep entries whose service name contains this (empty = no filter).
}

func (f AccessFilter) keep(e AccessEntry) bool {
	if f.MinStatus > 0 && e.Status < f.MinStatus {
		return false
	}
	if f.Service != "" && !strings.Contains(e.Service, f.Service) {
		return false
	}
	return true
}

// TailAccess parses r as a JSON access log and returns up to the last n entries
// matching the filter, in file order. Non-JSON lines are skipped.
func TailAccess(r io.Reader, n int, filter AccessFilter) ([]AccessEntry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var matched []AccessEntry
	for scanner.Scan() {
		entry, ok := parseAccessLine(scanner.Bytes())
		if !ok || !filter.keep(entry) {
			continue
		}
		matched = append(matched, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if n > 0 && len(matched) > n {
		matched = matched[len(matched)-n:]
	}
	return matched, nil
}

// parseAccessLine decodes one JSON line. It pulls values by the upstream field
// constants so the keys never drift from what Traefik writes.
func parseAccessLine(line []byte) (AccessEntry, bool) {
	var m map[string]any
	if err := json.Unmarshal(line, &m); err != nil {
		return AccessEntry{}, false
	}

	entry := AccessEntry{
		Status:     toInt(m[accesslog.DownstreamStatus]),
		Method:     toString(m[accesslog.RequestMethod]),
		Path:       toString(m[accesslog.RequestPath]),
		Host:       toString(m[accesslog.RequestHost]),
		Router:     toString(m[accesslog.RouterName]),
		Service:    toString(m[accesslog.ServiceName]),
		ClientHost: toString(m[accesslog.ClientHost]),
		DurationMs: toFloat(m[accesslog.Duration]) / 1e6, // Duration is nanoseconds.
	}
	if entry.Time = toString(m[accesslog.StartUTC]); entry.Time == "" {
		entry.Time = toString(m["time"])
	}
	return entry, true
}

func toString(v any) string {
	s, _ := v.(string)
	return s
}

func toFloat(v any) float64 {
	f, _ := v.(float64)
	return f
}

func toInt(v any) int {
	return int(toFloat(v))
}
