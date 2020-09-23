package sdk

import (
	"fmt"
)

type AlertSettings struct {
	AggregationPeriodMs int `json:"aggregationPeriodMs"`
}

func (c *Client) GetAlertSettings() (*AlertSettings, error) {
	req, err := c.newRequest("GET", "settings/alerts", nil)
	if err != nil {
		return nil, err
	}

	alertSettings := AlertSettings{}
	_, err = c.do(req, &alertSettings)
	if err != nil {
		return nil, fmt.Errorf("alertSetting not found")
	}

	return &alertSettings, nil
}

func (c *Client) SetAlertSettings(spec *AlertSettings) error {
	req, err := c.newRequest("POST", "settings/alerts", spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
