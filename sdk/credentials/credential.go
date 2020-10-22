package credentials

import (
	"fmt"
	"strings"
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type ProviderCredential struct {
	ID           string     `json:"_id"`
	Type         string     `json:"type"`
	AccountID    string     `json:"accountID"`
	AccountGUID  string     `json:"accountGUID"`
	Secret       sdk.Secret `json:"secret"`
	APIToken     APIToken   `json:"apiToken"`
	LastModified time.Time  `json:"lastModified"`
	Owner        string     `json:"owner"`
}

type APIToken struct {
	Encrypted string `json:"encrypted"`
}

func GetProviderCredentials(c sdk.Client) ([]ProviderCredential, error) {
	req, err := c.NewRequest("GET", "credentials", nil)
	if err != nil {
		return nil, err
	}

	providerCredentials := []ProviderCredential{}
	_, err = c.Do(req, &providerCredentials)
	if err != nil {
		return nil, err
	}

	return providerCredentials, nil
}

func GetProviderCredential(c sdk.Client, providerCredentialName string) (*ProviderCredential, error) {
	resp, err := GetProviderCredentials(c)
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

func SetProviderCredentials(c sdk.Client, spec *ProviderCredential) error {
	req, err := c.NewRequest("POST", "credentials", spec)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProviderCredential(c sdk.Client, providerCredentialName string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("credentials/%s", providerCredentialName), nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
