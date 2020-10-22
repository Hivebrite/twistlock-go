package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
)

func resourceAlertProfile() *schema.Resource {
	return &schema.Resource{
		Create: setAlertProfile,
		Read:   readAlertProfile,
		Update: setAlertProfile,
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
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,

				MinItems:    1,
				MaxItems:    1,
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
			"pagerduty": {
				Type:     schema.TypeSet,
				MinItems: 1,
				MaxItems: 1,
				Optional: true,
				Computed: true,

				Description: "Pager duty parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"routing_key": {
							Required: true,
							Type:     schema.TypeSet,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plain": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
									},
								},
							},
						},
						"enabled": {
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "",
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
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Webhook parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credential_id": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"enabled": {
							Optional:    true,
							Type:        schema.TypeBool,
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
				Type:        schema.TypeSet,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Description: "which events to alert on",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admission":                  policySchema(),
						"app_embedded_app_firewall":  policySchema(),
						"app_embedded_runtime":       policySchema(),
						"cloud_discovery":            policySchema(),
						"container_app_firewall":     policySchema(),
						"container_compliance":       policySchema(),
						"container_network_firewall": policySchema(),
						"container_runtime":          policySchema(),
						"container_vulnerability":    policySchema(),
						"defender":                   policySchema(),
						"docker":                     policySchema(),
						"host_app_firewall":          policySchema(),
						"host_compliance":            policySchema(),
						"host_runtime":               policySchema(),
						"incident":                   policySchema(),
						"kubernetes_audit":           policySchema(),
						"serverless_app_firewall":    policySchema(),
						"serverless_runtime":         policySchema(),
					},
				},
			},
		},
	}
}

