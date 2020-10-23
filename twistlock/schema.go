package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk/policies"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func alertProfilePolicySchema() *schema.Schema {
	model := &schema.Resource{
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
				Default:     true,
			},
			"rules": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "List of rules to be alerted on",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}

	return &schema.Schema{
		Optional:    true,
		Computed:    true,
		Type:        schema.TypeSet,
		MinItems:    1,
		MaxItems:    1,
		Description: "Policy definition",
		Elem:        model,
	}
}

func policiesExpirationSchema() *schema.Schema {
	return &schema.Schema{
		Optional: true,
		Type:     schema.TypeSet,
		MinItems: 0,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"date": {
					Optional: true,
					Type:     schema.TypeString,
				},
				"enabled": {
					Required: true,
					Type:     schema.TypeBool,
				},
			},
		},
	}
}

func policiesTagsSchema() *schema.Schema {
	model := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						policies.EffectAlert,
						policies.EffectBlock,
						policies.EffectIgnore,
					}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expiration": policiesExpirationSchema(),
		},
	}

	return &schema.Schema{
		Optional:    true,
		Computed:    true,
		Type:        schema.TypeSet,
		Description: "Tag Exception",
		Elem:        model,
	}
}

func policiesCveRulesSchema() *schema.Schema {
	model := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						policies.EffectAlert,
						policies.EffectBlock,
						policies.EffectIgnore,
					}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expiration": policiesExpirationSchema(),
		},
	}

	return &schema.Schema{
		Optional:    true,
		Computed:    true,
		Type:        schema.TypeSet,
		Description: "CVE Exception",
		Elem:        model,
	}
}

func policiesBlockThresholdSchema() *schema.Schema {
	var model = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						policies.Disable,
						policies.Low,
						policies.Medium,
						policies.High,
						policies.Critical,
					},
					false,
				),
			},
		},
	}
	return &schema.Schema{
		Required:    true,
		Type:        schema.TypeList,
		Description: "Policy to block",
		MinItems:    1,
		MaxItems:    1,
		Elem:        model,
	}
}

func policiesAlertThresholdSchema() *schema.Schema {
	var model = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						policies.Disable,
						policies.Low,
						policies.Medium,
						policies.High,
						policies.Critical,
					},
					false,
				),
			},
		},
	}
	return &schema.Schema{
		Required:    true,
		Type:        schema.TypeList,
		Description: "Policy to alert",
		MinItems:    1,
		MaxItems:    1,
		Elem:        model,
	}
}

func policiesResourcesSchema(keys []string) *schema.Schema {
	ruleSchema := map[string]*schema.Schema{}

	for _, i := range keys {
		ruleSchema[i] = &schema.Schema{
			Required:    true,
			Type:        schema.TypeList,
			Description: "",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		}
	}

	var model = &schema.Resource{
		Schema: ruleSchema,
	}

	return &schema.Schema{
		Type:        schema.TypeSet,
		Required:    true,
		Description: "",
		MinItems:    1,
		MaxItems:    1,
		Elem:        model,
	}
}
