package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/policies"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var admissionEffects = []string{Allow, Block, Alert}

func resourceAdmissionPolicies() *schema.Resource {
	return &schema.Resource{
		Create: createAdmissionPolicies,
		Read:   readAdmissionPolicies,
		Update: createAdmissionPolicies,
		Delete: deleteAdmissionPolicies,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"effect": {
							Optional: true,
							Type:     schema.TypeString,
							Default:  Alert,
							ValidateFunc: validation.StringInSlice(
								admissionEffects,
								false,
							),
						},
						"script": {
							Required: true,
							Type:     schema.TypeString,
						},
						"name": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"description": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"skip_raw_req": {
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "",
							Default:     false,
						},
					},
				},
			},
		},
	}
}

func parseAdmissionPolicies(d *schema.ResourceData) *policies.AdmissionRules {
	policiesObject := policies.AdmissionRules{
		ID: "admission",
	}

	rules := d.Get("rules").([]interface{})

	for _, i := range rules {
		rule := i.(map[string]interface{})

		ruleObject := policies.AdmissionRule{
			Name:        rule["name"].(string),
			Script:      rule["script"].(string),
			Effect:      rule["effect"].(string),
			Description: rule["description"].(string),
			SkipRawReq:  rule["skip_raw_req"].(bool),
		}

		policiesObject.Rules = append(
			policiesObject.Rules,
			ruleObject,
		)
	}

	return &policiesObject
}

func saveAdmissionPolicies(d *schema.ResourceData, policiesObject *policies.AdmissionRules) error {
	rules := make([]interface{}, 0, len(policiesObject.Rules))

	for _, i := range policiesObject.Rules {

		rules = append(
			rules,
			map[string]interface{}{
				"name":         i.Name,
				"effect":       i.Effect,
				"script":       i.Script,
				"description":  i.Description,
				"skip_raw_req": i.SkipRawReq,
			},
		)
	}

	d.SetId("admission")

	err := d.Set("rules", rules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	return nil
}

func createAdmissionPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetAdmissionRules(*client, parseAdmissionPolicies(d))
	if err != nil {
		return err
	}

	return readAdmissionPolicies(d, meta)
}

func readAdmissionPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	policies, err := policies.GetAdmissionRules(*client)
	if err != nil {
		return err
	}

	return saveAdmissionPolicies(d, policies)
}

func deleteAdmissionPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetAdmissionRules(*client, &policies.AdmissionRules{})
	if err != nil {
		return err
	}

	return nil
}
