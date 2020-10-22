package twistlock

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func policySchema() *schema.Schema {
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
		Type:        schema.TypeSet,
		MinItems:    1,
		MaxItems:    1,
		Description: "Policy definition",
		Elem:        model,
	}
}
