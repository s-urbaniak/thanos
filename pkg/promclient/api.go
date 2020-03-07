package promclient

import (
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/rules"
)

type AlertAPI struct {
	Alerts []*Alert `json:"alerts"`
}

type Alert struct {
	Labels                  labels.Labels `json:"labels"`
	Annotations             labels.Labels `json:"annotations"`
	State                   string        `json:"state"`
	ActiveAt                *time.Time    `json:"activeAt,omitempty"`
	Value                   string        `json:"value"`
	PartialResponseStrategy string        `json:"partial_response_strategy"`
}

type RuleAPI struct {
	RuleGroups []*RuleGroup `json:"groups"`
}

type RuleGroup struct {
	Name string `json:"name"`
	File string `json:"file"`
	// In order to preserve rule ordering, while exposing type (alerting or recording)
	// specific properties, both alerting and recording rules are exposed in the
	// same array.
	Rules                   []Rule  `json:"rules"`
	Interval                float64 `json:"interval"`
	PartialResponseStrategy string  `json:"partial_response_strategy"`

	EvaluationTime float64   `json:"evaluationTime"`
	LastEvaluation time.Time `json:"lastEvaluation"`
}

type Rule struct {
	// Common.
	Name           string           `json:"name"`
	Labels         labels.Labels    `json:"labels"`
	Query          string           `json:"query"`
	Health         rules.RuleHealth `json:"health"`
	LastError      string           `json:"lastError,omitempty"`
	Type           string           `json:"type"` // recording or alerting.
	EvaluationTime float64          `json:"evaluationTime"`
	LastEvaluation time.Time        `json:"lastEvaluation"`

	// Only for alerts.
	State                   string        `json:"state"` // Is state really in both places?
	Duration                float64       `json:"duration"`
	Annotations             labels.Labels `json:"annotations"`
	Alerts                  []*Alert      `json:"alerts"`
	PartialResponseStrategy string        `json:"partial_response_strategy"`
}
