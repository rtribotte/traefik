package logs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// AppEntry is a parsed line from Traefik's JSON application log. level, time and
// message are zerolog's standard fields; error holds the error text Traefik logs
// on failures (often with no message), and Fields keeps any remaining contextual
// keys (routerName, middlewareName, providerName, entryPointName, ...).
type AppEntry struct {
	Time    string            `json:"time"`
	Level   string            `json:"level"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

// appLevelRank orders zerolog levels by severity for MinLevel filtering.
var appLevelRank = map[string]int{
	"trace": 0,
	"debug": 1,
	"info":  2,
	"warn":  3,
	"error": 4,
	"fatal": 5,
	"panic": 6,
}

// AppFilter narrows which app-log entries are returned. Every field is optional;
// a zero value matches everything. Multiple fields combine with AND.
type AppFilter struct {
	Level    string // level equals this exactly (e.g. "error").
	MinLevel string // level severity >= this (e.g. "warn" for warnings and errors).
	Contains string // message or error contains this (case-insensitive).
}

func (f AppFilter) keep(e AppEntry) bool {
	if f.Level != "" && !strings.EqualFold(e.Level, f.Level) {
		return false
	}
	if f.MinLevel != "" && appLevelRank[strings.ToLower(e.Level)] < appLevelRank[strings.ToLower(f.MinLevel)] {
		return false
	}
	if f.Contains != "" {
		hay := strings.ToLower(e.Message + " " + e.Error)
		if !strings.Contains(hay, strings.ToLower(f.Contains)) {
			return false
		}
	}
	return true
}

// TailApp parses r as a JSON application log and returns up to the last n
// entries matching the filter, in file order. Non-JSON lines are skipped.
func TailApp(r io.Reader, n int, filter AppFilter) ([]AppEntry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var matched []AppEntry
	for scanner.Scan() {
		entry, ok := parseAppLine(scanner.Bytes())
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

func parseAppLine(line []byte) (AppEntry, bool) {
	var m map[string]any
	if err := json.Unmarshal(line, &m); err != nil {
		return AppEntry{}, false
	}

	entry := AppEntry{
		Time:    toString(m["time"]),
		Level:   toString(m["level"]),
		Message: toString(m["message"]),
		Error:   toString(m["error"]),
	}

	for k, v := range m {
		switch k {
		case "time", "level", "message", "error":
			continue
		}
		if entry.Fields == nil {
			entry.Fields = map[string]string{}
		}
		entry.Fields[k] = fmt.Sprintf("%v", v)
	}

	return entry, true
}
