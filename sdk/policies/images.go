package policies

import (
	"strings"
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Images struct {
	Rules      []ImageRules `json:"rules"`
	PolicyType string       `json:"policyType"`
	ID         string       `json:"_id"`
}
type ImageRules struct {
	Modified       time.Time        `json:"modified"`
	Owner          string           `json:"owner"`
	Name           string           `json:"name"`
	PreviousName   string           `json:"previousName"`
	Effect         string           `json:"effect"`
	Collections    []sdk.Collection `json:"collections"`
	BlockMsg       string           `json:"blockMsg,omitempty"`
	Verbose        bool             `json:"verbose,omitempty"`
	AlertThreshold AlertThreshold   `json:"alertThreshold"`
	BlockThreshold BlockThreshold   `json:"blockThreshold"`
	CveRules       []CveRules       `json:"cveRules,omitempty"`
	Tags           []Tags           `json:"tags,omitempty"`
	GraceDays      int              `json:"graceDays"`
	OnlyFixed      bool             `json:"onlyFixed"`
}

func GetImages(c sdk.Client) (*Images, error) {
	req, err := c.NewRequest("GET", "policies/vulnerability/images", nil)
	if err != nil {
		return nil, err
	}
	var policies Images
	_, err = c.Do(req, &policies)
	if err != nil {
		return nil, err
	}

	return &policies, nil
}

func SetImages(c sdk.Client, policies *Images) error {
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

func (rule *ImageRules) ComputeEffect() {
	effects := []string{}

	if !rule.AlertThreshold.Disabled {
		effects = append(effects, "alert")
	}

	if rule.BlockThreshold.Enabled {
		effects = append(effects, "block")
	}

	rule.Effect = strings.Join(effects, ", ")
}
