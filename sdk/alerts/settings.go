package alerts

import (
	"fmt"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Settings struct {
	AggregationPeriodMs int `json:"aggregationPeriodMs"`
}

func GetAlertSettings(c sdk.Client) (*Settings, error) {
	req, err := c.NewRequest("GET", "settings/alerts", nil)
	if err != nil {
		return nil, err
	}

	alertSettings := Settings{}
	_, err = c.Do(req, &alertSettings)
	if err != nil {
		return nil, fmt.Errorf("alertSetting not found")
	}

	return &alertSettings, nil
}

func SetAlertSettings(c sdk.Client, spec *Settings) error {
	req, err := c.NewRequest("POST", "settings/alerts", spec)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
