package vulnerabilities

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Policies struct {
	Rules      []Rules `json:"rules"`
	PolicyType string  `json:"policyType"`
	ID         string  `json:"_id"`
}
type Resources struct {
	Hosts      []string `json:"hosts"`
	Images     []string `json:"images"`
	Labels     []string `json:"labels"`
	Containers []string `json:"containers"`
	Namespaces []string `json:"namespaces"`
	AccountIDs []string `json:"accountIDs"`
	Clusters   []string `json:"clusters"`
}

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
	Date    time.Time `json:"date"`
}
type CveRules struct {
	Effect      string     `json:"effect"`
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Expiration  Expiration `json:"expiration"`
}
type Tags struct {
	Effect      string     `json:"effect"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Expiration  Expiration `json:"expiration"`
}
type Rules struct {
	Modified       time.Time      `json:"modified"`
	Owner          string         `json:"owner"`
	Name           string         `json:"name"`
	PreviousName   string         `json:"previousName"`
	Effect         string         `json:"effect"`
	Resources      Resources      `json:"resources"`
	BlockMsg       string         `json:"blockMsg,omitempty"`
	Verbose        bool           `json:"verbose,omitempty"`
	AlertThreshold AlertThreshold `json:"alertThreshold"`
	BlockThreshold BlockThreshold `json:"blockThreshold"`
	CveRules       []CveRules     `json:"cveRules,omitempty"`
	Tags           []Tags         `json:"tags,omitempty"`
	GraceDays      int            `json:"graceDays"`
}

func GetPolicies(c sdk.Client) (*Policies, error) {
	req, err := c.NewRequest("GET", "policies/vulnerability/images", nil)
	if err != nil {
		return nil, err
	}
	var policies Policies
	_, err = c.Do(req, &policies)
	if err != nil {
		return nil, err
	}

	return &policies, nil
}

func SetPolicies(c sdk.Client, policies *Policies) error {
	req, err := c.NewRequest("PUT", "policies/vulnerability/images", policies)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
