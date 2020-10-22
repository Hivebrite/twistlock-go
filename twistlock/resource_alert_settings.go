package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/alerts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAlertSettings() *schema.Resource {
	return &schema.Resource{
		Create: SetSettings,
		Read:   readAlertSettings,
		Update: SetSettings,
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

func parseAlertSettings(d *schema.ResourceData) *alerts.Settings {
	d.SetId("config")
	return &alerts.Settings{
		AggregationPeriodMs: d.Get("aggregation_period_ms").(int),
	}
}

func saveAlertSettings(d *schema.ResourceData, alertSettings *alerts.Settings) error {
	err := d.Set("aggregation_period_ms", alertSettings.AggregationPeriodMs)
	if err != nil {
		log.Printf("[ERROR] aggregation_period_ms setting caused by: %s", err)
		return err
	}

	return nil
}

func SetSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := alerts.SetSettings(*client, parseAlertSettings(d))

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
	alertSettings, err := alerts.GetSettings(*client)
	if err != nil {
		return err
	}

	return saveAlertSettings(d, alertSettings)
}

func deleteAlertSettings(d *schema.ResourceData, meta interface{}) error {
	return nil
}
