package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/settings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDefenderSettings() *schema.Resource {
	return &schema.Resource{
		Create: createDefenderSettings,
		Read:   readDefenderSettings,
		Update: createDefenderSettings,
		Delete: deleteDefenderSettings,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"admission_control_enabled": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"automatic_upgrade": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     true,
			},
			"host_custom_compliance_enabled": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"disconnect_period_days": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     1,
			},
		},
	}
}

func parseDefenderSettings(d *schema.ResourceData) *settings.DefenderSettings {
	return &settings.DefenderSettings{
		AdmissionControlEnabled:     d.Get("admission_control_enabled").(bool),
		AutomaticUpgrade:            d.Get("automatic_upgrade").(bool),
		DisconnectPeriodDays:        d.Get("disconnect_period_days").(int),
		HostCustomComplianceEnabled: d.Get("host_custom_compliance_enabled").(bool),
	}
}

func saveDefenderSettings(d *schema.ResourceData, settings *settings.DefenderSettings) error {
	d.SetId("defenderSettings")
	d.Set("admission_control_enabled", settings.AdmissionControlEnabled)
	d.Set("automatic_upgrade", settings.AutomaticUpgrade)
	d.Set("disconnect_period_days", settings.DisconnectPeriodDays)
	d.Set("host_custom_compliance_enabled", settings.HostCustomComplianceEnabled)
	return nil
}

func createDefenderSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := settings.UpdateDefenderSettings(*client, parseDefenderSettings(d))
	if err != nil {
		return err
	}

	return readDefenderSettings(d, meta)
}

func readDefenderSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	settings, err := settings.GetDefenderSettings(*client)
	if err != nil {
		return err
	}

	return saveDefenderSettings(d, settings)
}

func deleteDefenderSettings(d *schema.ResourceData, meta interface{}) error {
	return nil
}
