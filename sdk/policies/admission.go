package policies

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type AdmissionRules struct {
	ID    string          `json:"_id"`
	Rules []AdmissionRule `json:"rules"`
}
type AdmissionRule struct {
	Modified     time.Time `json:"modified"`
	Owner        string    `json:"owner"`
	Name         string    `json:"name"`
	PreviousName string    `json:"previousName"`
	Effect       string    `json:"effect"`
	Script       string    `json:"script"`
	Description  string    `json:"description"`
	SkipRawReq   bool      `json:"skipRawReq"`
}

func GetAdmissionRules(c sdk.Client) (*AdmissionRules, error) {
	req, err := c.NewRequest("GET", "policies/admission", nil)
	if err != nil {
		return nil, err
	}
	var policies AdmissionRules
	_, err = c.Do(req, &policies)
	if err != nil {
		return nil, err
	}

	return &policies, nil
}

func SetAdmissionRules(c sdk.Client, policies *AdmissionRules) error {
	req, err := c.NewRequest("PUT", "policies/admission", policies)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
