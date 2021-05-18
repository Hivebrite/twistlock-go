package settings

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/mitchellh/mapstructure"
)

type DefenderSettings struct {
	DisconnectPeriodDays          int    `json:"disconnectPeriodDays"`
	ListeningPort                 int    `json:"listeningPort"`
	AutomaticUpgrade              bool   `json:"automaticUpgrade"`
	AdmissionControlEnabled       bool   `json:"admissionControlEnabled"`
	AdmissionControlWebhookSuffix string `json:"admissionControlWebhookSuffix"`
	HostCustomComplianceEnabled   bool   `json:"hostCustomComplianceEnabled"`
}

func GetDefenderSettings(c sdk.Client) (*DefenderSettings, error) {
	var unpacker interface{}

	req, err := c.NewRequest("GET", "settings/system", nil)
	if err != nil {
		return nil, err
	}

	settings := DefenderSettings{}

	_, err = c.Do(req, &unpacker)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(unpacker.(map[string]interface{})["defenderSettings"], &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func UpdateDefenderSettings(c sdk.Client, settings *DefenderSettings) error {
	req, err := c.NewRequest("POST", "settings/defender", settings)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
