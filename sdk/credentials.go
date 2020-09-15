package sdk

import (
	"fmt"
	"strings"
	"time"
)

type ProviderCredential struct {
	ID           string    `json:"_id"`
	Type         string    `json:"type"`
	AccountID    string    `json:"accountID"`
	AccountGUID  string    `json:"accountGUID"`
	Secret       Secret    `json:"secret"`
	APIToken     APIToken  `json:"apiToken"`
	LastModified time.Time `json:"lastModified"`
	Owner        string    `json:"owner"`
}

type Secret struct {
	Encrypted string `json:"encrypted"`
	Plain     string `json:"plain"`
}

type APIToken struct {
	Encrypted string `json:"encrypted"`
}

func (c *Client) GetProviderCredentials() ([]ProviderCredential, error) {
	req, err := c.newRequest("GET", "credentials", nil)
	if err != nil {
		return nil, err
	}

	providerCredentials := []ProviderCredential{}
	_, err = c.do(req, &providerCredentials)
	if err != nil {
		return nil, err
	}

	return providerCredentials, nil
}

func (c *Client) GetProviderCredential(providerCredentialName string) (*ProviderCredential, error) {
	resp, err := c.GetProviderCredentials()
	if err != nil {
		return nil, err
	}

	for _, i := range resp {
		if strings.Compare(providerCredentialName, i.ID) == 0 {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("providerCredential: %s not found", providerCredentialName)
}

func (c *Client) SetProviderCredentials(spec *ProviderCredential) error {
	req, err := c.newRequest("POST", "credentials", spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteProviderCredential(providerCredentialName string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("credentials/%s", providerCredentialName), nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
