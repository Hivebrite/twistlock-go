package sdk

import (
	"time"
)

type Waas struct {
	ID      string `json:"_id"`
	Rules   []Rule `json:"rules"`
	MinPort int    `json:"minPort"`
	MaxPort int    `json:"maxPort"`
}

type Rule struct {
	Modified         time.Time         `json:"modified"`
	Owner            string            `json:"owner"`
	Name             string            `json:"name"`
	PreviousName     string            `json:"previousName"`
	Resources        Resources         `json:"resources"`
	ApplicationsSpec []ApplicationSpec `json:"applicationsSpec"`
}

type Resources struct {
	Hosts      []string `json:"hosts"`
	Images     []string `json:"images"`
	Labels     []string `json:"labels"`
	Containers []string `json:"containers"`
	Namespaces []string `json:"namespaces"`
	AccountIDs []string `json:"accountIDs"`
}

type ApplicationSpec struct {
	Certificate         Certificate     `json:"certificate"`
	APISpec             ApiSpec         `json:"apiSpec"`
	NetworkControls     NetworkControls `json:"networkControls"`
	Libinject           LibInject       `json:"libinject"`
	Body                Body            `json:"body"`
	IntelGathering      IntelGathering  `json:"intelGathering"`
	MaliciousUpload     MaliciousUpload `json:"maliciousUpload"`
	AttackToolsEffect   string          `json:"attackToolsEffect"`
	ShellshockEffect    string          `json:"shellshockEffect"`
	MalformedReqEffect  string          `json:"malformedReqEffect"`
	CmdiEffect          string          `json:"cmdiEffect"`
	LfiEffect           string          `json:"lfiEffect"`
	CodeInjectionEffect string          `json:"codeInjectionEffect"`
	CsrfEnabled         bool            `json:"csrfEnabled"`
	ClickjackingEnabled bool            `json:"clickjackingEnabled"`
	HeaderSpecs         []HeaderSpec    `json:"headerSpecs"`
}

type Certificate struct {
	Encrypted string `json:"encrypted"`
}

type ApiSpec struct {
	Endpoints []Endpoint `json:"endpoints"`
	Effect    string     `json:"effect"`
}

type Endpoint struct {
	Host         string `json:"host"`
	BasePath     string `json:"basePath"`
	ExposedPort  int    `json:"exposedPort"`
	InternalPort int    `json:"internalPort"`
	TLS          bool   `json:"tls"`
	HTTP2        bool   `json:"http2"`
}

type NetworkControls struct {
	AdvancedProtectionEffect string   `json:"advancedProtectionEffect"`
	DeniedSubnetsEffect      string   `json:"deniedSubnetsEffect"`
	DeniedCountriesEffect    string   `json:"deniedCountriesEffect"`
	AllowedCountriesEffect   string   `json:"allowedCountriesEffect"`
	DeniedSubnets            []string `json:"deniedSubnets"`
	AllowedSubnets           []string `json:"allowedSubnets"`
	DeniedCountries          []string `json:"deniedCountries"`
	AllowedCountries         []string `json:"allowedCountries"`
}

type LibInject struct {
	SqliEffect string `json:"sqliEffect"`
	XSSEffect  string `json:"xssEffect"`
}

type Body struct {
	InspectionSizeBytes int  `json:"inspectionSizeBytes"`
	Skip                bool `json:"skip"`
}

type IntelGathering struct {
	BruteforceEnabled         bool   `json:"bruteforceEnabled"`
	TrackErrorsEnabled        bool   `json:"trackErrorsEnabled"`
	InfoLeakageEffect         string `json:"infoLeakageEffect"`
	RemoveFingerprintsEnabled bool   `json:"removeFingerprintsEnabled"`
}

type MaliciousUpload struct {
	Effect            string   `json:"effect"`
	AllowedFileTypes  []string `json:"allowedFileTypes"`
	AllowedExtensions []string `json:"allowedExtensions"`
}
type HeaderSpec struct {
	Allow    bool     `json:"allow"`
	Required bool     `json:"required"`
	Effect   string   `json:"effect"`
	Name     string   `json:"name"`
	Values   []string `json:"values"`
}

func (c *Client) GetWaasRules() (*Waas, error) {
	req, err := c.newRequest("GET", "policies/firewall/app/container", nil)
	if err != nil {
		return nil, err
	}
	var waas Waas
	_, err = c.do(req, &waas)
	if err != nil {
		return nil, err
	}

	return &waas, nil
}

func (c *Client) SetWaasRules(waas *Waas) error {
	req, err := c.newRequest("PUT", "policies/firewall/app/container", waas)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
