package waas

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type NetworkWaas struct {
	ID               string        `json:"_id"`
	Modified         time.Time     `json:"modified"`
	Owner            string        `json:"owner"`
	HostEnabled      bool          `json:"hostEnabled"`
	HostRules        []interface{} `json:"hostRules"`
	ContainerEnabled bool          `json:"containerEnabled"`
	ContainerRules   []interface{} `json:"containerRules"`
}

func GetNetworkWaas(c sdk.Client) (*NetworkWaas, error) {
	req, err := c.NewRequest("GET", "policies/firewall/network", nil)
	if err != nil {
		return nil, err
	}
	var networkWaas NetworkWaas
	_, err = c.Do(req, &networkWaas)
	if err != nil {
		return nil, err
	}

	return &networkWaas, nil
}

func SetNetworkWaas(c sdk.Client, networkWaas *NetworkWaas) error {
	req, err := c.NewRequest("PUT", "policies/firewall/network", networkWaas)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
