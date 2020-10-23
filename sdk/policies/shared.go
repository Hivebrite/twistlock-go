package policies

import "time"

type AlertThreshold struct {
	Disabled bool `json:"disabled"`
	Value    int  `json:"value"`
}
type BlockThreshold struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

type Expiration struct {
	Enabled bool      `json:"enabled"`
	Date    time.Time `json:"date,omitempty"`
}
type CveRules struct {
	Effect      string     `json:"effect"`
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Expiration  Expiration `json:"expiration,omitempty"`
}
type Tags struct {
	Effect      string     `json:"effect"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Expiration  Expiration `json:"expiration,omitempty"`
}

const (
	Disable  = "disable"
	Low      = "low"
	Medium   = "medium"
	High     = "high"
	Critical = "critical"
)

const (
	EffectIgnore  = "ignore"
	EffectAlert   = "alert"
	EffectBlock   = "block"
	EffectPrevent = "prevent"
	EffectDisable = "disable"
)

func AlertingLevelToInt(level string) int {
	return map[string]int{
		Disable:  0,
		Low:      1,
		Medium:   4,
		High:     7,
		Critical: 9,
	}[level]
}

func AlertingIntToLevel(level int) string {
	return map[int]string{
		0: Disable,
		1: Low,
		4: Medium,
		7: High,
		9: Critical,
	}[level]
}
