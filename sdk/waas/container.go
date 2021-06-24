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

type Rule struct {
	Modified           time.Time         `json:"modified"`
	Owner              string            `json:"owner"`
	Name               string            `json:"name"`
	PreviousName       string            `json:"previousName"`
	Disabled           bool              `json:"disabled"`
	Collections        []sdk.Collection  `json:"collections"`
	ApplicationsSpec   []ApplicationSpec `json:"applicationsSpec"`
	ReadTimeoutSeconds int               `json:"readTimeoutSeconds"`
}

type ApplicationSpec struct {
	AppID                string                 `json:"appID"`
	SessionCookieEnabled bool                   `json:"sessionCookieEnabled"`
	SessionCookieBan     bool                   `json:"sessionCookieBan"`
	CustomBlockResponse  CustomBlockResponse    `json:"customBlockResponse"`
	BanDurationMinutes   int                    `json:"banDurationMinutes"`
	Certificate          sdk.Secret             `json:"certificate"`
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
	Sqli                 ApplicationSpecEffects `json:"sqli"`
	XSS                  ApplicationSpecEffects `json:"xss"`
	AttackTools          ApplicationSpecEffects `json:"attackTools"`
	Shellshock           ApplicationSpecEffects `json:"shellshock"`
	MalformedReq         ApplicationSpecEffects `json:"malformedReq"`
	Cmdi                 ApplicationSpecEffects `json:"cmdi"`
	Lfi                  ApplicationSpecEffects `json:"lfi"`
	CodeInjection        ApplicationSpecEffects `json:"codeInjection"`
	RemoteHostForwarding RemoteHostForwarding   `json:"remoteHostForwarding"`
}

type CustomBlockResponse struct {
}

type DosConfig struct {
	Enabled              bool              `json:"enabled"`
	TrackSession         bool              `json:"trackSession"`
	Alert                DosConfigRate     `json:"alert"`
	Ban                  DosConfigRate     `json:"ban"`
	MatchConditions      []MatchConditions `json:"matchConditions"`
	ExcludedNetworkLists []string          `json:"excludedNetworkLists"`
}

type DosConfigRate struct {
	Burst   int `json:"burst"`
	Average int `json:"average"`
}

type MatchConditions struct {
	Methods            []string            `json:"methods"`
	FileTypes          []string            `json:"fileTypes"`
	ResponseCodeRanges []ResponseCodeRange `json:"responseCodeRanges"`
}

type ResponseCodeRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type APISpec struct {
	Endpoints      []Endpoint `json:"endpoints"`
	Paths          []Path     `json:"paths"`
	Effect         string     `json:"effect"`
	FallbackEffect string     `json:"fallbackEffect"`
	SkipLearning   bool       `json:"skipLearning"`
}

type Endpoint struct {
	Host         string `json:"host"`
	BasePath     string `json:"basePath"`
	ExposedPort  int    `json:"exposedPort"`
	InternalPort int    `json:"internalPort"`
	TLS          bool   `json:"tls"`
	HTTP2        bool   `json:"http2"`
}

type Path struct {
	Path    string   `json:"path"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Method string `json:"method"`
}

type BotProtectionSpec struct {
	UserDefinedBots          []UserDefinedBot         `json:"userDefinedBots"`
	KnownBotProtectionsSpec  KnownBotProtectionsSpec  `json:"knownBotProtectionsSpec"`
	UnknownBotProtectionSpec UnknownBotProtectionSpec `json:"unknownBotProtectionSpec"`
	SessionValidation        string                   `json:"sessionValidation"`
	InterstitialPage         bool                     `json:"interstitialPage"`
	JsInjectionSpec          JsInjectionSpec          `json:"jsInjectionSpec"`
	ReCAPTCHASpec            ReCAPTCHASpec            `json:"reCAPTCHASpec"`
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

type RequestAnomalies struct {
	Threshold int    `json:"threshold"`
	Effect    string `json:"effect"`
}

type JsInjectionSpec struct {
	Enabled       bool   `json:"enabled"`
	TimeoutEffect string `json:"timeoutEffect"`
}
type ReCAPTCHASpec struct {
	Enabled                bool       `json:"enabled"`
	SiteKey                string     `json:"siteKey"`
	SecretKey              sdk.Secret `json:"secretKey"`
	Type                   string     `json:"type"`
	AllSessions            bool       `json:"allSessions"`
	SuccessExpirationHours int        `json:"successExpirationHours"`
}

type NetworkControls struct {
	AdvancedProtectionEffect string                `json:"advancedProtectionEffect"`
	Subnets                  NetworkControlsEffect `json:"subnets"`
	Countries                NetworkControlsEffect `json:"countries"`
	ExceptionSubnets         []string              `json:"exceptionSubnets"`
}

type NetworkControlsEffect struct {
	Enabled        bool     `json:"enabled"`
	AllowMode      bool     `json:"allowMode"`
	FallbackEffect string   `json:"fallbackEffect"`
	Allow          []string `json:"allow"`
	Alert          []string `json:"alert"`
	Prevent        []string `json:"prevent"`
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

type ApplicationSpecEffects struct {
	Effect          string            `json:"effect"`
	ExceptionFields []ExceptionFields `json:"exceptionFields"`
}

type ExceptionFields struct {
	Location string `json:"location"`
	Key      string `json:"key"`
}

type RemoteHostForwarding struct {
}

func GetContainerWaas(c sdk.Client) (*Waas, error) {
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

func SetContainerWaas(c sdk.Client, waas *Waas) error {
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
