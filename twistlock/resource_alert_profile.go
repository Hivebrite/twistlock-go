package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
)

func resourceAlertProfile() *schema.Resource {
	return &schema.Resource{
		Create: SetAlertProfile,
		Read:   readAlertProfile,
		Update: SetAlertProfile,
		Delete: deleteAlertProfile,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of Alert Profile",
				Required:    true,
				ForceNew:    true,
			},
			"slack": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Slack parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"webhook_url": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "URL of the endpoint to post the messages",
						},
						"enabled": {
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"channels": {
							Required:    true,
							Type:        schema.TypeList,
							Description: "List of channels to notify",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"pager_duty": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Pager duty parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"routing_key": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"encrypted": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
									},
								},
							},
						},
						"summary": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"severity": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
					},
				},
			},
			"webhook": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Webhook parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credential_id": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"url": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
					},
				},
			},
			"policy": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "which events to alert on",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admission": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"app_embedded_app_firewall": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"app_embedded_runtime": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"cloud_discovery": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_app_firewall": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_compliance": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_network_firewall": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_runtime": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_vulnerability": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"defender": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"docker": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_app_firewall": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_compliance": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_runtime": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"incident": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"kubernetes_audit": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"serverless_app_firewall": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"serverless_runtime": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"all_rules": {
										Optional:    true,
										Type:        schema.TypeBool,
										Description: "",
										Default:     false,
									},
									"rules": {
										Required:    false,
										Type:        schema.TypeList,
										Description: "List of rules to be alerted on",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func parseAlertProfile(d *schema.ResourceData) *sdk.AlertProfile {
	slack := d.Get("slack").(map[string]interface{})
	webhookURL := d.Get("webhook").(map[string]interface{})
	policy := d.Get("policy").(map[string]interface{})

	log.Printf("[DEBUG] Slack value: %s", slack)
	log.Printf("[ERROR] Enabled type: %T", slack["enabled"])

	alertProfile := sdk.AlertProfile{
		ID:   d.Get("name").(string),
		Name: d.Get("name").(string),
		Slack: sdk.Slack{
			Enabled:    slack["enabled"].(bool),
			WebhookURL: slack["webhook_url"].(string),
			Channels:   cast.ToStringSlice(slack["channels"]),
		},
		Pagerduty: sdk.Pagerduty{
			Severity: slack["severity"].(string),
			Summary:  slack["summary"].(string),
			RoutingKey: sdk.APIToken{
				Encrypted: slack["routing_key"].(map[string]interface{})["encrypted"].(string),
			},
		},
		Webhook: sdk.Webhook{
			CredentialID: webhookURL["credential_id"].(string),
			URL:          webhookURL["url"].(string),
		},
		Policy: sdk.Policy{
			Admission:                *policyRuleSchemaToInterface(policy["admission"]),
			AppEmbeddedAppFirewall:   *policyRuleSchemaToInterface(policy["app_embedded_app_firewall"]),
			AppEmbeddedRuntime:       *policyRuleSchemaToInterface(policy["app_embedded_runtime"]),
			CloudDiscovery:           *policyRuleSchemaToInterface(policy["cloud_discovery"]),
			ContainerAppFirewall:     *policyRuleSchemaToInterface(policy["container_app_firewall"]),
			ContainerCompliance:      *policyRuleSchemaToInterface(policy["container_compliance"]),
			ContainerNetworkFirewall: *policyRuleSchemaToInterface(policy["container_network_firewall"]),
			ContainerRuntime:         *policyRuleSchemaToInterface(policy["container_runtime"]),
			ContainerVulnerability:   *policyRuleSchemaToInterface(policy["container_vulnerability"]),
			Defender:                 *policyRuleSchemaToInterface(policy["defender"]),
			Docker:                   *policyRuleSchemaToInterface(policy["docker"]),
			HostAppFirewall:          *policyRuleSchemaToInterface(policy["host_app_firewall"]),
			HostCompliance:           *policyRuleSchemaToInterface(policy["host_compliance"]),
			HostRuntime:              *policyRuleSchemaToInterface(policy["host_runtime"]),
			Incident:                 *policyRuleSchemaToInterface(policy["incident"]),
			KubernetesAudit:          *policyRuleSchemaToInterface(policy["kubernetes_audit"]),
			ServerlessAppFirewall:    *policyRuleSchemaToInterface(policy["serverless_app_firewall"]),
			ServerlessRuntime:        *policyRuleSchemaToInterface(policy["serverless_runtime"]),
		},
	}

	return &alertProfile
}

func saveAlertProfile(d *schema.ResourceData, alertProfile *sdk.AlertProfile) error {
	d.SetId(alertProfile.ID)

	err := d.Set("name", alertProfile.ID)
	if err != nil {
		log.Printf("[ERROR] id setting caused by: %s", err)
		return err
	}

	err = d.Set("name", alertProfile.Name)
	if err != nil {
		log.Printf("[ERROR] name setting caused by: %s", err)
		return err
	}

	webhook := map[string]interface{}{
		"credential_id": alertProfile.Webhook.CredentialID,
		"url":           alertProfile.Webhook.URL,
	}

	err = d.Set("webhook", webhook)
	if err != nil {
		log.Printf("[ERROR] webhook setting caused by: %s", err)
		return err
	}

	slack := map[string]interface{}{
		"enabled":     alertProfile.Slack.Enabled,
		"webhook_url": alertProfile.Slack.WebhookURL,
		"channels":    alertProfile.Slack.Channels,
	}

	err = d.Set("slack", slack)
	if err != nil {
		log.Printf("[ERROR] slack setting caused by: %s", err)
		return err
	}

	pagerDuty := map[string]interface{}{
		"severity": alertProfile.Pagerduty.Severity,
		"summary":  alertProfile.Pagerduty.Summary,
		"routing_key": map[string]interface{}{
			"encrypted": alertProfile.Pagerduty.RoutingKey.Encrypted,
		},
	}

	err = d.Set("pager_duty", pagerDuty)
	if err != nil {
		log.Printf("[ERROR] slack setting caused by: %s", err)
		return err
	}

	policy := map[string]interface{}{
		"admission":                policyRuleToInterface(&alertProfile.Policy.Admission),
		"app_embedded_app_firewal": policyRuleToInterface(&alertProfile.Policy.AppEmbeddedAppFirewall),
		"app_embedded_runtime":     policyRuleToInterface(&alertProfile.Policy.AppEmbeddedRuntime),
		"cloud_discovery":          policyRuleToInterface(&alertProfile.Policy.CloudDiscovery),
		"container_app_firewall":   policyRuleToInterface(&alertProfile.Policy.ContainerAppFirewall),
		"container_compliance":     policyRuleToInterface(&alertProfile.Policy.ContainerCompliance),
		"container_network_firewa": policyRuleToInterface(&alertProfile.Policy.ContainerNetworkFirewall),
		"container_runtime":        policyRuleToInterface(&alertProfile.Policy.ContainerRuntime),
		"container_vulnerability":  policyRuleToInterface(&alertProfile.Policy.ContainerVulnerability),
		"defender":                 policyRuleToInterface(&alertProfile.Policy.Defender),
		"docker":                   policyRuleToInterface(&alertProfile.Policy.Docker),
		"host_app_firewall":        policyRuleToInterface(&alertProfile.Policy.HostAppFirewall),
		"host_compliance":          policyRuleToInterface(&alertProfile.Policy.HostCompliance),
		"host_runtime":             policyRuleToInterface(&alertProfile.Policy.HostRuntime),
		"incident":                 policyRuleToInterface(&alertProfile.Policy.Incident),
		"kubernetes_audit":         policyRuleToInterface(&alertProfile.Policy.KubernetesAudit),
		"serverless_app_firewall":  policyRuleToInterface(&alertProfile.Policy.ServerlessAppFirewall),
		"serverless_runtime":       policyRuleToInterface(&alertProfile.Policy.ServerlessRuntime),
	}

	err = d.Set("policy", policy)
	if err != nil {
		log.Printf("[ERROR] policy setting caused by: %s", err)
		return err
	}

	return nil
}

func policyRuleToInterface(policyRule *sdk.PolicyRule) map[string]interface{} {
	return map[string]interface{}{
		"all_rules": policyRule.AllRules,
		"enabled":   policyRule.Enabled,
		"rules":     policyRule.Rules,
	}
}

func policyRuleSchemaToInterface(policyRule interface{}) *sdk.PolicyRule {
	ruleRetyped := policyRule.(map[string]interface{})
	return &sdk.PolicyRule{
		Enabled:  ruleRetyped["enabled"].(bool),
		AllRules: ruleRetyped["all_rules"].(bool),
		Rules:    ruleRetyped["rules"].([]string),
	}
}

func SetAlertProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetAlertProfiles(parseAlertProfile(d))
	if err != nil {
		return err
	}

	if err := readAlertProfile(d, meta); err != nil {
		log.Printf("[ERROR] readAlertProfile func caused by: %s", err)
		return err
	}

	return nil
}

func readAlertProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	alertProfile, err := client.GetAlertProfile(d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveAlertProfile(d, alertProfile)
}

func deleteAlertProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return client.DeleteAlertProfile(d.Id())
}
