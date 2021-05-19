package twistlock

import (
	"log"
	"strconv"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/policies"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRuntimeCustomRule() *schema.Resource {
	return &schema.Resource{
		Create: createRuntimeCustomRule,
		Read:   readRuntimeCustomRule,
		Update: createRuntimeCustomRule,
		Delete: deleteRuntimeCustomRule,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "",
			},
			"type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "",
				ValidateFunc: validation.StringInSlice(
					[]string{
						policies.TypeFileSystem,
						policies.TypeNetworkOutgoing,
						policies.TypeProcess,
					},
					false,
				),
			},
			"message": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"script": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "",
			},
		},
	}
}

func parseRuntimeCustomRule(d *schema.ResourceData) *policies.CustomRuleRuntime {
	return &policies.CustomRuleRuntime{
		ID:          d.Get("identifier").(int),
		Type:        d.Get("type").(string),
		Message:     d.Get("message").(string),
		Description: d.Get("description").(string),
		Name:        d.Get("name").(string),
		Script:      d.Get("script").(string),
	}
}

func saveRuntimeCustomRule(d *schema.ResourceData, customRuleRuntime *policies.CustomRuleRuntime) error {
	d.SetId(strconv.Itoa(customRuleRuntime.ID))

	err := d.Set("name", customRuleRuntime.Name)
	if err != nil {
		log.Printf("[ERROR] name caused by: %s", err)
		return err
	}

	err = d.Set("identifier", customRuleRuntime.ID)
	if err != nil {
		log.Printf("[ERROR] id caused by: %s", err)
		return err
	}

	err = d.Set("message", customRuleRuntime.Message)
	if err != nil {
		log.Printf("[ERROR] message caused by: %s", err)
		return err
	}

	err = d.Set("type", customRuleRuntime.Type)
	if err != nil {
		log.Printf("[ERROR] type caused by: %s", err)
		return err
	}

	err = d.Set("description", customRuleRuntime.Description)
	if err != nil {
		log.Printf("[ERROR] description caused by: %s", err)
		return err
	}

	err = d.Set("script", customRuleRuntime.Script)
	if err != nil {
		log.Printf("[ERROR] script caused by: %s", err)
		return err
	}

	return nil
}

func createRuntimeCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := policies.SetCustomRuleRuntime(*client, parseRuntimeCustomRule(d))
	if err != nil {
		return err
	}

	return readRuntimeCustomRule(d, meta)
}

func readRuntimeCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	customRule, err := policies.GetCustomRuleRuntime(*client, d.Get("identifier").(int))
	if err != nil {
		return err
	}

	return saveRuntimeCustomRule(d, customRule)
}

func deleteRuntimeCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetCustomRuleRuntime(*client, parseRuntimeCustomRule(d))
	if err != nil {
		return err
	}

	return nil
}
