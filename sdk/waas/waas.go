package waas

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Waas struct {
	ID      string `json:"_id"`
	Rules   []Rule `json:"rules"`
	MinPort int    `json:"minPort"`
	MaxPort int    `json:"maxPort"`
}
type Collection struct {
	Hosts       []string  `json:"hosts"`
	Images      []string  `json:"images"`
	Labels      []string  `json:"labels"`
	Containers  []string  `json:"containers"`
	Functions   []string  `json:"functions"`
	Namespaces  []string  `json:"namespaces"`
	AppIDs      []string  `json:"appIDs"`
	AccountIDs  []string  `json:"accountIDs"`
	CodeRepos   []string  `json:"codeRepos"`
	Clusters    []string  `json:"clusters"`
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Modified    time.Time `json:"modified"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	System      bool      `json:"system"`
}
type Certificate struct {
	Encrypted string `json:"encrypted"`
}

type ResponseCodeRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
type MatchConditions struct {
	Methods            []string            `json:"methods,omitempty"`
	ResponseCodeRanges []ResponseCodeRange `json:"responseCodeRanges,omitempty"`
	FileTypes          []string            `json:"fileTypes,omitempty"`
}
type DosConfig struct {
	TrackSession    bool              `json:"trackSession"`
	Effect          string            `json:"effect"`
	BurstRate       int               `json:"burstRate"`
	AverageRate     int               `json:"averageRate"`
	MatchConditions []MatchConditions `json:"matchConditions"`
}
type Endpoint struct {
	Host         string `json:"host"`
	BasePath     string `json:"basePath"`
	ExposedPort  int    `json:"exposedPort"`
	InternalPort int    `json:"internalPort"`
	TLS          bool   `json:"tls"`
	HTTP2        bool   `json:"http2"`
}
type APISpec struct {
	Endpoints      []Endpoint    `json:"endpoints"`
	Effect         string        `json:"effect"`
	FallbackEffect string        `json:"fallbackEffect"`
	Paths          []interface{} `json:"paths"`
}
type UserDefinedBot struct {
	Name         string   `json:"name"`
	HeaderName   string   `json:"headerName"`
	HeaderValues []string `json:"headerValues"`
	Subnets      []string `json:"subnets"`
	Effect       string   `json:"effect"`
}
type KnownBotProtectionsSpec struct {
	SearchEngineCrawlers string `json:"searchEngineCrawlers"`
	BusinessAnalytics    string `json:"businessAnalytics"`
	Educational          string `json:"educational"`
	News                 string `json:"news"`
	Financial            string `json:"financial"`
	ContentFeedClients   string `json:"contentFeedClients"`
	Archiving            string `json:"archiving"`
	CareerSearch         string `json:"careerSearch"`
	MediaSearch          string `json:"mediaSearch"`
}
type RequestAnomalies struct {
	Threshold int    `json:"threshold"`
	Effect    string `json:"effect"`
}
type UnknownBotProtectionSpec struct {
	Generic              string           `json:"generic"`
	WebAutomationTools   string           `json:"webAutomationTools"`
	WebScrapers          string           `json:"webScrapers"`
	APILibraries         string           `json:"apiLibraries"`
	HTTPLibraries        string           `json:"httpLibraries"`
	BotImpersonation     string           `json:"botImpersonation"`
	BrowserImpersonation string           `json:"browserImpersonation"`
	RequestAnomalies     RequestAnomalies `json:"requestAnomalies"`
}
type JsInjectionSpec struct {
	Enabled       bool   `json:"enabled"`
	TimeoutEffect string `json:"timeoutEffect"`
}
type BotProtectionSpec struct {
	UserDefinedBots           []UserDefinedBot         `json:"userDefinedBots"`
	KnownBotProtectionsSpec   KnownBotProtectionsSpec  `json:"knownBotProtectionsSpec"`
	UnknownBotProtectionsSpec UnknownBotProtectionSpec `json:"unknownBotProtectionSpec"`
	SessionValidation         string                   `json:"sessionValidation"`
	InterstitialPage          bool                     `json:"interstitialPage"`
	JsInjectionSpec           JsInjectionSpec          `json:"jsInjectionSpec"`
}
type NetworkControls struct {
	AdvancedProtectionEffect string   `json:"advancedProtectionEffect"`
	DeniedSubnetsEffect      string   `json:"deniedSubnetsEffect"`
	DeniedSubnets            []string `json:"deniedSubnets"`
	AllowedSubnets           []string `json:"allowedSubnets"`
	DeniedCountries          []string `json:"deniedCountries"`
	AllowedCountries         []string `json:"allowedCountries"`
	DeniedCountriesEffect    string   `json:"deniedCountriesEffect"`
	AllowedCountriesEffect   string   `json:"allowedCountriesEffect"`
}
type Body struct {
	InspectionSizeBytes int  `json:"inspectionSizeBytes"`
	Skip                bool `json:"skip"`
}
type HeaderSpec struct {
	Allow    bool     `json:"allow"`
	Required bool     `json:"required"`
	Effect   string   `json:"effect"`
	Name     string   `json:"name"`
	Values   []string `json:"values"`
}
type IntelGathering struct {
	InfoLeakageEffect         string `json:"infoLeakageEffect"`
	RemoveFingerprintsEnabled bool   `json:"removeFingerprintsEnabled"`
}
type MaliciousUpload struct {
	Effect            string   `json:"effect"`
	AllowedFileTypes  []string `json:"allowedFileTypes"`
	AllowedExtensions []string `json:"allowedExtensions"`
}
type ExceptionFields struct {
	Location string `json:"location"`
	Key      string `json:"key"`
}
type ApplicationSpecEffects struct {
	Effect          string            `json:"effect"`
	ExceptionFields []ExceptionFields `json:"exceptionFields"`
}
type RemoteHostForwarding struct {
}
type ApplicationSpec struct {
	BanDurationMinutes   int                    `json:"banDurationMinutes"`
	Certificate          Certificate            `json:"certificate"`
	DosConfig            DosConfig              `json:"dosConfig"`
	APISpec              APISpec                `json:"apiSpec"`
	BotProtectionSpec    BotProtectionSpec      `json:"botProtectionSpec"`
	NetworkControls      NetworkControls        `json:"networkControls"`
	Body                 Body                   `json:"body"`
	HeaderSpecs          []HeaderSpec           `json:"headerSpecs"`
	IntelGathering       IntelGathering         `json:"intelGathering"`
	MaliciousUpload      MaliciousUpload        `json:"maliciousUpload"`
	CsrfEnabled          bool                   `json:"csrfEnabled"`
	ClickjackingEnabled  bool                   `json:"clickjackingEnabled"`
	SessionCookieEnabled bool                   `json:"sessionCookieEnabled"`
	SessionCookieBan     bool                   `json:"sessionCookieBan"`
	Sqli                 ApplicationSpecEffects `json:"sqli"`
	XSS                  ApplicationSpecEffects `json:"xss"`
	AttackTools          ApplicationSpecEffects `json:"attackTools"`
	Shellshock           ApplicationSpecEffects `json:"shellshock"`
	MalformedReq         ApplicationSpecEffects `json:"malformedReq"`
	Cmdi                 ApplicationSpecEffects `json:"cmdi"`
	Lfi                  ApplicationSpecEffects `json:"lfi"`
	CodeInjection        ApplicationSpecEffects `json:"codeInjection"`
	RemoteHostForwarding RemoteHostForwarding   `json:"remoteHostForwarding"`
	Selected             bool                   `json:"selected"`
}
type Rule struct {
	Modified         time.Time         `json:"modified"`
	Owner            string            `json:"owner"`
	Name             string            `json:"name"`
	PreviousName     string            `json:"previousName"`
	Collections      []Collection      `json:"collections"`
	ApplicationsSpec []ApplicationSpec `json:"applicationsSpec"`
	ExpandDetails    bool              `json:"expandDetails"`
}

func Index(c sdk.Client) (*Waas, error) {
	req, err := c.NewRequest("GET", "policies/firewall/app/container", nil)
	if err != nil {
		return nil, err
	}
	var waas Waas
	_, err = c.Do(req, &waas)
	if err != nil {
		return nil, err
	}

	return &waas, nil
}

func Set(c sdk.Client, waas *Waas) error {
	req, err := c.NewRequest("PUT", "policies/firewall/app/container", waas)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
