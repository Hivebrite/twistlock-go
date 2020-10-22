package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The API URL without the leading protocol",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_TWISTLOCK_BASE_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access key ID",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret key",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_PASSWORD", nil),
				Sensitive:   true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"twistlock_waas_container":    resourceWaasContainer(),
			"twistlock_registry_settings": resourceRegistrySettings(),
			"twistlock_tag":               resourceTag(),
			"twistlock_credential":        resourceCredentialProvider(),
			"twistlock_alert_profile":     resourceAlertProfile(),
			"twistlock_alert_settings":    resourceAlertSettings(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client, err := sdk.NewClient(d.Get("base_url").(string))
	if err != nil {
		return nil, err
	}

	err = client.Authentication(d.Get("username").(string), d.Get("password").(string))
	if err != nil {
		return nil, err
	}

	return client, nil
}
