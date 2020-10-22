package registry

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/credentials"
)

type Specifications struct {
	Settings []Setting `json:"specifications"`
}

type Setting struct {
	Version        string                         `json:"version"`
	Registry       string                         `json:"registry"`
	Repository     string                         `json:"repository"`
	Tag            string                         `json:"tag"`
	Os             string                         `json:"os"`
	Cap            int                            `json:"cap"`
	Hostname       string                         `json:"hostname"`
	Scanners       int                            `json:"scanners"`
	Namespace      string                         `json:"namespace"`
	UseAWSRole     bool                           `json:"useAWSRole"`
	Credential     credentials.ProviderCredential `json:"credential"`
	RoleArn        string                         `json:"roleArn"`
	VersionPattern string                         `json:"versionPattern"`
}

func GetRegistries(c sdk.Client) (*Specifications, error) {
	req, err := c.NewRequest("GET", "settings/registry", nil)
	if err != nil {
		return nil, err
	}

	registries := Specifications{}
	_, err = c.Do(req, &registries)
	if err != nil {
		return nil, err
	}

	return &registries, nil
}

func SetRegistries(c sdk.Client, spec *Specifications) error {
	req, err := c.NewRequest("PUT", "settings/registry", spec)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
