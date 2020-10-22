package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAlertSettings() *schema.Resource {
	return &schema.Resource{
		Create: SetAlertSettings,
		Read:   readAlertSettings,
		Update: SetAlertSettings,
		Delete: deleteAlertSettings,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aggregation_period_ms": {
				Type:        schema.TypeInt,
				Description: "number of ms delay",
				Required:    true,
				ValidateFunc: validation.IntInSlice(
					[]int{
						1000,     // second
						60000,    // minute
						3600000,  // hour
						86400000, // day
					},
				),
			},
		},
	}
}

func parseAlertSettings(d *schema.ResourceData) *sdk.AlertSettings {
	d.SetId("config")
	return &sdk.AlertSettings{
		AggregationPeriodMs: d.Get("aggregation_period_ms").(int),
	}
}

func saveAlertSettings(d *schema.ResourceData, alertSettings *sdk.AlertSettings) error {
	err := d.Set("aggregation_period_ms", alertSettings.AggregationPeriodMs)
	if err != nil {
		log.Printf("[ERROR] aggregation_period_ms setting caused by: %s", err)
		return err
	}

	return nil
}

func SetAlertSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetAlertSettings(parseAlertSettings(d))

	if err != nil {
		return err
	}

	if err := readAlertSettings(d, meta); err != nil {
		log.Printf("[ERROR] readAlertSettings func caused by: %s", err)
		return err
	}

	return nil
}

func readAlertSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	alertSettings, err := client.GetAlertSettings()
	if err != nil {
		return err
	}

	return saveAlertSettings(d, alertSettings)
}

func deleteAlertSettings(d *schema.ResourceData, meta interface{}) error {
	return nil
}