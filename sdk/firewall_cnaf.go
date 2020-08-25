package sdk

import (
	"time"
)

type Cnaf struct {
	ID      string  `json:"_id"`
	Rules   []Rules `json:"rules"`
	MinPort int     `json:"minPort"`
	MaxPort int     `json:"maxPort"`
}

type Rules struct {
	Modified             time.Time       `json:"modified"`
	Owner                string          `json:"owner"`
	Name                 string          `json:"name"`
	PreviousName         string          `json:"previousName"`
	Effect               string          `json:"effect"`
	Blacklist            Blacklist       `json:"blacklist"`
	WhitelistSubnets     []string        `json:"whitelistSubnets"`
	Libinject            Libinject       `json:"libinject"`
	Headers              Headers         `json:"headers"`
	Resources            Resources       `json:"resources"`
	Certificate          Certificate     `json:"certificate"`
	CsrfEnabled          bool            `json:"csrfEnabled"`
	ClickjackingEnabled  bool            `json:"clickjackingEnabled"`
	AttackToolsEnabled   bool            `json:"attackToolsEnabled"`
	IntelGathering       IntelGathering  `json:"intelGathering"`
	ShellshockEnabled    bool            `json:"shellshockEnabled"`
	MalformedReqEnabled  bool            `json:"malformedReqEnabled"`
	MaliciousUpload      MaliciousUpload `json:"maliciousUpload"`
	PortMaps             []PortMaps      `json:"portMaps"`
	CmdiEnabled          bool            `json:"cmdiEnabled"`
	LfiEnabled           bool            `json:"lfiEnabled"`
	CodeInjectionEnabled bool            `json:"codeInjectionEnabled"`
}

type Blacklist struct {
	AdvancedProtection bool     `json:"advancedProtection"`
	Subnets            []string `json:"subnets"`
}

type Libinject struct {
	SqliEnabled bool `json:"sqliEnabled"`
	XSSEnabled  bool `json:"xssEnabled"`
}

type Headers struct {
	Specs []Specs `json:"specs"`
}

type Specs struct {
	Allow  bool     `json:"allow"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type Resources struct {
	Hosts      []string `json:"hosts"`
	Images     []string `json:"images"`
	Labels     []string `json:"labels"`
	Containers []string `json:"containers"`
	Namespaces []string `json:"namespaces"`
	AccountIDs []string `json:"accountIDs"`
}

type Certificate struct {
	Encrypted string `json:"encrypted"`
}

type IntelGathering struct {
	BruteforceEnabled         bool `json:"bruteforceEnabled"`
	DirTraversalEnabled       bool `json:"dirTraversalEnabled"`
	TrackErrorsEnabled        bool `json:"trackErrorsEnabled"`
	InfoLeakageEnabled        bool `json:"infoLeakageEnabled"`
	RemoveFingerprintsEnabled bool `json:"removeFingerprintsEnabled"`
}

type MaliciousUpload struct {
	Enabled           bool     `json:"enabled"`
	AllowedFileTypes  []string `json:"allowedFileTypes"`
	AllowedExtensions []string `json:"allowedExtensions"`
}

type PortMaps struct {
	Exposed  int  `json:"exposed"`
	Internal int  `json:"internal"`
	TLS      bool `json:"tls"`
}

func (c *Client) GetCnafRules() (*Cnaf, error) {
	req, err := c.newRequest("GET", "policies/firewall/app/container", nil)
	if err != nil {
		return nil, err
	}
	var cnaf Cnaf
	_, err = c.do(req, &cnaf)
	if err != nil {
		return nil, err
	}

	return &cnaf, nil
}

func (c *Client) SetCnafRules(cnaf *Cnaf) error {
	req, err := c.newRequest("PUT", "policies/firewall/app/container", cnaf)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
