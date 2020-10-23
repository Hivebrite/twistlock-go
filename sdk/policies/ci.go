package policies

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Ci struct {
	Rules      []CiRules `json:"rules"`
	PolicyType string    `json:"policyType"`
	ID         string    `json:"_id"`
}
type CiResources struct {
	Images []string `json:"images"`
	Labels []string `json:"labels"`
}
type CiRules struct {
	Modified       time.Time      `json:"modified"`
	Owner          string         `json:"owner"`
	Name           string         `json:"name"`
	PreviousName   string         `json:"previousName"`
	Effect         string         `json:"effect"`
	Resources      CiResources    `json:"resources"`
	Verbose        bool           `json:"verbose,omitempty"`
	AlertThreshold AlertThreshold `json:"alertThreshold"`
	BlockThreshold BlockThreshold `json:"blockThreshold"`
	CveRules       []CveRules     `json:"cveRules,omitempty"`
	Tags           []Tags         `json:"tags,omitempty"`
	GraceDays      int            `json:"graceDays"`
	OnlyFixed      bool           `json:"onlyFixed"`
}

func GetCi(c sdk.Client) (*Ci, error) {
	req, err := c.NewRequest("GET", "policies/vulnerability/ci/images", nil)
	if err != nil {
		return nil, err
	}
	var policies Ci
	_, err = c.Do(req, &policies)
	if err != nil {
		return nil, err
	}

	return &policies, nil
}

func SetCi(c sdk.Client, policies *Ci) error {
	req, err := c.NewRequest("PUT", "policies/vulnerability/ci/images", policies)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
