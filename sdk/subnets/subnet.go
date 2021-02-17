package subnet

import (
	"fmt"
	"strings"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Subnet struct {
	ID          string   `json:"_id"`
	Description string   `json:"description"`
	Subnets     []string `json:"subnets"`
}

func Index(c sdk.Client) ([]Subnet, error) {
	req, err := c.NewRequest("GET", "policies/firewall/app/network-list", nil)
	if err != nil {
		return nil, err
	}

	subnets := []Subnet{}
	_, err = c.Do(req, &subnets)
	if err != nil {
		return nil, err
	}

	return subnets, nil
}

func Get(c sdk.Client, subnetName string) (*Subnet, error) {
	resp, err := Index(c)
	if err != nil {
		return nil, err
	}

	for _, i := range resp {
		if strings.Compare(subnetName, i.ID) == 0 {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("subnet: %s not found", subnetName)
}

func Update(c sdk.Client, subnet *Subnet) error {
	req, err := c.NewRequest("PUT", "policies/firewall/app/network-list", subnet)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func Create(c sdk.Client, subnet *Subnet) error {
	req, err := c.NewRequest("POST", "policies/firewall/app/network-list", subnet)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func Delete(c sdk.Client, subnetID string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("policies/firewall/app/network-list/%s", subnetID), nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