func parseAlertProfile(d *schema.ResourceData) *sdk.AlertProfile {
	slackList := d.Get("slack").(*schema.Set).List()
	webhookList := d.Get("webhook").(*schema.Set).List()
	pagerdutyList := d.Get("pagerduty").(*schema.Set).List()
	policyList := d.Get("policy").(*schema.Set).List()

	slack := sdk.Slack{Enabled: false}
	webhook := sdk.Webhook{Enabled: false}
	pagerduty := sdk.Pagerduty{Enabled: false}
	policy := sdk.Policy{}

	if len(slackList) > 0 {
		slackConfig := slackList[0].(map[string]interface{})
		slack.Enabled = slackConfig["enabled"].(bool)
		slack.WebhookURL = slackConfig["webhook_url"].(string)
		slack.Channels = cast.ToStringSlice(slackConfig["channels"])
	}

	if len(webhookList) > 0 {
		webhookConfig := webhookList[0].(map[string]interface{})
		webhook.CredentialID = webhookConfig["credential_id"].(string)
		webhook.URL = webhookConfig["url"].(string)
		webhook.Enabled = webhookConfig["enabled"].(bool)
	}

	if len(policyList) > 0 {
		policyConfig := policyList[0].(map[string]interface{})

		policy.Admission = *policyRuleSchemaToInterface(policyConfig["admission"])
		policy.AppEmbeddedAppFirewall = *policyRuleSchemaToInterface(policyConfig["app_embedded_app_firewall"])
		policy.AppEmbeddedRuntime = *policyRuleSchemaToInterface(policyConfig["app_embedded_runtime"])
		policy.CloudDiscovery = *policyRuleSchemaToInterface(policyConfig["cloud_discovery"])
		policy.ContainerAppFirewall = *policyRuleSchemaToInterface(policyConfig["container_app_firewall"])
		policy.ContainerCompliance = *policyRuleSchemaToInterface(policyConfig["container_compliance"])
		policy.ContainerRuntime = *policyRuleSchemaToInterface(policyConfig["container_runtime"])
		policy.ContainerVulnerability = *policyRuleSchemaToInterface(policyConfig["container_vulnerability"])
		policy.Defender = *policyRuleSchemaToInterface(policyConfig["defender"])
		policy.Docker = *policyRuleSchemaToInterface(policyConfig["docker"])
		policy.HostAppFirewall = *policyRuleSchemaToInterface(policyConfig["host_app_firewall"])
		policy.HostCompliance = *policyRuleSchemaToInterface(policyConfig["host_compliance"])
		policy.HostRuntime = *policyRuleSchemaToInterface(policyConfig["host_runtime"])
		policy.Incident = *policyRuleSchemaToInterface(policyConfig["incident"])
		policy.KubernetesAudit = *policyRuleSchemaToInterface(policyConfig["kubernetes_audit"])
		policy.ServerlessAppFirewall = *policyRuleSchemaToInterface(policyConfig["serverless_app_firewall"])
		policy.ServerlessRuntime = *policyRuleSchemaToInterface(policyConfig["serverless_runtime"])
	}

	if len(pagerdutyList) > 0 {
		pagerdutyConfig := pagerdutyList[0].(map[string]interface{})
		pagerdutyRoutingKeyConfig := pagerdutyConfig["routing_key"].(*schema.Set).List()[0]

		pagerduty.Severity = pagerdutyConfig["severity"].(string)
		pagerduty.Summary = pagerdutyConfig["summary"].(string)
		pagerduty.RoutingKey = sdk.Secret{
			Plain: pagerdutyRoutingKeyConfig.(map[string]interface{})["plain"].(string),
		}
		pagerduty.Enabled = pagerdutyConfig["enabled"].(bool)
	}

	return &sdk.AlertProfile{
		ID:        d.Get("name").(string),
		Name:      d.Get("name").(string),
		Slack:     slack,
		Webhook:   webhook,
		Pagerduty: pagerduty,
		Policy:    policy,
	}
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

	var webhook []map[string]interface{}
	webhook = append(webhook, map[string]interface{}{
		"credential_id": alertProfile.Webhook.CredentialID,
		"url":           alertProfile.Webhook.URL,
		"enabled":       alertProfile.Webhook.Enabled,
	})

	err = d.Set("webhook", webhook)
	if err != nil {
		log.Printf("[ERROR] webhook setting caused by: %s", err)
		return err
	}

	var slack []map[string]interface{}
	slack = append(slack, map[string]interface{}{
		"enabled":     alertProfile.Slack.Enabled,
		"webhook_url": alertProfile.Slack.WebhookURL,
		"channels":    alertProfile.Slack.Channels,
	})

	err = d.Set("slack", slack)
	if err != nil {
		log.Printf("[ERROR] slack setting caused by: %s", err)
		return err
	}

	var pagerduty []map[string]interface{}
	var pagerdutyRoutingKey []map[string]interface{}

	pagerduty = append(pagerduty, map[string]interface{}{
		"severity": alertProfile.Pagerduty.Severity,
		"summary":  alertProfile.Pagerduty.Summary,
		"enabled":  alertProfile.Pagerduty.Enabled,
		"routing_key": append(pagerdutyRoutingKey, map[string]interface{}{
			"plain": alertProfile.Pagerduty.RoutingKey.Plain,
		}),
	})

	err = d.Set("pagerduty", pagerduty)
	if err != nil {
		log.Printf("[ERROR] pagerduty setting caused by: %s", err)
		return err
	}

	var policy []map[string]interface{}
	policy = append(policy, map[string]interface{}{
		"admission":                 policyRuleToInterface(&alertProfile.Policy.Admission),
		"app_embedded_app_firewall": policyRuleToInterface(&alertProfile.Policy.AppEmbeddedAppFirewall),
		"app_embedded_runtime":      policyRuleToInterface(&alertProfile.Policy.AppEmbeddedRuntime),
		"cloud_discovery":           policyRuleToInterface(&alertProfile.Policy.CloudDiscovery),
		"container_app_firewall":    policyRuleToInterface(&alertProfile.Policy.ContainerAppFirewall),
		"container_compliance":      policyRuleToInterface(&alertProfile.Policy.ContainerCompliance),
		"container_runtime":         policyRuleToInterface(&alertProfile.Policy.ContainerRuntime),
		"container_vulnerability":   policyRuleToInterface(&alertProfile.Policy.ContainerVulnerability),
		"defender":                  policyRuleToInterface(&alertProfile.Policy.Defender),
		"docker":                    policyRuleToInterface(&alertProfile.Policy.Docker),
		"host_app_firewall":         policyRuleToInterface(&alertProfile.Policy.HostAppFirewall),
		"host_compliance":           policyRuleToInterface(&alertProfile.Policy.HostCompliance),
		"host_runtime":              policyRuleToInterface(&alertProfile.Policy.HostRuntime),
		"incident":                  policyRuleToInterface(&alertProfile.Policy.Incident),
		"kubernetes_audit":          policyRuleToInterface(&alertProfile.Policy.KubernetesAudit),
		"serverless_app_firewall":   policyRuleToInterface(&alertProfile.Policy.ServerlessAppFirewall),
		"serverless_runtime":        policyRuleToInterface(&alertProfile.Policy.ServerlessRuntime),
	})

	err = d.Set("policy", policy)
	if err != nil {
		log.Printf("[ERROR] policy setting caused by: %s", err)
		return err
	}

	return nil
}

func policyRuleToInterface(policyRule *sdk.PolicyRule) []map[string]interface{} {
	var policyRuleArray []map[string]interface{}
	policyRuleArray = append(policyRuleArray, map[string]interface{}{
		"all_rules": policyRule.AllRules,
		"enabled":   policyRule.Enabled,
		"rules":     policyRule.Rules,
	})

	return policyRuleArray
}

func policyRuleSchemaToInterface(policyRule interface{}) *sdk.PolicyRule {
	policyRuleList := policyRule.(*schema.Set).List()
	rule := sdk.PolicyRule{}
	if len(policyRuleList) > 0 {
		ruleRetyped := policyRuleList[0].(map[string]interface{})
		rule.Enabled = ruleRetyped["enabled"].(bool)
		rule.AllRules = ruleRetyped["all_rules"].(bool)
		rule.Rules = cast.ToStringSlice(ruleRetyped["rules"])
	}
	return &rule
}

func setAlertProfile(d *schema.ResourceData, meta interface{}) error {
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
