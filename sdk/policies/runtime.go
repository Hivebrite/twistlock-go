package policies

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Runtime struct {
	ID               string         `json:"_id"`
	Rules            []RuntimeRules `json:"rules"`
	LearningDisabled bool           `json:"learningDisabled"`
}

type Processes struct {
	Effect               string   `json:"effect"`
	Blacklist            []string `json:"blacklist"`
	Whitelist            []string `json:"whitelist"`
	SkipModified         bool     `json:"skipModified"`
	CheckCryptoMiners    bool     `json:"checkCryptoMiners"`
	CheckLateralMovement bool     `json:"checkLateralMovement"`
	CheckParentChild     bool     `json:"checkParentChild"`
}
type Ports struct {
	Start int  `json:"start"`
	End   int  `json:"end"`
	Deny  bool `json:"deny"`
}
type Network struct {
	Effect                  string   `json:"effect"`
	BlacklistIPs            []string `json:"blacklistIPs"`
	BlacklistListeningPorts []Ports  `json:"blacklistListeningPorts"`
	WhitelistListeningPorts []Ports  `json:"whitelistListeningPorts"`
	BlacklistOutboundPorts  []Ports  `json:"blacklistOutboundPorts"`
	WhitelistOutboundPorts  []Ports  `json:"whitelistOutboundPorts"`
	WhitelistIPs            []string `json:"whitelistIPs"`
	SkipModifiedProc        bool     `json:"skipModifiedProc"`
	DetectPortScan          bool     `json:"detectPortScan"`
	SkipRawSockets          bool     `json:"skipRawSockets"`
}
type DNS struct {
	Effect    string   `json:"effect"`
	Whitelist []string `json:"whitelist"`
	Blacklist []string `json:"blacklist"`
}
type Filesystem struct {
	Effect        string   `json:"effect"`
	Blacklist     []string `json:"blacklist"`
	Whitelist     []string `json:"whitelist"`
	CheckNewFiles bool     `json:"checkNewFiles"`
	BackdoorFiles bool     `json:"backdoorFiles"`
}
type CustomRules struct {
	ID     int    `json:"_id"`
	Action string `json:"action"`
	Effect string `json:"effect"`
}
type RuntimeRules struct {
	Modified                 time.Time        `json:"modified"`
	Owner                    string           `json:"owner"`
	Name                     string           `json:"name"`
	Notes                    string           `json:"notes"`
	PreviousName             string           `json:"previousName"`
	Collections              []sdk.Collection `json:"collections"`
	AdvancedProtection       bool             `json:"advancedProtection"`
	Processes                Processes        `json:"processes"`
	Network                  Network          `json:"network"`
	DNS                      DNS              `json:"dns"`
	Filesystem               Filesystem       `json:"filesystem"`
	KubernetesEnforcement    bool             `json:"kubernetesEnforcement"`
	CloudMetadataEnforcement bool             `json:"cloudMetadataEnforcement"`
	CustomRules              []CustomRules    `json:"customRules"`
}

const (
	ActionIncident = "incident"
	ActionAudit    = "audit"
)

func GetRuntime(c sdk.Client) (*Runtime, error) {
	req, err := c.NewRequest("GET", "policies/runtime/container", nil)
	if err != nil {
		return nil, err
	}
	var policies Runtime
	_, err = c.Do(req, &policies)
	if err != nil {
		return nil, err
	}

	return &policies, nil
}

func SetRuntime(c sdk.Client, policies *Runtime) error {
	req, err := c.NewRequest("PUT", "policies/runtime/container", policies)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
