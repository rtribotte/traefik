// Package metrics parses Traefik's Prometheus /metrics text into the few
// runtime signals an operator cares about: whether the last config reload
// succeeded and how traffic breaks down by service and status code.
package metrics

import (
	"fmt"
	"io"
	"math"
	"time"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
)

// MetricFamily aliases the Prometheus type so callers in this module need not
// import the client_model package directly.
type MetricFamily = dto.MetricFamily

const (
	metricLastReload = "traefik_config_last_reload_success"
	metricReloads    = "traefik_config_reloads_total"
	metricSvcReqs    = "traefik_service_requests_total"
	metricEpReqs     = "traefik_entrypoint_requests_total"
)

// Parse reads Prometheus text-format metrics into families keyed by metric name.
func Parse(r io.Reader) (map[string]*MetricFamily, error) {
	parser := expfmt.NewTextParser(model.UTF8Validation)
	fams, err := parser.TextToMetricFamilies(r)
	if err != nil {
		return nil, fmt.Errorf("parsing prometheus metrics: %w", err)
	}
	return fams, nil
}

// ReloadStatus summarises Traefik's configuration reload state.
type ReloadStatus struct {
	// Success is true when the last reload succeeded (the metric carries the
	// success timestamp; a non-zero value means success).
	Success bool `json:"success"`
	// LastReload is when the last successful reload happened; zero if unknown.
	LastReload time.Time `json:"lastReload"`
	// Reloads is the total number of reloads since start.
	Reloads int `json:"reloads"`
}

// ReloadStatusFrom derives the reload state from parsed metric families.
func ReloadStatusFrom(fams map[string]*MetricFamily) ReloadStatus {
	var rs ReloadStatus

	if ts := gaugeValue(fams, metricLastReload); ts > 0 {
		rs.Success = true
		sec, frac := math.Modf(ts)
		rs.LastReload = time.Unix(int64(sec), int64(frac*1e9)).UTC()
	}
	rs.Reloads = int(counterValue(fams, metricReloads))

	return rs
}

// RequestCount is one request-total series: a count for a given scope (service
// or entrypoint), name, status code and method.
type RequestCount struct {
	Scope  string `json:"scope"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Method string `json:"method"`
	Value  int    `json:"value"`
}

// RequestCounts flattens the service and entrypoint request-total metrics into a
// list of counts, tagged by scope.
func RequestCounts(fams map[string]*MetricFamily) []RequestCount {
	var out []RequestCount
	out = append(out, requestCountsFor(fams[metricSvcReqs], "service", "service")...)
	out = append(out, requestCountsFor(fams[metricEpReqs], "entrypoint", "entrypoint")...)
	return out
}

func requestCountsFor(fam *MetricFamily, scope, nameLabel string) []RequestCount {
	if fam == nil {
		return nil
	}

	var out []RequestCount
	for _, m := range fam.GetMetric() {
		labels := labelMap(m)
		out = append(out, RequestCount{
			Scope:  scope,
			Name:   labels[nameLabel],
			Code:   labels["code"],
			Method: labels["method"],
			Value:  int(m.GetCounter().GetValue()),
		})
	}
	return out
}

func gaugeValue(fams map[string]*MetricFamily, name string) float64 {
	fam := fams[name]
	if fam == nil || len(fam.GetMetric()) == 0 {
		return 0
	}
	return fam.GetMetric()[0].GetGauge().GetValue()
}

func counterValue(fams map[string]*MetricFamily, name string) float64 {
	fam := fams[name]
	if fam == nil || len(fam.GetMetric()) == 0 {
		return 0
	}
	return fam.GetMetric()[0].GetCounter().GetValue()
}

func labelMap(m *dto.Metric) map[string]string {
	labels := make(map[string]string, len(m.GetLabel()))
	for _, l := range m.GetLabel() {
		labels[l.GetName()] = l.GetValue()
	}
	return labels
}
