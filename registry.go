package twistlock

type Specifications struct {
	RegistrySettings []RegistrySetting `json:"specifications"`
}
type RegistrySetting struct {
	Version    string `json:"version"`
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Os         string `json:"os"`
	Cap        int    `json:"cap"`
	Hostname   string `json:"hostname"`
	Scanners   int    `json:"scanners"`
}

func (c *Client) GetRegistries() (*Specifications, error) {
	req, err := c.newRequest("GET", "settings/registry", nil)
	if err != nil {
		return nil, err
	}

	registries := Specifications{}
	_, err = c.do(req, &registries)
	if err != nil {
		return nil, err
	}

	return &registries, nil
}

func (c *Client) SetRegistries(spec *Specifications) error {
	req, err := c.newRequest("PUT", "settings/registry", spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
