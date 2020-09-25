package sdk

import (
	"fmt"
	"strings"
)

type AlertProfile struct {
	Name            string          `json:"name"`
	PreviousName    string          `json:"previousName"`
	ID              string          `json:"_id"`
	Email           Email           `json:"email"`
	Slack           Slack           `json:"slack"`
	Jira            Jira            `json:"jira"`
	SecurityCenter  SecurityCenter  `json:"securityCenter"`
	GcpPubsub       GcpPubsub       `json:"gcpPubsub"`
	SecurityHub     SecurityHub     `json:"securityHub"`
	SecurityAdvisor SecurityAdvisor `json:"securityAdvisor"`
	Pagerduty       Pagerduty       `json:"pagerduty"`
	Webhook         Webhook         `json:"webhook"`
	Policy          Policy          `json:"policy"`
}

type Policy struct {
	Admission                PolicyRule `json:"admission"`
	AppEmbeddedAppFirewall   PolicyRule `json:"appEmbeddedAppFirewall"`
	AppEmbeddedRuntime       PolicyRule `json:"appEmbeddedRuntime"`
	CloudDiscovery           PolicyRule `json:"cloudDiscovery"`
	ContainerAppFirewall     PolicyRule `json:"containerAppFirewall"`
	ContainerCompliance      PolicyRule `json:"containerCompliance"`
	ContainerNetworkFirewall PolicyRule `json:"containerNetworkFirewall"`
	ContainerRuntime         PolicyRule `json:"containerRuntime"`
	ContainerVulnerability   PolicyRule `json:"containerVulnerability"`
	Defender                 PolicyRule `json:"defender"`
	Docker                   PolicyRule `json:"docker"`
	HostAppFirewall          PolicyRule `json:"hostAppFirewall"`
	HostCompliance           PolicyRule `json:"hostCompliance"`
	HostRuntime              PolicyRule `json:"hostRuntime"`
	Incident                 PolicyRule `json:"incident"`
	KubernetesAudit          PolicyRule `json:"kubernetesAudit"`
	ServerlessAppFirewall    PolicyRule `json:"serverlessAppFirewall"`
	ServerlessRuntime        PolicyRule `json:"serverlessRuntime"`
}

type PolicyRule struct {
	Enabled  bool     `json:"enabled"`
	AllRules bool     `json:"allRules"`
	Rules    []string `json:"rules"`
}
type SecurityCenter struct {
	Enabled      bool   `json:"enabled"`
	CredentialID string `json:"credentialId"`
	SourceID     string `json:"sourceID"`
}
type GcpPubsub struct {
	Enabled      bool   `json:"enabled"`
	CredentialID string `json:"credentialId"`
	Topic        string `json:"topic"`
}
type SecurityHub struct {
	Enabled      bool   `json:"enabled"`
	Region       string `json:"region"`
	AccountID    string `json:"accountID"`
	CredentialID string `json:"credentialId"`
	UseAWSRole   bool   `json:"useAWSRole"`
	RoleArn      string `json:"roleArn"`
}
type SecurityAdvisor struct {
	Enabled      bool   `json:"enabled"`
	CredentialID string `json:"credentialID"`
	ProviderID   string `json:"providerId"`
	FindingsURL  string `json:"findingsURL"`
	TokenURL     string `json:"tokenURL"`
}
type Pagerduty struct {
	Enabled    bool   `json:"enabled"`
	RoutingKey Secret `json:"routingKey"`
	Summary    string `json:"summary"`
	Severity   string `json:"severity"`
}
type Webhook struct {
	CredentialID string `json:"credentialId"`
	URL          string `json:"url"`
	Enabled      bool   `json:"enabled"`
}

type Email struct {
	Enabled      bool   `json:"enabled"`
	SMTPAddress  string `json:"smtpAddress"`
	Port         int    `json:"port"`
	CredentialID string `json:"credentialId"`
	From         string `json:"from"`
	Ssl          bool   `json:"ssl"`
}

type Slack struct {
	Enabled    bool     `json:"enabled"`
	WebhookURL string   `json:"webhookUrl"`
	Channels   []string `json:"channels"`
}

type Jira struct {
	Enabled      bool   `json:"enabled"`
	BaseURL      string `json:"baseUrl"`
	CredentialID string `json:"credentialId"`
	CaCert       string `json:"caCert"`
	ProjectKey   struct {
	} `json:"projectKey"`
	IssueType string `json:"issueType"`
	Priority  string `json:"priority"`
	Labels    struct {
	} `json:"labels"`
	Assignee struct {
	} `json:"assignee"`
}

func (c *Client) GetAlertProfiles() ([]AlertProfile, error) {
	req, err := c.newRequest("GET", "alert-profiles", nil)
	if err != nil {
		return nil, err
	}

	alertProfiles := []AlertProfile{}
	_, err = c.do(req, &alertProfiles)
	if err != nil {
		return nil, err
	}

	return alertProfiles, nil
}

func (c *Client) GetAlertProfile(alertProfileName string) (*AlertProfile, error) {
	resp, err := c.GetAlertProfiles()
	if err != nil {
		return nil, err
	}

	for _, i := range resp {
		if strings.Compare(alertProfileName, i.ID) == 0 {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("alertProfile: %s not found", alertProfileName)
}

func (c *Client) SetAlertProfiles(spec *AlertProfile) error {
	req, err := c.newRequest("POST", "alert-profiles", spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteAlertProfile(alertProfileName string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("alert-profiles/%s", alertProfileName), nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
