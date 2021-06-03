package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/policies"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCiPolicies() *schema.Resource {
	return &schema.Resource{
		Create: createCiPolicies,
		Read:   readCiPolicies,
		Update: createCiPolicies,
		Delete: deleteCiPolicies,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"policy_type": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"effect": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"name": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"only_fixed": {
							Optional:    true,
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"collections": collectionSchema(),
						"verbose": {
							Computed:    true,
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"grace_days": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeInt,
						},
						"alert_threshold": policiesAlertThresholdSchema(),
						"block_threshold": policiesBlockThresholdSchema(),
						"cve_rules":       policiesCveRulesSchema(),
						"tags":            policiesTagsSchema(),
					},
				},
			},
		},
	}
}

func parseCiPolicies(d *schema.ResourceData) *policies.Ci {
	policiesObject := policies.Ci{
		PolicyType: "ciImagesVulnerability",
		ID:         "ciImagesVulnerability",
	}

	rules := d.Get("rules").([]interface{})

	for _, i := range rules {
		rule := i.(map[string]interface{})
		tags := rule["tags"].(*schema.Set).List()
		cveRules := rule["cve_rules"].(*schema.Set).List()

		ruleObject := policies.CiRules{
			Name:        rule["name"].(string),
			OnlyFixed:   rule["only_fixed"].(bool),
			GraceDays:   rule["grace_days"].(int),
			Collections: parseCollections(rule["collections"].(*schema.Set).List()),

			Verbose:        rule["verbose"].(bool),
			AlertThreshold: *alertThresholdFromRule(rule),
			BlockThreshold: *blockThresholdFromRule(rule),
		}

		ruleObject.ComputeEffect()

		for _, j := range cveRules {
			cveRule := j.(map[string]interface{})

			ruleObject.CveRules = append(
				ruleObject.CveRules,
				*cveRuleObjectFromInterface(cveRule),
			)
		}

		for _, j := range tags {
			tag := j.(map[string]interface{})

			ruleObject.Tags = append(
				ruleObject.Tags,
				*tagObjectFromInterface(tag),
			)
		}

		policiesObject.Rules = append(
			policiesObject.Rules,
			ruleObject,
		)
	}

	return &policiesObject
}

func saveCiPolicies(d *schema.ResourceData, policiesObject *policies.Ci) error {
	rules := make([]interface{}, 0, len(policiesObject.Rules))

	for _, i := range policiesObject.Rules {

		var tags []map[string]interface{}

		for _, tag := range i.Tags {
			var expiration []map[string]interface{}

			if tag.Expiration.Enabled {
				expiration = []map[string]interface{}{
					*expirationMapFromObject(tag.Expiration),
				}
			}

			tags = append(tags, map[string]interface{}{
				"effect":      tag.Effect,
				"name":        tag.Name,
				"description": tag.Description,
				"expiration":  expiration,
			})
		}

		var cveRules []map[string]interface{}

		for _, cveRule := range i.CveRules {
			var expiration []map[string]interface{}

			if cveRule.Expiration.Enabled {
				expiration = []map[string]interface{}{
					*expirationMapFromObject(cveRule.Expiration),
				}
			}

			cveRules = append(cveRules, map[string]interface{}{
				"effect":      cveRule.Effect,
				"id":          cveRule.ID,
				"description": cveRule.Description,
				"expiration":  expiration,
			})
		}

		rules = append(
			rules,
			map[string]interface{}{
				"name":        i.Name,
				"only_fixed":  i.OnlyFixed,
				"collections": collectionSliceToInterface(i.Collections),
				"verbose":     i.Verbose,
				"effect":      i.Effect,
				"alert_threshold": []map[string]interface{}{
					{
						"disabled": i.AlertThreshold.Disabled,
						"value":    policies.AlertingIntToLevel(i.AlertThreshold.Value),
					},
				},
				"block_threshold": []map[string]interface{}{
					{
						"enabled": i.BlockThreshold.Enabled,
						"value":   policies.AlertingIntToLevel(i.BlockThreshold.Value),
					},
				},
				"cve_rules":  cveRules,
				"tags":       tags,
				"grace_days": i.GraceDays,
			},
		)
	}

	d.SetId("ciImagesVulnerability")

	err := d.Set("rules", rules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	err = d.Set("policy_type", policiesObject.PolicyType)
	if err != nil {
		log.Printf("[ERROR] policy_type caused by: %s", err)
		return err
	}
	return nil
}

func createCiPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetCi(*client, parseCiPolicies(d))
	if err != nil {
		return err
	}

	return readCiPolicies(d, meta)
}

func readCiPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	policies, err := policies.GetCi(*client)
	if err != nil {
		return err
	}

	return saveCiPolicies(d, policies)
}

func deleteCiPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetCi(*client, &policies.Ci{
		PolicyType: "ciImagesVulnerability",
	})
	if err != nil {
		return err
	}

	return nil
}
