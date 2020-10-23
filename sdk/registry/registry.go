package registry

import (
	"github.com/Hivebrite/twistlock-go/sdk"
)

type Specifications struct {
	Settings []Setting `json:"specifications"`
}

type Setting struct {
	Version        string `json:"version"`
	Registry       string `json:"registry"`
	Repository     string `json:"repository"`
	Tag            string `json:"tag"`
	Os             string `json:"os"`
	Cap            int    `json:"cap"`
	Hostname       string `json:"hostname"`
	Scanners       int    `json:"scanners"`
	Namespace      string `json:"namespace"`
	UseAWSRole     bool   `json:"useAWSRole"`
	CredentialID   string `json:"credentialID"`
	RoleArn        string `json:"roleArn"`
	VersionPattern string `json:"versionPattern"`
}

func Index(c sdk.Client) (*Specifications, error) {
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

func Set(c sdk.Client, spec *Specifications) error {
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
